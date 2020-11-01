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

package project_dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// PromoteProjectDashboardReader is a Reader for the PromoteProjectDashboard structure.
type PromoteProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PromoteProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPromoteProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewPromoteProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewPromoteProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPromoteProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPromoteProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPromoteProjectDashboardOK creates a PromoteProjectDashboardOK with default headers values
func NewPromoteProjectDashboardOK() *PromoteProjectDashboardOK {
	return &PromoteProjectDashboardOK{}
}

/*PromoteProjectDashboardOK handles this case with default header values.

A successful response.
*/
type PromoteProjectDashboardOK struct {
}

func (o *PromoteProjectDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/dashboards/{uuid}/promote][%d] promoteProjectDashboardOK ", 200)
}

func (o *PromoteProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPromoteProjectDashboardNoContent creates a PromoteProjectDashboardNoContent with default headers values
func NewPromoteProjectDashboardNoContent() *PromoteProjectDashboardNoContent {
	return &PromoteProjectDashboardNoContent{}
}

/*PromoteProjectDashboardNoContent handles this case with default header values.

No content.
*/
type PromoteProjectDashboardNoContent struct {
	Payload interface{}
}

func (o *PromoteProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *PromoteProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardForbidden creates a PromoteProjectDashboardForbidden with default headers values
func NewPromoteProjectDashboardForbidden() *PromoteProjectDashboardForbidden {
	return &PromoteProjectDashboardForbidden{}
}

/*PromoteProjectDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type PromoteProjectDashboardForbidden struct {
	Payload interface{}
}

func (o *PromoteProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/dashboards/{uuid}/promote][%d] promoteProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *PromoteProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardNotFound creates a PromoteProjectDashboardNotFound with default headers values
func NewPromoteProjectDashboardNotFound() *PromoteProjectDashboardNotFound {
	return &PromoteProjectDashboardNotFound{}
}

/*PromoteProjectDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type PromoteProjectDashboardNotFound struct {
	Payload interface{}
}

func (o *PromoteProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/dashboards/{uuid}/promote][%d] promoteProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *PromoteProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PromoteProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPromoteProjectDashboardDefault creates a PromoteProjectDashboardDefault with default headers values
func NewPromoteProjectDashboardDefault(code int) *PromoteProjectDashboardDefault {
	return &PromoteProjectDashboardDefault{
		_statusCode: code,
	}
}

/*PromoteProjectDashboardDefault handles this case with default header values.

An unexpected error response.
*/
type PromoteProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the promote project dashboard default response
func (o *PromoteProjectDashboardDefault) Code() int {
	return o._statusCode
}

func (o *PromoteProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/{owner}/{project}/dashboards/{uuid}/promote][%d] PromoteProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *PromoteProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *PromoteProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
