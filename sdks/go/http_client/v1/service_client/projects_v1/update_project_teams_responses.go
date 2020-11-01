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

package projects_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdateProjectTeamsReader is a Reader for the UpdateProjectTeams structure.
type UpdateProjectTeamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateProjectTeamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateProjectTeamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdateProjectTeamsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateProjectTeamsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateProjectTeamsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateProjectTeamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateProjectTeamsOK creates a UpdateProjectTeamsOK with default headers values
func NewUpdateProjectTeamsOK() *UpdateProjectTeamsOK {
	return &UpdateProjectTeamsOK{}
}

/*UpdateProjectTeamsOK handles this case with default header values.

A successful response.
*/
type UpdateProjectTeamsOK struct {
	Payload *service_model.V1ProjectTeams
}

func (o *UpdateProjectTeamsOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/teams][%d] updateProjectTeamsOK  %+v", 200, o.Payload)
}

func (o *UpdateProjectTeamsOK) GetPayload() *service_model.V1ProjectTeams {
	return o.Payload
}

func (o *UpdateProjectTeamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ProjectTeams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectTeamsNoContent creates a UpdateProjectTeamsNoContent with default headers values
func NewUpdateProjectTeamsNoContent() *UpdateProjectTeamsNoContent {
	return &UpdateProjectTeamsNoContent{}
}

/*UpdateProjectTeamsNoContent handles this case with default header values.

No content.
*/
type UpdateProjectTeamsNoContent struct {
	Payload interface{}
}

func (o *UpdateProjectTeamsNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/teams][%d] updateProjectTeamsNoContent  %+v", 204, o.Payload)
}

func (o *UpdateProjectTeamsNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectTeamsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectTeamsForbidden creates a UpdateProjectTeamsForbidden with default headers values
func NewUpdateProjectTeamsForbidden() *UpdateProjectTeamsForbidden {
	return &UpdateProjectTeamsForbidden{}
}

/*UpdateProjectTeamsForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdateProjectTeamsForbidden struct {
	Payload interface{}
}

func (o *UpdateProjectTeamsForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/teams][%d] updateProjectTeamsForbidden  %+v", 403, o.Payload)
}

func (o *UpdateProjectTeamsForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectTeamsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectTeamsNotFound creates a UpdateProjectTeamsNotFound with default headers values
func NewUpdateProjectTeamsNotFound() *UpdateProjectTeamsNotFound {
	return &UpdateProjectTeamsNotFound{}
}

/*UpdateProjectTeamsNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdateProjectTeamsNotFound struct {
	Payload interface{}
}

func (o *UpdateProjectTeamsNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/teams][%d] updateProjectTeamsNotFound  %+v", 404, o.Payload)
}

func (o *UpdateProjectTeamsNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateProjectTeamsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateProjectTeamsDefault creates a UpdateProjectTeamsDefault with default headers values
func NewUpdateProjectTeamsDefault(code int) *UpdateProjectTeamsDefault {
	return &UpdateProjectTeamsDefault{
		_statusCode: code,
	}
}

/*UpdateProjectTeamsDefault handles this case with default header values.

An unexpected error response.
*/
type UpdateProjectTeamsDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update project teams default response
func (o *UpdateProjectTeamsDefault) Code() int {
	return o._statusCode
}

func (o *UpdateProjectTeamsDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/{owner}/{project}/teams][%d] UpdateProjectTeams default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateProjectTeamsDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdateProjectTeamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
