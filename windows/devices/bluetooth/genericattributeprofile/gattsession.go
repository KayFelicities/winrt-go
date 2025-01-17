// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package genericattributeprofile

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/KayFelicities/winrt-go/windows/devices/bluetooth"
	"github.com/KayFelicities/winrt-go/windows/foundation"
)

const SignatureGattSession string = "rc(Windows.Devices.Bluetooth.GenericAttributeProfile.GattSession;{d23b5143-e04e-4c24-999c-9c256f9856b1})"

type GattSession struct {
	ole.IUnknown
}

func (impl *GattSession) GetCanMaintainConnection() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.GetCanMaintainConnection()
}

func (impl *GattSession) SetMaintainConnection(value bool) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.SetMaintainConnection(value)
}

func (impl *GattSession) GetMaintainConnection() (bool, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.GetMaintainConnection()
}

func (impl *GattSession) GetMaxPduSize() (uint16, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.GetMaxPduSize()
}

func (impl *GattSession) AddMaxPduSizeChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.AddMaxPduSizeChanged(handler)
}

func (impl *GattSession) RemoveMaxPduSizeChanged(token foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattSession))
	defer itf.Release()
	v := (*iGattSession)(unsafe.Pointer(itf))
	return v.RemoveMaxPduSizeChanged(token)
}

func (impl *GattSession) Close() error {
	itf := impl.MustQueryInterface(ole.NewGUID(foundation.GUIDIClosable))
	defer itf.Release()
	v := (*foundation.IClosable)(unsafe.Pointer(itf))
	return v.Close()
}

const GUIDiGattSession string = "d23b5143-e04e-4c24-999c-9c256f9856b1"
const SignatureiGattSession string = "{d23b5143-e04e-4c24-999c-9c256f9856b1}"

type iGattSession struct {
	ole.IInspectable
}

type iGattSessionVtbl struct {
	ole.IInspectableVtbl

	GetDeviceId                uintptr
	GetCanMaintainConnection   uintptr
	SetMaintainConnection      uintptr
	GetMaintainConnection      uintptr
	GetMaxPduSize              uintptr
	GetSessionStatus           uintptr
	AddMaxPduSizeChanged       uintptr
	RemoveMaxPduSizeChanged    uintptr
	AddSessionStatusChanged    uintptr
	RemoveSessionStatusChanged uintptr
}

func (v *iGattSession) VTable() *iGattSessionVtbl {
	return (*iGattSessionVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattSession) GetCanMaintainConnection() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCanMaintainConnection,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSession) SetMaintainConnection(value bool) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetMaintainConnection,
		uintptr(unsafe.Pointer(v)),                // this
		uintptr(*(*byte)(unsafe.Pointer(&value))), // in bool
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iGattSession) GetMaintainConnection() (bool, error) {
	var out bool
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMaintainConnection,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out bool
	)

	if hr != 0 {
		return false, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSession) GetMaxPduSize() (uint16, error) {
	var out uint16
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetMaxPduSize,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out uint16
	)

	if hr != 0 {
		return 0, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSession) AddMaxPduSizeChanged(handler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddMaxPduSizeChanged,
		uintptr(unsafe.Pointer(v)),       // this
		uintptr(unsafe.Pointer(handler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),    // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattSession) RemoveMaxPduSizeChanged(token foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveMaxPduSizeChanged,
		uintptr(unsafe.Pointer(v)),      // this
		uintptr(unsafe.Pointer(&token)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiGattSessionStatics string = "2e65b95c-539f-4db7-82a8-73bdbbf73ebf"
const SignatureiGattSessionStatics string = "{2e65b95c-539f-4db7-82a8-73bdbbf73ebf}"

type iGattSessionStatics struct {
	ole.IInspectable
}

type iGattSessionStaticsVtbl struct {
	ole.IInspectableVtbl

	FromDeviceIdAsync uintptr
}

func (v *iGattSessionStatics) VTable() *iGattSessionStaticsVtbl {
	return (*iGattSessionStaticsVtbl)(unsafe.Pointer(v.RawVTable))
}

func FromDeviceIdAsync(deviceId *bluetooth.BluetoothDeviceId) (*foundation.IAsyncOperation, error) {
	inspectable, err := ole.RoGetActivationFactory("Windows.Devices.Bluetooth.GenericAttributeProfile.GattSession", ole.NewGUID(GUIDiGattSessionStatics))
	if err != nil {
		return nil, err
	}
	v := (*iGattSessionStatics)(unsafe.Pointer(inspectable))

	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().FromDeviceIdAsync,
		0,                                 // this is a static func, so there's no this
		uintptr(unsafe.Pointer(deviceId)), // in bluetooth.BluetoothDeviceId
		uintptr(unsafe.Pointer(&out)),     // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
