package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QCameraCaptureDestinationControl struct {
	QMediaControl
}

type QCameraCaptureDestinationControl_ITF interface {
	QMediaControl_ITF
	QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl
}

func PointerFromQCameraCaptureDestinationControl(ptr QCameraCaptureDestinationControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCameraCaptureDestinationControl_PTR().Pointer()
	}
	return nil
}

func NewQCameraCaptureDestinationControlFromPointer(ptr unsafe.Pointer) *QCameraCaptureDestinationControl {
	var n = new(QCameraCaptureDestinationControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QCameraCaptureDestinationControl_") {
		n.SetObjectName("QCameraCaptureDestinationControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QCameraCaptureDestinationControl) QCameraCaptureDestinationControl_PTR() *QCameraCaptureDestinationControl {
	return ptr
}

func (ptr *QCameraCaptureDestinationControl) CaptureDestination() QCameraImageCapture__CaptureDestination {
	defer qt.Recovering("QCameraCaptureDestinationControl::captureDestination")

	if ptr.Pointer() != nil {
		return QCameraImageCapture__CaptureDestination(C.QCameraCaptureDestinationControl_CaptureDestination(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCameraCaptureDestinationControl) ConnectCaptureDestinationChanged(f func(destination QCameraImageCapture__CaptureDestination)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ConnectCaptureDestinationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "captureDestinationChanged", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectCaptureDestinationChanged() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DisconnectCaptureDestinationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "captureDestinationChanged")
	}
}

//export callbackQCameraCaptureDestinationControlCaptureDestinationChanged
func callbackQCameraCaptureDestinationControlCaptureDestinationChanged(ptr unsafe.Pointer, ptrName *C.char, destination C.int) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::captureDestinationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "captureDestinationChanged"); signal != nil {
		signal.(func(QCameraImageCapture__CaptureDestination))(QCameraImageCapture__CaptureDestination(destination))
	}

}

func (ptr *QCameraCaptureDestinationControl) CaptureDestinationChanged(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraCaptureDestinationControl::captureDestinationChanged")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CaptureDestinationChanged(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraCaptureDestinationControl) IsCaptureDestinationSupported(destination QCameraImageCapture__CaptureDestination) bool {
	defer qt.Recovering("QCameraCaptureDestinationControl::isCaptureDestinationSupported")

	if ptr.Pointer() != nil {
		return C.QCameraCaptureDestinationControl_IsCaptureDestinationSupported(ptr.Pointer(), C.int(destination)) != 0
	}
	return false
}

func (ptr *QCameraCaptureDestinationControl) SetCaptureDestination(destination QCameraImageCapture__CaptureDestination) {
	defer qt.Recovering("QCameraCaptureDestinationControl::setCaptureDestination")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_SetCaptureDestination(ptr.Pointer(), C.int(destination))
	}
}

func (ptr *QCameraCaptureDestinationControl) DestroyQCameraCaptureDestinationControl() {
	defer qt.Recovering("QCameraCaptureDestinationControl::~QCameraCaptureDestinationControl")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_DestroyQCameraCaptureDestinationControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQCameraCaptureDestinationControlTimerEvent
func callbackQCameraCaptureDestinationControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQCameraCaptureDestinationControlChildEvent
func callbackQCameraCaptureDestinationControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::childEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCameraCaptureDestinationControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQCameraCaptureDestinationControlCustomEvent
func callbackQCameraCaptureDestinationControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCameraCaptureDestinationControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCameraCaptureDestinationControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCameraCaptureDestinationControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCameraCaptureDestinationControl::customEvent")

	if ptr.Pointer() != nil {
		C.QCameraCaptureDestinationControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
