package controller

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service) *AccountController {
	return &AccountController{Controller: service.NewController("AccountController")}
}

// UpdatePassword runs the update password action.
func (c *AccountController) UpdatePassword(ctx *app.UpdatePasswordAccountContext) error {
	// AccountController_UpdatePassword: start_implement

	// Put your logic here

	return ctx.NoContent()
	// AccountController_UpdatePassword: end_implement
}
