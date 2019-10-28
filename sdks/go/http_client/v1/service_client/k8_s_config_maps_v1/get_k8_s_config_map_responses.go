// Copyright 2019 Polyaxon, Inc.
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

package k8_s_config_maps_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	service_model "github.com/polyaxon/polyaxon/sdks/go/http_client/v1/service_model"
)

// GetK8SConfigMapReader is a Reader for the GetK8SConfigMap structure.
type GetK8SConfigMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetK8SConfigMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetK8SConfigMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetK8SConfigMapForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetK8SConfigMapNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetK8SConfigMapOK creates a GetK8SConfigMapOK with default headers values
func NewGetK8SConfigMapOK() *GetK8SConfigMapOK {
	return &GetK8SConfigMapOK{}
}

/*GetK8SConfigMapOK handles this case with default header values.

A successful response.
*/
type GetK8SConfigMapOK struct {
	Payload *service_model.V1K8SResource
}

func (o *GetK8SConfigMapOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_config_maps/{uuid}][%d] getK8SConfigMapOK  %+v", 200, o.Payload)
}

func (o *GetK8SConfigMapOK) GetPayload() *service_model.V1K8SResource {
	return o.Payload
}

func (o *GetK8SConfigMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(service_model.V1K8SResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8SConfigMapForbidden creates a GetK8SConfigMapForbidden with default headers values
func NewGetK8SConfigMapForbidden() *GetK8SConfigMapForbidden {
	return &GetK8SConfigMapForbidden{}
}

/*GetK8SConfigMapForbidden handles this case with default header values.

You don't have permission to access the resource.
*/
type GetK8SConfigMapForbidden struct {
	Payload interface{}
}

func (o *GetK8SConfigMapForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_config_maps/{uuid}][%d] getK8SConfigMapForbidden  %+v", 403, o.Payload)
}

func (o *GetK8SConfigMapForbidden) GetPayload() interface{} {
	return o.Payload
}

func (o *GetK8SConfigMapForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetK8SConfigMapNotFound creates a GetK8SConfigMapNotFound with default headers values
func NewGetK8SConfigMapNotFound() *GetK8SConfigMapNotFound {
	return &GetK8SConfigMapNotFound{}
}

/*GetK8SConfigMapNotFound handles this case with default header values.

Resource does not exist.
*/
type GetK8SConfigMapNotFound struct {
	Payload string
}

func (o *GetK8SConfigMapNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v1/{owner}/k8s_config_maps/{uuid}][%d] getK8SConfigMapNotFound  %+v", 404, o.Payload)
}

func (o *GetK8SConfigMapNotFound) GetPayload() string {
	return o.Payload
}

func (o *GetK8SConfigMapNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}