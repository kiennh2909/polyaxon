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

package users_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// DeleteTokenReader is a Reader for the DeleteToken structure.
type DeleteTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewDeleteTokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteTokenForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteTokenNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTokenOK creates a DeleteTokenOK with default headers values
func NewDeleteTokenOK() *DeleteTokenOK {
	return &DeleteTokenOK{}
}

/*DeleteTokenOK handles this case with default header values.

A successful response.
*/
type DeleteTokenOK struct {
}

func (o *DeleteTokenOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/users/tokens/{uuid}][%d] deleteTokenOK ", 200)
}

func (o *DeleteTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTokenNoContent creates a DeleteTokenNoContent with default headers values
func NewDeleteTokenNoContent() *DeleteTokenNoContent {
	return &DeleteTokenNoContent{}
}

/*DeleteTokenNoContent handles this case with default header values.

No content.
*/
type DeleteTokenNoContent struct {
	Payload interface{}
}

func (o *DeleteTokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/users/tokens/{uuid}][%d] deleteTokenNoContent  %+v", 204, o.Payload)
}

func (o *DeleteTokenNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenForbidden creates a DeleteTokenForbidden with default headers values
func NewDeleteTokenForbidden() *DeleteTokenForbidden {
	return &DeleteTokenForbidden{}
}

/*DeleteTokenForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type DeleteTokenForbidden struct {
	Payload interface{}
}

func (o *DeleteTokenForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/users/tokens/{uuid}][%d] deleteTokenForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTokenForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTokenForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenNotFound creates a DeleteTokenNotFound with default headers values
func NewDeleteTokenNotFound() *DeleteTokenNotFound {
	return &DeleteTokenNotFound{}
}

/*DeleteTokenNotFound handles this case with default header values.

Resource does not exist.
*/
type DeleteTokenNotFound struct {
	Payload interface{}
}

func (o *DeleteTokenNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/users/tokens/{uuid}][%d] deleteTokenNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTokenNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteTokenNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteTokenDefault creates a DeleteTokenDefault with default headers values
func NewDeleteTokenDefault(code int) *DeleteTokenDefault {
	return &DeleteTokenDefault{
		_statusCode: code,
	}
}

/*DeleteTokenDefault handles this case with default header values.

An unexpected error response.
*/
type DeleteTokenDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the delete token default response
func (o *DeleteTokenDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTokenDefault) Error() string {
	return fmt.Sprintf("[DELETE /api/v1/users/tokens/{uuid}][%d] DeleteToken default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTokenDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *DeleteTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
