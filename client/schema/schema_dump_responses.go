//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// SchemaDumpReader is a Reader for the SchemaDump structure.
type SchemaDumpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaDumpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSchemaDumpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSchemaDumpUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSchemaDumpForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSchemaDumpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaDumpOK creates a SchemaDumpOK with default headers values
func NewSchemaDumpOK() *SchemaDumpOK {
	return &SchemaDumpOK{}
}

/*SchemaDumpOK handles this case with default header values.

Successfully dumped the database schema.
*/
type SchemaDumpOK struct {
	Payload *SchemaDumpOKBody
}

func (o *SchemaDumpOK) Error() string {
	return fmt.Sprintf("[GET /schema][%d] schemaDumpOK  %+v", 200, o.Payload)
}

func (o *SchemaDumpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SchemaDumpOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaDumpUnauthorized creates a SchemaDumpUnauthorized with default headers values
func NewSchemaDumpUnauthorized() *SchemaDumpUnauthorized {
	return &SchemaDumpUnauthorized{}
}

/*SchemaDumpUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaDumpUnauthorized struct {
}

func (o *SchemaDumpUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schema][%d] schemaDumpUnauthorized ", 401)
}

func (o *SchemaDumpUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaDumpForbidden creates a SchemaDumpForbidden with default headers values
func NewSchemaDumpForbidden() *SchemaDumpForbidden {
	return &SchemaDumpForbidden{}
}

/*SchemaDumpForbidden handles this case with default header values.

Forbidden
*/
type SchemaDumpForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaDumpForbidden) Error() string {
	return fmt.Sprintf("[GET /schema][%d] schemaDumpForbidden  %+v", 403, o.Payload)
}

func (o *SchemaDumpForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaDumpInternalServerError creates a SchemaDumpInternalServerError with default headers values
func NewSchemaDumpInternalServerError() *SchemaDumpInternalServerError {
	return &SchemaDumpInternalServerError{}
}

/*SchemaDumpInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaDumpInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaDumpInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schema][%d] schemaDumpInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaDumpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SchemaDumpOKBody schema dump o k body
swagger:model SchemaDumpOKBody
*/
type SchemaDumpOKBody struct {

	// actions
	Actions *models.Schema `json:"actions,omitempty"`

	// things
	Things *models.Schema `json:"things,omitempty"`
}

// Validate validates this schema dump o k body
func (o *SchemaDumpOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SchemaDumpOKBody) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(o.Actions) { // not required
		return nil
	}

	if o.Actions != nil {
		if err := o.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemaDumpOK" + "." + "actions")
			}
			return err
		}
	}

	return nil
}

func (o *SchemaDumpOKBody) validateThings(formats strfmt.Registry) error {

	if swag.IsZero(o.Things) { // not required
		return nil
	}

	if o.Things != nil {
		if err := o.Things.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemaDumpOK" + "." + "things")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SchemaDumpOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SchemaDumpOKBody) UnmarshalBinary(b []byte) error {
	var res SchemaDumpOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
