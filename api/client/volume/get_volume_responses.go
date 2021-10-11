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

// GetVolumeReader is a Reader for the GetVolume structure.
type GetVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVolumeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVolumeOK creates a GetVolumeOK with default headers values
func NewGetVolumeOK() *GetVolumeOK {
	return &GetVolumeOK{}
}

/* GetVolumeOK describes a response with status code 200, with default header values.

OK
*/
type GetVolumeOK struct {
	Payload *models.V1VolumeResponse
}

func (o *GetVolumeOK) Error() string {
	return fmt.Sprintf("[GET /v1/volume/{id}][%d] getVolumeOK  %+v", 200, o.Payload)
}
func (o *GetVolumeOK) GetPayload() *models.V1VolumeResponse {
	return o.Payload
}

func (o *GetVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VolumeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumeDefault creates a GetVolumeDefault with default headers values
func NewGetVolumeDefault(code int) *GetVolumeDefault {
	return &GetVolumeDefault{
		_statusCode: code,
	}
}

/* GetVolumeDefault describes a response with status code -1, with default header values.

Error
*/
type GetVolumeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get volume default response
func (o *GetVolumeDefault) Code() int {
	return o._statusCode
}

func (o *GetVolumeDefault) Error() string {
	return fmt.Sprintf("[GET /v1/volume/{id}][%d] getVolume default  %+v", o._statusCode, o.Payload)
}
func (o *GetVolumeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetVolumeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}