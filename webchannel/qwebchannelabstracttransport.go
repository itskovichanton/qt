package webchannel

//#include "webchannel.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWebChannelAbstractTransport struct {
	core.QObject
}

type QWebChannelAbstractTransport_ITF interface {
	core.QObject_ITF
	QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport
}

func PointerFromQWebChannelAbstractTransport(ptr QWebChannelAbstractTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWebChannelAbstractTransport_PTR().Pointer()
	}
	return nil
}

func NewQWebChannelAbstractTransportFromPointer(ptr unsafe.Pointer) *QWebChannelAbstractTransport {
	var n = new(QWebChannelAbstractTransport)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QWebChannelAbstractTransport_") {
		n.SetObjectName("QWebChannelAbstractTransport_" + qt.Identifier())
	}
	return n
}

func (ptr *QWebChannelAbstractTransport) QWebChannelAbstractTransport_PTR() *QWebChannelAbstractTransport {
	return ptr
}

func (ptr *QWebChannelAbstractTransport) ConnectMessageReceived(f func(message *core.QJsonObject, transport *QWebChannelAbstractTransport)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ConnectMessageReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "messageReceived", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectMessageReceived() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DisconnectMessageReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "messageReceived")
	}
}

//export callbackQWebChannelAbstractTransportMessageReceived
func callbackQWebChannelAbstractTransportMessageReceived(ptr unsafe.Pointer, ptrName *C.char, message unsafe.Pointer, transport unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::messageReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "messageReceived"); signal != nil {
		signal.(func(*core.QJsonObject, *QWebChannelAbstractTransport))(core.NewQJsonObjectFromPointer(message), NewQWebChannelAbstractTransportFromPointer(transport))
	}

}

func (ptr *QWebChannelAbstractTransport) MessageReceived(message core.QJsonObject_ITF, transport QWebChannelAbstractTransport_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::messageReceived")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_MessageReceived(ptr.Pointer(), core.PointerFromQJsonObject(message), PointerFromQWebChannelAbstractTransport(transport))
	}
}

func (ptr *QWebChannelAbstractTransport) SendMessage(message core.QJsonObject_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::sendMessage")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_SendMessage(ptr.Pointer(), core.PointerFromQJsonObject(message))
	}
}

func (ptr *QWebChannelAbstractTransport) DestroyQWebChannelAbstractTransport() {
	defer qt.Recovering("QWebChannelAbstractTransport::~QWebChannelAbstractTransport")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQWebChannelAbstractTransportTimerEvent
func callbackQWebChannelAbstractTransportTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::timerEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQWebChannelAbstractTransportChildEvent
func callbackQWebChannelAbstractTransportChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::childEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QWebChannelAbstractTransport) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQWebChannelAbstractTransportCustomEvent
func callbackQWebChannelAbstractTransportCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QWebChannelAbstractTransport::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQWebChannelAbstractTransportFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QWebChannelAbstractTransport) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QWebChannelAbstractTransport) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QWebChannelAbstractTransport::customEvent")

	if ptr.Pointer() != nil {
		C.QWebChannelAbstractTransport_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
