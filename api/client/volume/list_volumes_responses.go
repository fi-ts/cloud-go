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

// ListVolumesReader is a Reader for the ListVolumes structure.
type ListVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListVolumesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListVolumesOK creates a ListVolumesOK with default headers values
func NewListVolumesOK() *ListVolumesOK {
	return &ListVolumesOK{}
}

/*ListVolumesOK handles this case with default header values.

OK
*/
type ListVolumesOK struct {
	Payload []*models.V1VolumeResponse
}

func (o *ListVolumesOK) Error() string {
	return fmt.Sprintf("[GET /v1/volume][%d] listVolumesOK  %+v", 200, o.Payload)
}

func (o *ListVolumesOK) GetPayload() []*models.V1VolumeResponse {
	return o.Payload
}

func (o *ListVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVolumesDefault creates a ListVolumesDefault with default headers values
func NewListVolumesDefault(code int) *ListVolumesDefault {
	return &ListVolumesDefault{
		_statusCode: code,
	}
}

/*ListVolumesDefault handles this case with default header values.

Error
*/
type ListVolumesDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the list volumes default response
func (o *ListVolumesDefault) Code() int {
	return o._statusCode
}

func (o *ListVolumesDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volume][%d] listVolumes default  %+v", o._statusCode, o.Payload)
}

func (o *ListVolumesDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListVolumesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}