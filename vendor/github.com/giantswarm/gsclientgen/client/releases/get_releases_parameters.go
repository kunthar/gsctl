// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetReleasesParams creates a new GetReleasesParams object
// with the default values initialized.
func NewGetReleasesParams() *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReleasesParamsWithTimeout creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReleasesParamsWithTimeout(timeout time.Duration) *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		timeout: timeout,
	}
}

// NewGetReleasesParamsWithContext creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReleasesParamsWithContext(ctx context.Context) *GetReleasesParams {
	var ()
	return &GetReleasesParams{

		Context: ctx,
	}
}

// NewGetReleasesParamsWithHTTPClient creates a new GetReleasesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReleasesParamsWithHTTPClient(client *http.Client) *GetReleasesParams {
	var ()
	return &GetReleasesParams{
		HTTPClient: client,
	}
}

/*GetReleasesParams contains all the parameters to send to the API endpoint
for the get releases operation typically these are written to a http.Request
*/
type GetReleasesParams struct {

	/*Authorization
	  As described in the [authentication](#section/Authentication) section


	*/
	Authorization string
	/*XGiantSwarmActivity
	  Name of an activity to track, like "list-clusters". This allows to
	analyze several API requests sent in context and gives an idea on
	the purpose.


	*/
	XGiantSwarmActivity *string
	/*XGiantSwarmCmdLine
	  If activity has been issued by a CLI, this header can contain the
	command line


	*/
	XGiantSwarmCmdLine *string
	/*XRequestID
	  A randomly generated key that can be used to track a request throughout
	services of Giant Swarm.


	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) WithTimeout(timeout time.Duration) *GetReleasesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get releases params
func (o *GetReleasesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get releases params
func (o *GetReleasesParams) WithContext(ctx context.Context) *GetReleasesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get releases params
func (o *GetReleasesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) WithHTTPClient(client *http.Client) *GetReleasesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get releases params
func (o *GetReleasesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get releases params
func (o *GetReleasesParams) WithAuthorization(authorization string) *GetReleasesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get releases params
func (o *GetReleasesParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get releases params
func (o *GetReleasesParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetReleasesParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get releases params
func (o *GetReleasesParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get releases params
func (o *GetReleasesParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetReleasesParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get releases params
func (o *GetReleasesParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get releases params
func (o *GetReleasesParams) WithXRequestID(xRequestID *string) *GetReleasesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get releases params
func (o *GetReleasesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetReleasesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	if o.XGiantSwarmActivity != nil {

		// header param X-Giant-Swarm-Activity
		if err := r.SetHeaderParam("X-Giant-Swarm-Activity", *o.XGiantSwarmActivity); err != nil {
			return err
		}

	}

	if o.XGiantSwarmCmdLine != nil {

		// header param X-Giant-Swarm-CmdLine
		if err := r.SetHeaderParam("X-Giant-Swarm-CmdLine", *o.XGiantSwarmCmdLine); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
