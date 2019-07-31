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

package actions

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
)

// NewActionsDeleteParams creates a new ActionsDeleteParams object
// with the default values initialized.
func NewActionsDeleteParams() *ActionsDeleteParams {
	var ()
	return &ActionsDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActionsDeleteParamsWithTimeout creates a new ActionsDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActionsDeleteParamsWithTimeout(timeout time.Duration) *ActionsDeleteParams {
	var ()
	return &ActionsDeleteParams{

		timeout: timeout,
	}
}

// NewActionsDeleteParamsWithContext creates a new ActionsDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewActionsDeleteParamsWithContext(ctx context.Context) *ActionsDeleteParams {
	var ()
	return &ActionsDeleteParams{

		Context: ctx,
	}
}

// NewActionsDeleteParamsWithHTTPClient creates a new ActionsDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActionsDeleteParamsWithHTTPClient(client *http.Client) *ActionsDeleteParams {
	var ()
	return &ActionsDeleteParams{
		HTTPClient: client,
	}
}

/*ActionsDeleteParams contains all the parameters to send to the API endpoint
for the actions delete operation typically these are written to a http.Request
*/
type ActionsDeleteParams struct {

	/*ID
	  Unique ID of the Thing.

	*/
	ID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the actions delete params
func (o *ActionsDeleteParams) WithTimeout(timeout time.Duration) *ActionsDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the actions delete params
func (o *ActionsDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the actions delete params
func (o *ActionsDeleteParams) WithContext(ctx context.Context) *ActionsDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the actions delete params
func (o *ActionsDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the actions delete params
func (o *ActionsDeleteParams) WithHTTPClient(client *http.Client) *ActionsDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the actions delete params
func (o *ActionsDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the actions delete params
func (o *ActionsDeleteParams) WithID(id strfmt.UUID) *ActionsDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the actions delete params
func (o *ActionsDeleteParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ActionsDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
