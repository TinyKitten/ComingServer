package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/goadesign/goa"
	"golang.org/x/crypto/bcrypt"
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
	at := time.Now()
	hashed, err := bcrypt.GenerateFromPassword([]byte(ctx.Payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}
	user := &models.User{
		ScreenName:  ctx.Payload.ScreenName,
		Password:    string(hashed),
		PrivilegeID: uint64(ctx.Payload.PrivilegeID),
		CreatedAt:   at,
		UpdatedAt:   at,
	}

	err = user.Insert(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	res := &app.User{
		ID:         user.ID,
		ScreenName: user.ScreenName,
		CreatedAt:  at.Unix(),
		UpdatedAt:  at.Unix(),
	}

	return ctx.Created(res)
	// UsersController_Add: end_implement
}

// List runs the list action.
func (c *UsersController) List(ctx *app.ListUsersContext) error {
	// UsersController_List: start_implement

	// Put your logic here
	users, err := models.UserList(c.db, ctx.Offset, ctx.Limit)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	res := app.UserCollection{}

	for _, user := range users {
		userMedia := &app.User{
			ID:         user.ID,
			ScreenName: user.ScreenName,
			CreatedAt:  user.CreatedAt.Unix(),
			UpdatedAt:  user.UpdatedAt.Unix(),
		}
		res = append(res, userMedia)
	}

	return ctx.OK(res)
	// UsersController_List: end_implement
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here
	user, err := models.UserByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrInternal(err))
	}
	res := &app.User{
		ID:         user.ID,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt.Unix(),
		UpdatedAt:  user.UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// UsersController_Show: end_implement
}
