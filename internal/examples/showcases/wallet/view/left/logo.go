package left

import (
	"github.com/itskovichanton/qt/quick"

	_ "github.com/itskovichanton/qt/internal/examples/showcases/wallet/view/left/controller"
)

func init() { logoTemplate_QmlRegisterType2("LeftTemplate", 1, 0, "LogoTemplate") }

type logoTemplate struct {
	quick.QQuickItem

	_ func() `signal:"clicked,->(controller.NewLogoController(nil))"`
}
