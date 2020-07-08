// Code generated by go-swagger; DO NOT EDIT.

package accounting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/cloud-go/api/models"
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

/*VolumeUsageOK handles this case with default header values.

OK
*/
type VolumeUsageOK struct {
	Payload *models.V1VolumeUsageResponse
}

func (o *VolumeUsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsageOK  %+v", 200, o.Payload)
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

/*VolumeUsageDefault handles this case with default header values.

Error
*/
type VolumeUsageDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the volume usage default response
func (o *VolumeUsageDefault) Code() int {
	return o._statusCode
}

func (o *VolumeUsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/volume-usage][%d] volumeUsage default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
