package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/middleware"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewFindPetsByStatusParams creates a new FindPetsByStatusParams object
// with the default values initialized.
func NewFindPetsByStatusParams() FindPetsByStatusParams {
	return FindPetsByStatusParams{}
}

// FindPetsByStatusParams contains all the bound params for the find pets by status operation
// typically these are obtained from a http.Request
//
// swagger:parameters findPetsByStatus
type FindPetsByStatusParams struct {
	/*Status values that need to be considered for filter
	  Required: true
	  In: query
	  Collection Format: csv
	*/
	Status []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *FindPetsByStatusParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	qs := httpkit.Values(r.URL.Query())

	qStatus, qhkStatus, _ := qs.GetOK("status")
	if err := o.bindStatus(qStatus, qhkStatus, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *FindPetsByStatusParams) bindStatus(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("status", "query")
	}

	var qvStatus string
	if len(rawData) > 0 {
		qvStatus = rawData[len(rawData)-1]
	}

	raw := swag.SplitByFormat(qvStatus, "csv")
	size := len(raw)

	if size == 0 {
		return errors.Required("status", "query")
	}

	if size == 0 {
		return nil
	}

	ic := raw
	isz := size
	var ir []string
	iValidateElement := func(i int, statusI string) *errors.Validation {

		if err := validate.Enum(fmt.Sprintf("%s.%v", "status", i), "query", o.Status[i], []interface{}{"available", "pending", "sold"}); err != nil {
			return err
		}

		return nil
	}

	for i := 0; i < isz; i++ {

		if err := iValidateElement(i, ic[i]); err != nil {
			return err
		}
		ir = append(ir, ic[i])
	}

	o.Status = ir

	return nil
}
