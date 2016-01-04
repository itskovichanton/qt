package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAudioInputSelectorControl struct {
	QMediaControl
}

type QAudioInputSelectorControl_ITF interface {
	QMediaControl_ITF
	QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl
}

func PointerFromQAudioInputSelectorControl(ptr QAudioInputSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAudioInputSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQAudioInputSelectorControlFromPointer(ptr unsafe.Pointer) *QAudioInputSelectorControl {
	var n = new(QAudioInputSelectorControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAudioInputSelectorControl_") {
		n.SetObjectName("QAudioInputSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QAudioInputSelectorControl) QAudioInputSelectorControl_PTR() *QAudioInputSelectorControl {
	return ptr
}

func (ptr *QAudioInputSelectorControl) ActiveInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::activeInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_ActiveInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) ConnectActiveInputChanged(f func(name string)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectActiveInputChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "activeInputChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectActiveInputChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectActiveInputChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "activeInputChanged")
	}
}

//export callbackQAudioInputSelectorControlActiveInputChanged
func callbackQAudioInputSelectorControlActiveInputChanged(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::activeInputChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "activeInputChanged"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QAudioInputSelectorControl) ActiveInputChanged(name string) {
	defer qt.Recovering("QAudioInputSelectorControl::activeInputChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ActiveInputChanged(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectAvailableInputsChanged(f func()) {
	defer qt.Recovering("connect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ConnectAvailableInputsChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "availableInputsChanged", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectAvailableInputsChanged() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DisconnectAvailableInputsChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "availableInputsChanged")
	}
}

//export callbackQAudioInputSelectorControlAvailableInputsChanged
func callbackQAudioInputSelectorControlAvailableInputsChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QAudioInputSelectorControl::availableInputsChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "availableInputsChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QAudioInputSelectorControl) AvailableInputsChanged() {
	defer qt.Recovering("QAudioInputSelectorControl::availableInputsChanged")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_AvailableInputsChanged(ptr.Pointer())
	}
}

func (ptr *QAudioInputSelectorControl) DefaultInput() string {
	defer qt.Recovering("QAudioInputSelectorControl::defaultInput")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_DefaultInput(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) InputDescription(name string) string {
	defer qt.Recovering("QAudioInputSelectorControl::inputDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAudioInputSelectorControl_InputDescription(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QAudioInputSelectorControl) SetActiveInput(name string) {
	defer qt.Recovering("QAudioInputSelectorControl::setActiveInput")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_SetActiveInput(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QAudioInputSelectorControl) DestroyQAudioInputSelectorControl() {
	defer qt.Recovering("QAudioInputSelectorControl::~QAudioInputSelectorControl")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAudioInputSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAudioInputSelectorControlTimerEvent
func callbackQAudioInputSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAudioInputSelectorControlChildEvent
func callbackQAudioInputSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAudioInputSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAudioInputSelectorControlCustomEvent
func callbackQAudioInputSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAudioInputSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAudioInputSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAudioInputSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAudioInputSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAudioInputSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QAudioInputSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
