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

// IPUsageReader is a Reader for the IPUsage structure.
type IPUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewIPUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIPUsageOK creates a IPUsageOK with default headers values
func NewIPUsageOK() *IPUsageOK {
	return &IPUsageOK{}
}

/*IPUsageOK handles this case with default header values.

OK
*/
type IPUsageOK struct {
	Payload *models.V1IPUsageResponse
}

func (o *IPUsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/ip-usage][%d] ipUsageOK  %+v", 200, o.Payload)
}

func (o *IPUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1IPUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIPUsageDefault creates a IPUsageDefault with default headers values
func NewIPUsageDefault(code int) *IPUsageDefault {
	return &IPUsageDefault{
		_statusCode: code,
	}
}

/*IPUsageDefault handles this case with default header values.

Error
*/
type IPUsageDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the ip usage default response
func (o *IPUsageDefault) Code() int {
	return o._statusCode
}

func (o *IPUsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/ip-usage][%d] ipUsage default  %+v", o._statusCode, o.Payload)
}

func (o *IPUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
