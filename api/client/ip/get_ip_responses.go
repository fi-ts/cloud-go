// Code generated by go-swagger; DO NOT EDIT.

package ip

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

// GetIPReader is a Reader for the GetIP structure.
type GetIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetIPOK creates a GetIPOK with default headers values
func NewGetIPOK() *GetIPOK {
	return &GetIPOK{}
}

/*
	GetIPOK describes a response with status code 200, with default header values.

OK
*/
type GetIPOK struct {
	Payload *models.ModelsV1IPResponse
}

func (o *GetIPOK) Error() string {
	return fmt.Sprintf("[GET /v1/ip/{ip}][%d] getIpOK  %+v", 200, o.Payload)
}
func (o *GetIPOK) GetPayload() *models.ModelsV1IPResponse {
	return o.Payload
}

func (o *GetIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModelsV1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIPDefault creates a GetIPDefault with default headers values
func NewGetIPDefault(code int) *GetIPDefault {
	return &GetIPDefault{
		_statusCode: code,
	}
}

/*
	GetIPDefault describes a response with status code -1, with default header values.

Error
*/
type GetIPDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get IP default response
func (o *GetIPDefault) Code() int {
	return o._statusCode
}

func (o *GetIPDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ip/{ip}][%d] getIP default  %+v", o._statusCode, o.Payload)
}
func (o *GetIPDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
