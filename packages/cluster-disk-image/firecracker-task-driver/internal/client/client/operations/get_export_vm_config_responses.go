// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/devbookhq/devbook-api/packages/cluster-disk-image/firecracker-task-driver/internal/client/models"
)

// GetExportVMConfigReader is a Reader for the GetExportVMConfig structure.
type GetExportVMConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetExportVMConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetExportVMConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetExportVMConfigDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetExportVMConfigOK creates a GetExportVMConfigOK with default headers values
func NewGetExportVMConfigOK() *GetExportVMConfigOK {
	return &GetExportVMConfigOK{}
}

/* GetExportVMConfigOK describes a response with status code 200, with default header values.

OK
*/
type GetExportVMConfigOK struct {
	Payload *models.FullVMConfiguration
}

func (o *GetExportVMConfigOK) Error() string {
	return fmt.Sprintf("[GET /vm/config][%d] getExportVmConfigOK  %+v", 200, o.Payload)
}
func (o *GetExportVMConfigOK) GetPayload() *models.FullVMConfiguration {
	return o.Payload
}

func (o *GetExportVMConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FullVMConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetExportVMConfigDefault creates a GetExportVMConfigDefault with default headers values
func NewGetExportVMConfigDefault(code int) *GetExportVMConfigDefault {
	return &GetExportVMConfigDefault{
		_statusCode: code,
	}
}

/* GetExportVMConfigDefault describes a response with status code -1, with default header values.

Internal server error
*/
type GetExportVMConfigDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get export Vm config default response
func (o *GetExportVMConfigDefault) Code() int {
	return o._statusCode
}

func (o *GetExportVMConfigDefault) Error() string {
	return fmt.Sprintf("[GET /vm/config][%d] getExportVmConfig default  %+v", o._statusCode, o.Payload)
}
func (o *GetExportVMConfigDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetExportVMConfigDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}