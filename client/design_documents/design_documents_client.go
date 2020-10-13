// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new design documents API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for design documents API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DesignGetDoc(params *DesignGetDocParams) (*DesignGetDocOK, error)

	DesignGetDocSearch(params *DesignGetDocSearchParams) (*DesignGetDocSearchOK, error)

	DesignGetDocSearchInfo(params *DesignGetDocSearchInfoParams) (*DesignGetDocSearchInfoOK, error)

	DesignGetDocView(params *DesignGetDocViewParams) (*DesignGetDocViewOK, error)

	DesignHeadDelete(params *DesignHeadDeleteParams) (*DesignHeadDeleteOK, *DesignHeadDeleteAccepted, error)

	DesignHeadDoc(params *DesignHeadDocParams) (*DesignHeadDocOK, error)

	DesignHeadDocInfo(params *DesignHeadDocInfoParams) (*DesignHeadDocInfoOK, error)

	DesignPostDocView(params *DesignPostDocViewParams) (*DesignPostDocViewOK, error)

	DesignPuttDoc(params *DesignPuttDocParams) (*DesignPuttDocOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DesignGetDoc returns the contents of the design document specified with the name of the design document and from the specified database from the URL

  Unless you request a specific revision, the latest revision of the document will always be returned.

*/
func (a *Client) DesignGetDoc(params *DesignGetDocParams) (*DesignGetDocOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignGetDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designGetDoc",
		Method:             "GET",
		PathPattern:        "/{db}/_design/{ddoc}",
		ProducesMediaTypes: []string{"application/json", "multipart/mixed", "multipart/related", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "multipart/mixed", "multipart/related", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignGetDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignGetDocOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designGetDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignGetDocSearch executes a search request against the named index in the specified design document

  *Warning*
Search endpoints require a running search plugin connected to each cluster node. See Search Plugin Installation for details.

*Note*
You must enable faceting before you can use the counts, drilldown, and ranges parameters.

*Note*
Faceting and grouping are not supported on partitioned searches, so the following query parameters should not be used on those requests: counts, drilldown, ranges, and group_field, group_limit, group_sort``.

*Note*
Do not combine the bookmark and stale options. These options constrain the choice of shard replicas to use for the response. When used together, the options might cause problems when contact is attempted with replicas that are slow or not available.

*/
func (a *Client) DesignGetDocSearch(params *DesignGetDocSearchParams) (*DesignGetDocSearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignGetDocSearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designGetDocSearch",
		Method:             "GET",
		PathPattern:        "/{db}/_design/{ddoc}/_search/{index}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignGetDocSearchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignGetDocSearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designGetDocSearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignGetDocSearchInfo executes a search request against the named index in the specified design document

  *Warning*
Search endpoints require a running search plugin connected to each cluster node. See Search Plugin Installation for details.

*/
func (a *Client) DesignGetDocSearchInfo(params *DesignGetDocSearchInfoParams) (*DesignGetDocSearchInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignGetDocSearchInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designGetDocSearchInfo",
		Method:             "GET",
		PathPattern:        "/{db}/_design/{ddoc}/_search_info/{index}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignGetDocSearchInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignGetDocSearchInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designGetDocSearchInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignGetDocView executes the specified view function from the specified design document
*/
func (a *Client) DesignGetDocView(params *DesignGetDocViewParams) (*DesignGetDocViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignGetDocViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designGetDocView",
		Method:             "GET",
		PathPattern:        "/{db}/_design/{ddoc}/_view/{view}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignGetDocViewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignGetDocViewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designGetDocView: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignHeadDelete deletes the specified document from the database you must supply the current latest revision either by using the rev parameter to specify the revision
*/
func (a *Client) DesignHeadDelete(params *DesignHeadDeleteParams) (*DesignHeadDeleteOK, *DesignHeadDeleteAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignHeadDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designHeadDelete",
		Method:             "DELETE",
		PathPattern:        "/{db}/_design/{ddoc}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignHeadDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DesignHeadDeleteOK:
		return value, nil, nil
	case *DesignHeadDeleteAccepted:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for design_documents: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignHeadDoc returns the HTTP headers containing a minimal amount of information about the specified design document
*/
func (a *Client) DesignHeadDoc(params *DesignHeadDocParams) (*DesignHeadDocOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignHeadDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designHeadDoc",
		Method:             "HEAD",
		PathPattern:        "/{db}/_design/{ddoc}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignHeadDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignHeadDocOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designHeadDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignHeadDocInfo obtains information about the specified design document including the index index size and current status of the design document and associated index information
*/
func (a *Client) DesignHeadDocInfo(params *DesignHeadDocInfoParams) (*DesignHeadDocInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignHeadDocInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designHeadDocInfo",
		Method:             "HEAD",
		PathPattern:        "/{db}/_design/{ddoc}/_info",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignHeadDocInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignHeadDocInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designHeadDocInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignPostDocView executes the specified view function from the specified design document

  POST view functionality supports identical parameters and behavior as specified in the GET /{db}/_design/{ddoc}/_view/{view} API but allows for the query string parameters to be supplied as keys in a JSON object in the body of the POST request.

*/
func (a *Client) DesignPostDocView(params *DesignPostDocViewParams) (*DesignPostDocViewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignPostDocViewParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designPostDocView",
		Method:             "POST",
		PathPattern:        "/{db}/_design/{ddoc}/_view/{view}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignPostDocViewReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignPostDocViewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designPostDocView: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DesignPuttDoc thes p u t method creates a new named design document or creates a new revision of the existing design document

  *Note*
that for filters, lists, shows and updates fields objects are mapping of function name to string function source code. For views mapping is the same except that values are objects with map and reduce (optional) keys which also contains functions source code.

*/
func (a *Client) DesignPuttDoc(params *DesignPuttDocParams) (*DesignPuttDocOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDesignPuttDocParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "designPuttDoc",
		Method:             "PUT",
		PathPattern:        "/{db}/_design/{ddoc}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DesignPuttDocReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DesignPuttDocOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for designPuttDoc: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
