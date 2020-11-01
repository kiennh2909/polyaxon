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

// ListSearchesReader is a Reader for the ListSearches structure.
type ListSearchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSearchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSearchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewListSearchesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewListSearchesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListSearchesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewListSearchesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSearchesOK creates a ListSearchesOK with default headers values
func NewListSearchesOK() *ListSearchesOK {
	return &ListSearchesOK{}
}

/*ListSearchesOK handles this case with default header values.

A successful response.
*/
type ListSearchesOK struct {
	Payload *service_model.V1ListSearchesResponse
}

func (o *ListSearchesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches][%d] listSearchesOK  %+v", 200, o.Payload)
}

func (o *ListSearchesOK) GetPayload() *service_model.V1ListSearchesResponse {
	return o.Payload
}

func (o *ListSearchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1ListSearchesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchesNoContent creates a ListSearchesNoContent with default headers values
func NewListSearchesNoContent() *ListSearchesNoContent {
	return &ListSearchesNoContent{}
}

/*ListSearchesNoContent handles this case with default header values.

No content.
*/
type ListSearchesNoContent struct {
	Payload interface{}
}

func (o *ListSearchesNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches][%d] listSearchesNoContent  %+v", 204, o.Payload)
}

func (o *ListSearchesNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchesForbidden creates a ListSearchesForbidden with default headers values
func NewListSearchesForbidden() *ListSearchesForbidden {
	return &ListSearchesForbidden{}
}

/*ListSearchesForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type ListSearchesForbidden struct {
	Payload interface{}
}

func (o *ListSearchesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches][%d] listSearchesForbidden  %+v", 403, o.Payload)
}

func (o *ListSearchesForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchesNotFound creates a ListSearchesNotFound with default headers values
func NewListSearchesNotFound() *ListSearchesNotFound {
	return &ListSearchesNotFound{}
}

/*ListSearchesNotFound handles this case with default header values.

Resource does not exist.
*/
type ListSearchesNotFound struct {
	Payload interface{}
}

func (o *ListSearchesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches][%d] listSearchesNotFound  %+v", 404, o.Payload)
}

func (o *ListSearchesNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *ListSearchesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSearchesDefault creates a ListSearchesDefault with default headers values
func NewListSearchesDefault(code int) *ListSearchesDefault {
	return &ListSearchesDefault{
		_statusCode: code,
	}
}

/*ListSearchesDefault handles this case with default header values.

An unexpected error response.
*/
type ListSearchesDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the list searches default response
func (o *ListSearchesDefault) Code() int {
	return o._statusCode
}

func (o *ListSearchesDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/orgs/{owner}/searches][%d] ListSearches default  %+v", o._statusCode, o.Payload)
}

func (o *ListSearchesDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *ListSearchesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
