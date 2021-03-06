// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package component_hub_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewBookmarkComponentHubParams creates a new BookmarkComponentHubParams object
// with the default values initialized.
func NewBookmarkComponentHubParams() *BookmarkComponentHubParams {
	var ()
	return &BookmarkComponentHubParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewBookmarkComponentHubParamsWithTimeout creates a new BookmarkComponentHubParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewBookmarkComponentHubParamsWithTimeout(timeout time.Duration) *BookmarkComponentHubParams {
	var ()
	return &BookmarkComponentHubParams{

		timeout: timeout,
	}
}

// NewBookmarkComponentHubParamsWithContext creates a new BookmarkComponentHubParams object
// with the default values initialized, and the ability to set a context for a request
func NewBookmarkComponentHubParamsWithContext(ctx context.Context) *BookmarkComponentHubParams {
	var ()
	return &BookmarkComponentHubParams{

		Context: ctx,
	}
}

// NewBookmarkComponentHubParamsWithHTTPClient creates a new BookmarkComponentHubParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewBookmarkComponentHubParamsWithHTTPClient(client *http.Client) *BookmarkComponentHubParams {
	var ()
	return &BookmarkComponentHubParams{
		HTTPClient: client,
	}
}

/*BookmarkComponentHubParams contains all the parameters to send to the API endpoint
for the bookmark component hub operation typically these are written to a http.Request
*/
type BookmarkComponentHubParams struct {

	/*Name
	  Component under namesapce

	*/
	Name string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the bookmark component hub params
func (o *BookmarkComponentHubParams) WithTimeout(timeout time.Duration) *BookmarkComponentHubParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the bookmark component hub params
func (o *BookmarkComponentHubParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the bookmark component hub params
func (o *BookmarkComponentHubParams) WithContext(ctx context.Context) *BookmarkComponentHubParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the bookmark component hub params
func (o *BookmarkComponentHubParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the bookmark component hub params
func (o *BookmarkComponentHubParams) WithHTTPClient(client *http.Client) *BookmarkComponentHubParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the bookmark component hub params
func (o *BookmarkComponentHubParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the bookmark component hub params
func (o *BookmarkComponentHubParams) WithName(name string) *BookmarkComponentHubParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the bookmark component hub params
func (o *BookmarkComponentHubParams) SetName(name string) {
	o.Name = name
}

// WithOwner adds the owner to the bookmark component hub params
func (o *BookmarkComponentHubParams) WithOwner(owner string) *BookmarkComponentHubParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the bookmark component hub params
func (o *BookmarkComponentHubParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *BookmarkComponentHubParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
