package main

import (
	"os"

	"github.com/itskovichanton/qt/core"
	"github.com/itskovichanton/qt/gui"
	"github.com/itskovichanton/qt/quick"

	_ "github.com/itskovichanton/qt/internal/examples/qml/extending/components/test_qml/component"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	gui.NewQGuiApplication(len(os.Args), os.Args)

	view := quick.NewQQuickView(nil)
	view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	view.SetSource(core.NewQUrl3("qrc:/app.qml", 0))
	view.Show()

	gui.QGuiApplication_Exec()
}
