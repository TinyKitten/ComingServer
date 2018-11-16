package controller

import (
	"database/sql"
	"log"
	"time"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/goadesign/goa"
)

// PodsController implements the pods resource.
type PodsController struct {
	*goa.Controller
	db *sql.DB
}

// NewPodsController creates a pods controller.
func NewPodsController(service *goa.Service, db *sql.DB) *PodsController {
	return &PodsController{
		Controller: service.NewController("PodsController"),
		db:         db,
	}
}

// Add runs the add action.
func (c *PodsController) Add(ctx *app.AddPodsContext) error {
	// PodsController_Add: start_implement

	// Put your logic here
	at := time.Now()
	token, err := generateToken()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	pod := models.Pod{
		Code:      ctx.Payload.Code,
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
		Token:     token,
		CreatedAt: at,
		UpdatedAt: at,
	}

	err = pod.Insert(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(err))
	}

	created := &app.PodCreated{
		ID:          int64(pod.ID),
		Code:        ctx.Payload.Code,
		CreatedAt:   at.Unix(),
		UpdatedAt:   at.Unix(),
		Approaching: false,
		Token:       token,
		Latitude:    ctx.Payload.Latitude,
		Longitude:   ctx.Payload.Longitude,
	}

	return ctx.Created(created)
	// PodsController_Add: end_implement
}

// List runs the list action.
func (c *PodsController) List(ctx *app.ListPodsContext) error {
	// PodsController_List: start_implement

	// Put your logic here

	res := app.PodCollection{}
	pods, err := models.PodList(c.db, ctx.Offset, ctx.Limit)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	for _, pod := range pods {
		podMedia := &app.Pod{
			ID:        int64(pod.ID),
			Code:      pod.Code,
			CreatedAt: pod.CreatedAt.Unix(),
			UpdatedAt: pod.UpdatedAt.Unix(),
			Latitude:  pod.Latitude,
			Longitude: pod.Longitude,
		}
		res = append(res, podMedia)
	}
	return ctx.OK(res)
	// PodsController_List: end_implement
}

// RegenerateToken runs the regenerate token action.
func (c *PodsController) RegenerateToken(ctx *app.RegenerateTokenPodsContext) error {
	// PodsController_RegenerateToken: start_implement

	// Put your logic here
	at := time.Now().Unix()
	token, err := generateToken()
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	pod, err := models.PodByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound(ErrPodNotFound))
	}

	pod.Token = token

	pod.Update(c.db)

	res := &app.Token{
		ID:        uint64(ctx.ID),
		Token:     token,
		UpdatedAt: at,
	}
	return ctx.OK(res)
	// PodsController_RegenerateToken: end_implement
}

// Show runs the show action.
func (c *PodsController) Show(ctx *app.ShowPodsContext) error {
	// PodsController_Show: start_implement

	// Put your logic here
	pod, err := models.PodByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}
	res := &app.Pod{
		ID:        int64(pod.ID),
		Code:      pod.Code,
		Latitude:  pod.Latitude,
		Longitude: pod.Longitude,
		CreatedAt: pod.CreatedAt.Unix(),
		UpdatedAt: pod.UpdatedAt.Unix(),
	}
	return ctx.OK(res)
	// PodsController_Show: end_implement
}

// Update runs the update action.
func (c *PodsController) Update(ctx *app.UpdatePodsContext) error {
	// PodsController_Update: start_implement

	// Put your logic here
	pod, err := models.PodByID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound(ErrPodNotFound))
	}
	if ctx.Payload.Latitude != nil {
		pod.Latitude = *ctx.Payload.Latitude
	}
	if ctx.Payload.Longitude != nil {
		pod.Longitude = *ctx.Payload.Longitude
	}
	if ctx.Payload.Code != nil {
		pod.Code = *ctx.Payload.Code
	}

	err = pod.Update(c.db)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	return ctx.NoContent()
	// PodsController_Update: end_implement
}

// PeersList runs the peers list action.
func (c *PodsController) PeersList(ctx *app.PeersListPodsContext) error {
	// PodsController_PeersList: start_implement

	// Put your logic here
	peersMap, err := models.PeersMapsByPodID(c.db, uint64(ctx.ID))
	if err != nil {
		log.Println(err)
		return ctx.NotFound(goa.ErrNotFound("Pod not found"))
	}

	if len(peersMap) == 0 {
		return ctx.NotFound(goa.ErrNotFound("Peers not recorded"))
	}

	res := app.PeerCollection{}

	for _, p := range peersMap {
		peer, err := p.Peer(c.db)
		if err != nil {
			log.Println(err)
			return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
		}
		peerMedia := &app.Peer{
			ID:        int64(peer.ID),
			Code:      peer.Code,
			CreatedAt: peer.CreatedAt.Unix(),
			UpdatedAt: peer.UpdatedAt.Unix(),
		}
		res = append(res, peerMedia)
	}

	return ctx.OK(res)
	// PodsController_PeersList: end_implement
}
