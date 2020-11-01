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

package presets_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// UpdatePresetReader is a Reader for the UpdatePreset structure.
type UpdatePresetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePresetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdatePresetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewUpdatePresetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdatePresetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdatePresetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdatePresetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdatePresetOK creates a UpdatePresetOK with default headers values
func NewUpdatePresetOK() *UpdatePresetOK {
	return &UpdatePresetOK{}
}

/*UpdatePresetOK handles this case with default header values.

A successful response.
*/
type UpdatePresetOK struct {
	Payload *service_model.V1Preset
}

func (o *UpdatePresetOK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] updatePresetOK  %+v", 200, o.Payload)
}

func (o *UpdatePresetOK) GetPayload() *service_model.V1Preset {
	return o.Payload
}

func (o *UpdatePresetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1Preset)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePresetNoContent creates a UpdatePresetNoContent with default headers values
func NewUpdatePresetNoContent() *UpdatePresetNoContent {
	return &UpdatePresetNoContent{}
}

/*UpdatePresetNoContent handles this case with default header values.

No content.
*/
type UpdatePresetNoContent struct {
	Payload interface{}
}

func (o *UpdatePresetNoContent) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] updatePresetNoContent  %+v", 204, o.Payload)
}

func (o *UpdatePresetNoContent) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePresetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePresetForbidden creates a UpdatePresetForbidden with default headers values
func NewUpdatePresetForbidden() *UpdatePresetForbidden {
	return &UpdatePresetForbidden{}
}

/*UpdatePresetForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type UpdatePresetForbidden struct {
	Payload interface{}
}

func (o *UpdatePresetForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] updatePresetForbidden  %+v", 403, o.Payload)
}

func (o *UpdatePresetForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePresetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePresetNotFound creates a UpdatePresetNotFound with default headers values
func NewUpdatePresetNotFound() *UpdatePresetNotFound {
	return &UpdatePresetNotFound{}
}

/*UpdatePresetNotFound handles this case with default header values.

Resource does not exist.
*/
type UpdatePresetNotFound struct {
	Payload interface{}
}

func (o *UpdatePresetNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] updatePresetNotFound  %+v", 404, o.Payload)
}

func (o *UpdatePresetNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdatePresetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdatePresetDefault creates a UpdatePresetDefault with default headers values
func NewUpdatePresetDefault(code int) *UpdatePresetDefault {
	return &UpdatePresetDefault{
		_statusCode: code,
	}
}

/*UpdatePresetDefault handles this case with default header values.

An unexpected error response.
*/
type UpdatePresetDefault struct {
	_statusCode int

	Payload *service_model.RuntimeError
}

// Code gets the status code for the update preset default response
func (o *UpdatePresetDefault) Code() int {
	return o._statusCode
}

func (o *UpdatePresetDefault) Error() string {
	return fmt.Sprintf("[PUT /api/v1/orgs/{owner}/presets/{preset.uuid}][%d] UpdatePreset default  %+v", o._statusCode, o.Payload)
}

func (o *UpdatePresetDefault) GetPayload() *service_model.RuntimeError {
	return o.Payload
}

func (o *UpdatePresetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.RuntimeError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
