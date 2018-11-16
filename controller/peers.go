package controller

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
)

// PeersController implements the peers resource.
type PeersController struct {
	*goa.Controller
}

// NewPeersController creates a peers controller.
func NewPeersController(service *goa.Service) *PeersController {
	return &PeersController{Controller: service.NewController("PeersController")}
}

// Add runs the add action.
func (c *PeersController) Add(ctx *app.AddPeersContext) error {
	// PeersController_Add: start_implement

	// Put your logic here

	return nil
	// PeersController_Add: end_implement
}

// CurrentLocation runs the current location action.
func (c *PeersController) CurrentLocation(ctx *app.CurrentLocationPeersContext) error {
	// PeersController_CurrentLocation: start_implement

	// Put your logic here

	res := &app.PeerLocation{}
	return ctx.OK(res)
	// PeersController_CurrentLocation: end_implement
}

// List runs the list action.
func (c *PeersController) List(ctx *app.ListPeersContext) error {
	// PeersController_List: start_implement

	// Put your logic here

	res := app.PeerCollection{}
	return ctx.OK(res)
	// PeersController_List: end_implement
}

// Locations runs the locations action.
func (c *PeersController) Locations(ctx *app.LocationsPeersContext) error {
	// PeersController_Locations: start_implement

	// Put your logic here

	res := app.PeerLocationCollection{}
	return ctx.OK(res)
	// PeersController_Locations: end_implement
}

// RegenerateToken runs the regenerate token action.
func (c *PeersController) RegenerateToken(ctx *app.RegenerateTokenPeersContext) error {
	// PeersController_RegenerateToken: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// PeersController_RegenerateToken: end_implement
}

// SendLocation runs the send location action.
func (c *PeersController) SendLocation(ctx *app.SendLocationPeersContext) error {
	// PeersController_SendLocation: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// PeersController_SendLocation: end_implement
}

// Show runs the show action.
func (c *PeersController) Show(ctx *app.ShowPeersContext) error {
	// PeersController_Show: start_implement

	// Put your logic here

	res := &app.Peer{}
	return ctx.OK(res)
	// PeersController_Show: end_implement
}

// Update runs the update action.
func (c *PeersController) Update(ctx *app.UpdatePeersContext) error {
	// PeersController_Update: start_implement

	// Put your logic here

	return nil
	// PeersController_Update: end_implement
}
