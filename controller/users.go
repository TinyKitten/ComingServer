package controller

import (
	"database/sql"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *sql.DB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *sql.DB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// Add runs the add action.
func (c *UsersController) Add(ctx *app.AddUsersContext) error {
	// UsersController_Add: start_implement

	// Put your logic here

	return nil
	// UsersController_Add: end_implement
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here

	res := app.UserCollection{}
	return ctx.OK(res)
	// UsersController_List: end_implement
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UsersController_Show: end_implement
}
