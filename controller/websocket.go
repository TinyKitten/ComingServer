package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"time"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/math"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/go-redis/redis"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

const (
	WsErrInternalServerError = "ERR_INTERNAL_SERVER_ERROR"
	WsErrPodNotFound         = "ERR_POD_NOT_FOUND"
)

// WebsocketController implements the websocket resource.
type WebsocketController struct {
	*goa.Controller
	db          *sql.DB
	redisClient *redis.Client
}

type locationSendRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// NewWebsocketController creates a websocket controller.
func NewWebsocketController(service *goa.Service, db *sql.DB, redisClient *redis.Client) *WebsocketController {
	return &WebsocketController{
		Controller:  service.NewController("WebsocketController"),
		db:          db,
		redisClient: redisClient,
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
		pod, err := models.PodByToken(c.db, ctx.Token)
		if err != nil {
			log.Println(err)
			ws.Write([]byte(WsErrPodNotFound))
			ws.Close()
		}

		pubsub := c.redisClient.Subscribe(pod.Code)
		defer pubsub.Close()
		for {
			msgi, err := pubsub.Receive()
			if err != nil {
				log.Println(err)
				ws.Write([]byte(WsErrPodNotFound))
				return
			}
			switch msg := msgi.(type) {
			case *redis.Subscription:
				log.Println("subscribed to", msg.Channel)
			case *redis.Message:
				var peerLoc locationSendRequest
				err := json.Unmarshal([]byte(msg.Payload), &peerLoc)
				if err != nil {
					log.Println(err)
					ws.Write([]byte(WsErrPodNotFound))
				}
				podCoords := math.Coordinate{
					Latitude:  pod.Latitude,
					Longitude: pod.Longitude,
				}
				peerCoords := math.Coordinate{
					Latitude:  peerLoc.Latitude,
					Longitude: peerLoc.Longitude,
				}
				gap := math.HubenyDistance(podCoords, peerCoords)
				if gap < 1000.0 && !pod.Approaching.Bool {
					ws.Write([]byte(msg.Payload))
					pod.Approaching = sql.NullBool{
						Valid: true,
						Bool:  true,
					}
					err = pod.Update(c.db)
					if err != nil {
						log.Println(err)
						continue
					}
				}
				if gap > 1500.0 && pod.Approaching.Bool {
					pod.Approaching = sql.NullBool{
						Valid: true,
						Bool:  false,
					}
					err = pod.Update(c.db)
					if err != nil {
						log.Println(err)
						continue
					}
				}
			default:
				ws.Write([]byte(WsErrInternalServerError))
				return
			}

		}

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
		peer, err := models.PeerByToken(c.db, ctx.Token)
		if err != nil {
			log.Println(err)
			ws.Write([]byte(WsErrPodNotFound))
			return
		}

		pm, err := models.PeersMapByPeerID(c.db, peer.ID)
		if err != nil {
			log.Println(err)
			ws.Write([]byte(WsErrInternalServerError))
			return
		}

		pod, err := pm.Pod(c.db)
		if err != nil {
			log.Println(err)
			ws.Write([]byte(WsErrInternalServerError))
			return
		}

		for {
			var req locationSendRequest

			if err := websocket.JSON.Receive(ws, &req); err != nil {
				if err == io.EOF {
					return
				}
				log.Println(err)
				ws.Write([]byte(WsErrInternalServerError))
				continue
			}

			at := time.Now()
			loc := models.PeerLocation{
				PeerID:    peer.ID,
				Latitude:  req.Latitude,
				Longitude: req.Longitude,
				CreatedAt: at,
				UpdatedAt: at,
			}

			err = loc.Insert(c.db)
			if err != nil {
				// 挿入できなくても処理は続けたい
				log.Println(err)
			}

			approachingMedia := app.PeerApproaching{
				Type:      "APPROACHING",
				Code:      peer.Code,
				Latitude:  req.Latitude,
				Longitude: req.Longitude,
				CreatedAt: at.Unix(),
			}
			bytes, err := json.Marshal(approachingMedia)
			if err != nil {
				log.Println(err)
				ws.Write([]byte(WsErrInternalServerError))
				continue
			}

			err = c.redisClient.Publish(pod.Code, bytes).Err()
			if err != nil {
				log.Println(err)
				ws.Write([]byte(WsErrInternalServerError))
				continue
			}

		}

		// WebsocketController_SendCurrentPeerLocation: end_implement
	}
}
