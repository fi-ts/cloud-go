// Code generated by go-swagger; DO NOT EDIT.

package volume

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

// GetSnapshotReader is a Reader for the GetSnapshot structure.
type GetSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetSnapshotDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSnapshotOK creates a GetSnapshotOK with default headers values
func NewGetSnapshotOK() *GetSnapshotOK {
	return &GetSnapshotOK{}
}

/*
GetSnapshotOK describes a response with status code 200, with default header values.

OK
*/
type GetSnapshotOK struct {
	Payload *models.V1SnapshotResponse
}

// IsSuccess returns true when this get snapshot o k response has a 2xx status code
func (o *GetSnapshotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get snapshot o k response has a 3xx status code
func (o *GetSnapshotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get snapshot o k response has a 4xx status code
func (o *GetSnapshotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get snapshot o k response has a 5xx status code
func (o *GetSnapshotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get snapshot o k response a status code equal to that given
func (o *GetSnapshotOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get snapshot o k response
func (o *GetSnapshotOK) Code() int {
	return 200
}

func (o *GetSnapshotOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume/snapshot/{id}][%d] getSnapshotOK %s", 200, payload)
}

func (o *GetSnapshotOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume/snapshot/{id}][%d] getSnapshotOK %s", 200, payload)
}

func (o *GetSnapshotOK) GetPayload() *models.V1SnapshotResponse {
	return o.Payload
}

func (o *GetSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1SnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSnapshotDefault creates a GetSnapshotDefault with default headers values
func NewGetSnapshotDefault(code int) *GetSnapshotDefault {
	return &GetSnapshotDefault{
		_statusCode: code,
	}
}

/*
GetSnapshotDefault describes a response with status code -1, with default header values.

Error
*/
type GetSnapshotDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this get snapshot default response has a 2xx status code
func (o *GetSnapshotDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get snapshot default response has a 3xx status code
func (o *GetSnapshotDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get snapshot default response has a 4xx status code
func (o *GetSnapshotDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get snapshot default response has a 5xx status code
func (o *GetSnapshotDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get snapshot default response a status code equal to that given
func (o *GetSnapshotDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the get snapshot default response
func (o *GetSnapshotDefault) Code() int {
	return o._statusCode
}

func (o *GetSnapshotDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume/snapshot/{id}][%d] getSnapshot default %s", o._statusCode, payload)
}

func (o *GetSnapshotDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /v1/volume/snapshot/{id}][%d] getSnapshot default %s", o._statusCode, payload)
}

func (o *GetSnapshotDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetSnapshotDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
