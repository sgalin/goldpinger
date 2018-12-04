// Copyright 2018 Bloomberg Finance L.P.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckAllPods Queries the API server for all other pods in this service, and makes all of them query all of their neighbours, using their pods IPs. Calls their /check endpoint.
*/
func (a *Client) CheckAllPods(params *CheckAllPodsParams) (*CheckAllPodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckAllPodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkAllPods",
		Method:             "GET",
		PathPattern:        "/check_all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckAllPodsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckAllPodsOK), nil

}

/*
CheckServicePods Queries the API server for all other pods in this service, and pings them via their pods IPs. Calls their /ping endpoint
*/
func (a *Client) CheckServicePods(params *CheckServicePodsParams) (*CheckServicePodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckServicePodsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkServicePods",
		Method:             "GET",
		PathPattern:        "/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CheckServicePodsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckServicePodsOK), nil

}

/*
Ping return query stats
*/
func (a *Client) Ping(params *PingParams) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ping",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PingOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}