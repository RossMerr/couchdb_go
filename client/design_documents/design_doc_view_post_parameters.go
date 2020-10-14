// Code generated by go-swagger; DO NOT EDIT.

package design_documents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/RossMerr/couchdb_go/models"
)

// NewDesignDocViewPostParams creates a new DesignDocViewPostParams object
// with the default values initialized.
func NewDesignDocViewPostParams() *DesignDocViewPostParams {
	var ()
	return &DesignDocViewPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDesignDocViewPostParamsWithTimeout creates a new DesignDocViewPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDesignDocViewPostParamsWithTimeout(timeout time.Duration) *DesignDocViewPostParams {
	var ()
	return &DesignDocViewPostParams{

		timeout: timeout,
	}
}

// NewDesignDocViewPostParamsWithContext creates a new DesignDocViewPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewDesignDocViewPostParamsWithContext(ctx context.Context) *DesignDocViewPostParams {
	var ()
	return &DesignDocViewPostParams{

		Context: ctx,
	}
}

// NewDesignDocViewPostParamsWithHTTPClient creates a new DesignDocViewPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDesignDocViewPostParamsWithHTTPClient(client *http.Client) *DesignDocViewPostParams {
	var ()
	return &DesignDocViewPostParams{
		HTTPClient: client,
	}
}

/*DesignDocViewPostParams contains all the parameters to send to the API endpoint
for the design doc view post operation typically these are written to a http.Request
*/
type DesignDocViewPostParams struct {

	/*AttEncodingInfo
	  Include encoding information in attachment stubs if include_docs is true and the particular attachment is compressed. Ignored if include_docs isn’t true. Default is false.

	*/
	AttEncodingInfo *bool
	/*Attachments
	  Include the Base64-encoded content of attachments in the documents that are included if include_docs is true. Ignored if include_docs isn’t true. Default is false.

	*/
	Attachments *bool
	/*Body*/
	Body *models.Keys
	/*Conflicts
	  Include conflicts information in response. Ignored if include_docs isn’t true. Default is false.

	*/
	Conflicts *bool
	/*Db
	  Database name

	*/
	Db string
	/*Ddoc
	  Design document id

	*/
	Ddoc string
	/*Descending
	  Return the documents in descending order by key. Default is false.

	*/
	Descending *bool
	/*EndKey
	  Alias for endkey param

	*/
	EndKey *string
	/*EndKeyDocID
	  Alias for endkey_docid.

	*/
	EndKeyDocID *string
	/*Endkey
	  Stop returning records when the specified key is reached.

	*/
	Endkey *string
	/*EndkeyDocid
	  Stop returning records when the specified document ID is reached. Ignored if endkey is not set.

	*/
	EndkeyDocid *string
	/*Group
	  Group the results using the reduce function to a group or single row. Implies reduce is true and the maximum group_level. Default is false.

	*/
	Group *bool
	/*GroupLevel
	  Specify the group level to be used. Implies group is true.

	*/
	GroupLevel *int64
	/*IncludeDocs
	  Include the associated document with each row. Default is false.

	*/
	IncludeDocs *bool
	/*InclusiveEnd
	  Specifies whether the specified end key should be included in the result. Default is true.

	*/
	InclusiveEnd *bool
	/*Key
	  eturn only documents that match the specified key.

	*/
	Key *string
	/*Keys
	  Return only documents where the key matches one of the keys specified in the array.

	*/
	Keys []string
	/*Limit
	  Limit the number of the returned documents to the specified number.

	*/
	Limit *int64
	/*Reduce
	  Use the reduction function. Default is true when a reduce function is defined.

	*/
	Reduce *bool
	/*RevsInfo
	  Includes detailed information for all known document revisions. Default is false

	*/
	RevsInfo *bool
	/*Skip
	  Skip this number of records before starting to return the results. Default is 0.

	*/
	Skip *int64
	/*Sorted
	  Sort returned rows (see Sorting Returned Rows). Setting this to false offers a performance boost. The total_rows and offset fields are not available when this is set to false. Default is true.

	*/
	Sorted *bool
	/*Stable
	  Whether or not the view results should be returned from a stable set of shards. Default is false.

	*/
	Stable *bool
	/*Stale
	  Allow the results from a stale view to be used. Supported values: ok and update_after. ok is equivalent to stable=true&update=false. update_after is equivalent to stable=true&update=lazy. The default behavior is equivalent to stable=false&update=true. Note that this parameter is deprecated. Use stable and update instead. See Views Generation for more details.


	*/
	Stale *string
	/*StartKey
	  Alias for startkey.

	*/
	StartKey *string
	/*StartKeyDocID
	  Alias for startkey_docid param

	*/
	StartKeyDocID *string
	/*Startkey
	  Return records starting with the specified key.

	*/
	Startkey *string
	/*StartkeyDocid
	  Return records starting with the specified document ID. Ignored if startkey is not set.

	*/
	StartkeyDocid *string
	/*Update
	  Whether or not the view in question should be updated prior to responding to the user. Supported values: true, false, lazy. Default is true.


	*/
	Update *string
	/*UpdateSeq
	  Whether to include in the response an update_seq value indicating the sequence id of the database the view reflects. Default is false.

	*/
	UpdateSeq *bool
	/*View
	  View function name

	*/
	View string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the design doc view post params
func (o *DesignDocViewPostParams) WithTimeout(timeout time.Duration) *DesignDocViewPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the design doc view post params
func (o *DesignDocViewPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the design doc view post params
func (o *DesignDocViewPostParams) WithContext(ctx context.Context) *DesignDocViewPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the design doc view post params
func (o *DesignDocViewPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the design doc view post params
func (o *DesignDocViewPostParams) WithHTTPClient(client *http.Client) *DesignDocViewPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the design doc view post params
func (o *DesignDocViewPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttEncodingInfo adds the attEncodingInfo to the design doc view post params
func (o *DesignDocViewPostParams) WithAttEncodingInfo(attEncodingInfo *bool) *DesignDocViewPostParams {
	o.SetAttEncodingInfo(attEncodingInfo)
	return o
}

// SetAttEncodingInfo adds the attEncodingInfo to the design doc view post params
func (o *DesignDocViewPostParams) SetAttEncodingInfo(attEncodingInfo *bool) {
	o.AttEncodingInfo = attEncodingInfo
}

// WithAttachments adds the attachments to the design doc view post params
func (o *DesignDocViewPostParams) WithAttachments(attachments *bool) *DesignDocViewPostParams {
	o.SetAttachments(attachments)
	return o
}

// SetAttachments adds the attachments to the design doc view post params
func (o *DesignDocViewPostParams) SetAttachments(attachments *bool) {
	o.Attachments = attachments
}

// WithBody adds the body to the design doc view post params
func (o *DesignDocViewPostParams) WithBody(body *models.Keys) *DesignDocViewPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the design doc view post params
func (o *DesignDocViewPostParams) SetBody(body *models.Keys) {
	o.Body = body
}

// WithConflicts adds the conflicts to the design doc view post params
func (o *DesignDocViewPostParams) WithConflicts(conflicts *bool) *DesignDocViewPostParams {
	o.SetConflicts(conflicts)
	return o
}

// SetConflicts adds the conflicts to the design doc view post params
func (o *DesignDocViewPostParams) SetConflicts(conflicts *bool) {
	o.Conflicts = conflicts
}

// WithDb adds the db to the design doc view post params
func (o *DesignDocViewPostParams) WithDb(db string) *DesignDocViewPostParams {
	o.SetDb(db)
	return o
}

// SetDb adds the db to the design doc view post params
func (o *DesignDocViewPostParams) SetDb(db string) {
	o.Db = db
}

// WithDdoc adds the ddoc to the design doc view post params
func (o *DesignDocViewPostParams) WithDdoc(ddoc string) *DesignDocViewPostParams {
	o.SetDdoc(ddoc)
	return o
}

// SetDdoc adds the ddoc to the design doc view post params
func (o *DesignDocViewPostParams) SetDdoc(ddoc string) {
	o.Ddoc = ddoc
}

// WithDescending adds the descending to the design doc view post params
func (o *DesignDocViewPostParams) WithDescending(descending *bool) *DesignDocViewPostParams {
	o.SetDescending(descending)
	return o
}

// SetDescending adds the descending to the design doc view post params
func (o *DesignDocViewPostParams) SetDescending(descending *bool) {
	o.Descending = descending
}

// WithEndKey adds the endKey to the design doc view post params
func (o *DesignDocViewPostParams) WithEndKey(endKey *string) *DesignDocViewPostParams {
	o.SetEndKey(endKey)
	return o
}

// SetEndKey adds the endKey to the design doc view post params
func (o *DesignDocViewPostParams) SetEndKey(endKey *string) {
	o.EndKey = endKey
}

// WithEndKeyDocID adds the endKeyDocID to the design doc view post params
func (o *DesignDocViewPostParams) WithEndKeyDocID(endKeyDocID *string) *DesignDocViewPostParams {
	o.SetEndKeyDocID(endKeyDocID)
	return o
}

// SetEndKeyDocID adds the endKeyDocId to the design doc view post params
func (o *DesignDocViewPostParams) SetEndKeyDocID(endKeyDocID *string) {
	o.EndKeyDocID = endKeyDocID
}

// WithEndkey adds the endkey to the design doc view post params
func (o *DesignDocViewPostParams) WithEndkey(endkey *string) *DesignDocViewPostParams {
	o.SetEndkey(endkey)
	return o
}

// SetEndkey adds the endkey to the design doc view post params
func (o *DesignDocViewPostParams) SetEndkey(endkey *string) {
	o.Endkey = endkey
}

// WithEndkeyDocid adds the endkeyDocid to the design doc view post params
func (o *DesignDocViewPostParams) WithEndkeyDocid(endkeyDocid *string) *DesignDocViewPostParams {
	o.SetEndkeyDocid(endkeyDocid)
	return o
}

// SetEndkeyDocid adds the endkeyDocid to the design doc view post params
func (o *DesignDocViewPostParams) SetEndkeyDocid(endkeyDocid *string) {
	o.EndkeyDocid = endkeyDocid
}

// WithGroup adds the group to the design doc view post params
func (o *DesignDocViewPostParams) WithGroup(group *bool) *DesignDocViewPostParams {
	o.SetGroup(group)
	return o
}

// SetGroup adds the group to the design doc view post params
func (o *DesignDocViewPostParams) SetGroup(group *bool) {
	o.Group = group
}

// WithGroupLevel adds the groupLevel to the design doc view post params
func (o *DesignDocViewPostParams) WithGroupLevel(groupLevel *int64) *DesignDocViewPostParams {
	o.SetGroupLevel(groupLevel)
	return o
}

// SetGroupLevel adds the groupLevel to the design doc view post params
func (o *DesignDocViewPostParams) SetGroupLevel(groupLevel *int64) {
	o.GroupLevel = groupLevel
}

// WithIncludeDocs adds the includeDocs to the design doc view post params
func (o *DesignDocViewPostParams) WithIncludeDocs(includeDocs *bool) *DesignDocViewPostParams {
	o.SetIncludeDocs(includeDocs)
	return o
}

// SetIncludeDocs adds the includeDocs to the design doc view post params
func (o *DesignDocViewPostParams) SetIncludeDocs(includeDocs *bool) {
	o.IncludeDocs = includeDocs
}

// WithInclusiveEnd adds the inclusiveEnd to the design doc view post params
func (o *DesignDocViewPostParams) WithInclusiveEnd(inclusiveEnd *bool) *DesignDocViewPostParams {
	o.SetInclusiveEnd(inclusiveEnd)
	return o
}

// SetInclusiveEnd adds the inclusiveEnd to the design doc view post params
func (o *DesignDocViewPostParams) SetInclusiveEnd(inclusiveEnd *bool) {
	o.InclusiveEnd = inclusiveEnd
}

// WithKey adds the key to the design doc view post params
func (o *DesignDocViewPostParams) WithKey(key *string) *DesignDocViewPostParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the design doc view post params
func (o *DesignDocViewPostParams) SetKey(key *string) {
	o.Key = key
}

// WithKeys adds the keys to the design doc view post params
func (o *DesignDocViewPostParams) WithKeys(keys []string) *DesignDocViewPostParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the design doc view post params
func (o *DesignDocViewPostParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithLimit adds the limit to the design doc view post params
func (o *DesignDocViewPostParams) WithLimit(limit *int64) *DesignDocViewPostParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the design doc view post params
func (o *DesignDocViewPostParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithReduce adds the reduce to the design doc view post params
func (o *DesignDocViewPostParams) WithReduce(reduce *bool) *DesignDocViewPostParams {
	o.SetReduce(reduce)
	return o
}

// SetReduce adds the reduce to the design doc view post params
func (o *DesignDocViewPostParams) SetReduce(reduce *bool) {
	o.Reduce = reduce
}

// WithRevsInfo adds the revsInfo to the design doc view post params
func (o *DesignDocViewPostParams) WithRevsInfo(revsInfo *bool) *DesignDocViewPostParams {
	o.SetRevsInfo(revsInfo)
	return o
}

// SetRevsInfo adds the revsInfo to the design doc view post params
func (o *DesignDocViewPostParams) SetRevsInfo(revsInfo *bool) {
	o.RevsInfo = revsInfo
}

// WithSkip adds the skip to the design doc view post params
func (o *DesignDocViewPostParams) WithSkip(skip *int64) *DesignDocViewPostParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the design doc view post params
func (o *DesignDocViewPostParams) SetSkip(skip *int64) {
	o.Skip = skip
}

// WithSorted adds the sorted to the design doc view post params
func (o *DesignDocViewPostParams) WithSorted(sorted *bool) *DesignDocViewPostParams {
	o.SetSorted(sorted)
	return o
}

// SetSorted adds the sorted to the design doc view post params
func (o *DesignDocViewPostParams) SetSorted(sorted *bool) {
	o.Sorted = sorted
}

// WithStable adds the stable to the design doc view post params
func (o *DesignDocViewPostParams) WithStable(stable *bool) *DesignDocViewPostParams {
	o.SetStable(stable)
	return o
}

// SetStable adds the stable to the design doc view post params
func (o *DesignDocViewPostParams) SetStable(stable *bool) {
	o.Stable = stable
}

// WithStale adds the stale to the design doc view post params
func (o *DesignDocViewPostParams) WithStale(stale *string) *DesignDocViewPostParams {
	o.SetStale(stale)
	return o
}

// SetStale adds the stale to the design doc view post params
func (o *DesignDocViewPostParams) SetStale(stale *string) {
	o.Stale = stale
}

// WithStartKey adds the startKey to the design doc view post params
func (o *DesignDocViewPostParams) WithStartKey(startKey *string) *DesignDocViewPostParams {
	o.SetStartKey(startKey)
	return o
}

// SetStartKey adds the startKey to the design doc view post params
func (o *DesignDocViewPostParams) SetStartKey(startKey *string) {
	o.StartKey = startKey
}

// WithStartKeyDocID adds the startKeyDocID to the design doc view post params
func (o *DesignDocViewPostParams) WithStartKeyDocID(startKeyDocID *string) *DesignDocViewPostParams {
	o.SetStartKeyDocID(startKeyDocID)
	return o
}

// SetStartKeyDocID adds the startKeyDocId to the design doc view post params
func (o *DesignDocViewPostParams) SetStartKeyDocID(startKeyDocID *string) {
	o.StartKeyDocID = startKeyDocID
}

// WithStartkey adds the startkey to the design doc view post params
func (o *DesignDocViewPostParams) WithStartkey(startkey *string) *DesignDocViewPostParams {
	o.SetStartkey(startkey)
	return o
}

// SetStartkey adds the startkey to the design doc view post params
func (o *DesignDocViewPostParams) SetStartkey(startkey *string) {
	o.Startkey = startkey
}

// WithStartkeyDocid adds the startkeyDocid to the design doc view post params
func (o *DesignDocViewPostParams) WithStartkeyDocid(startkeyDocid *string) *DesignDocViewPostParams {
	o.SetStartkeyDocid(startkeyDocid)
	return o
}

// SetStartkeyDocid adds the startkeyDocid to the design doc view post params
func (o *DesignDocViewPostParams) SetStartkeyDocid(startkeyDocid *string) {
	o.StartkeyDocid = startkeyDocid
}

// WithUpdate adds the update to the design doc view post params
func (o *DesignDocViewPostParams) WithUpdate(update *string) *DesignDocViewPostParams {
	o.SetUpdate(update)
	return o
}

// SetUpdate adds the update to the design doc view post params
func (o *DesignDocViewPostParams) SetUpdate(update *string) {
	o.Update = update
}

// WithUpdateSeq adds the updateSeq to the design doc view post params
func (o *DesignDocViewPostParams) WithUpdateSeq(updateSeq *bool) *DesignDocViewPostParams {
	o.SetUpdateSeq(updateSeq)
	return o
}

// SetUpdateSeq adds the updateSeq to the design doc view post params
func (o *DesignDocViewPostParams) SetUpdateSeq(updateSeq *bool) {
	o.UpdateSeq = updateSeq
}

// WithView adds the view to the design doc view post params
func (o *DesignDocViewPostParams) WithView(view string) *DesignDocViewPostParams {
	o.SetView(view)
	return o
}

// SetView adds the view to the design doc view post params
func (o *DesignDocViewPostParams) SetView(view string) {
	o.View = view
}

// WriteToRequest writes these params to a swagger request
func (o *DesignDocViewPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AttEncodingInfo != nil {

		// query param att_encoding_info
		var qrAttEncodingInfo bool
		if o.AttEncodingInfo != nil {
			qrAttEncodingInfo = *o.AttEncodingInfo
		}
		qAttEncodingInfo := swag.FormatBool(qrAttEncodingInfo)
		if qAttEncodingInfo != "" {
			if err := r.SetQueryParam("att_encoding_info", qAttEncodingInfo); err != nil {
				return err
			}
		}

	}

	if o.Attachments != nil {

		// query param attachments
		var qrAttachments bool
		if o.Attachments != nil {
			qrAttachments = *o.Attachments
		}
		qAttachments := swag.FormatBool(qrAttachments)
		if qAttachments != "" {
			if err := r.SetQueryParam("attachments", qAttachments); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Conflicts != nil {

		// query param conflicts
		var qrConflicts bool
		if o.Conflicts != nil {
			qrConflicts = *o.Conflicts
		}
		qConflicts := swag.FormatBool(qrConflicts)
		if qConflicts != "" {
			if err := r.SetQueryParam("conflicts", qConflicts); err != nil {
				return err
			}
		}

	}

	// path param db
	if err := r.SetPathParam("db", o.Db); err != nil {
		return err
	}

	// path param ddoc
	if err := r.SetPathParam("ddoc", o.Ddoc); err != nil {
		return err
	}

	if o.Descending != nil {

		// query param descending
		var qrDescending bool
		if o.Descending != nil {
			qrDescending = *o.Descending
		}
		qDescending := swag.FormatBool(qrDescending)
		if qDescending != "" {
			if err := r.SetQueryParam("descending", qDescending); err != nil {
				return err
			}
		}

	}

	if o.EndKey != nil {

		// query param end_key
		var qrEndKey string
		if o.EndKey != nil {
			qrEndKey = *o.EndKey
		}
		qEndKey := qrEndKey
		if qEndKey != "" {
			if err := r.SetQueryParam("end_key", qEndKey); err != nil {
				return err
			}
		}

	}

	if o.EndKeyDocID != nil {

		// query param end_key_doc_id
		var qrEndKeyDocID string
		if o.EndKeyDocID != nil {
			qrEndKeyDocID = *o.EndKeyDocID
		}
		qEndKeyDocID := qrEndKeyDocID
		if qEndKeyDocID != "" {
			if err := r.SetQueryParam("end_key_doc_id", qEndKeyDocID); err != nil {
				return err
			}
		}

	}

	if o.Endkey != nil {

		// query param endkey
		var qrEndkey string
		if o.Endkey != nil {
			qrEndkey = *o.Endkey
		}
		qEndkey := qrEndkey
		if qEndkey != "" {
			if err := r.SetQueryParam("endkey", qEndkey); err != nil {
				return err
			}
		}

	}

	if o.EndkeyDocid != nil {

		// query param endkey_docid
		var qrEndkeyDocid string
		if o.EndkeyDocid != nil {
			qrEndkeyDocid = *o.EndkeyDocid
		}
		qEndkeyDocid := qrEndkeyDocid
		if qEndkeyDocid != "" {
			if err := r.SetQueryParam("endkey_docid", qEndkeyDocid); err != nil {
				return err
			}
		}

	}

	if o.Group != nil {

		// query param group
		var qrGroup bool
		if o.Group != nil {
			qrGroup = *o.Group
		}
		qGroup := swag.FormatBool(qrGroup)
		if qGroup != "" {
			if err := r.SetQueryParam("group", qGroup); err != nil {
				return err
			}
		}

	}

	if o.GroupLevel != nil {

		// query param group_level
		var qrGroupLevel int64
		if o.GroupLevel != nil {
			qrGroupLevel = *o.GroupLevel
		}
		qGroupLevel := swag.FormatInt64(qrGroupLevel)
		if qGroupLevel != "" {
			if err := r.SetQueryParam("group_level", qGroupLevel); err != nil {
				return err
			}
		}

	}

	if o.IncludeDocs != nil {

		// query param include_docs
		var qrIncludeDocs bool
		if o.IncludeDocs != nil {
			qrIncludeDocs = *o.IncludeDocs
		}
		qIncludeDocs := swag.FormatBool(qrIncludeDocs)
		if qIncludeDocs != "" {
			if err := r.SetQueryParam("include_docs", qIncludeDocs); err != nil {
				return err
			}
		}

	}

	if o.InclusiveEnd != nil {

		// query param inclusive_end
		var qrInclusiveEnd bool
		if o.InclusiveEnd != nil {
			qrInclusiveEnd = *o.InclusiveEnd
		}
		qInclusiveEnd := swag.FormatBool(qrInclusiveEnd)
		if qInclusiveEnd != "" {
			if err := r.SetQueryParam("inclusive_end", qInclusiveEnd); err != nil {
				return err
			}
		}

	}

	if o.Key != nil {

		// query param key
		var qrKey string
		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {
			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}

	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Reduce != nil {

		// query param reduce
		var qrReduce bool
		if o.Reduce != nil {
			qrReduce = *o.Reduce
		}
		qReduce := swag.FormatBool(qrReduce)
		if qReduce != "" {
			if err := r.SetQueryParam("reduce", qReduce); err != nil {
				return err
			}
		}

	}

	if o.RevsInfo != nil {

		// query param revs_info
		var qrRevsInfo bool
		if o.RevsInfo != nil {
			qrRevsInfo = *o.RevsInfo
		}
		qRevsInfo := swag.FormatBool(qrRevsInfo)
		if qRevsInfo != "" {
			if err := r.SetQueryParam("revs_info", qRevsInfo); err != nil {
				return err
			}
		}

	}

	if o.Skip != nil {

		// query param skip
		var qrSkip int64
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt64(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Sorted != nil {

		// query param sorted
		var qrSorted bool
		if o.Sorted != nil {
			qrSorted = *o.Sorted
		}
		qSorted := swag.FormatBool(qrSorted)
		if qSorted != "" {
			if err := r.SetQueryParam("sorted", qSorted); err != nil {
				return err
			}
		}

	}

	if o.Stable != nil {

		// query param stable
		var qrStable bool
		if o.Stable != nil {
			qrStable = *o.Stable
		}
		qStable := swag.FormatBool(qrStable)
		if qStable != "" {
			if err := r.SetQueryParam("stable", qStable); err != nil {
				return err
			}
		}

	}

	if o.Stale != nil {

		// query param stale
		var qrStale string
		if o.Stale != nil {
			qrStale = *o.Stale
		}
		qStale := qrStale
		if qStale != "" {
			if err := r.SetQueryParam("stale", qStale); err != nil {
				return err
			}
		}

	}

	if o.StartKey != nil {

		// query param start_key
		var qrStartKey string
		if o.StartKey != nil {
			qrStartKey = *o.StartKey
		}
		qStartKey := qrStartKey
		if qStartKey != "" {
			if err := r.SetQueryParam("start_key", qStartKey); err != nil {
				return err
			}
		}

	}

	if o.StartKeyDocID != nil {

		// query param start_key_doc_id
		var qrStartKeyDocID string
		if o.StartKeyDocID != nil {
			qrStartKeyDocID = *o.StartKeyDocID
		}
		qStartKeyDocID := qrStartKeyDocID
		if qStartKeyDocID != "" {
			if err := r.SetQueryParam("start_key_doc_id", qStartKeyDocID); err != nil {
				return err
			}
		}

	}

	if o.Startkey != nil {

		// query param startkey
		var qrStartkey string
		if o.Startkey != nil {
			qrStartkey = *o.Startkey
		}
		qStartkey := qrStartkey
		if qStartkey != "" {
			if err := r.SetQueryParam("startkey", qStartkey); err != nil {
				return err
			}
		}

	}

	if o.StartkeyDocid != nil {

		// query param startkey_docid
		var qrStartkeyDocid string
		if o.StartkeyDocid != nil {
			qrStartkeyDocid = *o.StartkeyDocid
		}
		qStartkeyDocid := qrStartkeyDocid
		if qStartkeyDocid != "" {
			if err := r.SetQueryParam("startkey_docid", qStartkeyDocid); err != nil {
				return err
			}
		}

	}

	if o.Update != nil {

		// query param update
		var qrUpdate string
		if o.Update != nil {
			qrUpdate = *o.Update
		}
		qUpdate := qrUpdate
		if qUpdate != "" {
			if err := r.SetQueryParam("update", qUpdate); err != nil {
				return err
			}
		}

	}

	if o.UpdateSeq != nil {

		// query param update_seq
		var qrUpdateSeq bool
		if o.UpdateSeq != nil {
			qrUpdateSeq = *o.UpdateSeq
		}
		qUpdateSeq := swag.FormatBool(qrUpdateSeq)
		if qUpdateSeq != "" {
			if err := r.SetQueryParam("update_seq", qUpdateSeq); err != nil {
				return err
			}
		}

	}

	// path param view
	if err := r.SetPathParam("view", o.View); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
