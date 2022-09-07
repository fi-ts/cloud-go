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

// ProjectsReader is a Reader for the Projects structure.
type ProjectsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProjectsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProjectsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProjectsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProjectsOK creates a ProjectsOK with default headers values
func NewProjectsOK() *ProjectsOK {
	return &ProjectsOK{}
}

/*
	ProjectsOK describes a response with status code 200, with default header values.

OK
*/
type ProjectsOK struct {
	Payload []*models.V1ProjectInfoResponse
}

func (o *ProjectsOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projectsOK  %+v", 200, o.Payload)
}
func (o *ProjectsOK) GetPayload() []*models.V1ProjectInfoResponse {
	return o.Payload
}

func (o *ProjectsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProjectsDefault creates a ProjectsDefault with default headers values
func NewProjectsDefault(code int) *ProjectsDefault {
	return &ProjectsDefault{
		_statusCode: code,
	}
}

/*
	ProjectsDefault describes a response with status code -1, with default header values.

Error
*/
type ProjectsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the projects default response
func (o *ProjectsDefault) Code() int {
	return o._statusCode
}

func (o *ProjectsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/projects][%d] projects default  %+v", o._statusCode, o.Payload)
}
func (o *ProjectsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ProjectsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
