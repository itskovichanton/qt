package dialog

import (
	"github.com/itskovichanton/qt/core"

	_ "github.com/itskovichanton/qt/internal/examples/showcases/wallet/wallet/dialog/controller"
)

func init() { recoverTemplate_QmlRegisterType2("DialogTemplate", 1, 0, "RecoverTemplate") }

type recoverTemplate struct {
	dialogTemplate

	_ func(string) *core.QVariant `slot:"recover,->(controller.Controller)"`
}
