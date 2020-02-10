// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/treeverse/lakefs/api/gen/models"
)

// GetCommitHandlerFunc turns a function with the right signature into a get commit handler
type GetCommitHandlerFunc func(GetCommitParams, *models.User) middleware.Responder

// Handle executing the request and returning a response
func (fn GetCommitHandlerFunc) Handle(params GetCommitParams, principal *models.User) middleware.Responder {
	return fn(params, principal)
}

// GetCommitHandler interface for that can handle valid get commit params
type GetCommitHandler interface {
	Handle(GetCommitParams, *models.User) middleware.Responder
}

// NewGetCommit creates a new http.Handler for the get commit operation
func NewGetCommit(ctx *middleware.Context, handler GetCommitHandler) *GetCommit {
	return &GetCommit{Context: ctx, Handler: handler}
}

/*GetCommit swagger:route GET /repositories/{repositoryId}/commits/{commitId} getCommit

get commit

*/
type GetCommit struct {
	Context *middleware.Context
	Handler GetCommitHandler
}

func (o *GetCommit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetCommitParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal *models.User
	if uprinc != nil {
		principal = uprinc.(*models.User) // this is really a models.User, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
