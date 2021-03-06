// Code generated by go-swagger; DO NOT EDIT.

package masterdata

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

// GetMasterdataReader is a Reader for the GetMasterdata structure.
type GetMasterdataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMasterdataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMasterdataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetMasterdataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMasterdataDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMasterdataOK creates a GetMasterdataOK with default headers values
func NewGetMasterdataOK() *GetMasterdataOK {
	return &GetMasterdataOK{}
}

/* GetMasterdataOK describes a response with status code 200, with default header values.

Ok
*/
type GetMasterdataOK struct {
	Payload *models.V1MasterdataLookupResponse
}

func (o *GetMasterdataOK) Error() string {
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataOK  %+v", 200, o.Payload)
}
func (o *GetMasterdataOK) GetPayload() *models.V1MasterdataLookupResponse {
	return o.Payload
}

func (o *GetMasterdataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MasterdataLookupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMasterdataNotFound creates a GetMasterdataNotFound with default headers values
func NewGetMasterdataNotFound() *GetMasterdataNotFound {
	return &GetMasterdataNotFound{}
}

/* GetMasterdataNotFound describes a response with status code 404, with default header values.

NotFound
*/
type GetMasterdataNotFound struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *GetMasterdataNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdataNotFound  %+v", 404, o.Payload)
}
func (o *GetMasterdataNotFound) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMasterdataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMasterdataDefault creates a GetMasterdataDefault with default headers values
func NewGetMasterdataDefault(code int) *GetMasterdataDefault {
	return &GetMasterdataDefault{
		_statusCode: code,
	}
}

/* GetMasterdataDefault describes a response with status code -1, with default header values.

Error
*/
type GetMasterdataDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get masterdata default response
func (o *GetMasterdataDefault) Code() int {
	return o._statusCode
}

func (o *GetMasterdataDefault) Error() string {
	return fmt.Sprintf("[GET /v1/masterdata][%d] getMasterdata default  %+v", o._statusCode, o.Payload)
}
func (o *GetMasterdataDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMasterdataDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
