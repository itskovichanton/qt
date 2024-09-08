package mpv

//it's also possible to directly place this file beside the files from mpv-examples/libmpv/qml
//to get this working, change init.go to just run `make mocables`
//and remove the cgo LDFLAGS line used to link against the static lib below
//more info: https://github.com/itskovichanton/qt/issues/1162

/*
#cgo pkg-config: mpv
#cgo linux,amd64 LDFLAGS: -L ${SRCDIR}/mpv-examples/libmpv/qml -lmpvtest

void initMpv();
*/
import "C"
import (
	"github.com/itskovichanton/qt/core"
	_ "github.com/itskovichanton/qt/quick"
)

type stub struct{ core.QObject } //TODO: needed for linking at the moment

func InitMpv() {
	C.initMpv()
}
