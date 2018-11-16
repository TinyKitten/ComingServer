//go:generate goagen bootstrap -d github.com/TinyKitten/ComingServer/design

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/controller"
	"github.com/TinyKitten/ComingServer/database"
	"github.com/TinyKitten/ComingServer/security"
	"github.com/TinyKitten/ComingServer/utils"
	"github.com/go-redis/redis"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	utils.LoadEnv()

	cs, err := database.NewConfigsFromFile("dbconfig.yml")
	if err != nil {
		log.Fatalf("cannot open database configuration. exit. %s", err)
	}
	dbConn, err := cs.Open(utils.GetEnv())
	if err != nil {
		log.Fatalf("database initialization failed: %s", err)
	}

	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		exitOnFailure(err)
	}

	redisdb := redis.NewClient(opt)
	// Create service
	service := goa.New("ComingServer")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	jwtMiddleware, err := security.NewJWTMiddleware()
	exitOnFailure(err)
	app.UseJWTMiddleware(service, jwtMiddleware)

	optionalJwtMiddleware, err := security.NewOptionalJWTMiddleware()
	exitOnFailure(err)
	app.UseOptionalJWTMiddleware(service, optionalJwtMiddleware)

	// Mount "auth" controller
	c, err := controller.NewAuthController(service, dbConn)
	if err != nil {
		exitOnFailure(err)
	}
	app.MountAuthController(service, c)
	// Mount "peers" controller
	c2 := controller.NewPeersController(service, dbConn, redisdb)
	app.MountPeersController(service, c2)
	// Mount "pods" controller
	c3 := controller.NewPodsController(service, dbConn)
	app.MountPodsController(service, c3)
	// Mount "send current peer location" controller
	c4 := controller.NewWebsocketController(service, dbConn, redisdb)
	app.MountWebsocketController(service, c4)
	// Mount "swagger" controller
	c5 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c5)
	// Mount "users" controller
	c6 := controller.NewUsersController(service, dbConn)
	app.MountUsersController(service, c6)
	// Mount "account" controller
	c7 := controller.NewAccountController(service, dbConn)
	app.MountAccountController(service, c7)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

// exitOnFailure prints a fatal error message and exits the process with status 1.
func exitOnFailure(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "[CRIT] %s", err.Error())
	os.Exit(1)
}
