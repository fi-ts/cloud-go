// Code generated by go-swagger; DO NOT EDIT.

package s3

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

// Lists3partitionsReader is a Reader for the Lists3partitions structure.
type Lists3partitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *Lists3partitionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLists3partitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLists3partitionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLists3partitionsOK creates a Lists3partitionsOK with default headers values
func NewLists3partitionsOK() *Lists3partitionsOK {
	return &Lists3partitionsOK{}
}

/*
Lists3partitionsOK describes a response with status code 200, with default header values.

OK
*/
type Lists3partitionsOK struct {
	Payload []*models.V1S3PartitionResponse
}

// IsSuccess returns true when this lists3partitions o k response has a 2xx status code
func (o *Lists3partitionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lists3partitions o k response has a 3xx status code
func (o *Lists3partitionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lists3partitions o k response has a 4xx status code
func (o *Lists3partitionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lists3partitions o k response has a 5xx status code
func (o *Lists3partitionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lists3partitions o k response a status code equal to that given
func (o *Lists3partitionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the lists3partitions o k response
func (o *Lists3partitionsOK) Code() int {
	return 200
}

func (o *Lists3partitionsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/s3/partitions][%d] lists3partitionsOK %s", 200, payload)
}

func (o *Lists3partitionsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/s3/partitions][%d] lists3partitionsOK %s", 200, payload)
}

func (o *Lists3partitionsOK) GetPayload() []*models.V1S3PartitionResponse {
	return o.Payload
}

func (o *Lists3partitionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLists3partitionsDefault creates a Lists3partitionsDefault with default headers values
func NewLists3partitionsDefault(code int) *Lists3partitionsDefault {
	return &Lists3partitionsDefault{
		_statusCode: code,
	}
}

/*
Lists3partitionsDefault describes a response with status code -1, with default header values.

Error
*/
type Lists3partitionsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this lists3partitions default response has a 2xx status code
func (o *Lists3partitionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lists3partitions default response has a 3xx status code
func (o *Lists3partitionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lists3partitions default response has a 4xx status code
func (o *Lists3partitionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lists3partitions default response has a 5xx status code
func (o *Lists3partitionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lists3partitions default response a status code equal to that given
func (o *Lists3partitionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the lists3partitions default response
func (o *Lists3partitionsDefault) Code() int {
	return o._statusCode
}

func (o *Lists3partitionsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/s3/partitions][%d] lists3partitions default %s", o._statusCode, payload)
}

func (o *Lists3partitionsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/s3/partitions][%d] lists3partitions default %s", o._statusCode, payload)
}

func (o *Lists3partitionsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *Lists3partitionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
