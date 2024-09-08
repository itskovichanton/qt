package top

import (
	"github.com/itskovichanton/qt/quick"

	_ "github.com/itskovichanton/qt/internal/examples/showcases/wallet/view/top/controller"
)

func init() { searchTemplate_QmlRegisterType2("TopTemplate", 1, 0, "SearchTemplate") }

type searchTemplate struct {
	quick.QQuickItem

	_ func(string) `signal:"search,->(controller.NewSearchController(nil))"`
}
