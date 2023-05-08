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

// ClusterUsageCSVReader is a Reader for the ClusterUsageCSV structure.
type ClusterUsageCSVReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterUsageCSVReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClusterUsageCSVOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterUsageCSVDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterUsageCSVOK creates a ClusterUsageCSVOK with default headers values
func NewClusterUsageCSVOK() *ClusterUsageCSVOK {
	return &ClusterUsageCSVOK{}
}

/*
ClusterUsageCSVOK describes a response with status code 200, with default header values.

OK
*/
type ClusterUsageCSVOK struct {
	Payload string
}

// IsSuccess returns true when this cluster usage c s v o k response has a 2xx status code
func (o *ClusterUsageCSVOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster usage c s v o k response has a 3xx status code
func (o *ClusterUsageCSVOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster usage c s v o k response has a 4xx status code
func (o *ClusterUsageCSVOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster usage c s v o k response has a 5xx status code
func (o *ClusterUsageCSVOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster usage c s v o k response a status code equal to that given
func (o *ClusterUsageCSVOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cluster usage c s v o k response
func (o *ClusterUsageCSVOK) Code() int {
	return 200
}

func (o *ClusterUsageCSVOK) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage-csv][%d] clusterUsageCSVOK  %+v", 200, o.Payload)
}

func (o *ClusterUsageCSVOK) String() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage-csv][%d] clusterUsageCSVOK  %+v", 200, o.Payload)
}

func (o *ClusterUsageCSVOK) GetPayload() string {
	return o.Payload
}

func (o *ClusterUsageCSVOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterUsageCSVDefault creates a ClusterUsageCSVDefault with default headers values
func NewClusterUsageCSVDefault(code int) *ClusterUsageCSVDefault {
	return &ClusterUsageCSVDefault{
		_statusCode: code,
	}
}

/*
ClusterUsageCSVDefault describes a response with status code -1, with default header values.

Error
*/
type ClusterUsageCSVDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this cluster usage c s v default response has a 2xx status code
func (o *ClusterUsageCSVDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster usage c s v default response has a 3xx status code
func (o *ClusterUsageCSVDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster usage c s v default response has a 4xx status code
func (o *ClusterUsageCSVDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster usage c s v default response has a 5xx status code
func (o *ClusterUsageCSVDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster usage c s v default response a status code equal to that given
func (o *ClusterUsageCSVDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cluster usage c s v default response
func (o *ClusterUsageCSVDefault) Code() int {
	return o._statusCode
}

func (o *ClusterUsageCSVDefault) Error() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage-csv][%d] clusterUsageCSV default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterUsageCSVDefault) String() string {
	return fmt.Sprintf("[POST /v1/accounting/cluster-usage-csv][%d] clusterUsageCSV default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterUsageCSVDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ClusterUsageCSVDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
