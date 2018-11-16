package controller

import (
	"database/sql"
	"io"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

// WebsocketController implements the websocket resource.
type WebsocketController struct {
	*goa.Controller
	db *sql.DB
}

// NewWebsocketController creates a websocket controller.
func NewWebsocketController(service *goa.Service, db *sql.DB) *WebsocketController {
	return &WebsocketController{
		Controller: service.NewController("WebsocketController"),
		db:         db,
	}
}

// ReceivePeerLocation runs the receive peer location action.
func (c *WebsocketController) ReceivePeerLocation(ctx *app.ReceivePeerLocationWebsocketContext) error {
	c.ReceivePeerLocationWSHandler(ctx).ServeHTTP(ctx.ResponseWriter, ctx.Request)
	return nil
}

// ReceivePeerLocationWSHandler establishes a websocket connection to run the receive peer location action.
func (c *WebsocketController) ReceivePeerLocationWSHandler(ctx *app.ReceivePeerLocationWebsocketContext) websocket.Handler {
	return func(ws *websocket.Conn) {
		// WebsocketController_ReceivePeerLocation: start_implement

		// Put your logic here

		ws.Write([]byte("receive peer location websocket"))
		// Dummy echo websocket server
		io.Copy(ws, ws)
		// WebsocketController_ReceivePeerLocation: end_implement
	}
} // SendCurrentPeerLocation runs the send current peer location action.
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
