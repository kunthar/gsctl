// Code generated by go-swagger; DO NOT EDIT.

package clusters

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

// NewGetClusterStatusParams creates a new GetClusterStatusParams object
// with the default values initialized.
func NewGetClusterStatusParams() *GetClusterStatusParams {
	var ()
	return &GetClusterStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterStatusParamsWithTimeout creates a new GetClusterStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterStatusParamsWithTimeout(timeout time.Duration) *GetClusterStatusParams {
	var ()
	return &GetClusterStatusParams{

		timeout: timeout,
	}
}

// NewGetClusterStatusParamsWithContext creates a new GetClusterStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterStatusParamsWithContext(ctx context.Context) *GetClusterStatusParams {
	var ()
	return &GetClusterStatusParams{

		Context: ctx,
	}
}

// NewGetClusterStatusParamsWithHTTPClient creates a new GetClusterStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterStatusParamsWithHTTPClient(client *http.Client) *GetClusterStatusParams {
	var ()
	return &GetClusterStatusParams{
		HTTPClient: client,
	}
}

/*GetClusterStatusParams contains all the parameters to send to the API endpoint
for the get cluster status operation typically these are written to a http.Request
*/
type GetClusterStatusParams struct {

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
	/*ClusterID
	  Cluster ID

	*/
	ClusterID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster status params
func (o *GetClusterStatusParams) WithTimeout(timeout time.Duration) *GetClusterStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster status params
func (o *GetClusterStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster status params
func (o *GetClusterStatusParams) WithContext(ctx context.Context) *GetClusterStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster status params
func (o *GetClusterStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster status params
func (o *GetClusterStatusParams) WithHTTPClient(client *http.Client) *GetClusterStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster status params
func (o *GetClusterStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get cluster status params
func (o *GetClusterStatusParams) WithAuthorization(authorization string) *GetClusterStatusParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get cluster status params
func (o *GetClusterStatusParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithXGiantSwarmActivity adds the xGiantSwarmActivity to the get cluster status params
func (o *GetClusterStatusParams) WithXGiantSwarmActivity(xGiantSwarmActivity *string) *GetClusterStatusParams {
	o.SetXGiantSwarmActivity(xGiantSwarmActivity)
	return o
}

// SetXGiantSwarmActivity adds the xGiantSwarmActivity to the get cluster status params
func (o *GetClusterStatusParams) SetXGiantSwarmActivity(xGiantSwarmActivity *string) {
	o.XGiantSwarmActivity = xGiantSwarmActivity
}

// WithXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get cluster status params
func (o *GetClusterStatusParams) WithXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) *GetClusterStatusParams {
	o.SetXGiantSwarmCmdLine(xGiantSwarmCmdLine)
	return o
}

// SetXGiantSwarmCmdLine adds the xGiantSwarmCmdLine to the get cluster status params
func (o *GetClusterStatusParams) SetXGiantSwarmCmdLine(xGiantSwarmCmdLine *string) {
	o.XGiantSwarmCmdLine = xGiantSwarmCmdLine
}

// WithXRequestID adds the xRequestID to the get cluster status params
func (o *GetClusterStatusParams) WithXRequestID(xRequestID *string) *GetClusterStatusParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get cluster status params
func (o *GetClusterStatusParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithClusterID adds the clusterID to the get cluster status params
func (o *GetClusterStatusParams) WithClusterID(clusterID string) *GetClusterStatusParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get cluster status params
func (o *GetClusterStatusParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
