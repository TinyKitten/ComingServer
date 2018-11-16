package controller

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/TinyKitten/ComingServer/app"
	"github.com/TinyKitten/ComingServer/models"
	"github.com/goadesign/goa"
	hashids "github.com/speps/go-hashids"
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

	created := &app.Pod{
		ID:        int64(pod.ID),
		Code:      ctx.Payload.Code,
		CreatedAt: at.Unix(),
		UpdatedAt: at.Unix(),
		Rumbling:  false,
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
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
	err = models.PodUpdateToken(c.db, uint64(ctx.ID), token)
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
	id := uint64(ctx.ID)
	pID := &id
	updateInput := &models.PodUpdateInput{
		ID:        pID,
		Latitude:  ctx.Payload.Latitude,
		Longitude: ctx.Payload.Longitude,
		Code:      ctx.Payload.Code,
	}

	err := models.PodUpdate(c.db, updateInput)
	if err != nil {
		log.Println(err)
		return ctx.InternalServerError(goa.ErrInternal(ErrInternalServerError))
	}

	return ctx.NoContent()
	// PodsController_Update: end_implement
}

func generateToken() (string, error) {
	now := time.Now()
	hd := hashids.NewData()
	hd.Salt = strconv.FormatInt(now.Unix(), 10)
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}
	id, err := h.Encode([]int{1, 2, 3})
	if err != nil {
		return "", err
	}
	return id, err
}
