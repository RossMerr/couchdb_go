// Code generated by go-swagger; DO NOT EDIT.

package database

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/rossmerr/couchdb_go/models"
)

// DocBulkPostAllReader is a Reader for the DocBulkPostAll structure.
type DocBulkPostAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DocBulkPostAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDocBulkPostAllCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDocBulkPostAllBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDocBulkPostAllNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDocBulkPostAllCreated creates a DocBulkPostAllCreated with default headers values
func NewDocBulkPostAllCreated() *DocBulkPostAllCreated {
	return &DocBulkPostAllCreated{}
}

/*DocBulkPostAllCreated handles this case with default header values.

Document(s) have been created or updated
*/
type DocBulkPostAllCreated struct {
	Payload []*models.Results
}

func (o *DocBulkPostAllCreated) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] docBulkPostAllCreated  %+v", 201, o.Payload)
}

func (o *DocBulkPostAllCreated) GetPayload() []*models.Results {
	return o.Payload
}

func (o *DocBulkPostAllCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocBulkPostAllBadRequest creates a DocBulkPostAllBadRequest with default headers values
func NewDocBulkPostAllBadRequest() *DocBulkPostAllBadRequest {
	return &DocBulkPostAllBadRequest{}
}

/*DocBulkPostAllBadRequest handles this case with default header values.

The request provided invalid JSON data
*/
type DocBulkPostAllBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *DocBulkPostAllBadRequest) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] docBulkPostAllBadRequest  %+v", 400, o.Payload)
}

func (o *DocBulkPostAllBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocBulkPostAllBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDocBulkPostAllNotFound creates a DocBulkPostAllNotFound with default headers values
func NewDocBulkPostAllNotFound() *DocBulkPostAllNotFound {
	return &DocBulkPostAllNotFound{}
}

/*DocBulkPostAllNotFound handles this case with default header values.

Requested database not found
*/
type DocBulkPostAllNotFound struct {
	Payload *models.ErrorResponse
}

func (o *DocBulkPostAllNotFound) Error() string {
	return fmt.Sprintf("[POST /{db}/_bulk_get][%d] docBulkPostAllNotFound  %+v", 404, o.Payload)
}

func (o *DocBulkPostAllNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DocBulkPostAllNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DocBulkPostAllBody doc bulk post all body
swagger:model DocBulkPostAllBody
*/
type DocBulkPostAllBody struct {

	// docs
	Docs []*DocBulkPostAllParamsBodyDocsItems0 `json:"docs"`
}

// Validate validates this doc bulk post all body
func (o *DocBulkPostAllBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDocs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DocBulkPostAllBody) validateDocs(formats strfmt.Registry) error {

	if swag.IsZero(o.Docs) { // not required
		return nil
	}

	for i := 0; i < len(o.Docs); i++ {
		if swag.IsZero(o.Docs[i]) { // not required
			continue
		}

		if o.Docs[i] != nil {
			if err := o.Docs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("body" + "." + "docs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DocBulkPostAllBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocBulkPostAllBody) UnmarshalBinary(b []byte) error {
	var res DocBulkPostAllBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DocBulkPostAllParamsBodyDocsItems0 doc bulk post all params body docs items0
swagger:model DocBulkPostAllParamsBodyDocsItems0
*/
type DocBulkPostAllParamsBodyDocsItems0 struct {

	// Document ID
	ID string `json:"id,omitempty"`

	// Revision MVCC token
	Rev string `json:"rev,omitempty"`
}

// Validate validates this doc bulk post all params body docs items0
func (o *DocBulkPostAllParamsBodyDocsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DocBulkPostAllParamsBodyDocsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DocBulkPostAllParamsBodyDocsItems0) UnmarshalBinary(b []byte) error {
	var res DocBulkPostAllParamsBodyDocsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
