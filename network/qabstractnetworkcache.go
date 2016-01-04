package network

//#include "network.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAbstractNetworkCache struct {
	core.QObject
}

type QAbstractNetworkCache_ITF interface {
	core.QObject_ITF
	QAbstractNetworkCache_PTR() *QAbstractNetworkCache
}

func PointerFromQAbstractNetworkCache(ptr QAbstractNetworkCache_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractNetworkCache_PTR().Pointer()
	}
	return nil
}

func NewQAbstractNetworkCacheFromPointer(ptr unsafe.Pointer) *QAbstractNetworkCache {
	var n = new(QAbstractNetworkCache)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QAbstractNetworkCache_") {
		n.SetObjectName("QAbstractNetworkCache_" + qt.Identifier())
	}
	return n
}

func (ptr *QAbstractNetworkCache) QAbstractNetworkCache_PTR() *QAbstractNetworkCache {
	return ptr
}

func (ptr *QAbstractNetworkCache) CacheSize() int64 {
	defer qt.Recovering("QAbstractNetworkCache::cacheSize")

	if ptr.Pointer() != nil {
		return int64(C.QAbstractNetworkCache_CacheSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAbstractNetworkCache) Clear() {
	defer qt.Recovering("QAbstractNetworkCache::clear")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_Clear(ptr.Pointer())
	}
}

func (ptr *QAbstractNetworkCache) Data(url core.QUrl_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::data")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Data(ptr.Pointer(), core.PointerFromQUrl(url)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) Prepare(metaData QNetworkCacheMetaData_ITF) *core.QIODevice {
	defer qt.Recovering("QAbstractNetworkCache::prepare")

	if ptr.Pointer() != nil {
		return core.NewQIODeviceFromPointer(C.QAbstractNetworkCache_Prepare(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData)))
	}
	return nil
}

func (ptr *QAbstractNetworkCache) UpdateMetaData(metaData QNetworkCacheMetaData_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::updateMetaData")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_UpdateMetaData(ptr.Pointer(), PointerFromQNetworkCacheMetaData(metaData))
	}
}

func (ptr *QAbstractNetworkCache) DestroyQAbstractNetworkCache() {
	defer qt.Recovering("QAbstractNetworkCache::~QAbstractNetworkCache")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_DestroyQAbstractNetworkCache(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QAbstractNetworkCache) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQAbstractNetworkCacheTimerEvent
func callbackQAbstractNetworkCacheTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::timerEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQAbstractNetworkCacheChildEvent
func callbackQAbstractNetworkCacheChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::childEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QAbstractNetworkCache) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQAbstractNetworkCacheCustomEvent
func callbackQAbstractNetworkCacheCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QAbstractNetworkCache::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQAbstractNetworkCacheFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QAbstractNetworkCache) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QAbstractNetworkCache::customEvent")

	if ptr.Pointer() != nil {
		C.QAbstractNetworkCache_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
