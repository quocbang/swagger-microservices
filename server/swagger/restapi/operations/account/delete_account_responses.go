// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"quocbang/swagger-microservices/swagger/models"
)

// DeleteAccountOKCode is the HTTP code returned for type DeleteAccountOK
const DeleteAccountOKCode int = 200

/*DeleteAccountOK OK

swagger:response deleteAccountOK
*/
type DeleteAccountOK struct {
}

// NewDeleteAccountOK creates DeleteAccountOK with default headers values
func NewDeleteAccountOK() *DeleteAccountOK {

	return &DeleteAccountOK{}
}

// WriteResponse to the client
func (o *DeleteAccountOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

/*DeleteAccountDefault Unexpected error

swagger:response deleteAccountDefault
*/
type DeleteAccountDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteAccountDefault creates DeleteAccountDefault with default headers values
func NewDeleteAccountDefault(code int) *DeleteAccountDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteAccountDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete account default response
func (o *DeleteAccountDefault) WithStatusCode(code int) *DeleteAccountDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete account default response
func (o *DeleteAccountDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete account default response
func (o *DeleteAccountDefault) WithPayload(payload *models.Error) *DeleteAccountDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete account default response
func (o *DeleteAccountDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteAccountDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
