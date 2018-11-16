package controller

import (
	"io"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

// WebsocketController implements the websocket resource.
type WebsocketController struct {
	*goa.Controller
}

// NewWebsocketController creates a websocket controller.
func NewWebsocketController(service *goa.Service) *WebsocketController {
	return &WebsocketController{Controller: service.NewController("WebsocketController")}
}

// SendCurrentPeerLocation runs the send current peer location action.
func (c *WebsocketController) SendCurrentPeerLocation(ctx *app.SendCurrentPeerLocationWebsocketContext) error {
	c.SendCurrentPeerLocationWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// SendCurrentPeerLocationWSHandler establishes a websocket connection to run the send current peer location action.
func (c *WebsocketController) SendCurrentPeerLocationWSHandler(ctx *app.SendCurrentPeerLocationWebsocketContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// WebsocketController_SendCurrentPeerLocation: start_implement

		// Put your logic here

		ws.Write([]byte("send current peer location websocket"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// WebsocketController_SendCurrentPeerLocation: end_implement
	}
}
