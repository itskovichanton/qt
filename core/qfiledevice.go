package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QFileDevice struct {
	QIODevice
}

type QFileDevice_ITF interface {
	QIODevice_ITF
	QFileDevice_PTR() *QFileDevice
}

func PointerFromQFileDevice(ptr QFileDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFileDevice_PTR().Pointer()
	}
	return nil
}

func NewQFileDeviceFromPointer(ptr unsafe.Pointer) *QFileDevice {
	var n = new(QFileDevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QFileDevice_") {
		n.SetObjectName("QFileDevice_" + qt.Identifier())
	}
	return n
}

func (ptr *QFileDevice) QFileDevice_PTR() *QFileDevice {
	return ptr
}

//QFileDevice::FileError
type QFileDevice__FileError int64

const (
	QFileDevice__NoError          = QFileDevice__FileError(0)
	QFileDevice__ReadError        = QFileDevice__FileError(1)
	QFileDevice__WriteError       = QFileDevice__FileError(2)
	QFileDevice__FatalError       = QFileDevice__FileError(3)
	QFileDevice__ResourceError    = QFileDevice__FileError(4)
	QFileDevice__OpenError        = QFileDevice__FileError(5)
	QFileDevice__AbortError       = QFileDevice__FileError(6)
	QFileDevice__TimeOutError     = QFileDevice__FileError(7)
	QFileDevice__UnspecifiedError = QFileDevice__FileError(8)
	QFileDevice__RemoveError      = QFileDevice__FileError(9)
	QFileDevice__RenameError      = QFileDevice__FileError(10)
	QFileDevice__PositionError    = QFileDevice__FileError(11)
	QFileDevice__ResizeError      = QFileDevice__FileError(12)
	QFileDevice__PermissionsError = QFileDevice__FileError(13)
	QFileDevice__CopyError        = QFileDevice__FileError(14)
)

//QFileDevice::FileHandleFlag
type QFileDevice__FileHandleFlag int64

const (
	QFileDevice__AutoCloseHandle = QFileDevice__FileHandleFlag(0x0001)
	QFileDevice__DontCloseHandle = QFileDevice__FileHandleFlag(0)
)

//QFileDevice::MemoryMapFlags
type QFileDevice__MemoryMapFlags int64

const (
	QFileDevice__NoOptions        = QFileDevice__MemoryMapFlags(0)
	QFileDevice__MapPrivateOption = QFileDevice__MemoryMapFlags(0x0001)
)

//QFileDevice::Permission
type QFileDevice__Permission int64

const (
	QFileDevice__ReadOwner  = QFileDevice__Permission(0x4000)
	QFileDevice__WriteOwner = QFileDevice__Permission(0x2000)
	QFileDevice__ExeOwner   = QFileDevice__Permission(0x1000)
	QFileDevice__ReadUser   = QFileDevice__Permission(0x0400)
	QFileDevice__WriteUser  = QFileDevice__Permission(0x0200)
	QFileDevice__ExeUser    = QFileDevice__Permission(0x0100)
	QFileDevice__ReadGroup  = QFileDevice__Permission(0x0040)
	QFileDevice__WriteGroup = QFileDevice__Permission(0x0020)
	QFileDevice__ExeGroup   = QFileDevice__Permission(0x0010)
	QFileDevice__ReadOther  = QFileDevice__Permission(0x0004)
	QFileDevice__WriteOther = QFileDevice__Permission(0x0002)
	QFileDevice__ExeOther   = QFileDevice__Permission(0x0001)
)

func (ptr *QFileDevice) Seek(pos int64) bool {
	defer qt.Recovering("QFileDevice::seek")

	if ptr.Pointer() != nil {
		return C.QFileDevice_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QFileDevice) AtEnd() bool {
	defer qt.Recovering("QFileDevice::atEnd")

	if ptr.Pointer() != nil {
		return C.QFileDevice_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) ConnectClose(f func()) {
	defer qt.Recovering("connect QFileDevice::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QFileDevice) DisconnectClose() {
	defer qt.Recovering("disconnect QFileDevice::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

//export callbackQFileDeviceClose
func callbackQFileDeviceClose(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QFileDevice::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQFileDeviceFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QFileDevice) Close() {
	defer qt.Recovering("QFileDevice::close")

	if ptr.Pointer() != nil {
		C.QFileDevice_Close(ptr.Pointer())
	}
}

func (ptr *QFileDevice) CloseDefault() {
	defer qt.Recovering("QFileDevice::close")

	if ptr.Pointer() != nil {
		C.QFileDevice_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QFileDevice) Error() QFileDevice__FileError {
	defer qt.Recovering("QFileDevice::error")

	if ptr.Pointer() != nil {
		return QFileDevice__FileError(C.QFileDevice_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) FileName() string {
	defer qt.Recovering("QFileDevice::fileName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QFileDevice_FileName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QFileDevice) Flush() bool {
	defer qt.Recovering("QFileDevice::flush")

	if ptr.Pointer() != nil {
		return C.QFileDevice_Flush(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) Handle() int {
	defer qt.Recovering("QFileDevice::handle")

	if ptr.Pointer() != nil {
		return int(C.QFileDevice_Handle(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) IsSequential() bool {
	defer qt.Recovering("QFileDevice::isSequential")

	if ptr.Pointer() != nil {
		return C.QFileDevice_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QFileDevice) Permissions() QFileDevice__Permission {
	defer qt.Recovering("QFileDevice::permissions")

	if ptr.Pointer() != nil {
		return QFileDevice__Permission(C.QFileDevice_Permissions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) Pos() int64 {
	defer qt.Recovering("QFileDevice::pos")

	if ptr.Pointer() != nil {
		return int64(C.QFileDevice_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) ReadData(data string, len int64) int64 {
	defer qt.Recovering("QFileDevice::readData")

	if ptr.Pointer() != nil {
		return int64(C.QFileDevice_ReadData(ptr.Pointer(), C.CString(data), C.longlong(len)))
	}
	return 0
}

func (ptr *QFileDevice) ReadLineData(data string, maxlen int64) int64 {
	defer qt.Recovering("QFileDevice::readLineData")

	if ptr.Pointer() != nil {
		return int64(C.QFileDevice_ReadLineData(ptr.Pointer(), C.CString(data), C.longlong(maxlen)))
	}
	return 0
}

func (ptr *QFileDevice) Resize(sz int64) bool {
	defer qt.Recovering("QFileDevice::resize")

	if ptr.Pointer() != nil {
		return C.QFileDevice_Resize(ptr.Pointer(), C.longlong(sz)) != 0
	}
	return false
}

func (ptr *QFileDevice) SetPermissions(permissions QFileDevice__Permission) bool {
	defer qt.Recovering("QFileDevice::setPermissions")

	if ptr.Pointer() != nil {
		return C.QFileDevice_SetPermissions(ptr.Pointer(), C.int(permissions)) != 0
	}
	return false
}

func (ptr *QFileDevice) Size() int64 {
	defer qt.Recovering("QFileDevice::size")

	if ptr.Pointer() != nil {
		return int64(C.QFileDevice_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QFileDevice) UnsetError() {
	defer qt.Recovering("QFileDevice::unsetError")

	if ptr.Pointer() != nil {
		C.QFileDevice_UnsetError(ptr.Pointer())
	}
}

func (ptr *QFileDevice) WriteData(data string, len int64) int64 {
	defer qt.Recovering("QFileDevice::writeData")

	if ptr.Pointer() != nil {
		return int64(C.QFileDevice_WriteData(ptr.Pointer(), C.CString(data), C.longlong(len)))
	}
	return 0
}

func (ptr *QFileDevice) DestroyQFileDevice() {
	defer qt.Recovering("QFileDevice::~QFileDevice")

	if ptr.Pointer() != nil {
		C.QFileDevice_DestroyQFileDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QFileDevice) ConnectTimerEvent(f func(event *QTimerEvent)) {
	defer qt.Recovering("connect QFileDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QFileDevice) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QFileDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQFileDeviceTimerEvent
func callbackQFileDeviceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*QTimerEvent))(NewQTimerEventFromPointer(event))
	} else {
		NewQFileDeviceFromPointer(ptr).TimerEventDefault(NewQTimerEventFromPointer(event))
	}
}

func (ptr *QFileDevice) TimerEvent(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_TimerEvent(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileDevice) TimerEventDefault(event QTimerEvent_ITF) {
	defer qt.Recovering("QFileDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_TimerEventDefault(ptr.Pointer(), PointerFromQTimerEvent(event))
	}
}

func (ptr *QFileDevice) ConnectChildEvent(f func(event *QChildEvent)) {
	defer qt.Recovering("connect QFileDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QFileDevice) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QFileDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQFileDeviceChildEvent
func callbackQFileDeviceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*QChildEvent))(NewQChildEventFromPointer(event))
	} else {
		NewQFileDeviceFromPointer(ptr).ChildEventDefault(NewQChildEventFromPointer(event))
	}
}

func (ptr *QFileDevice) ChildEvent(event QChildEvent_ITF) {
	defer qt.Recovering("QFileDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_ChildEvent(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileDevice) ChildEventDefault(event QChildEvent_ITF) {
	defer qt.Recovering("QFileDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_ChildEventDefault(ptr.Pointer(), PointerFromQChildEvent(event))
	}
}

func (ptr *QFileDevice) ConnectCustomEvent(f func(event *QEvent)) {
	defer qt.Recovering("connect QFileDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QFileDevice) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QFileDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQFileDeviceCustomEvent
func callbackQFileDeviceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QFileDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*QEvent))(NewQEventFromPointer(event))
	} else {
		NewQFileDeviceFromPointer(ptr).CustomEventDefault(NewQEventFromPointer(event))
	}
}

func (ptr *QFileDevice) CustomEvent(event QEvent_ITF) {
	defer qt.Recovering("QFileDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_CustomEvent(ptr.Pointer(), PointerFromQEvent(event))
	}
}

func (ptr *QFileDevice) CustomEventDefault(event QEvent_ITF) {
	defer qt.Recovering("QFileDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QFileDevice_CustomEventDefault(ptr.Pointer(), PointerFromQEvent(event))
	}
}
