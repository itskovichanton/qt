package main

import (
	"os"

	"github.com/itskovichanton/qt/widgets"

	"github.com/itskovichanton/qt/internal/examples/uitools/calculator/ui"
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	ui.NewCalculatorForm(nil).Show()

	widgets.QApplication_Exec()
}
