// Code generated by go-swagger; DO NOT EDIT.

package partition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new partition API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for partition API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	PartitionDesignDocSearch(params *PartitionDesignDocSearchParams) (*PartitionDesignDocSearchOK, error)

	PartitionDesignDocView(params *PartitionDesignDocViewParams) (*PartitionDesignDocViewOK, error)

	PartitionDesignDocViewPost(params *PartitionDesignDocViewPostParams) (*PartitionDesignDocViewPostOK, error)

	PartitionDocGetAll(params *PartitionDocGetAllParams) (*PartitionDocGetAllOK, error)

	PartitionInfo(params *PartitionInfoParams) (*PartitionInfoOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  PartitionDesignDocSearch executes a search request against the named index in the specified design document

  *Warning*
Search endpoints require a running search plugin connected to each cluster node. See Search Plugin Installation for details.

*Note*
You must enable faceting before you can use the counts, drilldown, and ranges parameters.

*Note*
Faceting and grouping are not supported on partitioned searches, so the following query parameters should not be used on those requests: counts, drilldown, ranges, and group_field, group_limit, group_sort``.

*Note*
Do not combine the bookmark and stale options. These options constrain the choice of shard replicas to use for the response. When used together, the options might cause problems when contact is attempted with replicas that are slow or not available.

*/
func (a *Client) PartitionDesignDocSearch(params *PartitionDesignDocSearchParams) (*PartitionDesignDocSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionDesignDocSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionDesignDocSearch",
		Method:             "GET",
		PathPattern:        "/{db}/_partition/{partition}/_design/{ddoc}/_search/{index}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PartitionDesignDocSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PartitionDesignDocSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for partitionDesignDocSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartitionDesignDocView executes the specified view function from the specified design document
*/
func (a *Client) PartitionDesignDocView(params *PartitionDesignDocViewParams) (*PartitionDesignDocViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionDesignDocViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionDesignDocView",
		Method:             "GET",
		PathPattern:        "/{db}/_partition/{partition}/_design/{ddoc}/_view/{view}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PartitionDesignDocViewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PartitionDesignDocViewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for partitionDesignDocView: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartitionDesignDocViewPost executes the specified view function from the specified design document
*/
func (a *Client) PartitionDesignDocViewPost(params *PartitionDesignDocViewPostParams) (*PartitionDesignDocViewPostOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionDesignDocViewPostParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionDesignDocViewPost",
		Method:             "POST",
		PathPattern:        "/{db}/_partition/{partition}/_design/{ddoc}/_view/{view}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PartitionDesignDocViewPostReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PartitionDesignDocViewPostOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for partitionDesignDocViewPost: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartitionDocGetAll executes the built in all docs view

  This endpoint is a convenience endpoint for automatically setting bounds on the provided partition range. Similar results can be had by using the global /db/_all_docs endpoint with appropriately configured values for start_key and end_key.

Refer to the view endpoint documentation for a complete description of the available query parameters and the format of the returned data.

*/
func (a *Client) PartitionDocGetAll(params *PartitionDocGetAllParams) (*PartitionDocGetAllOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionDocGetAllParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionDocGetAll",
		Method:             "GET",
		PathPattern:        "/{db}/_partition/{partition}/_all_docs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PartitionDocGetAllReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PartitionDocGetAllOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for partitionDocGetAll: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PartitionInfo this endpoint returns information describing the provided partition it includes document and deleted document counts along with external and active data sizes
*/
func (a *Client) PartitionInfo(params *PartitionInfoParams) (*PartitionInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPartitionInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "partitionInfo",
		Method:             "GET",
		PathPattern:        "/{db}/_partition/{partition}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PartitionInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PartitionInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for partitionInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
