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

// CycleMachineReader is a Reader for the CycleMachine structure.
type CycleMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CycleMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCycleMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCycleMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCycleMachineOK creates a CycleMachineOK with default headers values
func NewCycleMachineOK() *CycleMachineOK {
	return &CycleMachineOK{}
}

/*
CycleMachineOK describes a response with status code 200, with default header values.

OK
*/
type CycleMachineOK struct {
	Payload *models.V1ClusterResponse
}

// IsSuccess returns true when this cycle machine o k response has a 2xx status code
func (o *CycleMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cycle machine o k response has a 3xx status code
func (o *CycleMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cycle machine o k response has a 4xx status code
func (o *CycleMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cycle machine o k response has a 5xx status code
func (o *CycleMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cycle machine o k response a status code equal to that given
func (o *CycleMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cycle machine o k response
func (o *CycleMachineOK) Code() int {
	return 200
}

func (o *CycleMachineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachineOK %s", 200, payload)
}

func (o *CycleMachineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachineOK %s", 200, payload)
}

func (o *CycleMachineOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *CycleMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCycleMachineDefault creates a CycleMachineDefault with default headers values
func NewCycleMachineDefault(code int) *CycleMachineDefault {
	return &CycleMachineDefault{
		_statusCode: code,
	}
}

/*
CycleMachineDefault describes a response with status code -1, with default header values.

Error
*/
type CycleMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this cycle machine default response has a 2xx status code
func (o *CycleMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cycle machine default response has a 3xx status code
func (o *CycleMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cycle machine default response has a 4xx status code
func (o *CycleMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cycle machine default response has a 5xx status code
func (o *CycleMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cycle machine default response a status code equal to that given
func (o *CycleMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cycle machine default response
func (o *CycleMachineDefault) Code() int {
	return o._statusCode
}

func (o *CycleMachineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachine default %s", o._statusCode, payload)
}

func (o *CycleMachineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster/{id}/cyclemachine][%d] cycleMachine default %s", o._statusCode, payload)
}

func (o *CycleMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *CycleMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
