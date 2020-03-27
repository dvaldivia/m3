// Code generated by go-swagger; DO NOT EDIT.

// This file is part of MinIO Console Server
// Copyright (c) 2020 MinIO, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.
//

package user_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ListBucketEventsHandlerFunc turns a function with the right signature into a list bucket events handler
type ListBucketEventsHandlerFunc func(ListBucketEventsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ListBucketEventsHandlerFunc) Handle(params ListBucketEventsParams) middleware.Responder {
	return fn(params)
}

// ListBucketEventsHandler interface for that can handle valid list bucket events params
type ListBucketEventsHandler interface {
	Handle(ListBucketEventsParams) middleware.Responder
}

// NewListBucketEvents creates a new http.Handler for the list bucket events operation
func NewListBucketEvents(ctx *middleware.Context, handler ListBucketEventsHandler) *ListBucketEvents {
	return &ListBucketEvents{Context: ctx, Handler: handler}
}

/*ListBucketEvents swagger:route GET /api/v1/buckets/{bucket_name}/events UserAPI listBucketEvents

List Bucket Events

*/
type ListBucketEvents struct {
	Context *middleware.Context
	Handler ListBucketEventsHandler
}

func (o *ListBucketEvents) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewListBucketEventsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}