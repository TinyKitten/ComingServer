package controller

import (
	"github.com/TinyKitten/ComingServer/app"
	"github.com/goadesign/goa"
)

// PodsController implements the pods resource.
type PodsController struct {
	*goa.Controller
}

// NewPodsController creates a pods controller.
func NewPodsController(service *goa.Service) *PodsController {
	return &PodsController{Controller: service.NewController("PodsController")}
}

// Add runs the add action.
func (c *PodsController) Add(ctx *app.AddPodsContext) error {
	// PodsController_Add: start_implement

	// Put your logic here

	return nil
	// PodsController_Add: end_implement
}

// List runs the list action.
func (c *PodsController) List(ctx *app.ListPodsContext) error {
	// PodsController_List: start_implement

	// Put your logic here

	res := app.PodCollection{}
	return ctx.OK(res)
	// PodsController_List: end_implement
}

// RegenerateToken runs the regenerate token action.
func (c *PodsController) RegenerateToken(ctx *app.RegenerateTokenPodsContext) error {
	// PodsController_RegenerateToken: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// PodsController_RegenerateToken: end_implement
}

// Show runs the show action.
func (c *PodsController) Show(ctx *app.ShowPodsContext) error {
	// PodsController_Show: start_implement

	// Put your logic here

	res := &app.Pod{}
	return ctx.OK(res)
	// PodsController_Show: end_implement
}

// Update runs the update action.
func (c *PodsController) Update(ctx *app.UpdatePodsContext) error {
	// PodsController_Update: start_implement

	// Put your logic here

	return nil
	// PodsController_Update: end_implement
}
