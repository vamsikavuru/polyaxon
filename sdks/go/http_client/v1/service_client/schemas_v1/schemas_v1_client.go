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

package schemas_v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new schemas v1 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for schemas v1 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	NoOp(params *NoOpParams, authInfo runtime.ClientAuthInfoWriter) (*NoOpOK, *NoOpNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  NoOp lists teams names
*/
func (a *Client) NoOp(params *NoOpParams, authInfo runtime.ClientAuthInfoWriter) (*NoOpOK, *NoOpNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNoOpParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NoOp",
		Method:             "GET",
		PathPattern:        "/schemas",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NoOpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *NoOpOK:
		return value, nil, nil
	case *NoOpNoContent:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*NoOpDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}