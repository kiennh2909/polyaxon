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

package searches_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// ListSearchNamesReader is a Reader for the ListSearchNames structure.
type ListSearchNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSearchNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSearchNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListSearchNamesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListSearchNamesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListSearchNamesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListSearchNamesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSearchNamesOK creates a ListSearchNamesOK with default headers values
func NewListSearchNamesOK() *ListSearchNamesOK {
	return &ListSearchNamesOK{}
}

/*ListSearchNamesOK handles this case with default header values.

A successful response.
*/
type ListSearchNamesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

func (o *ListSearchNamesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] listSearchNamesOK  %+v", 200, o.Payload)
}

func (o *ListSearchNamesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *ListSearchNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchNamesNoContent creates a ListSearchNamesNoContent with default headers values
func NewListSearchNamesNoContent() *ListSearchNamesNoContent {
	return &ListSearchNamesNoContent{}
}

/*ListSearchNamesNoContent handles this case with default header values.

No content.
*/
type ListSearchNamesNoContent struct {
	Payload interface{}
}

func (o *ListSearchNamesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] listSearchNamesNoContent  %+v", 204, o.Payload)
}

func (o *ListSearchNamesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchNamesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchNamesForbidden creates a ListSearchNamesForbidden with default headers values
func NewListSearchNamesForbidden() *ListSearchNamesForbidden {
	return &ListSearchNamesForbidden{}
}

/*ListSearchNamesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListSearchNamesForbidden struct {
	Payload interface{}
}

func (o *ListSearchNamesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] listSearchNamesForbidden  %+v", 403, o.Payload)
}

func (o *ListSearchNamesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchNamesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchNamesNotFound creates a ListSearchNamesNotFound with default headers values
func NewListSearchNamesNotFound() *ListSearchNamesNotFound {
	return &ListSearchNamesNotFound{}
}

/*ListSearchNamesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListSearchNamesNotFound struct {
	Payload interface{}
}

func (o *ListSearchNamesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] listSearchNamesNotFound  %+v", 404, o.Payload)
}

func (o *ListSearchNamesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchNamesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchNamesDefault creates a ListSearchNamesDefault with default headers values
func NewListSearchNamesDefault(code int) *ListSearchNamesDefault {
	return &ListSearchNamesDefault{
		_statusCode: code,
	}
}

/*ListSearchNamesDefault handles this case with default header values.

An unexpected error response.
*/
type ListSearchNamesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list search names default response
func (o *ListSearchNamesDefault) Code() int {
	return o._statusCode
}

func (o *ListSearchNamesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches/names][%d] ListSearchNames default  %+v", o._statusCode, o.Payload)
}

func (o *ListSearchNamesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListSearchNamesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
