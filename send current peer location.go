package main

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
	"io"
)

// SendCurrentPeerLocationController implements the send current peer location resource.
type SendCurrentPeerLocationController struct {
	*goa.Controller
}

// NewSendCurrentPeerLocationController creates a send current peer location controller.
func NewSendCurrentPeerLocationController(service *goa.Service) *SendCurrentPeerLocationController {
	return &SendCurrentPeerLocationController{Controller: service.NewController("SendCurrentPeerLocationController")}
}

// Connect runs the connect action.
func (c *SendCurrentPeerLocationController) Connect(ctx *app.ConnectSendCurrentPeerLocationContext) error {
	c.ConnectWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// ConnectWSHandler establishes a websocket connection to run the connect action.
func (c *SendCurrentPeerLocationController) ConnectWSHandler(ctx *app.ConnectSendCurrentPeerLocationContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// SendCurrentPeerLocationController_Connect: start_implement

		// Put your logic here

		ws.Write([]byte("connect send current peer location"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// SendCurrentPeerLocationController_Connect: end_implement
	}
}
