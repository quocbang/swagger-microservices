// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"quocbang/swagger-microservices/swagger/models"
)

// UpdateAccountAuthorizationOKCode is the HTTP code returned for type UpdateAccountAuthorizationOK
const UpdateAccountAuthorizationOKCode int = 200

/*UpdateAccountAuthorizationOK OK

swagger:response updateAccountAuthorizationOK
*/
type UpdateAccountAuthorizationOK struct {
}

// NewUpdateAccountAuthorizationOK creates UpdateAccountAuthorizationOK with default headers values
func NewUpdateAccountAuthorizationOK() *UpdateAccountAuthorizationOK {

	return &UpdateAccountAuthorizationOK{}
}

// WriteResponse to the client
func (o *UpdateAccountAuthorizationOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*UpdateAccountAuthorizationDefault Unexpected error

swagger:response updateAccountAuthorizationDefault
*/
type UpdateAccountAuthorizationDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateAccountAuthorizationDefault creates UpdateAccountAuthorizationDefault with default headers values
func NewUpdateAccountAuthorizationDefault(code int) *UpdateAccountAuthorizationDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateAccountAuthorizationDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update account authorization default response
func (o *UpdateAccountAuthorizationDefault) WithStatusCode(code int) *UpdateAccountAuthorizationDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update account authorization default response
func (o *UpdateAccountAuthorizationDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update account authorization default response
func (o *UpdateAccountAuthorizationDefault) WithPayload(payload *models.Error) *UpdateAccountAuthorizationDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update account authorization default response
func (o *UpdateAccountAuthorizationDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateAccountAuthorizationDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}