// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package advertisement

import (
	"syscall"
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/KayFelicities/winrt-go/windows/foundation/collections"
)

const SignatureBluetoothLEAdvertisement string = "rc(Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisement;{066fb2b7-33d1-4e7d-8367-cf81d0f79653})"

type BluetoothLEAdvertisement struct {
	ole.IUnknown
}

func NewBluetoothLEAdvertisement() (*BluetoothLEAdvertisement, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Devices.Bluetooth.Advertisement.BluetoothLEAdvertisement")
	if err != nil {
		return nil, err
	}
	return (*BluetoothLEAdvertisement)(unsafe.Pointer(inspectable)), nil
}

func (impl *BluetoothLEAdvertisement) GetLocalName() (string, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisement))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisement)(unsafe.Pointer(itf))
	return v.GetLocalName()
}

func (impl *BluetoothLEAdvertisement) SetLocalName(value string) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisement))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisement)(unsafe.Pointer(itf))
	return v.SetLocalName(value)
}

func (impl *BluetoothLEAdvertisement) GetServiceUuids() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisement))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisement)(unsafe.Pointer(itf))
	return v.GetServiceUuids()
}

func (impl *BluetoothLEAdvertisement) GetManufacturerData() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisement))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisement)(unsafe.Pointer(itf))
	return v.GetManufacturerData()
}

func (impl *BluetoothLEAdvertisement) GetDataSections() (*collections.IVector, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDiBluetoothLEAdvertisement))
	defer itf.Release()
	v := (*iBluetoothLEAdvertisement)(unsafe.Pointer(itf))
	return v.GetDataSections()
}

const GUIDiBluetoothLEAdvertisement string = "066fb2b7-33d1-4e7d-8367-cf81d0f79653"
const SignatureiBluetoothLEAdvertisement string = "{066fb2b7-33d1-4e7d-8367-cf81d0f79653}"

type iBluetoothLEAdvertisement struct {
	ole.IInspectable
}

type iBluetoothLEAdvertisementVtbl struct {
	ole.IInspectableVtbl

	GetFlags                       uintptr
	SetFlags                       uintptr
	GetLocalName                   uintptr
	SetLocalName                   uintptr
	GetServiceUuids                uintptr
	GetManufacturerData            uintptr
	GetDataSections                uintptr
	GetManufacturerDataByCompanyId uintptr
	GetSectionsByType              uintptr
}

func (v *iBluetoothLEAdvertisement) VTable() *iBluetoothLEAdvertisementVtbl {
	return (*iBluetoothLEAdvertisementVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *iBluetoothLEAdvertisement) GetLocalName() (string, error) {
	var outHStr ole.HString
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetLocalName,
		uintptr(unsafe.Pointer(v)),        // this
		uintptr(unsafe.Pointer(&outHStr)), // out string
	)

	if hr != 0 {
		return "", ole.NewError(hr)
	}

	out := outHStr.String()
	ole.DeleteHString(outHStr)
	return out, nil
}

func (v *iBluetoothLEAdvertisement) SetLocalName(value string) error {
	valueHStr, err := ole.NewHString(value)
	if err != nil {
		return err
	}
	hr, _, _ := syscall.SyscallN(
		v.VTable().SetLocalName,
		uintptr(unsafe.Pointer(v)), // this
		uintptr(valueHStr),         // in string
	)

	if hr != 0 {
		return ole.NewError(hr)
	}

	return nil
}

func (v *iBluetoothLEAdvertisement) GetServiceUuids() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetServiceUuids,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEAdvertisement) GetManufacturerData() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetManufacturerData,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}

func (v *iBluetoothLEAdvertisement) GetDataSections() (*collections.IVector, error) {
	var out *collections.IVector
	hr, _, _ := syscall.SyscallN(
		v.VTable().GetDataSections,
		uintptr(unsafe.Pointer(v)),    // this
		uintptr(unsafe.Pointer(&out)), // out collections.IVector
	)

	if hr != 0 {
		return nil, ole.NewError(hr)
	}

	return out, nil
}
