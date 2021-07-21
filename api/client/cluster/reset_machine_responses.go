// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// ResetMachineReader is a Reader for the ResetMachine structure.
type ResetMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResetMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResetMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewResetMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewResetMachineOK creates a ResetMachineOK with default headers values
func NewResetMachineOK() *ResetMachineOK {
	return &ResetMachineOK{}
}

/* ResetMachineOK describes a response with status code 200, with default header values.

OK
*/
type ResetMachineOK struct {
	Payload *models.V1ClusterResponse
}

func (o *ResetMachineOK) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/{id}/resetmachine][%d] resetMachineOK  %+v", 200, o.Payload)
}
func (o *ResetMachineOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *ResetMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewResetMachineDefault creates a ResetMachineDefault with default headers values
func NewResetMachineDefault(code int) *ResetMachineDefault {
	return &ResetMachineDefault{
		_statusCode: code,
	}
}

/* ResetMachineDefault describes a response with status code -1, with default header values.

Error
*/
type ResetMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the reset machine default response
func (o *ResetMachineDefault) Code() int {
	return o._statusCode
}

func (o *ResetMachineDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cluster/{id}/resetmachine][%d] resetMachine default  %+v", o._statusCode, o.Payload)
}
func (o *ResetMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ResetMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
