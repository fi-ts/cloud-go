// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/fi-ts/cloud-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// UpdateMachineReservationReader is a Reader for the UpdateMachineReservation structure.
type UpdateMachineReservationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateMachineReservationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateMachineReservationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateMachineReservationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateMachineReservationOK creates a UpdateMachineReservationOK with default headers values
func NewUpdateMachineReservationOK() *UpdateMachineReservationOK {
	return &UpdateMachineReservationOK{}
}

/*
UpdateMachineReservationOK describes a response with status code 200, with default header values.

Updated
*/
type UpdateMachineReservationOK struct {
	Payload *models.V1MachineReservationResponse
}

// IsSuccess returns true when this update machine reservation o k response has a 2xx status code
func (o *UpdateMachineReservationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update machine reservation o k response has a 3xx status code
func (o *UpdateMachineReservationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update machine reservation o k response has a 4xx status code
func (o *UpdateMachineReservationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update machine reservation o k response has a 5xx status code
func (o *UpdateMachineReservationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update machine reservation o k response a status code equal to that given
func (o *UpdateMachineReservationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update machine reservation o k response
func (o *UpdateMachineReservationOK) Code() int {
	return 200
}

func (o *UpdateMachineReservationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine][%d] updateMachineReservationOK %s", 200, payload)
}

func (o *UpdateMachineReservationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine][%d] updateMachineReservationOK %s", 200, payload)
}

func (o *UpdateMachineReservationOK) GetPayload() *models.V1MachineReservationResponse {
	return o.Payload
}

func (o *UpdateMachineReservationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineReservationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateMachineReservationDefault creates a UpdateMachineReservationDefault with default headers values
func NewUpdateMachineReservationDefault(code int) *UpdateMachineReservationDefault {
	return &UpdateMachineReservationDefault{
		_statusCode: code,
	}
}

/*
UpdateMachineReservationDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateMachineReservationDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update machine reservation default response has a 2xx status code
func (o *UpdateMachineReservationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update machine reservation default response has a 3xx status code
func (o *UpdateMachineReservationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update machine reservation default response has a 4xx status code
func (o *UpdateMachineReservationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update machine reservation default response has a 5xx status code
func (o *UpdateMachineReservationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update machine reservation default response a status code equal to that given
func (o *UpdateMachineReservationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update machine reservation default response
func (o *UpdateMachineReservationDefault) Code() int {
	return o._statusCode
}

func (o *UpdateMachineReservationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine][%d] updateMachineReservation default %s", o._statusCode, payload)
}

func (o *UpdateMachineReservationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/project/reservation/machine][%d] updateMachineReservation default %s", o._statusCode, payload)
}

func (o *UpdateMachineReservationDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateMachineReservationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}