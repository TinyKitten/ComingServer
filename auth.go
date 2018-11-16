package main

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// Auth runs the auth action.
func (c *AuthController) Auth(ctx *app.AuthAuthContext) error {
	// AuthController_Auth: start_implement

	// Put your logic here

	res := &app.AuthSucces{}
	return ctx.OK(res)
	// AuthController_Auth: end_implement
}
