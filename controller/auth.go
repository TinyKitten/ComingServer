package controller

import (
	"crypto/rsa"
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/goadesign/goa"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db         *sql.DB
	privateKey *rsa.PrivateKey
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *sql.DB) (*AuthController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         db,
		privateKey: privKey,
	}, nil
}

// Auth runs the auth action.
func (c *AuthController) Auth(ctx *app.AuthAuthContext) error {
	// AuthController_Auth: start_implement

	// Put your logic here
	user, err := models.UserByScreenName(c.db, ctx.Payload.ScreenName)
	if err != nil {
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(ctx.Payload.Password))
	if err != nil {
		return ctx.Unauthorized(goa.ErrUnauthorized(ErrUnauthorized))
	}

	priv, err := user.Privilege(c.db)
	if err != nil {
		return ctx.Unauthorized(goa.ErrUnauthorized(ErrUnauthorized))
	}

	var scopes []string
	switch priv.Code {
	case PrivilegeViewer:
		scopes = []string{"api:read"}
	case PrivilegeModerator:
		scopes = []string{"api:read", "api:write"}
	case PrivilegeADMIN:
		scopes = []string{"api:read", "api:write", "api:admin"}
	}

	token := jwtgo.New(jwtgo.SigningMethodRS512)
	expire := time.Now().Add(time.Hour * 1).Unix()
	token.Claims = jwtgo.MapClaims{
		"iss":    "Coming",                         // who creates the token and signs it
		"aud":    user.ID,                          // to whom the token is intended to be sent
		"exp":    expire,                           // time when the token will expire (10 minutes from now)
		"jti":    uuid.Must(uuid.NewV4()).String(), // a unique identifier for the token
		"iat":    time.Now().Unix(),                // when the token was issued/created (now)
		"nbf":    2,                                // time before which the token is not yet valid (2 minutes ago)
		"sub":    "AccessToken",                    // the subject/principal is whom the token is about
		"scopes": scopes,                           // token scope - not a standard claim
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return ctx.InternalServerError(fmt.Errorf("failed to sign token: %s", err))
	}

	// Set auth header for client retrieval
	ctx.ResponseData.Header().Set("Authorization", "Bearer "+signedToken)

	res := &app.AuthSucces{
		ScreenName: user.ScreenName,
		Token:      signedToken,
	}

	return ctx.OK(res)
	// AuthController_Auth: end_implement
}
