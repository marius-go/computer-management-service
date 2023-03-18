// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/marius-go/computer-management-service/gen/controller/rest/models"
)

// UpdateComputerOKCode is the HTTP code returned for type UpdateComputerOK
const UpdateComputerOKCode int = 200

/*
UpdateComputerOK updated

swagger:response updateComputerOK
*/
type UpdateComputerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Computer `json:"body,omitempty"`
}

// NewUpdateComputerOK creates UpdateComputerOK with default headers values
func NewUpdateComputerOK() *UpdateComputerOK {

	return &UpdateComputerOK{}
}

// WithPayload adds the payload to the update computer o k response
func (o *UpdateComputerOK) WithPayload(payload *models.Computer) *UpdateComputerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update computer o k response
func (o *UpdateComputerOK) SetPayload(payload *models.Computer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateComputerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateComputerBadRequestCode is the HTTP code returned for type UpdateComputerBadRequest
const UpdateComputerBadRequestCode int = 400

/*
UpdateComputerBadRequest bad request

swagger:response updateComputerBadRequest
*/
type UpdateComputerBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateComputerBadRequest creates UpdateComputerBadRequest with default headers values
func NewUpdateComputerBadRequest() *UpdateComputerBadRequest {

	return &UpdateComputerBadRequest{}
}

// WithPayload adds the payload to the update computer bad request response
func (o *UpdateComputerBadRequest) WithPayload(payload *models.Error) *UpdateComputerBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update computer bad request response
func (o *UpdateComputerBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateComputerBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdateComputerNotFoundCode is the HTTP code returned for type UpdateComputerNotFound
const UpdateComputerNotFoundCode int = 404

/*
UpdateComputerNotFound the specified resource was not found

swagger:response updateComputerNotFound
*/
type UpdateComputerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateComputerNotFound creates UpdateComputerNotFound with default headers values
func NewUpdateComputerNotFound() *UpdateComputerNotFound {

	return &UpdateComputerNotFound{}
}

// WithPayload adds the payload to the update computer not found response
func (o *UpdateComputerNotFound) WithPayload(payload *models.Error) *UpdateComputerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update computer not found response
func (o *UpdateComputerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateComputerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
UpdateComputerDefault error

swagger:response updateComputerDefault
*/
type UpdateComputerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateComputerDefault creates UpdateComputerDefault with default headers values
func NewUpdateComputerDefault(code int) *UpdateComputerDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateComputerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update computer default response
func (o *UpdateComputerDefault) WithStatusCode(code int) *UpdateComputerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update computer default response
func (o *UpdateComputerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update computer default response
func (o *UpdateComputerDefault) WithPayload(payload *models.Error) *UpdateComputerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update computer default response
func (o *UpdateComputerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateComputerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
