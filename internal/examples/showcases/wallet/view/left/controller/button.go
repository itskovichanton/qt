package controller

import (
	"github.com/itskovichanton/qt/core"

	"github.com/itskovichanton/qt/internal/examples/showcases/wallet/view/controller"
)

type buttonController struct {
	core.QObject

	_ func(code string) `signal:"clicked,auto"`
}

// lazy binding to the view/stack controller
func (c *buttonController) clicked(code string) { controller.StackController.Clicked(code) }
