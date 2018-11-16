//go:generate goagen bootstrap -d github.com/TinyKitten/ComingServer/design

package main

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/controller"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Create service
	service := goa.New("ComingServer")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "auth" controller
	c := controller.NewAuthController(service)
	app.MountAuthController(service, c)
	// Mount "peers" controller
	c2 := controller.NewPeersController(service)
	app.MountPeersController(service, c2)
	// Mount "pods" controller
	c3 := controller.NewPodsController(service)
	app.MountPodsController(service, c3)
	// Mount "send current peer location" controller
	c4 := controller.NewSendCurrentPeerLocationController(service)
	app.MountSendCurrentPeerLocationController(service, c4)
	// Mount "swagger" controller
	c5 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c5)
	// Mount "users" controller
	c6 := controller.NewUsersController(service)
	app.MountUsersController(service, c6)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
