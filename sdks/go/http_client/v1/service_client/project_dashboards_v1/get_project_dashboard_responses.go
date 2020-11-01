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

// GetProjectDashboardReader is a Reader for the GetProjectDashboard structure.
type GetProjectDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetProjectDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetProjectDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetProjectDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProjectDashboardOK creates a GetProjectDashboardOK with default headers values
func NewGetProjectDashboardOK() *GetProjectDashboardOK {
	return &GetProjectDashboardOK{}
}

/*GetProjectDashboardOK handles this case with default header values.

A successful response.
*/
type GetProjectDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *GetProjectDashboardOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards/{uuid}][%d] getProjectDashboardOK  %+v", 200, o.Payload)
}

func (o *GetProjectDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *GetProjectDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDashboardNoContent creates a GetProjectDashboardNoContent with default headers values
func NewGetProjectDashboardNoContent() *GetProjectDashboardNoContent {
	return &GetProjectDashboardNoContent{}
}

/*GetProjectDashboardNoContent handles this case with default header values.

No content.
*/
type GetProjectDashboardNoContent struct {
	Payload interface{}
}

func (o *GetProjectDashboardNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards/{uuid}][%d] getProjectDashboardNoContent  %+v", 204, o.Payload)
}

func (o *GetProjectDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDashboardForbidden creates a GetProjectDashboardForbidden with default headers values
func NewGetProjectDashboardForbidden() *GetProjectDashboardForbidden {
	return &GetProjectDashboardForbidden{}
}

/*GetProjectDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetProjectDashboardForbidden struct {
	Payload interface{}
}

func (o *GetProjectDashboardForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards/{uuid}][%d] getProjectDashboardForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDashboardNotFound creates a GetProjectDashboardNotFound with default headers values
func NewGetProjectDashboardNotFound() *GetProjectDashboardNotFound {
	return &GetProjectDashboardNotFound{}
}

/*GetProjectDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type GetProjectDashboardNotFound struct {
	Payload interface{}
}

func (o *GetProjectDashboardNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards/{uuid}][%d] getProjectDashboardNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetProjectDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectDashboardDefault creates a GetProjectDashboardDefault with default headers values
func NewGetProjectDashboardDefault(code int) *GetProjectDashboardDefault {
	return &GetProjectDashboardDefault{
		_statusCode: code,
	}
}

/*GetProjectDashboardDefault handles this case with default header values.

An unexpected error response.
*/
type GetProjectDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get project dashboard default response
func (o *GetProjectDashboardDefault) Code() int {
	return o._statusCode
}

func (o *GetProjectDashboardDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/{project}/dashboards/{uuid}][%d] GetProjectDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *GetProjectDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetProjectDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
