// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// VolumeUsageReader is a Reader for the VolumeUsage structure.
type VolumeUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeUsageOK creates a VolumeUsageOK with default headers values
func NewVolumeUsageOK() *VolumeUsageOK {
	return &VolumeUsageOK{}
}

/*
VolumeUsageOK describes a response with status code 200, with default header values.

OK
*/
type VolumeUsageOK struct {
	Payload *models.V1VolumeUsageResponse
}

// IsSuccess returns true when this volume usage o k response has a 2xx status code
func (o *VolumeUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume usage o k response has a 3xx status code
func (o *VolumeUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume usage o k response has a 4xx status code
func (o *VolumeUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume usage o k response has a 5xx status code
func (o *VolumeUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this volume usage o k response a status code equal to that given
func (o *VolumeUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the volume usage o k response
func (o *VolumeUsageOK) Code() int {
	return 200
}

func (o *VolumeUsageOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsageOK %s", 200, payload)
}

func (o *VolumeUsageOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsageOK %s", 200, payload)
}

func (o *VolumeUsageOK) GetPayload() *models.V1VolumeUsageResponse {
	return o.Payload
}

func (o *VolumeUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VolumeUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeUsageDefault creates a VolumeUsageDefault with default headers values
func NewVolumeUsageDefault(code int) *VolumeUsageDefault {
	return &VolumeUsageDefault{
		_statusCode: code,
	}
}

/*
VolumeUsageDefault describes a response with status code -1, with default header values.

Error
*/
type VolumeUsageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this volume usage default response has a 2xx status code
func (o *VolumeUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume usage default response has a 3xx status code
func (o *VolumeUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume usage default response has a 4xx status code
func (o *VolumeUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume usage default response has a 5xx status code
func (o *VolumeUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume usage default response a status code equal to that given
func (o *VolumeUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the volume usage default response
func (o *VolumeUsageDefault) Code() int {
	return o._statusCode
}

func (o *VolumeUsageDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsage default %s", o._statusCode, payload)
}

func (o *VolumeUsageDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsage default %s", o._statusCode, payload)
}

func (o *VolumeUsageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *VolumeUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
