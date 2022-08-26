// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint
package genericattributeprofile

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/saltosystems/winrt-go/windows/devices/bluetooth"
	"github.com/saltosystems/winrt-go/windows/foundation"
	"github.com/saltosystems/winrt-go/windows/storage/streams"
)

const SignatureGattCharacteristic string = "rc(Windows.Devices.Bluetooth.GenericAttributeProfile.GattCharacteristic;{59cb50c1-5934-4f68-a198-eb864fa44e6b})"

type GattCharacteristic struct {
	ole.IUnknown
}

func (impl *GattCharacteristic) GetCharacteristicProperties() (GattCharacteristicProperties, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.GetCharacteristicProperties()
}

func (impl *GattCharacteristic) GetUuid() (syscall.GUID, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.GetUuid()
}

func (impl *GattCharacteristic) ReadValueAsync() (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.ReadValueAsync()
}

func (impl *GattCharacteristic) ReadValueWithCacheModeAsync(cacheMode bluetooth.BluetoothCacheMode) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.ReadValueWithCacheModeAsync(cacheMode)
}

func (impl *GattCharacteristic) WriteValueAsync(value *streams.IBuffer) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.WriteValueAsync(value)
}

func (impl *GattCharacteristic) WriteValueWithOptionAsync(value *streams.IBuffer, writeOption GattWriteOption) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.WriteValueWithOptionAsync(value, writeOption)
}

func (impl *GattCharacteristic) WriteClientCharacteristicConfigurationDescriptorAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) (*foundation.IAsyncOperation, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.WriteClientCharacteristicConfigurationDescriptorAsync(clientCharacteristicConfigurationDescriptorValue)
}

func (impl *GattCharacteristic) AddValueChanged(valueChangedHandler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.AddValueChanged(valueChangedHandler)
}

func (impl *GattCharacteristic) RemoveValueChanged(valueChangedEventCookie foundation.EventRegistrationToken) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiGattCharacteristic))
	defer itf.Release()
	v := (*iGattCharacteristic)(unsafe.Pointer(itf))
	return v.RemoveValueChanged(valueChangedEventCookie)
}

const GUIDiGattCharacteristic string = "59cb50c1-5934-4f68-a198-eb864fa44e6b"
const SignatureiGattCharacteristic string = "{59cb50c1-5934-4f68-a198-eb864fa44e6b}"

type iGattCharacteristic struct {
	ole.IInspectable
}

type iGattCharacteristicVtbl struct {
	ole.IInspectableVtbl

	GetDescriptors                                        uintptr
	GetCharacteristicProperties                           uintptr
	GetProtectionLevel                                    uintptr
	SetProtectionLevel                                    uintptr
	GetUserDescription                                    uintptr
	GetUuid                                               uintptr
	GetAttributeHandle                                    uintptr
	GetPresentationFormats                                uintptr
	ReadValueAsync                                        uintptr
	ReadValueWithCacheModeAsync                           uintptr
	WriteValueAsync                                       uintptr
	WriteValueWithOptionAsync                             uintptr
	ReadClientCharacteristicConfigurationDescriptorAsync  uintptr
	WriteClientCharacteristicConfigurationDescriptorAsync uintptr
	AddValueChanged                                       uintptr
	RemoveValueChanged                                    uintptr
}

func (v *iGattCharacteristic) VTable() *iGattCharacteristicVtbl {
	return (*iGattCharacteristicVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iGattCharacteristic) GetCharacteristicProperties() (GattCharacteristicProperties, error) {
	var out GattCharacteristicProperties
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetCharacteristicProperties,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out GattCharacteristicProperties
	)

	if hr != 0 {
		return GattCharacteristicPropertiesNone, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) GetUuid() (syscall.GUID, error) {
	var out syscall.GUID
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetUuid,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out syscall.GUID
	)

	if hr != 0 {
		return syscall.GUID{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) ReadValueAsync() (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().ReadValueAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) ReadValueWithCacheModeAsync(cacheMode bluetooth.BluetoothCacheMode) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().ReadValueWithCacheModeAsync,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(cacheMode),            // in bluetooth.BluetoothCacheMode
		uintptr(unsafe.Pointer(&out)), // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) WriteValueAsync(value *streams.IBuffer) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().WriteValueAsync,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in streams.IBuffer
		uintptr(unsafe.Pointer(&out)),  // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) WriteValueWithOptionAsync(value *streams.IBuffer, writeOption GattWriteOption) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().WriteValueWithOptionAsync,
		uintptr(unsafe.Pointer(v)),     // this
		uintptr(unsafe.Pointer(value)), // in streams.IBuffer
		uintptr(writeOption),           // in GattWriteOption
		uintptr(unsafe.Pointer(&out)),  // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) WriteClientCharacteristicConfigurationDescriptorAsync(clientCharacteristicConfigurationDescriptorValue GattClientCharacteristicConfigurationDescriptorValue) (*foundation.IAsyncOperation, error) {
	var out *foundation.IAsyncOperation
	hr, _, _ := syscall.SyscallN(
		v.VTable().WriteClientCharacteristicConfigurationDescriptorAsync,
		uintptr(unsafe.Pointer(v)),                                // this
		uintptr(clientCharacteristicConfigurationDescriptorValue), // in GattClientCharacteristicConfigurationDescriptorValue
		uintptr(unsafe.Pointer(&out)),                             // out foundation.IAsyncOperation
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) AddValueChanged(valueChangedHandler *foundation.TypedEventHandler) (foundation.EventRegistrationToken, error) {
	var out foundation.EventRegistrationToken
	hr, _, _ := syscall.SyscallN(
		v.VTable().AddValueChanged,
		uintptr(unsafe.Pointer(v)),                   // this
		uintptr(unsafe.Pointer(valueChangedHandler)), // in foundation.TypedEventHandler
		uintptr(unsafe.Pointer(&out)),                // out foundation.EventRegistrationToken
	)

	if hr != 0 {
		return foundation.EventRegistrationToken{}, ole.NewError(hr)
	}

	return out, nil
}

func (v *iGattCharacteristic) RemoveValueChanged(valueChangedEventCookie foundation.EventRegistrationToken) error {
	hr, _, _ := syscall.SyscallN(
		v.VTable().RemoveValueChanged,
		uintptr(unsafe.Pointer(v)),                        // this
		uintptr(unsafe.Pointer(&valueChangedEventCookie)), // in foundation.EventRegistrationToken
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

const GUIDiGattCharacteristic2 string = "ae1ab578-ec06-4764-b780-9835a1d35d6e"
const SignatureiGattCharacteristic2 string = "{ae1ab578-ec06-4764-b780-9835a1d35d6e}"

type iGattCharacteristic2 struct {
	ole.IInspectable
}

type iGattCharacteristic2Vtbl struct {
	ole.IInspectableVtbl

	GetService        uintptr
	GetAllDescriptors uintptr
}

func (v *iGattCharacteristic2) VTable() *iGattCharacteristic2Vtbl {
	return (*iGattCharacteristic2Vtbl)(unsafe.Pointer(v.RawVTable))
}

const GUIDiGattCharacteristic3 string = "3f3c663e-93d4-406b-b817-db81f8ed53b3"
const SignatureiGattCharacteristic3 string = "{3f3c663e-93d4-406b-b817-db81f8ed53b3}"

type iGattCharacteristic3 struct {
	ole.IInspectable
}

type iGattCharacteristic3Vtbl struct {
	ole.IInspectableVtbl

	GetDescriptorsAsync                                             uintptr
	GetDescriptorsWithCacheModeAsync                                uintptr
	GetDescriptorsForUuidAsync                                      uintptr
	GetDescriptorsForUuidWithCacheModeAsync                         uintptr
	WriteValueWithResultAsync                                       uintptr
	WriteValueWithResultAndOptionAsync                              uintptr
	WriteClientCharacteristicConfigurationDescriptorWithResultAsync uintptr
}

func (v *iGattCharacteristic3) VTable() *iGattCharacteristic3Vtbl {
	return (*iGattCharacteristic3Vtbl)(unsafe.Pointer(v.RawVTable))
}
