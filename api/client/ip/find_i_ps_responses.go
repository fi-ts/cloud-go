// Code generated by go-swagger; DO NOT EDIT.

package ip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// FindIPsReader is a Reader for the FindIPs structure.
type FindIPsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindIPsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindIPsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindIPsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindIPsOK creates a FindIPsOK with default headers values
func NewFindIPsOK() *FindIPsOK {
	return &FindIPsOK{}
}

/*
FindIPsOK describes a response with status code 200, with default header values.

OK
*/
type FindIPsOK struct {
	Payload []*models.ModelsV1IPResponse
}

// IsSuccess returns true when this find i ps o k response has a 2xx status code
func (o *FindIPsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find i ps o k response has a 3xx status code
func (o *FindIPsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find i ps o k response has a 4xx status code
func (o *FindIPsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find i ps o k response has a 5xx status code
func (o *FindIPsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find i ps o k response a status code equal to that given
func (o *FindIPsOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindIPsOK) Error() string {
	return fmt.Sprintf("[POST /v1/ip/find][%d] findIPsOK  %+v", 200, o.Payload)
}

func (o *FindIPsOK) String() string {
	return fmt.Sprintf("[POST /v1/ip/find][%d] findIPsOK  %+v", 200, o.Payload)
}

func (o *FindIPsOK) GetPayload() []*models.ModelsV1IPResponse {
	return o.Payload
}

func (o *FindIPsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindIPsDefault creates a FindIPsDefault with default headers values
func NewFindIPsDefault(code int) *FindIPsDefault {
	return &FindIPsDefault{
		_statusCode: code,
	}
}

/*
FindIPsDefault describes a response with status code -1, with default header values.

Error
*/
type FindIPsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find i ps default response
func (o *FindIPsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find i ps default response has a 2xx status code
func (o *FindIPsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find i ps default response has a 3xx status code
func (o *FindIPsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find i ps default response has a 4xx status code
func (o *FindIPsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find i ps default response has a 5xx status code
func (o *FindIPsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find i ps default response a status code equal to that given
func (o *FindIPsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindIPsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ip/find][%d] findIPs default  %+v", o._statusCode, o.Payload)
}

func (o *FindIPsDefault) String() string {
	return fmt.Sprintf("[POST /v1/ip/find][%d] findIPs default  %+v", o._statusCode, o.Payload)
}

func (o *FindIPsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindIPsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
