// Code generated by winrt-go-gen. DO NOT EDIT.

//go:build windows

//nolint:all
package streams

import (
	"unsafe"

	"github.com/go-ole/go-ole"
	"github.com/KayFelicities/winrt-go/windows/foundation"
)

const SignatureDataWriter string = "rc(Windows.Storage.Streams.DataWriter;{64b89265-d341-4922-b38a-dd4af8808c4e})"

type DataWriter struct {
	ole.IUnknown
}

func NewDataWriter() (*DataWriter, error) {
	inspectable, err := ole.RoActivateInstance("Windows.Storage.Streams.DataWriter")
	if err != nil {
		return nil, err
	}
	return (*DataWriter)(unsafe.Pointer(inspectable)), nil
}

func (impl *DataWriter) WriteBytes(valueSize uint32, value []uint8) error {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIDataWriter))
	defer itf.Release()
	v := (*IDataWriter)(unsafe.Pointer(itf))
	return v.WriteBytes(valueSize, value)
}

func (impl *DataWriter) DetachBuffer() (*IBuffer, error) {
	itf := impl.MustQueryInterface(ole.NewGUID(GUIDIDataWriter))
	defer itf.Release()
	v := (*IDataWriter)(unsafe.Pointer(itf))
	return v.DetachBuffer()
}

func (impl *DataWriter) Close() error {
	itf := impl.MustQueryInterface(ole.NewGUID(foundation.GUIDIClosable))
	defer itf.Release()
	v := (*foundation.IClosable)(unsafe.Pointer(itf))
	return v.Close()
}
