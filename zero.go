// Package zero provides functions that ensure a variable is completely unreadable in memory.
package zero

import (
	"reflect"
	"unsafe"
)

// Zeros out all data making the variable unreadable. Useful if you need to eradicate a variable from memory.
func Zero(data interface{}) {
	PointerZero(data)
}

// Slower than PointerZero, might be safer.
func ReflectZero(data interface{}) {
	v := reflect.ValueOf(data).Elem()
	zeroValue := reflect.Zero(v.Type())
	v.Set(zeroValue)
}

// Faster than ReflectZero, might be more unstable.
func PointerZero(data interface{}) {
	v := reflect.ValueOf(data).Elem()
	size := v.Type().Size()
	ptr := unsafe.Pointer(uintptr(v.Addr().UnsafePointer()))

	for i := uintptr(0); i < size; i++ {
		*(*byte)(unsafe.Pointer(uintptr(ptr) + i)) = 0
	}
}
