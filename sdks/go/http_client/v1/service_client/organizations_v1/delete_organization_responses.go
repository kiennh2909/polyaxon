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

package organizations_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteOrganizationReader is a Reader for the DeleteOrganization structure.
type DeleteOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteOrganizationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteOrganizationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOrganizationOK creates a DeleteOrganizationOK with default headers values
func NewDeleteOrganizationOK() *DeleteOrganizationOK {
	return &DeleteOrganizationOK{}
}

/*DeleteOrganizationOK handles this case with default header values.

A successful response.
*/
type DeleteOrganizationOK struct {
}

func (o *DeleteOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}][%d] deleteOrganizationOK ", 200)
}

func (o *DeleteOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationNoContent creates a DeleteOrganizationNoContent with default headers values
func NewDeleteOrganizationNoContent() *DeleteOrganizationNoContent {
	return &DeleteOrganizationNoContent{}
}

/*DeleteOrganizationNoContent handles this case with default header values.

No content.
*/
type DeleteOrganizationNoContent struct {
	Payload interface{}
}

func (o *DeleteOrganizationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}][%d] deleteOrganizationNoContent  %+v", 204, o.Payload)
}

func (o *DeleteOrganizationNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationForbidden creates a DeleteOrganizationForbidden with default headers values
func NewDeleteOrganizationForbidden() *DeleteOrganizationForbidden {
	return &DeleteOrganizationForbidden{}
}

/*DeleteOrganizationForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteOrganizationForbidden struct {
	Payload interface{}
}

func (o *DeleteOrganizationForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}][%d] deleteOrganizationForbidden  %+v", 403, o.Payload)
}

func (o *DeleteOrganizationForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationNotFound creates a DeleteOrganizationNotFound with default headers values
func NewDeleteOrganizationNotFound() *DeleteOrganizationNotFound {
	return &DeleteOrganizationNotFound{}
}

/*DeleteOrganizationNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteOrganizationNotFound struct {
	Payload interface{}
}

func (o *DeleteOrganizationNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}][%d] deleteOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *DeleteOrganizationNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOrganizationDefault creates a DeleteOrganizationDefault with default headers values
func NewDeleteOrganizationDefault(code int) *DeleteOrganizationDefault {
	return &DeleteOrganizationDefault{
		_statusCode: code,
	}
}

/*DeleteOrganizationDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteOrganizationDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete organization default response
func (o *DeleteOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOrganizationDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/orgs/{owner}][%d] DeleteOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteOrganizationDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
