// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// FindVolumesReader is a Reader for the FindVolumes structure.
type FindVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindVolumesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindVolumesOK creates a FindVolumesOK with default headers values
func NewFindVolumesOK() *FindVolumesOK {
	return &FindVolumesOK{}
}

/*
FindVolumesOK describes a response with status code 200, with default header values.

OK
*/
type FindVolumesOK struct {
	Payload []*models.V1VolumeResponse
}

// IsSuccess returns true when this find volumes o k response has a 2xx status code
func (o *FindVolumesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this find volumes o k response has a 3xx status code
func (o *FindVolumesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this find volumes o k response has a 4xx status code
func (o *FindVolumesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this find volumes o k response has a 5xx status code
func (o *FindVolumesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this find volumes o k response a status code equal to that given
func (o *FindVolumesOK) IsCode(code int) bool {
	return code == 200
}

func (o *FindVolumesOK) Error() string {
	return fmt.Sprintf("[POST /v1/volume/find][%d] findVolumesOK  %+v", 200, o.Payload)
}

func (o *FindVolumesOK) String() string {
	return fmt.Sprintf("[POST /v1/volume/find][%d] findVolumesOK  %+v", 200, o.Payload)
}

func (o *FindVolumesOK) GetPayload() []*models.V1VolumeResponse {
	return o.Payload
}

func (o *FindVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindVolumesDefault creates a FindVolumesDefault with default headers values
func NewFindVolumesDefault(code int) *FindVolumesDefault {
	return &FindVolumesDefault{
		_statusCode: code,
	}
}

/*
FindVolumesDefault describes a response with status code -1, with default header values.

Error
*/
type FindVolumesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find volumes default response
func (o *FindVolumesDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this find volumes default response has a 2xx status code
func (o *FindVolumesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this find volumes default response has a 3xx status code
func (o *FindVolumesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this find volumes default response has a 4xx status code
func (o *FindVolumesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this find volumes default response has a 5xx status code
func (o *FindVolumesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this find volumes default response a status code equal to that given
func (o *FindVolumesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FindVolumesDefault) Error() string {
	return fmt.Sprintf("[POST /v1/volume/find][%d] findVolumes default  %+v", o._statusCode, o.Payload)
}

func (o *FindVolumesDefault) String() string {
	return fmt.Sprintf("[POST /v1/volume/find][%d] findVolumes default  %+v", o._statusCode, o.Payload)
}

func (o *FindVolumesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindVolumesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
