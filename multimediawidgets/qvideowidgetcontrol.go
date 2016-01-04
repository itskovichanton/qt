package multimediawidgets

//#include "multimediawidgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
	"unsafe"
)

type QVideoWidgetControl struct {
	multimedia.QMediaControl
}

type QVideoWidgetControl_ITF interface {
	multimedia.QMediaControl_ITF
	QVideoWidgetControl_PTR() *QVideoWidgetControl
}

func PointerFromQVideoWidgetControl(ptr QVideoWidgetControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoWidgetControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoWidgetControlFromPointer(ptr unsafe.Pointer) *QVideoWidgetControl {
	var n = new(QVideoWidgetControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoWidgetControl_") {
		n.SetObjectName("QVideoWidgetControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoWidgetControl) QVideoWidgetControl_PTR() *QVideoWidgetControl {
	return ptr
}

func (ptr *QVideoWidgetControl) AspectRatioMode() core.Qt__AspectRatioMode {
	defer qt.Recovering("QVideoWidgetControl::aspectRatioMode")

	if ptr.Pointer() != nil {
		return core.Qt__AspectRatioMode(C.QVideoWidgetControl_AspectRatioMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) Brightness() int {
	defer qt.Recovering("QVideoWidgetControl::brightness")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Brightness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectBrightnessChanged(f func(brightness int)) {
	defer qt.Recovering("connect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectBrightnessChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "brightnessChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectBrightnessChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectBrightnessChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "brightnessChanged")
	}
}

//export callbackQVideoWidgetControlBrightnessChanged
func callbackQVideoWidgetControlBrightnessChanged(ptr unsafe.Pointer, ptrName *C.char, brightness C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::brightnessChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "brightnessChanged"); signal != nil {
		signal.(func(int))(int(brightness))
	}

}

func (ptr *QVideoWidgetControl) BrightnessChanged(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::brightnessChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_BrightnessChanged(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) Contrast() int {
	defer qt.Recovering("QVideoWidgetControl::contrast")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Contrast(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectContrastChanged(f func(contrast int)) {
	defer qt.Recovering("connect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectContrastChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "contrastChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectContrastChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectContrastChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "contrastChanged")
	}
}

//export callbackQVideoWidgetControlContrastChanged
func callbackQVideoWidgetControlContrastChanged(ptr unsafe.Pointer, ptrName *C.char, contrast C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::contrastChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "contrastChanged"); signal != nil {
		signal.(func(int))(int(contrast))
	}

}

func (ptr *QVideoWidgetControl) ContrastChanged(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::contrastChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ContrastChanged(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) ConnectFullScreenChanged(f func(fullScreen bool)) {
	defer qt.Recovering("connect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectFullScreenChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "fullScreenChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectFullScreenChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectFullScreenChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "fullScreenChanged")
	}
}

//export callbackQVideoWidgetControlFullScreenChanged
func callbackQVideoWidgetControlFullScreenChanged(ptr unsafe.Pointer, ptrName *C.char, fullScreen C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::fullScreenChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "fullScreenChanged"); signal != nil {
		signal.(func(bool))(int(fullScreen) != 0)
	}

}

func (ptr *QVideoWidgetControl) FullScreenChanged(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::fullScreenChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_FullScreenChanged(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) Hue() int {
	defer qt.Recovering("QVideoWidgetControl::hue")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Hue(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectHueChanged(f func(hue int)) {
	defer qt.Recovering("connect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectHueChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hueChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectHueChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectHueChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hueChanged")
	}
}

//export callbackQVideoWidgetControlHueChanged
func callbackQVideoWidgetControlHueChanged(ptr unsafe.Pointer, ptrName *C.char, hue C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::hueChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hueChanged"); signal != nil {
		signal.(func(int))(int(hue))
	}

}

func (ptr *QVideoWidgetControl) HueChanged(hue int) {
	defer qt.Recovering("QVideoWidgetControl::hueChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_HueChanged(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) IsFullScreen() bool {
	defer qt.Recovering("QVideoWidgetControl::isFullScreen")

	if ptr.Pointer() != nil {
		return C.QVideoWidgetControl_IsFullScreen(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVideoWidgetControl) Saturation() int {
	defer qt.Recovering("QVideoWidgetControl::saturation")

	if ptr.Pointer() != nil {
		return int(C.QVideoWidgetControl_Saturation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoWidgetControl) ConnectSaturationChanged(f func(saturation int)) {
	defer qt.Recovering("connect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ConnectSaturationChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "saturationChanged", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectSaturationChanged() {
	defer qt.Recovering("disconnect QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DisconnectSaturationChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "saturationChanged")
	}
}

//export callbackQVideoWidgetControlSaturationChanged
func callbackQVideoWidgetControlSaturationChanged(ptr unsafe.Pointer, ptrName *C.char, saturation C.int) {
	defer qt.Recovering("callback QVideoWidgetControl::saturationChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "saturationChanged"); signal != nil {
		signal.(func(int))(int(saturation))
	}

}

func (ptr *QVideoWidgetControl) SaturationChanged(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::saturationChanged")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SaturationChanged(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) SetAspectRatioMode(mode core.Qt__AspectRatioMode) {
	defer qt.Recovering("QVideoWidgetControl::setAspectRatioMode")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetAspectRatioMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QVideoWidgetControl) SetBrightness(brightness int) {
	defer qt.Recovering("QVideoWidgetControl::setBrightness")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetBrightness(ptr.Pointer(), C.int(brightness))
	}
}

func (ptr *QVideoWidgetControl) SetContrast(contrast int) {
	defer qt.Recovering("QVideoWidgetControl::setContrast")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetContrast(ptr.Pointer(), C.int(contrast))
	}
}

func (ptr *QVideoWidgetControl) SetFullScreen(fullScreen bool) {
	defer qt.Recovering("QVideoWidgetControl::setFullScreen")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetFullScreen(ptr.Pointer(), C.int(qt.GoBoolToInt(fullScreen)))
	}
}

func (ptr *QVideoWidgetControl) SetHue(hue int) {
	defer qt.Recovering("QVideoWidgetControl::setHue")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetHue(ptr.Pointer(), C.int(hue))
	}
}

func (ptr *QVideoWidgetControl) SetSaturation(saturation int) {
	defer qt.Recovering("QVideoWidgetControl::setSaturation")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_SetSaturation(ptr.Pointer(), C.int(saturation))
	}
}

func (ptr *QVideoWidgetControl) VideoWidget() *widgets.QWidget {
	defer qt.Recovering("QVideoWidgetControl::videoWidget")

	if ptr.Pointer() != nil {
		return widgets.NewQWidgetFromPointer(C.QVideoWidgetControl_VideoWidget(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVideoWidgetControl) DestroyQVideoWidgetControl() {
	defer qt.Recovering("QVideoWidgetControl::~QVideoWidgetControl")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_DestroyQVideoWidgetControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QVideoWidgetControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoWidgetControlTimerEvent
func callbackQVideoWidgetControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidgetControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoWidgetControlChildEvent
func callbackQVideoWidgetControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoWidgetControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoWidgetControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoWidgetControlCustomEvent
func callbackQVideoWidgetControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoWidgetControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoWidgetControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoWidgetControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoWidgetControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoWidgetControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoWidgetControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
