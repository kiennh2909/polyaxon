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

package model_registry_v1

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

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// NewPatchModelRegistryParams creates a new PatchModelRegistryParams object
// with the default values initialized.
func NewPatchModelRegistryParams() *PatchModelRegistryParams {
	var ()
	return &PatchModelRegistryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchModelRegistryParamsWithTimeout creates a new PatchModelRegistryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchModelRegistryParamsWithTimeout(timeout time.Duration) *PatchModelRegistryParams {
	var ()
	return &PatchModelRegistryParams{

		timeout: timeout,
	}
}

// NewPatchModelRegistryParamsWithContext creates a new PatchModelRegistryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchModelRegistryParamsWithContext(ctx context.Context) *PatchModelRegistryParams {
	var ()
	return &PatchModelRegistryParams{

		Context: ctx,
	}
}

// NewPatchModelRegistryParamsWithHTTPClient creates a new PatchModelRegistryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchModelRegistryParamsWithHTTPClient(client *http.Client) *PatchModelRegistryParams {
	var ()
	return &PatchModelRegistryParams{
		HTTPClient: client,
	}
}

/*PatchModelRegistryParams contains all the parameters to send to the API endpoint
for the patch model registry operation typically these are written to a http.Request
*/
type PatchModelRegistryParams struct {

	/*Body
	  Model body

	*/
	Body *service_model.V1ModelRegistry
	/*ModelName
	  Optional component name, should be a valid fully qualified value: name[:version]

	*/
	ModelName string
	/*Owner
	  Owner of the namespace

	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch model registry params
func (o *PatchModelRegistryParams) WithTimeout(timeout time.Duration) *PatchModelRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch model registry params
func (o *PatchModelRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch model registry params
func (o *PatchModelRegistryParams) WithContext(ctx context.Context) *PatchModelRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch model registry params
func (o *PatchModelRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch model registry params
func (o *PatchModelRegistryParams) WithHTTPClient(client *http.Client) *PatchModelRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch model registry params
func (o *PatchModelRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch model registry params
func (o *PatchModelRegistryParams) WithBody(body *service_model.V1ModelRegistry) *PatchModelRegistryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch model registry params
func (o *PatchModelRegistryParams) SetBody(body *service_model.V1ModelRegistry) {
	o.Body = body
}

// WithModelName adds the modelName to the patch model registry params
func (o *PatchModelRegistryParams) WithModelName(modelName string) *PatchModelRegistryParams {
	o.SetModelName(modelName)
	return o
}

// SetModelName adds the modelName to the patch model registry params
func (o *PatchModelRegistryParams) SetModelName(modelName string) {
	o.ModelName = modelName
}

// WithOwner adds the owner to the patch model registry params
func (o *PatchModelRegistryParams) WithOwner(owner string) *PatchModelRegistryParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the patch model registry params
func (o *PatchModelRegistryParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *PatchModelRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param model.name
	if err := r.SetPathParam("model.name", o.ModelName); err != nil {
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
