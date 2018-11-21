package controller

import (
	"database/sql"
	"encoding/json"
	"log"
	"sort"
	"time"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/math"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/TinyKitten/ComingServer/utils"
	"github.com/go-redis/redis"
	"github.com/goadesign/goa"
)

type SortablePeerLocations []*models.PeerLocation

func (a SortablePeerLocations) Len() int           { return len(a) }
func (a SortablePeerLocations) Less(i, j int) bool { return a[i].ID < a[j].ID }
func (a SortablePeerLocations) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// PeersController implements the peers resource.
type PeersController struct {
	*goa.Controller
	db          *sql.DB
	redisClient *redis.Client
}

// NewPeersController creates a peers controller.
func NewPeersController(service *goa.Service, db *sql.DB, redisClient *redis.Client) *PeersController {
	return &PeersController{
		Controller:  service.NewController("PeersController"),
		db:          db,
		redisClient: redisClient,
	}
}

// Add runs the add action.
func (c *PeersController) Add(ctx *app.AddPeersContext) error {
	// PeersController_Add: start_implement

	// Put your logic here
	pod, err := models.PodByID(c.db, uint64(ctx.Payload.PodID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrInternal("Pod not found"))
	}

	at := time.Now()
	token, err := generateToken()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	approaching := sql.NullBool{
		Valid: true,
		Bool:  false,
	}
	peer := models.Peer{
		Code:        ctx.Payload.Code,
		Token:       token,
		Approaching: approaching,
		CreatedAt:   at,
		UpdatedAt:   at,
	}

	err = peer.Insert(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	pm := models.PeersMap{
		PodID:     pod.ID,
		PeerID:    peer.ID,
		CreatedAt: at,
		UpdatedAt: at,
	}

	err = pm.Insert(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	created := &app.PeerCreated{
		ID:        int64(peer.ID),
		Code:      ctx.Payload.Code,
		Token:     token,
		CreatedAt: at.Unix(),
		UpdatedAt: at.Unix(),
	}

	return ctx.Created(created)
	// PeersController_Add: end_implement
}

// CurrentLocation runs the current location action.
func (c *PeersController) CurrentLocation(ctx *app.CurrentLocationPeersContext) error {
	// PeersController_CurrentLocation: start_implement

	// Put your logic here
	locs, err := models.PeerLocationsByPeerID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	if len(locs) == 0 {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound("peer location is not recorded"))
	}

	res := &app.PeerLocation{
		Latitude:  locs[len(locs)-1].Latitude,
		Longitude: locs[len(locs)-1].Longitude,
		CreatedAt: locs[len(locs)-1].CreatedAt.Unix(),
		UpdatedAt: locs[len(locs)-1].UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// PeersController_CurrentLocation: end_implement
}

// List runs the list action.
func (c *PeersController) List(ctx *app.ListPeersContext) error {
	// PeersController_List: start_implement

	// Put your logic here
	res := app.PeerCollection{}
	peers, err := models.PeerList(c.db, ctx.Offset, ctx.Limit)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	for _, peer := range peers {
		peerMedia := &app.Peer{
			ID:          int64(peer.ID),
			Code:        peer.Code,
			Approaching: peer.Approaching.Bool,
			CreatedAt:   peer.CreatedAt.Unix(),
			UpdatedAt:   peer.UpdatedAt.Unix(),
		}
		res = append(res, peerMedia)
	}
	return ctx.OK(res)
	// PeersController_List: end_implement
}

// Locations runs the locations action.
func (c *PeersController) Locations(ctx *app.LocationsPeersContext) error {
	// PeersController_Locations: start_implement

	// Put your logic here
	locs, err := models.PeerLocationsByPeerID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	sort.Sort(sort.Reverse(SortablePeerLocations(locs)))

	res := app.PeerLocationCollection{}
	for _, loc := range locs {
		peerLocMedia := &app.PeerLocation{
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			CreatedAt: loc.CreatedAt.Unix(),
			UpdatedAt: loc.UpdatedAt.Unix(),
		}
		res = append(res, peerLocMedia)
	}
	return ctx.OK(res)
	// PeersController_Locations: end_implement
}

// RegenerateToken runs the regenerate token action.
func (c *PeersController) RegenerateToken(ctx *app.RegenerateTokenPeersContext) error {
	// PeersController_RegenerateToken: start_implement

	// Put your logic here

	at := time.Now().Unix()
	token, err := generateToken()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	peer, err := models.PeerByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound(ErrPeerNotFound))
	}

	peer.Token = token
	err = peer.Update(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	res := &app.Token{
		ID:        uint64(ctx.ID),
		Token:     token,
		UpdatedAt: at,
	}
	return ctx.OK(res)
	// PeersController_RegenerateToken: end_implement
}

// SendLocation runs the send location action.
func (c *PeersController) SendLocation(ctx *app.SendLocationPeersContext) error {
	// PeersController_SendLocation: start_implement

	// Put your logic here
	approachThreshold, err := utils.GetApproachedThreshold()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	leaveThreshold, err := utils.GetLeaveThreshold()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	peer, err := models.PeerByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	if peer.Token != ctx.Payload.Token {
		return ctx.Forbidden(goa.ErrBadRequest("Token not matched"))
	}

	at := time.Now()
	loc := models.PeerLocation{
		PeerID:    uint64(ctx.ID),
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
		CreatedAt: at,
		UpdatedAt: at,
	}

	err = loc.Insert(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	pm, err := models.PeersMapByPeerID(c.db, peer.ID)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	pod, err := pm.Pod(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	approachingMedia := app.PeerApproaching{
		Type:      ACCEPT,
		Code:      peer.Code,
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
		CreatedAt: at.Unix(),
	}
	bytes, err := json.Marshal(approachingMedia)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	err = c.redisClient.Publish(pod.Code, bytes).Err()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	podCoords := math.Coordinate{
		Latitude:  pod.Latitude,
		Longitude: pod.Longitude,
	}
	peerCoords := math.Coordinate{
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
	}

	gap := math.HubenyDistance(podCoords, peerCoords)
	if gap <= approachThreshold && !peer.Approaching.Bool {
		approachingMedia = app.PeerApproaching{
			Type:      APPROACHING,
			Code:      peer.Code,
			Latitude:  ctx.Payload.Latitude,
			Longitude: ctx.Payload.Longitude,
			CreatedAt: at.Unix(),
		}
		peer.Approaching = sql.NullBool{
			Valid: true,
			Bool:  true,
		}
		err = peer.Update(c.db)
		if err != nil {
			log.Println(err)
			return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
		}
	}
	if gap >= leaveThreshold && peer.Approaching.Bool {
		approachingMedia = app.PeerApproaching{
			Type:      LEAVED,
			Code:      peer.Code,
			Latitude:  ctx.Payload.Latitude,
			Longitude: ctx.Payload.Longitude,
			CreatedAt: at.Unix(),
		}

		peer.Approaching = sql.NullBool{
			Valid: true,
			Bool:  false,
		}
		err = peer.Update(c.db)
		if err != nil {
			log.Println(err)
			return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
		}
	}

	return ctx.Created(&approachingMedia)
	// PeersController_SendLocation: end_implement
}

// Show runs the show action.
func (c *PeersController) Show(ctx *app.ShowPeersContext) error {
	// PeersController_Show: start_implement

	// Put your logic here
	peer, err := models.PeerByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound(ErrPeerNotFound))
	}

	res := &app.Peer{
		ID:          int64(peer.ID),
		Code:        peer.Code,
		Approaching: peer.Approaching.Bool,
		CreatedAt:   peer.CreatedAt.Unix(),
		UpdatedAt:   peer.UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// PeersController_Show: end_implement
}

// Update runs the update action.
func (c *PeersController) Update(ctx *app.UpdatePeersContext) error {
	// PeersController_Update: start_implement

	// Put your logic here
	peer, err := models.PeerByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound(ErrPeerNotFound))
	}
	if ctx.Payload.Code != nil {
		peer.Code = *ctx.Payload.Code
	}
	err = peer.Update(c.db)
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrInternal(ErrInternalServerError))
	}

	return ctx.NoContent()
	// PeersController_Update: end_implement
}

// ShowByToken runs the show by token action.
func (c *PeersController) ShowByToken(ctx *app.ShowByTokenPeersContext) error {
	// PeersController_ShowByToken: start_implement

	// Put your logic here
	peer, err := models.PeerByToken(c.db, ctx.Token)
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound("Peer not found."))
	}

	res := &app.Peer{
		ID:          int64(peer.ID),
		Code:        peer.Code,
		Approaching: peer.Approaching.Bool,
		CreatedAt:   peer.CreatedAt.Unix(),
		UpdatedAt:   peer.UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// PeersController_ShowByToken: end_implement
}
