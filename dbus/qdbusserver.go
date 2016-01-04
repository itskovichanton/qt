package dbus

//#include "dbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDBusServer struct {
	core.QObject
}

type QDBusServer_ITF interface {
	core.QObject_ITF
	QDBusServer_PTR() *QDBusServer
}

func PointerFromQDBusServer(ptr QDBusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDBusServer_PTR().Pointer()
	}
	return nil
}

func NewQDBusServerFromPointer(ptr unsafe.Pointer) *QDBusServer {
	var n = new(QDBusServer)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QDBusServer_") {
		n.SetObjectName("QDBusServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QDBusServer) QDBusServer_PTR() *QDBusServer {
	return ptr
}

func NewQDBusServer2(parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer2(core.PointerFromQObject(parent)))
}

func NewQDBusServer(address string, parent core.QObject_ITF) *QDBusServer {
	defer qt.Recovering("QDBusServer::QDBusServer")

	return NewQDBusServerFromPointer(C.QDBusServer_NewQDBusServer(C.CString(address), core.PointerFromQObject(parent)))
}

func (ptr *QDBusServer) Address() string {
	defer qt.Recovering("QDBusServer::address")

	if ptr.Pointer() != nil {
		return C.GoString(C.QDBusServer_Address(ptr.Pointer()))
	}
	return ""
}

func (ptr *QDBusServer) IsAnonymousAuthenticationAllowed() bool {
	defer qt.Recovering("QDBusServer::isAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsAnonymousAuthenticationAllowed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) IsConnected() bool {
	defer qt.Recovering("QDBusServer::isConnected")

	if ptr.Pointer() != nil {
		return C.QDBusServer_IsConnected(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QDBusServer) SetAnonymousAuthenticationAllowed(value bool) {
	defer qt.Recovering("QDBusServer::setAnonymousAuthenticationAllowed")

	if ptr.Pointer() != nil {
		C.QDBusServer_SetAnonymousAuthenticationAllowed(ptr.Pointer(), C.int(qt.GoBoolToInt(value)))
	}
}

func (ptr *QDBusServer) DestroyQDBusServer() {
	defer qt.Recovering("QDBusServer::~QDBusServer")

	if ptr.Pointer() != nil {
		C.QDBusServer_DestroyQDBusServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QDBusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QDBusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQDBusServerTimerEvent
func callbackQDBusServerTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QDBusServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QDBusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QDBusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QDBusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQDBusServerChildEvent
func callbackQDBusServerChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QDBusServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QDBusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QDBusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QDBusServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QDBusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQDBusServerCustomEvent
func callbackQDBusServerCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QDBusServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQDBusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QDBusServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QDBusServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QDBusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QDBusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
