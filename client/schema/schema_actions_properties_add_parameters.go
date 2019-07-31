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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/semi-technologies/weaviate/entities/models"
)

// NewSchemaActionsPropertiesAddParams creates a new SchemaActionsPropertiesAddParams object
// with the default values initialized.
func NewSchemaActionsPropertiesAddParams() *SchemaActionsPropertiesAddParams {
	var ()
	return &SchemaActionsPropertiesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaActionsPropertiesAddParamsWithTimeout creates a new SchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaActionsPropertiesAddParamsWithTimeout(timeout time.Duration) *SchemaActionsPropertiesAddParams {
	var ()
	return &SchemaActionsPropertiesAddParams{

		timeout: timeout,
	}
}

// NewSchemaActionsPropertiesAddParamsWithContext creates a new SchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaActionsPropertiesAddParamsWithContext(ctx context.Context) *SchemaActionsPropertiesAddParams {
	var ()
	return &SchemaActionsPropertiesAddParams{

		Context: ctx,
	}
}

// NewSchemaActionsPropertiesAddParamsWithHTTPClient creates a new SchemaActionsPropertiesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaActionsPropertiesAddParamsWithHTTPClient(client *http.Client) *SchemaActionsPropertiesAddParams {
	var ()
	return &SchemaActionsPropertiesAddParams{
		HTTPClient: client,
	}
}

/*SchemaActionsPropertiesAddParams contains all the parameters to send to the API endpoint
for the schema actions properties add operation typically these are written to a http.Request
*/
type SchemaActionsPropertiesAddParams struct {

	/*Body*/
	Body *models.Property
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) WithTimeout(timeout time.Duration) *SchemaActionsPropertiesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) WithContext(ctx context.Context) *SchemaActionsPropertiesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) WithHTTPClient(client *http.Client) *SchemaActionsPropertiesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) WithBody(body *models.Property) *SchemaActionsPropertiesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) SetBody(body *models.Property) {
	o.Body = body
}

// WithClassName adds the className to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) WithClassName(className string) *SchemaActionsPropertiesAddParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema actions properties add params
func (o *SchemaActionsPropertiesAddParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaActionsPropertiesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
