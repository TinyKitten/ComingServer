package controller

import (
	"database/sql"
	"log"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/models"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/crypto/bcrypt"
)

// AccountController implements the account resource.
type AccountController struct {
	*goa.Controller
	db *sql.DB
}

// NewAccountController creates a account controller.
func NewAccountController(service *goa.Service, db *sql.DB) *AccountController {
	return &AccountController{
		Controller: service.NewController("AccountController"),
		db:         db,
	}
}

// UpdatePassword runs the update password action.
func (c *AccountController) UpdatePassword(ctx *app.UpdatePasswordAccountContext) error {
	// AccountController_UpdatePassword: start_implement

	// Put your logic here
	token := jwt.ContextJWT(ctx)
	claims := token.Claims.(jwtgo.MapClaims)
	userID := claims["aud"].(uint64)
	user, err := models.UserByID(c.db, userID)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(ctx.Payload.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}
	user.Password = string(hashed)
	err = user.Update(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	return ctx.NoContent()
	// AccountController_UpdatePassword: end_implement
}

// Profile runs the profile action.
func (c *AccountController) Profile(ctx *app.ProfileAccountContext) error {
	// AccountController_Profile: start_implement

	// Put your logic here
	token := jwt.ContextJWT(ctx)
	claims := token.Claims.(jwtgo.MapClaims)
	userID := claims["aud"].(uint64)
	user, err := models.UserByID(c.db, userID)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	res := &app.User{
		ID:         user.ID,
		ScreenName: user.ScreenName,
		CreatedAt:  user.CreatedAt.Unix(),
		UpdatedAt:  user.UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// AccountController_Profile: end_implement
}
