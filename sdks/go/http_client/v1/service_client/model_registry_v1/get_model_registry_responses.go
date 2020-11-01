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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetModelRegistryReader is a Reader for the GetModelRegistry structure.
type GetModelRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetModelRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetModelRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewGetModelRegistryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetModelRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetModelRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetModelRegistryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetModelRegistryOK creates a GetModelRegistryOK with default headers values
func NewGetModelRegistryOK() *GetModelRegistryOK {
	return &GetModelRegistryOK{}
}

/*GetModelRegistryOK handles this case with default header values.

A successful response.
*/
type GetModelRegistryOK struct {
	Payload *service_model.V1ModelRegistry
}

func (o *GetModelRegistryOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/{uuid}][%d] getModelRegistryOK  %+v", 200, o.Payload)
}

func (o *GetModelRegistryOK) GetPayload() *service_model.V1ModelRegistry {
	return o.Payload
}

func (o *GetModelRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ModelRegistry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRegistryNoContent creates a GetModelRegistryNoContent with default headers values
func NewGetModelRegistryNoContent() *GetModelRegistryNoContent {
	return &GetModelRegistryNoContent{}
}

/*GetModelRegistryNoContent handles this case with default header values.

No content.
*/
type GetModelRegistryNoContent struct {
	Payload interface{}
}

func (o *GetModelRegistryNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/{uuid}][%d] getModelRegistryNoContent  %+v", 204, o.Payload)
}

func (o *GetModelRegistryNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelRegistryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRegistryForbidden creates a GetModelRegistryForbidden with default headers values
func NewGetModelRegistryForbidden() *GetModelRegistryForbidden {
	return &GetModelRegistryForbidden{}
}

/*GetModelRegistryForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetModelRegistryForbidden struct {
	Payload interface{}
}

func (o *GetModelRegistryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/{uuid}][%d] getModelRegistryForbidden  %+v", 403, o.Payload)
}

func (o *GetModelRegistryForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRegistryNotFound creates a GetModelRegistryNotFound with default headers values
func NewGetModelRegistryNotFound() *GetModelRegistryNotFound {
	return &GetModelRegistryNotFound{}
}

/*GetModelRegistryNotFound handles this case with default header values.

Resource does not exist.
*/
type GetModelRegistryNotFound struct {
	Payload interface{}
}

func (o *GetModelRegistryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/{uuid}][%d] getModelRegistryNotFound  %+v", 404, o.Payload)
}

func (o *GetModelRegistryNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *GetModelRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetModelRegistryDefault creates a GetModelRegistryDefault with default headers values
func NewGetModelRegistryDefault(code int) *GetModelRegistryDefault {
	return &GetModelRegistryDefault{
		_statusCode: code,
	}
}

/*GetModelRegistryDefault handles this case with default header values.

An unexpected error response.
*/
type GetModelRegistryDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the get model registry default response
func (o *GetModelRegistryDefault) Code() int {
	return o._statusCode
}

func (o *GetModelRegistryDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/{uuid}][%d] GetModelRegistry default  %+v", o._statusCode, o.Payload)
}

func (o *GetModelRegistryDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *GetModelRegistryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
