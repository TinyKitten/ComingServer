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
	"github.com/TinyKitten/ComingServer/utils"
	"github.com/go-redis/redis"
	"github.com/goadesign/goa"
	"golang.org/x/net/websocket"
)

var (
	WsErrInternalServerError = app.WsError{
		Type:    "ERROR",
		Code:    "ERR_INTERNAL_SERVER_ERROR",
		Message: "Internal server error",
	}
	WsErrPodNotFound = app.WsError{
		Type:    "ERROR",
		Code:    "ERR_POD_NOT_FOUND",
		Message: "Pod not found",
	}
	WsErrPeerNotFound = app.WsError{
		Type:    "ERROR",
		Code:    "ERR_PEER_NOT_FOUND",
		Message: "Peer not found",
	}
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
		err500bytes, err := json.Marshal(WsErrInternalServerError)
		if err != nil {
			log.Println(err)
			ws.Close()
		}
		errPodNotFound, err := json.Marshal(WsErrPodNotFound)
		if err != nil {
			log.Println(err)
			ws.Close()
		}
		errPeerNotFound, err := json.Marshal(WsErrPeerNotFound)
		if err != nil {
			log.Println(err)
			ws.Close()
		}

		pod, err := models.PodByToken(c.db, ctx.Token)
		if err != nil {
			log.Println(err)
			ws.Write(errPodNotFound)
			ws.Close()
		}

		approachThreshold, err := utils.GetApproachedThreshold()
		if err != nil {
			log.Println(err)
			return
		}
		leaveThreshold, err := utils.GetLeaveThreshold()
		if err != nil {
			log.Println(err)
			return
		}

		pubsub := c.redisClient.Subscribe(pod.Code)
		defer pubsub.Close()
		for {
			msgi, err := pubsub.Receive()
			if err != nil {
				log.Println(err)
				ws.Write(err500bytes)
				return
			}
			switch msg := msgi.(type) {
			case *redis.Subscription:
				log.Println("subscribed to", msg.Channel)
			case *redis.Message:
				var peerLoc app.PeerApproaching
				err := json.Unmarshal([]byte(msg.Payload), &peerLoc)
				if err != nil {
					log.Println(err)
					ws.Write(err500bytes)
				}

				bytes, err := json.Marshal(peerLoc)
				if err != nil {
					log.Println(err)
					continue
				}
				ws.Write([]byte(bytes))

				peer, err := models.PeerByCode(c.db, peerLoc.Code)
				if err != nil {
					log.Println(err)
					ws.Write(errPeerNotFound)
					return
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
				if gap <= approachThreshold && !peer.Approaching.Bool {
					approachingMedia := app.PeerApproaching{
						Type:      APPROACHING,
						Code:      peerLoc.Code,
						Latitude:  peerLoc.Latitude,
						Longitude: peerLoc.Longitude,
						CreatedAt: peerLoc.CreatedAt,
					}
					bytes, err := json.Marshal(approachingMedia)
					if err != nil {
						log.Println(err)
						continue
					}

					ws.Write([]byte(bytes))
					continue
				}
				if gap >= leaveThreshold && peer.Approaching.Bool {
					approachingMedia := app.PeerApproaching{
						Type:      LEAVED,
						Code:      peerLoc.Code,
						Latitude:  peerLoc.Latitude,
						Longitude: peerLoc.Longitude,
						CreatedAt: peerLoc.CreatedAt,
					}
					bytes, err := json.Marshal(approachingMedia)
					if err != nil {
						log.Println(err)
						ws.Write(err500bytes)
						continue
					}

					ws.Write([]byte(bytes))
					continue
				}
			default:
				ws.Write(err500bytes)
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
		err500bytes, err := json.Marshal(WsErrInternalServerError)
		if err != nil {
			log.Println(err)
			ws.Close()
		}
		errPeerNotFound, err := json.Marshal(WsErrPeerNotFound)
		if err != nil {
			log.Println(err)
			ws.Close()
		}

		peer, err := models.PeerByToken(c.db, ctx.Token)
		if err != nil {
			log.Println(err)
			ws.Write(errPeerNotFound)
			return
		}

		pm, err := models.PeersMapByPeerID(c.db, peer.ID)
		if err != nil {
			log.Println(err)
			ws.Write(err500bytes)
			return
		}

		pod, err := pm.Pod(c.db)
		if err != nil {
			log.Println(err)
			ws.Write(err500bytes)
			return
		}

		approachThreshold, err := utils.GetApproachedThreshold()
		if err != nil {
			log.Println(err)
			return
		}
		leaveThreshold, err := utils.GetLeaveThreshold()
		if err != nil {
			log.Println(err)
			return
		}

		for {
			var req locationSendRequest

			if err := websocket.JSON.Receive(ws, &req); err != nil {
				if err == io.EOF {
					return
				}
				log.Println(err)
				ws.Write(err500bytes)
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
				Type:      ACCEPT,
				Code:      peer.Code,
				Latitude:  req.Latitude,
				Longitude: req.Longitude,
				CreatedAt: at.Unix(),
			}
			bytes, err := json.Marshal(approachingMedia)
			if err != nil {
				log.Println(err)
				ws.Write(err500bytes)
				continue
			}

			err = c.redisClient.Publish(pod.Code, bytes).Err()
			if err != nil {
				log.Println(err)
				ws.Write(err500bytes)
				continue
			}

			podCoords := math.Coordinate{
				Latitude:  pod.Latitude,
				Longitude: pod.Longitude,
			}
			peerCoords := math.Coordinate{
				Latitude:  req.Latitude,
				Longitude: req.Longitude,
			}
			gap := math.HubenyDistance(podCoords, peerCoords)
			if gap <= approachThreshold && !peer.Approaching.Bool {
				approachingMedia := app.PeerApproaching{
					Type:      APPROACHING,
					Code:      peer.Code,
					Latitude:  req.Latitude,
					Longitude: req.Longitude,
					CreatedAt: at.Unix(),
				}
				bytes, err := json.Marshal(approachingMedia)
				if err != nil {
					log.Println(err)
					ws.Write(err500bytes)
					continue
				}
				ws.Write(bytes)

				peer.Approaching = sql.NullBool{
					Valid: true,
					Bool:  true,
				}
				err = peer.Update(c.db)
				if err != nil {
					log.Println(err)
					continue
				}
				continue
			}
			log.Println(gap)
			if gap >= leaveThreshold && peer.Approaching.Bool {
				approachingMedia := app.PeerApproaching{
					Type:      LEAVED,
					Code:      peer.Code,
					Latitude:  req.Latitude,
					Longitude: req.Longitude,
					CreatedAt: at.Unix(),
				}
				bytes, err := json.Marshal(approachingMedia)
				if err != nil {
					log.Println(err)
					ws.Write(err500bytes)
					continue
				}
				ws.Write(bytes)

				peer.Approaching = sql.NullBool{
					Valid: true,
					Bool:  false,
				}
				err = peer.Update(c.db)
				if err != nil {
					log.Println(err)
					continue
				}
				continue
			}
		}

		// WebsocketController_SendCurrentPeerLocation: end_implement
	}
}
