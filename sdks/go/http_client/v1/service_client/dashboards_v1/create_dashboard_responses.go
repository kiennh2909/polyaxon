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

package dashboards_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// CreateDashboardReader is a Reader for the CreateDashboard structure.
type CreateDashboardReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateDashboardReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateDashboardOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewCreateDashboardNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCreateDashboardForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateDashboardNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCreateDashboardDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateDashboardOK creates a CreateDashboardOK with default headers values
func NewCreateDashboardOK() *CreateDashboardOK {
	return &CreateDashboardOK{}
}

/*CreateDashboardOK handles this case with default header values.

A successful response.
*/
type CreateDashboardOK struct {
	Payload *service_model.V1Dashboard
}

func (o *CreateDashboardOK) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardOK  %+v", 200, o.Payload)
}

func (o *CreateDashboardOK) GetPayload() *service_model.V1Dashboard {
	return o.Payload
}

func (o *CreateDashboardOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Dashboard)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNoContent creates a CreateDashboardNoContent with default headers values
func NewCreateDashboardNoContent() *CreateDashboardNoContent {
	return &CreateDashboardNoContent{}
}

/*CreateDashboardNoContent handles this case with default header values.

No content.
*/
type CreateDashboardNoContent struct {
	Payload interface{}
}

func (o *CreateDashboardNoContent) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNoContent  %+v", 204, o.Payload)
}

func (o *CreateDashboardNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardForbidden creates a CreateDashboardForbidden with default headers values
func NewCreateDashboardForbidden() *CreateDashboardForbidden {
	return &CreateDashboardForbidden{}
}

/*CreateDashboardForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type CreateDashboardForbidden struct {
	Payload interface{}
}

func (o *CreateDashboardForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardForbidden  %+v", 403, o.Payload)
}

func (o *CreateDashboardForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardNotFound creates a CreateDashboardNotFound with default headers values
func NewCreateDashboardNotFound() *CreateDashboardNotFound {
	return &CreateDashboardNotFound{}
}

/*CreateDashboardNotFound handles this case with default header values.

Resource does not exist.
*/
type CreateDashboardNotFound struct {
	Payload interface{}
}

func (o *CreateDashboardNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] createDashboardNotFound  %+v", 404, o.Payload)
}

func (o *CreateDashboardNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateDashboardNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateDashboardDefault creates a CreateDashboardDefault with default headers values
func NewCreateDashboardDefault(code int) *CreateDashboardDefault {
	return &CreateDashboardDefault{
		_statusCode: code,
	}
}

/*CreateDashboardDefault handles this case with default header values.

An unexpected error response.
*/
type CreateDashboardDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the create dashboard default response
func (o *CreateDashboardDefault) Code() int {
	return o._statusCode
}

func (o *CreateDashboardDefault) Error() string {
	return fmt.Sprintf("[POST /api/v1/orgs/{owner}/dashboards][%d] CreateDashboard default  %+v", o._statusCode, o.Payload)
}

func (o *CreateDashboardDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *CreateDashboardDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
