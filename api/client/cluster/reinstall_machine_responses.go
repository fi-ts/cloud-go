// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ReinstallMachineReader is a Reader for the ReinstallMachine structure.
type ReinstallMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReinstallMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReinstallMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReinstallMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReinstallMachineOK creates a ReinstallMachineOK with default headers values
func NewReinstallMachineOK() *ReinstallMachineOK {
	return &ReinstallMachineOK{}
}

/*
ReinstallMachineOK describes a response with status code 200, with default header values.

OK
*/
type ReinstallMachineOK struct {
	Payload *models.V1ClusterResponse
}

// IsSuccess returns true when this reinstall machine o k response has a 2xx status code
func (o *ReinstallMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reinstall machine o k response has a 3xx status code
func (o *ReinstallMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reinstall machine o k response has a 4xx status code
func (o *ReinstallMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reinstall machine o k response has a 5xx status code
func (o *ReinstallMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reinstall machine o k response a status code equal to that given
func (o *ReinstallMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reinstall machine o k response
func (o *ReinstallMachineOK) Code() int {
	return 200
}

func (o *ReinstallMachineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/reinstallmachine][%d] reinstallMachineOK %s", 200, payload)
}

func (o *ReinstallMachineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/reinstallmachine][%d] reinstallMachineOK %s", 200, payload)
}

func (o *ReinstallMachineOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *ReinstallMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReinstallMachineDefault creates a ReinstallMachineDefault with default headers values
func NewReinstallMachineDefault(code int) *ReinstallMachineDefault {
	return &ReinstallMachineDefault{
		_statusCode: code,
	}
}

/*
ReinstallMachineDefault describes a response with status code -1, with default header values.

Error
*/
type ReinstallMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this reinstall machine default response has a 2xx status code
func (o *ReinstallMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this reinstall machine default response has a 3xx status code
func (o *ReinstallMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this reinstall machine default response has a 4xx status code
func (o *ReinstallMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this reinstall machine default response has a 5xx status code
func (o *ReinstallMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this reinstall machine default response a status code equal to that given
func (o *ReinstallMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the reinstall machine default response
func (o *ReinstallMachineDefault) Code() int {
	return o._statusCode
}

func (o *ReinstallMachineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/reinstallmachine][%d] reinstallMachine default %s", o._statusCode, payload)
}

func (o *ReinstallMachineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/reinstallmachine][%d] reinstallMachine default %s", o._statusCode, payload)
}

func (o *ReinstallMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ReinstallMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
