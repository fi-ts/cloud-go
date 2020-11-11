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

/*UpdateClusterOK handles this case with default header values.

OK
*/
type UpdateClusterOK struct {
	Payload *models.V1ClusterResponse
}

func (o *UpdateClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterOK  %+v", 200, o.Payload)
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

/*UpdateClusterConflict handles this case with default header values.

Conflict
*/
type UpdateClusterConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *UpdateClusterConflict) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] updateClusterConflict  %+v", 409, o.Payload)
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

/*UpdateClusterDefault handles this case with default header values.

Error
*/
type UpdateClusterDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the update cluster default response
func (o *UpdateClusterDefault) Code() int {
	return o._statusCode
}

func (o *UpdateClusterDefault) Error() string {
	return fmt.Sprintf("[POST /v1/cluster][%d] updateCluster default  %+v", o._statusCode, o.Payload)
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
