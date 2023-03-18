// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/marius-go/computer-management-service/gen/controller/rest/models"
)

// DeleteComputerNoContentCode is the HTTP code returned for type DeleteComputerNoContent
const DeleteComputerNoContentCode int = 204

/*
DeleteComputerNoContent removed

swagger:response deleteComputerNoContent
*/
type DeleteComputerNoContent struct {
}

// NewDeleteComputerNoContent creates DeleteComputerNoContent with default headers values
func NewDeleteComputerNoContent() *DeleteComputerNoContent {

	return &DeleteComputerNoContent{}
}

// WriteResponse to the client
func (o *DeleteComputerNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteComputerNotFoundCode is the HTTP code returned for type DeleteComputerNotFound
const DeleteComputerNotFoundCode int = 404

/*
DeleteComputerNotFound the specified resource was not found

swagger:response deleteComputerNotFound
*/
type DeleteComputerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteComputerNotFound creates DeleteComputerNotFound with default headers values
func NewDeleteComputerNotFound() *DeleteComputerNotFound {

	return &DeleteComputerNotFound{}
}

// WithPayload adds the payload to the delete computer not found response
func (o *DeleteComputerNotFound) WithPayload(payload *models.Error) *DeleteComputerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete computer not found response
func (o *DeleteComputerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteComputerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*
DeleteComputerDefault error

swagger:response deleteComputerDefault
*/
type DeleteComputerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteComputerDefault creates DeleteComputerDefault with default headers values
func NewDeleteComputerDefault(code int) *DeleteComputerDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteComputerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete computer default response
func (o *DeleteComputerDefault) WithStatusCode(code int) *DeleteComputerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete computer default response
func (o *DeleteComputerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete computer default response
func (o *DeleteComputerDefault) WithPayload(payload *models.Error) *DeleteComputerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete computer default response
func (o *DeleteComputerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteComputerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
