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

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewUpdateClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateClusterOK creates a UpdateClusterOK with default headers values
func NewUpdateClusterOK() *UpdateClusterOK {
	return &UpdateClusterOK{}
}

/*
UpdateClusterOK describes a response with status code 200, with default header values.

OK
*/
type UpdateClusterOK struct {
	Payload *models.V1ClusterResponse
}

// IsSuccess returns true when this update cluster o k response has a 2xx status code
func (o *UpdateClusterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update cluster o k response has a 3xx status code
func (o *UpdateClusterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster o k response has a 4xx status code
func (o *UpdateClusterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update cluster o k response has a 5xx status code
func (o *UpdateClusterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster o k response a status code equal to that given
func (o *UpdateClusterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update cluster o k response
func (o *UpdateClusterOK) Code() int {
	return 200
}

func (o *UpdateClusterOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterOK %s", 200, payload)
}

func (o *UpdateClusterOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterOK %s", 200, payload)
}

func (o *UpdateClusterOK) GetPayload() *models.V1ClusterResponse {
	return o.Payload
}

func (o *UpdateClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ClusterResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterConflict creates a UpdateClusterConflict with default headers values
func NewUpdateClusterConflict() *UpdateClusterConflict {
	return &UpdateClusterConflict{}
}

/*
UpdateClusterConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateClusterConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update cluster conflict response has a 2xx status code
func (o *UpdateClusterConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update cluster conflict response has a 3xx status code
func (o *UpdateClusterConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update cluster conflict response has a 4xx status code
func (o *UpdateClusterConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update cluster conflict response has a 5xx status code
func (o *UpdateClusterConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update cluster conflict response a status code equal to that given
func (o *UpdateClusterConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the update cluster conflict response
func (o *UpdateClusterConflict) Code() int {
	return 409
}

func (o *UpdateClusterConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterConflict %s", 409, payload)
}

func (o *UpdateClusterConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterConflict %s", 409, payload)
}

func (o *UpdateClusterConflict) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterDefault creates a UpdateClusterDefault with default headers values
func NewUpdateClusterDefault(code int) *UpdateClusterDefault {
	return &UpdateClusterDefault{
		_statusCode: code,
	}
}

/*
UpdateClusterDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateClusterDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this update cluster default response has a 2xx status code
func (o *UpdateClusterDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this update cluster default response has a 3xx status code
func (o *UpdateClusterDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this update cluster default response has a 4xx status code
func (o *UpdateClusterDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this update cluster default response has a 5xx status code
func (o *UpdateClusterDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this update cluster default response a status code equal to that given
func (o *UpdateClusterDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the update cluster default response
func (o *UpdateClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateCluster default %s", o._statusCode, payload)
}

func (o *UpdateClusterDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/cluster][%d] updateCluster default %s", o._statusCode, payload)
}

func (o *UpdateClusterDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
