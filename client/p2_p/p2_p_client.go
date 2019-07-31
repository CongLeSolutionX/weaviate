//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2019 Weaviate. All rights reserved.
//  LICENSE: https://github.com/semi-technologies/weaviate/blob/develop/LICENSE.md
//  DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package p2_p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new p2 p API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for p2 p API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
P2pGenesisUpdate Receive an update from the Genesis server.
*/
func (a *Client) P2pGenesisUpdate(params *P2pGenesisUpdateParams) (*P2pGenesisUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewP2pGenesisUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "p2p.genesis_update",
		Method:             "PUT",
		PathPattern:        "/p2p/genesis",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &P2pGenesisUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*P2pGenesisUpdateOK), nil

}

/*
P2pHealth checks if a peer is alive

Check if a peer is alive and healthy.
*/
func (a *Client) P2pHealth(params *P2pHealthParams) (*P2pHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewP2pHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "p2p.health",
		Method:             "GET",
		PathPattern:        "/p2p/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/yaml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &P2pHealthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*P2pHealthOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
