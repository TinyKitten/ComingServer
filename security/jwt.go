package security

import (
	"context"
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/TinyKitten/ComingServer/app"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
)

// NewJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
// and validates its content. A real app would probably use goa's JWT security middleware instead.
//
// Note: the code below assumes the example is compiled against the master branch of goa.
// If compiling against goa v1 the call to jwt.New needs to be:
//
//    middleware := jwt.New(keys, ForceFail(), app.NewJWTSecurity())
func NewJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	return jwt.New(jwt.NewSimpleResolver(keys), ForceFail(), app.NewJWTSecurity()), nil
}

// JWTController implements the JWT resource.
type JWTController struct {
	*goa.Controller
	privateKey *rsa.PrivateKey
}

// NewJWTController creates a JWT controller.
func NewJWTController(service *goa.Service) (*JWTController, error) {
	b, err := ioutil.ReadFile("./jwtkey/jwt.key")
	if err != nil {
		return nil, err
	}
	privKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("jwt: failed to load private key: %s", err) // bug
	}
	return &JWTController{
		Controller: service.NewController("JWTController"),
		privateKey: privKey,
	}, nil
}

// LoadJWTPublicKeys loads PEM encoded RSA public keys used to validata and decrypt the JWT.
func LoadJWTPublicKeys() ([]jwt.Key, error) {
	keyFiles, err := filepath.Glob("./jwtkey/*.pub")
	if err != nil {
		return nil, err
	}
	keys := make([]jwt.Key, len(keyFiles))
	for i, keyFile := range keyFiles {
		pem, err := ioutil.ReadFile(keyFile)
		if err != nil {
			return nil, err
		}
		key, err := jwtgo.ParseRSAPublicKeyFromPEM([]byte(pem))
		if err != nil {
			return nil, fmt.Errorf("failed to load key %s: %s", keyFile, err)
		}
		keys[i] = key
	}
	if len(keys) == 0 {
		return nil, fmt.Errorf("couldn't load public keys for JWT security")
	}

	return keys, nil
}

// ForceFail is a middleware illustrating the use of validation middleware with JWT auth.  It checks
// for the presence of a "fail" query string and fails validation if set to the value "true".
func ForceFail() goa.Middleware {
	errValidationFailed := goa.NewErrorClass("validation_failed", 401)
	forceFail := func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			if f, ok := req.URL.Query()["fail"]; ok {
				if f[0] == "true" {
					return errValidationFailed("forcing failure to illustrate Validation middleware")
				}
			}
			return h(ctx, rw, req)
		}
	}
	fm, _ := goa.NewMiddleware(forceFail)
	return fm
}

// NewOptionalJWTMiddleware creates a middleware that checks for the presence of a JWT Authorization header
func NewOptionalJWTMiddleware() (goa.Middleware, error) {
	keys, err := LoadJWTPublicKeys()
	if err != nil {
		return nil, err
	}
	jwtMiddleware := jwt.New(jwt.NewSimpleResolver(keys), ForceFail(), app.NewJWTSecurity())
	return func(nextHandler goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			err := jwtMiddleware(nextHandler)(ctx, rw, req)
			if err != nil {
				return nextHandler(ctx, rw, req)
			}
			return nil
		}
	}, nil

}
