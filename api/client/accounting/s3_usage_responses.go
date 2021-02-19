// Code generated by go-swagger; DO NOT EDIT.

package accounting

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

// S3UsageReader is a Reader for the S3Usage structure.
type S3UsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3UsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3UsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3UsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3UsageOK creates a S3UsageOK with default headers values
func NewS3UsageOK() *S3UsageOK {
	return &S3UsageOK{}
}

/*S3UsageOK handles this case with default header values.

OK
*/
type S3UsageOK struct {
	Payload *models.V1S3UsageResponse
}

func (o *S3UsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/s3-usage][%d] s3UsageOK  %+v", 200, o.Payload)
}

func (o *S3UsageOK) GetPayload() *models.V1S3UsageResponse {
	return o.Payload
}

func (o *S3UsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1S3UsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3UsageDefault creates a S3UsageDefault with default headers values
func NewS3UsageDefault(code int) *S3UsageDefault {
	return &S3UsageDefault{
		_statusCode: code,
	}
}

/*S3UsageDefault handles this case with default header values.

Error
*/
type S3UsageDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the s3 usage default response
func (o *S3UsageDefault) Code() int {
	return o._statusCode
}

func (o *S3UsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/s3-usage][%d] s3Usage default  %+v", o._statusCode, o.Payload)
}

func (o *S3UsageDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *S3UsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
