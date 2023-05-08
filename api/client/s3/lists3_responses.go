// Code generated by go-swagger; DO NOT EDIT.

package s3

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

// Lists3Reader is a Reader for the Lists3 structure.
type Lists3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Lists3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLists3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLists3Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLists3OK creates a Lists3OK with default headers values
func NewLists3OK() *Lists3OK {
	return &Lists3OK{}
}

/*
Lists3OK describes a response with status code 200, with default header values.

OK
*/
type Lists3OK struct {
	Payload []*models.V1S3Response
}

// IsSuccess returns true when this lists3 o k response has a 2xx status code
func (o *Lists3OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lists3 o k response has a 3xx status code
func (o *Lists3OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists3 o k response has a 4xx status code
func (o *Lists3OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists3 o k response has a 5xx status code
func (o *Lists3OK) IsServerError() bool {
	return false
}

// IsCode returns true when this lists3 o k response a status code equal to that given
func (o *Lists3OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lists3 o k response
func (o *Lists3OK) Code() int {
	return 200
}

func (o *Lists3OK) Error() string {
	return fmt.Sprintf("[GET /v1/s3/list][%d] lists3OK  %+v", 200, o.Payload)
}

func (o *Lists3OK) String() string {
	return fmt.Sprintf("[GET /v1/s3/list][%d] lists3OK  %+v", 200, o.Payload)
}

func (o *Lists3OK) GetPayload() []*models.V1S3Response {
	return o.Payload
}

func (o *Lists3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLists3Default creates a Lists3Default with default headers values
func NewLists3Default(code int) *Lists3Default {
	return &Lists3Default{
		_statusCode: code,
	}
}

/*
Lists3Default describes a response with status code -1, with default header values.

Error
*/
type Lists3Default struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this lists3 default response has a 2xx status code
func (o *Lists3Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lists3 default response has a 3xx status code
func (o *Lists3Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lists3 default response has a 4xx status code
func (o *Lists3Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lists3 default response has a 5xx status code
func (o *Lists3Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lists3 default response a status code equal to that given
func (o *Lists3Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lists3 default response
func (o *Lists3Default) Code() int {
	return o._statusCode
}

func (o *Lists3Default) Error() string {
	return fmt.Sprintf("[GET /v1/s3/list][%d] lists3 default  %+v", o._statusCode, o.Payload)
}

func (o *Lists3Default) String() string {
	return fmt.Sprintf("[GET /v1/s3/list][%d] lists3 default  %+v", o._statusCode, o.Payload)
}

func (o *Lists3Default) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *Lists3Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
