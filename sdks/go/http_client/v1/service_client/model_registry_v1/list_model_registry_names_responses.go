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

// ListModelRegistryNamesReader is a Reader for the ListModelRegistryNames structure.
type ListModelRegistryNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListModelRegistryNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListModelRegistryNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListModelRegistryNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListModelRegistryNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListModelRegistryNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListModelRegistryNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListModelRegistryNamesOK creates a ListModelRegistryNamesOK with default headers values
func NewListModelRegistryNamesOK() *ListModelRegistryNamesOK {
	return &ListModelRegistryNamesOK{}
}

/*ListModelRegistryNamesOK handles this case with default header values.

A successful response.
*/
type ListModelRegistryNamesOK struct {
	Payload *service_model.V1ListModelRegistryResponse
}

func (o *ListModelRegistryNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/names][%d] listModelRegistryNamesOK  %+v", 200, o.Payload)
}

func (o *ListModelRegistryNamesOK) GetPayload() *service_model.V1ListModelRegistryResponse {
	return o.Payload
}

func (o *ListModelRegistryNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListModelRegistryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModelRegistryNamesNoContent creates a ListModelRegistryNamesNoContent with default headers values
func NewListModelRegistryNamesNoContent() *ListModelRegistryNamesNoContent {
	return &ListModelRegistryNamesNoContent{}
}

/*ListModelRegistryNamesNoContent handles this case with default header values.

No content.
*/
type ListModelRegistryNamesNoContent struct {
	Payload interface{}
}

func (o *ListModelRegistryNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/names][%d] listModelRegistryNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListModelRegistryNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListModelRegistryNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModelRegistryNamesForbidden creates a ListModelRegistryNamesForbidden with default headers values
func NewListModelRegistryNamesForbidden() *ListModelRegistryNamesForbidden {
	return &ListModelRegistryNamesForbidden{}
}

/*ListModelRegistryNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListModelRegistryNamesForbidden struct {
	Payload interface{}
}

func (o *ListModelRegistryNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/names][%d] listModelRegistryNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListModelRegistryNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListModelRegistryNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModelRegistryNamesNotFound creates a ListModelRegistryNamesNotFound with default headers values
func NewListModelRegistryNamesNotFound() *ListModelRegistryNamesNotFound {
	return &ListModelRegistryNamesNotFound{}
}

/*ListModelRegistryNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListModelRegistryNamesNotFound struct {
	Payload interface{}
}

func (o *ListModelRegistryNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/names][%d] listModelRegistryNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListModelRegistryNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListModelRegistryNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListModelRegistryNamesDefault creates a ListModelRegistryNamesDefault with default headers values
func NewListModelRegistryNamesDefault(code int) *ListModelRegistryNamesDefault {
	return &ListModelRegistryNamesDefault{
		_statusCode: code,
	}
}

/*ListModelRegistryNamesDefault handles this case with default header values.

An unexpected error response.
*/
type ListModelRegistryNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list model registry names default response
func (o *ListModelRegistryNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListModelRegistryNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/models/names][%d] ListModelRegistryNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListModelRegistryNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListModelRegistryNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
