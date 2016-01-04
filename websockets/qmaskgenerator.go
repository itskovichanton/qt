package websockets

//#include "websockets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMaskGenerator struct {
	core.QObject
}

type QMaskGenerator_ITF interface {
	core.QObject_ITF
	QMaskGenerator_PTR() *QMaskGenerator
}

func PointerFromQMaskGenerator(ptr QMaskGenerator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMaskGenerator_PTR().Pointer()
	}
	return nil
}

func NewQMaskGeneratorFromPointer(ptr unsafe.Pointer) *QMaskGenerator {
	var n = new(QMaskGenerator)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QMaskGenerator_") {
		n.SetObjectName("QMaskGenerator_" + qt.Identifier())
	}
	return n
}

func (ptr *QMaskGenerator) QMaskGenerator_PTR() *QMaskGenerator {
	return ptr
}

func (ptr *QMaskGenerator) Seed() bool {
	defer qt.Recovering("QMaskGenerator::seed")

	if ptr.Pointer() != nil {
		return C.QMaskGenerator_Seed(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMaskGenerator) DestroyQMaskGenerator() {
	defer qt.Recovering("QMaskGenerator::~QMaskGenerator")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_DestroyQMaskGenerator(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QMaskGenerator) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQMaskGeneratorTimerEvent
func callbackQMaskGeneratorTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::timerEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QMaskGenerator) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQMaskGeneratorChildEvent
func callbackQMaskGeneratorChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::childEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QMaskGenerator) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QMaskGenerator) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQMaskGeneratorCustomEvent
func callbackQMaskGeneratorCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QMaskGenerator::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQMaskGeneratorFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QMaskGenerator) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QMaskGenerator) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QMaskGenerator::customEvent")

	if ptr.Pointer() != nil {
		C.QMaskGenerator_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
