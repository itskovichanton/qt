package controller

import (
	"github.com/itskovichanton/qt/core"

	_ "github.com/itskovichanton/qt/internal/examples/showcases/wallet/files/dialog/controller"
)

var ActionButtonController *actionButtonController

type actionButtonController struct {
	core.QObject

	_ func() `constructor:"init"`

	_ func(string) `signal:"showDownload,->(controller.Controller)"`
	_ func(string) `signal:"deleteRequest,auto"`
}

func (c *actionButtonController) init() {
	ActionButtonController = c
}

func (c *actionButtonController) deleteRequest(name string) {
}
