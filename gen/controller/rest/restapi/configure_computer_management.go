// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	goerrors "errors"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"github.com/marius-go/computer-management-service/gen/controller/rest/models"
	"github.com/marius-go/computer-management-service/gen/controller/rest/restapi/operations"
	"github.com/marius-go/computer-management-service/gen/controller/rest/restapi/operations/computer"
	"github.com/marius-go/computer-management-service/internal/controller/rest"
	"github.com/marius-go/computer-management-service/internal/core/domain"
	computersrv "github.com/marius-go/computer-management-service/internal/core/service/computer"
	"github.com/marius-go/computer-management-service/internal/infrastructure/adminnotifier"
	"github.com/marius-go/computer-management-service/internal/infrastructure/memstorage"
)

//go:generate swagger generate server --target ../../gen --name ComputerManagement --spec ../../api/v1/computer-management.yaml --principal interface{}

var configurationFlags = struct {
	NotificationServiceAddress string `long:"notification-service" description:"optional: address of the admin notifications service" default:"http://localhost:8080"`
}{}

func configureFlags(api *operations.ComputerManagementAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Computer Service Flags",
			LongDescription:  "",
			Options:          &configurationFlags,
		},
	}
}

func configureAPI(api *operations.ComputerManagementAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	storage := memstorage.New()
	adminNotifier := adminnotifier.New(&http.Client{}, configurationFlags.NotificationServiceAddress)
	computerService := computersrv.New(storage, adminNotifier)

	api.ComputerCreateComputerHandler = computer.CreateComputerHandlerFunc(
		func(params computer.CreateComputerParams) middleware.Responder {

			comp := rest.NewComputerRestModelToDomain(*params.Computer)

			insertedComputer, err := computerService.CreateComputer(comp)
			if err != nil {
				errorMessage := err.Error()
				errPayload := models.Error{Message: &errorMessage}
				var errConflict domain.ErrConflict
				var errValidation domain.ErrValidation
				switch {
				case goerrors.As(err, &errConflict):
					return computer.NewCreateComputerConflict().WithPayload(&errPayload)
				case goerrors.As(err, &errValidation):
					return computer.NewCreateComputerBadRequest().WithPayload(&errPayload)
				default:
					return computer.NewCreateComputerDefault(500).WithPayload(&errPayload)
				}
			}

			responsePayload := rest.ComputerRestModelFromDomain(insertedComputer)
			return computer.NewCreateComputerCreated().WithPayload(&responsePayload)
		},
	)

	api.ComputerDeleteComputerHandler = computer.DeleteComputerHandlerFunc(
		func(params computer.DeleteComputerParams) middleware.Responder {

			err := computerService.DeleteComputer(params.ComputerName)
			if err != nil {
				errorMessage := err.Error()
				errPayload := models.Error{Message: &errorMessage}
				switch {
				case goerrors.Is(err, domain.ErrNotFound):
					return computer.NewDeleteComputerNotFound().WithPayload(&errPayload)
				default:
					return computer.NewDeleteComputerDefault(500).WithPayload(&errPayload)
				}
			}

			return computer.NewDeleteComputerNoContent()
		},
	)

	api.ComputerGetComputerHandler = computer.GetComputerHandlerFunc(
		func(params computer.GetComputerParams) middleware.Responder {

			comp, err := computerService.GetComputer(params.ComputerName)
			if err != nil {
				errorMessage := err.Error()
				errPayload := models.Error{Message: &errorMessage}
				switch {
				case goerrors.Is(err, domain.ErrNotFound):
					return computer.NewGetComputerNotFound().WithPayload(&errPayload)
				default:
					return computer.NewGetComputerDefault(500).WithPayload(&errPayload)
				}
			}

			responsePayload := rest.ComputerRestModelFromDomain(comp)
			return computer.NewGetComputerOK().WithPayload(&responsePayload)
		})

	api.ComputerListComputersHandler = computer.ListComputersHandlerFunc(func(params computer.ListComputersParams) middleware.Responder {

		employeeAbbreviation := params.EmployeeAbbreviation
		if employeeAbbreviation != nil && (*employeeAbbreviation == "\"\"" || *employeeAbbreviation == "''") { // "" and '' should be interpreted as empty string
			*employeeAbbreviation = ""
		}

		computers, err := computerService.ListComputers(employeeAbbreviation)
		if err != nil {
			errorMessage := err.Error()
			errPayload := models.Error{Message: &errorMessage}
			computer.NewCreateComputerDefault(500).WithPayload(&errPayload)
		}

		responsePayload := make([]*models.Computer, 0, len(computers))
		for _, comp := range computers {
			restModel := rest.ComputerRestModelFromDomain(comp)
			responsePayload = append(responsePayload, &restModel)
		}
		return computer.NewListComputersOK().WithPayload(responsePayload)
	})

	api.ComputerUpdateComputerHandler = computer.UpdateComputerHandlerFunc(func(params computer.UpdateComputerParams) middleware.Responder {

		var propertiesToUpdate []string
		if params.UpdateMask != nil {
			propertiesToUpdate = strings.Split(*params.UpdateMask, ",")
		}

		comp := rest.ComputerRestModelToDomain(*params.Computer, propertiesToUpdate)

		updatedComputer, err := computerService.UpdateComputer(comp)
		if err != nil {
			errorMessage := err.Error()
			errPayload := models.Error{Message: &errorMessage}
			var errValidation domain.ErrValidation
			switch {
			case goerrors.As(err, &errValidation):
				return computer.NewUpdateComputerBadRequest().WithPayload(&errPayload)
			case goerrors.Is(err, domain.ErrNotFound):
				return computer.NewUpdateComputerNotFound().WithPayload(&errPayload)
			default:
				return computer.NewUpdateComputerDefault(500).WithPayload(&errPayload)
			}

		}

		responsePayload := rest.ComputerRestModelFromDomain(updatedComputer)
		return computer.NewUpdateComputerOK().WithPayload(&responsePayload)
	})

	if api.ComputerCreateComputerHandler == nil {
		api.ComputerCreateComputerHandler = computer.CreateComputerHandlerFunc(func(params computer.CreateComputerParams) middleware.Responder {
			return middleware.NotImplemented("operation computer.CreateComputer has not yet been implemented")
		})
	}
	if api.ComputerDeleteComputerHandler == nil {
		api.ComputerDeleteComputerHandler = computer.DeleteComputerHandlerFunc(func(params computer.DeleteComputerParams) middleware.Responder {
			return middleware.NotImplemented("operation computer.DeleteComputer has not yet been implemented")
		})
	}
	if api.ComputerGetComputerHandler == nil {
		api.ComputerGetComputerHandler = computer.GetComputerHandlerFunc(func(params computer.GetComputerParams) middleware.Responder {
			return middleware.NotImplemented("operation computer.GetComputer has not yet been implemented")
		})
	}
	if api.ComputerListComputersHandler == nil {
		api.ComputerListComputersHandler = computer.ListComputersHandlerFunc(func(params computer.ListComputersParams) middleware.Responder {
			return middleware.NotImplemented("operation computer.ListComputers has not yet been implemented")
		})
	}
	if api.ComputerUpdateComputerHandler == nil {
		api.ComputerUpdateComputerHandler = computer.UpdateComputerHandlerFunc(func(params computer.UpdateComputerParams) middleware.Responder {
			return middleware.NotImplemented("operation computer.UpdateComputer has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
