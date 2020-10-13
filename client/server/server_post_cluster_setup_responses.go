// Code generated by go-swagger; DO NOT EDIT.

package server

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ServerPostClusterSetupReader is a Reader for the ServerPostClusterSetup structure.
type ServerPostClusterSetupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServerPostClusterSetupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServerPostClusterSetupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewServerPostClusterSetupOK creates a ServerPostClusterSetupOK with default headers values
func NewServerPostClusterSetupOK() *ServerPostClusterSetupOK {
	return &ServerPostClusterSetupOK{}
}

/*ServerPostClusterSetupOK handles this case with default header values.

Request completed successfully
*/
type ServerPostClusterSetupOK struct {
	Payload interface{}
}

func (o *ServerPostClusterSetupOK) Error() string {
	return fmt.Sprintf("[POST /_cluster_setup][%d] serverPostClusterSetupOK  %+v", 200, o.Payload)
}

func (o *ServerPostClusterSetupOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ServerPostClusterSetupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
