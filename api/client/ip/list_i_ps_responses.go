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

// ListIPsReader is a Reader for the ListIPs structure.
type ListIPsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIPsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIPsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListIPsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListIPsOK creates a ListIPsOK with default headers values
func NewListIPsOK() *ListIPsOK {
	return &ListIPsOK{}
}

/*ListIPsOK handles this case with default header values.

OK
*/
type ListIPsOK struct {
	Payload []*models.ModelsV1IPResponse
}

func (o *ListIPsOK) Error() string {
	return fmt.Sprintf("[GET /v1/ip][%d] listIPsOK  %+v", 200, o.Payload)
}

func (o *ListIPsOK) GetPayload() []*models.ModelsV1IPResponse {
	return o.Payload
}

func (o *ListIPsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIPsDefault creates a ListIPsDefault with default headers values
func NewListIPsDefault(code int) *ListIPsDefault {
	return &ListIPsDefault{
		_statusCode: code,
	}
}

/*ListIPsDefault handles this case with default header values.

Error
*/
type ListIPsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the list i ps default response
func (o *ListIPsDefault) Code() int {
	return o._statusCode
}

func (o *ListIPsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ip][%d] listIPs default  %+v", o._statusCode, o.Payload)
}

func (o *ListIPsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ListIPsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
