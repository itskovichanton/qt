//go:build qml
// +build qml

package main

import (
	"os"
	"path/filepath"

	"github.com/itskovichanton/qt/core"
	"github.com/itskovichanton/qt/qml"
	"github.com/itskovichanton/qt/widgets"

	"github.com/itskovichanton/qt/internal/examples/sql/masterdetail_qml/controller"

	_ "github.com/itskovichanton/qt/internal/examples/sql/masterdetail_qml/view"
)

const PRODUCTION = true

func init() {
	if !PRODUCTION {
		os.Setenv("QML_DISABLE_DISK_CACHE", "true")
	}
}

func main() {
	var path string
	if PRODUCTION {
		path = "qrc:/qml/view.qml"
	} else {
		path = filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "itskovichanton", "qt", "internal", "examples", "sql", "masterdetail_qml", "view", "qml", "view.qml")
	}

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	qApp := widgets.NewQApplication(len(os.Args), os.Args)

	controller.NewController(nil).InitWith(core.NewQFile2(":/albumdetails.xml"), qApp)

	app := qml.NewQQmlApplicationEngine(nil)
	if PRODUCTION {
		app.AddImportPath("qrc:/qml/")
	} else {
		app.AddImportPath("./view/detail/qml")
		app.AddImportPath("./view/album/qml")
		app.AddImportPath("./view/artist/qml")
		app.AddImportPath("./view/dialog/qml")
	}
	app.Load(core.NewQUrl3(path, 0))

	widgets.QApplication_Exec()
}
