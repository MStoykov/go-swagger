package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-swagger/go-swagger/httpkit/middleware"
)

// DeletePetHandlerFunc turns a function with the right signature into a delete pet handler
type DeletePetHandlerFunc func(DeletePetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeletePetHandlerFunc) Handle(params DeletePetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeletePetHandler interface for that can handle valid delete pet params
type DeletePetHandler interface {
	Handle(DeletePetParams, interface{}) middleware.Responder
}

// NewDeletePet creates a new http.Handler for the delete pet operation
func NewDeletePet(ctx *middleware.Context, handler DeletePetHandler) *DeletePet {
	return &DeletePet{Context: ctx, Handler: handler}
}

/*DeletePet swagger:route DELETE /pet/{petId} pet deletePet

Deletes a pet

*/
type DeletePet struct {
	Context *middleware.Context
	Params  DeletePetParams
	Handler DeletePetHandler
}

func (o *DeletePet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	o.Params = NewDeletePetParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &o.Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(o.Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
