// Code generated by go-swagger; DO NOT EDIT.

package power_edge_router

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewV1PoweredgerouterActionPostParams creates a new V1PoweredgerouterActionPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV1PoweredgerouterActionPostParams() *V1PoweredgerouterActionPostParams {
	return &V1PoweredgerouterActionPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV1PoweredgerouterActionPostParamsWithTimeout creates a new V1PoweredgerouterActionPostParams object
// with the ability to set a timeout on a request.
func NewV1PoweredgerouterActionPostParamsWithTimeout(timeout time.Duration) *V1PoweredgerouterActionPostParams {
	return &V1PoweredgerouterActionPostParams{
		timeout: timeout,
	}
}

// NewV1PoweredgerouterActionPostParamsWithContext creates a new V1PoweredgerouterActionPostParams object
// with the ability to set a context for a request.
func NewV1PoweredgerouterActionPostParamsWithContext(ctx context.Context) *V1PoweredgerouterActionPostParams {
	return &V1PoweredgerouterActionPostParams{
		Context: ctx,
	}
}

// NewV1PoweredgerouterActionPostParamsWithHTTPClient creates a new V1PoweredgerouterActionPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewV1PoweredgerouterActionPostParamsWithHTTPClient(client *http.Client) *V1PoweredgerouterActionPostParams {
	return &V1PoweredgerouterActionPostParams{
		HTTPClient: client,
	}
}

/*
V1PoweredgerouterActionPostParams contains all the parameters to send to the API endpoint

	for the v1 poweredgerouter action post operation.

	Typically these are written to a http.Request.
*/
type V1PoweredgerouterActionPostParams struct {

	/* Body.

	   Parameters for the desired action
	*/
	Body *models.PowerEdgeRouterAction

	/* WorkspaceID.

	   Workspace ID
	*/
	WorkspaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v1 poweredgerouter action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1PoweredgerouterActionPostParams) WithDefaults() *V1PoweredgerouterActionPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v1 poweredgerouter action post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V1PoweredgerouterActionPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) WithTimeout(timeout time.Duration) *V1PoweredgerouterActionPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) WithContext(ctx context.Context) *V1PoweredgerouterActionPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) WithHTTPClient(client *http.Client) *V1PoweredgerouterActionPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) WithBody(body *models.PowerEdgeRouterAction) *V1PoweredgerouterActionPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) SetBody(body *models.PowerEdgeRouterAction) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) WithWorkspaceID(workspaceID string) *V1PoweredgerouterActionPostParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the v1 poweredgerouter action post params
func (o *V1PoweredgerouterActionPostParams) SetWorkspaceID(workspaceID string) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *V1PoweredgerouterActionPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param workspace_id
	if err := r.SetPathParam("workspace_id", o.WorkspaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}