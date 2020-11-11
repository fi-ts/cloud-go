// Code generated by go-swagger; DO NOT EDIT.

package accounting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/httperrors"
)

// S3UsageCSVReader is a Reader for the S3UsageCSV structure.
type S3UsageCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3UsageCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3UsageCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3UsageCSVDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3UsageCSVOK creates a S3UsageCSVOK with default headers values
func NewS3UsageCSVOK() *S3UsageCSVOK {
	return &S3UsageCSVOK{}
}

/*S3UsageCSVOK handles this case with default header values.

OK
*/
type S3UsageCSVOK struct {
	Payload string
}

func (o *S3UsageCSVOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/s3-usage-csv][%d] s3UsageCSVOK  %+v", 200, o.Payload)
}

func (o *S3UsageCSVOK) GetPayload() string {
	return o.Payload
}

func (o *S3UsageCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3UsageCSVDefault creates a S3UsageCSVDefault with default headers values
func NewS3UsageCSVDefault(code int) *S3UsageCSVDefault {
	return &S3UsageCSVDefault{
		_statusCode: code,
	}
}

/*S3UsageCSVDefault handles this case with default header values.

Error
*/
type S3UsageCSVDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the s3 usage c s v default response
func (o *S3UsageCSVDefault) Code() int {
	return o._statusCode
}

func (o *S3UsageCSVDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/s3-usage-csv][%d] s3UsageCSV default  %+v", o._statusCode, o.Payload)
}

func (o *S3UsageCSVDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *S3UsageCSVDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
