// Code generated by go-swagger; DO NOT EDIT.

package health

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

// HealthReader is a Reader for the Health structure.
type HealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewHealthDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewHealthOK creates a HealthOK with default headers values
func NewHealthOK() *HealthOK {
	return &HealthOK{}
}

/*
	HealthOK describes a response with status code 200, with default header values.

OK
*/
type HealthOK struct {
	Payload *models.RestHealthResponse
}

func (o *HealthOK) Error() string {
	return fmt.Sprintf("[GET /v1/health][%d] healthOK  %+v", 200, o.Payload)
}
func (o *HealthOK) GetPayload() *models.RestHealthResponse {
	return o.Payload
}

func (o *HealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestHealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthInternalServerError creates a HealthInternalServerError with default headers values
func NewHealthInternalServerError() *HealthInternalServerError {
	return &HealthInternalServerError{}
}

/*
	HealthInternalServerError describes a response with status code 500, with default header values.

Unhealthy
*/
type HealthInternalServerError struct {
	Payload *models.RestHealthResponse
}

func (o *HealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/health][%d] healthInternalServerError  %+v", 500, o.Payload)
}
func (o *HealthInternalServerError) GetPayload() *models.RestHealthResponse {
	return o.Payload
}

func (o *HealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestHealthResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthDefault creates a HealthDefault with default headers values
func NewHealthDefault(code int) *HealthDefault {
	return &HealthDefault{
		_statusCode: code,
	}
}

/*
	HealthDefault describes a response with status code -1, with default header values.

Error
*/
type HealthDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the health default response
func (o *HealthDefault) Code() int {
	return o._statusCode
}

func (o *HealthDefault) Error() string {
	return fmt.Sprintf("[GET /v1/health][%d] health default  %+v", o._statusCode, o.Payload)
}
func (o *HealthDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *HealthDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
