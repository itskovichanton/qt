package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QPanGesture struct {
	QGesture
}

type QPanGesture_ITF interface {
	QGesture_ITF
	QPanGesture_PTR() *QPanGesture
}

func PointerFromQPanGesture(ptr QPanGesture_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPanGesture_PTR().Pointer()
	}
	return nil
}

func NewQPanGestureFromPointer(ptr unsafe.Pointer) *QPanGesture {
	var n = new(QPanGesture)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QPanGesture_") {
		n.SetObjectName("QPanGesture_" + qt.Identifier())
	}
	return n
}

func (ptr *QPanGesture) QPanGesture_PTR() *QPanGesture {
	return ptr
}

func (ptr *QPanGesture) Acceleration() float64 {
	defer qt.Recovering("QPanGesture::acceleration")

	if ptr.Pointer() != nil {
		return float64(C.QPanGesture_Acceleration(ptr.Pointer()))
	}
	return 0
}

func (ptr *QPanGesture) SetAcceleration(value float64) {
	defer qt.Recovering("QPanGesture::setAcceleration")

	if ptr.Pointer() != nil {
		C.QPanGesture_SetAcceleration(ptr.Pointer(), C.double(value))
	}
}

func (ptr *QPanGesture) SetLastOffset(value core.QPointF_ITF) {
	defer qt.Recovering("QPanGesture::setLastOffset")

	if ptr.Pointer() != nil {
		C.QPanGesture_SetLastOffset(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPanGesture) SetOffset(value core.QPointF_ITF) {
	defer qt.Recovering("QPanGesture::setOffset")

	if ptr.Pointer() != nil {
		C.QPanGesture_SetOffset(ptr.Pointer(), core.PointerFromQPointF(value))
	}
}

func (ptr *QPanGesture) DestroyQPanGesture() {
	defer qt.Recovering("QPanGesture::~QPanGesture")

	if ptr.Pointer() != nil {
		C.QPanGesture_DestroyQPanGesture(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QPanGesture) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QPanGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QPanGesture) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QPanGesture::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQPanGestureTimerEvent
func callbackQPanGestureTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPanGesture::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQPanGestureFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QPanGesture) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPanGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPanGesture) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QPanGesture::timerEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QPanGesture) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QPanGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QPanGesture) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QPanGesture::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQPanGestureChildEvent
func callbackQPanGestureChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPanGesture::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQPanGestureFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QPanGesture) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPanGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPanGesture) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QPanGesture::childEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QPanGesture) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QPanGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QPanGesture) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QPanGesture::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQPanGestureCustomEvent
func callbackQPanGestureCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QPanGesture::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQPanGestureFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QPanGesture) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QPanGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QPanGesture) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QPanGesture::customEvent")

	if ptr.Pointer() != nil {
		C.QPanGesture_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
