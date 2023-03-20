# Computer Management Service

A small service managing computers of a company. It allows a system administrator to keep track of the computers issued by the company.

Computers can be added, updated and removed. They can be assigned to an employee by providing his abbreviation.

## Prerequisites

Needed tools:
 - [docker](https://docs.docker.com/engine/install/)
 - [docker compose v2](https://docs.docker.com/compose/install/)
    
If you are having only compose v1 installed, just replace in all commands `docker compose` with `docker-compose`.

## Quickstart

To run all the required components execute

```sh
docker compose up --build --abort-on-container-exit
```
You can then access the api docs at http://localhost:5000/api/v1/docs.

## Local development

In order build the service on you local pc without docker you need to [install go](https://go.dev/doc/install). If you want to change the REST API you need to install [go-swagger](#code-generation-via-go-swagger) as well.

To build the service execute

```sh
go build -o computer-management-service gen/controller/rest/cmd/computer-management-server/main.go
```

Then you can run the service with 
```sh
./computer-management-service --port 5000 # run with flag --help for further command line options
```

You can then access the service as in the docker setup.

In doubt see the steps in the [Dockerfile](Dockerfile) (you can ignore the `COPY` commands as the files are already in their correct location).

### Run tests

Execute 

```sh
go test ./...
```

## Code generation via go-swagger

This project uses [go-swagger](https://github.com/go-swagger/go-swagger) to generate server code from the [OpenApi 2.0 specs](api/v1/computer-management.yaml).  Install according to the [official instructions](https://goswagger.io/install.html).

Afterwards generate code with 

```sh
swagger generate server -t gen/controller/rest -f api/v1/computer-management.yaml
```

This generates the required server code for the REST API. The glue code to connect the REST API calls to the core needs to be added to [configure_computer_management.go](gen/controller/rest/restapi/configure_computer_management.go).

Note that this file is not overwritten and needs to be deleted if you need to update it due to api changes. Then simply reapply your changes afterwards.

# Development Notes

## Assumptions regarding the Task

Additionally to the specifications from the task more assumptions were made:

The computer name is the unique identifier of the computer and can't be changed. It is assumed that there is no need for an ID beside the name. All other properties of the computer can change. E.g. a MAC address can change when the network interface card is replaced.

## General Architecture/Content

This service uses the Hexagonal Architecture (Ports & Adapters). The business logic is therefore not aware the REST-API or database details and interacts with them through ports. The interface definitions for the ports can be found in the package [port](internal/core/port). This allows to e.g. switch the underlying database or to make the service available via gRPC without touching the business logic in the core.

As driving adapter/controller exists a (mostly generated) [rest-api controller](gen/controller/rest/) with additionally minor added code in the package [rest](internal/controller/rest). On the infrastructure side (i.e. driven adapters) exists a simple in memory data storage using go-maps. An adapter for a persistent database was not implemented. Furthermore exists an adapter to send notifications to the provided [admin notification service](https://github.com/greenbone/exercise-admin-notification).

## REST API definition in OpenApi 2.0

The REST API is specified using OpenApi 2.0 as this allows to use use tooling to generate server and client code which makes it more easy to interact with the service. This also ensures that the API documentation is in sync with the implementation. 

### Validation

A basic validation (i.e. checking for required values) for a computer are defined in the API specs. Although strictly speaking this is part of the business rules (and is therefore validated in the core as well). However the duplicate validation comes as a small cost as it is done by the generated code and it greatly increases the usability for implementing clients, when required values are evident from the api docs right away.

Note that the validation when creating a new computer is actually not enforced by the generated code due to a [known bug](https://github.com/go-swagger/go-swagger/issues/1904). But as mentioned the intent here was more to provide helpful API docs.

## Omissions and room for improvements

The code is fully functional and complies with the requirements from the task. 

However some basic things were omitted as this task did not ask for a complete solution:
  - unit tests (were only done for the core service)
  - adapter for a persistent database
  - CI/CD
    - to create the docker image and to run tests and allow only a merge to master on passing
    - to deploy the docker image for the service on releases, i.e. triggered by pushing a tag (and a latest version on merges to master)

Several things were omitted which should be added if this service is intended to be used in production:

- API
    - add authentication and authorization
        - e.g. using Oauth2, access to the operations could then be defined by scopes, e.g. limited available operations for a user which is only supposed to have read access
    - add pagination for the `ListComputers` operations
        - if the number registered computers becomes big an unpaginated list would lead to big response payloads and therefore a poor performance
        - if there is no pagination in the initial api it is difficult to add it later: a client using the old api does not see the `nextPage` property and assumes that it got the complete list, whereas it just got the first page
- logging
    - in the current implementation logging is almost non-existent, logs should also be sent to a log and monitoring service
    - ideally structured logging should be used (i.e. using [zerolog](https://github.com/rs/zerolog)), so that the logs can be easily queried
    - alerts should be set up on internal server errors which show a disruption of proper operation
- notification
    - currently the admin notification is dropped if there occur any errors, if the errors are 'retryable' sending the notification should be retried with a backoff
- notification implementation http-client
    - the default settings of the http-client from the net library are not suitable for a production environment, e.g. it has no timeout configured
- testing
    - beside unit-tests there should be also integration tests und end-to-tests for testing the different components in interaction in various scenarios
- validation
    - validate if MAC and IP are valid values (could be done with the [net](https://pkg.go.dev/net#ParseMAC) package)
    - it could be also ensured that no MAC is duplicate
- employee
    - the only data about the user is currently the three character abbreviation, it might become useful to use a proper foreign-key/id to connect the information of the computer management service with a user service
- deployment
    - the service could e.g. be deployed in a kubernetes cluster, for this there could be a helm chart defining the deployment
