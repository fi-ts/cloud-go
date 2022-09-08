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

// FreeIPReader is a Reader for the FreeIP structure.
type FreeIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FreeIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFreeIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFreeIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFreeIPOK creates a FreeIPOK with default headers values
func NewFreeIPOK() *FreeIPOK {
	return &FreeIPOK{}
}

/*
FreeIPOK describes a response with status code 200, with default header values.

OK
*/
type FreeIPOK struct {
	Payload *models.ModelsV1IPResponse
}

// IsSuccess returns true when this free Ip o k response has a 2xx status code
func (o *FreeIPOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this free Ip o k response has a 3xx status code
func (o *FreeIPOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this free Ip o k response has a 4xx status code
func (o *FreeIPOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this free Ip o k response has a 5xx status code
func (o *FreeIPOK) IsServerError() bool {
	return false
}

// IsCode returns true when this free Ip o k response a status code equal to that given
func (o *FreeIPOK) IsCode(code int) bool {
	return code == 200
}

func (o *FreeIPOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/ip/{ip}][%d] freeIpOK  %+v", 200, o.Payload)
}

func (o *FreeIPOK) String() string {
	return fmt.Sprintf("[DELETE /v1/ip/{ip}][%d] freeIpOK  %+v", 200, o.Payload)
}

func (o *FreeIPOK) GetPayload() *models.ModelsV1IPResponse {
	return o.Payload
}

func (o *FreeIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsV1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeIPDefault creates a FreeIPDefault with default headers values
func NewFreeIPDefault(code int) *FreeIPDefault {
	return &FreeIPDefault{
		_statusCode: code,
	}
}

/*
FreeIPDefault describes a response with status code -1, with default header values.

Error
*/
type FreeIPDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the free IP default response
func (o *FreeIPDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this free IP default response has a 2xx status code
func (o *FreeIPDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this free IP default response has a 3xx status code
func (o *FreeIPDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this free IP default response has a 4xx status code
func (o *FreeIPDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this free IP default response has a 5xx status code
func (o *FreeIPDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this free IP default response a status code equal to that given
func (o *FreeIPDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FreeIPDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/ip/{ip}][%d] freeIP default  %+v", o._statusCode, o.Payload)
}

func (o *FreeIPDefault) String() string {
	return fmt.Sprintf("[DELETE /v1/ip/{ip}][%d] freeIP default  %+v", o._statusCode, o.Payload)
}

func (o *FreeIPDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FreeIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
