package zero

import (
	"reflect"
	"testing"
	"unsafe"
)

var testCases = []interface{}{
	true,
	"string",
	int(1),
	int8(1),
	int16(1),
	int32(1),
	int64(1),
	uint(1),
	uint8(1),
	uint16(1),
	uint32(1),
	uint64(1),
	uintptr(1),
	float32(1),
	float64(1),
	complex64(1i),
	complex128(1i),
	[]string{"a", "slice"},
	[]any{"a", "more", complex128(532 * 32i), "slice"},
	struct {
		Name  string
		Email string
	}{
		Name:  "John Doe",
		Email: "noreply@example.com",
	},
}

func TestReflectZero(t *testing.T) {
	for _, tc := range testCases {
		ReflectZero(&tc)
		mem := memory(&tc)
		for _, b := range mem {
			if b != 0 {
				t.Errorf("ReflectZero(%t) -> %s", tc, mem)
			}
		}
	}
}

func TestPointerZero(t *testing.T) {
	for _, tc := range testCases {
		PointerZero(&tc)
		mem := memory(&tc)
		for _, b := range mem {
			if b != 0 {
				t.Errorf("PointerZero(%t) -> %s", tc, mem)
			}
		}
	}
}

func TestPointerRand(t *testing.T) {
	for _, tc := range testCases {
		og := tc
		PointerRand(&tc)
		mem := memory(&tc)
		if og == tc {
			t.Errorf("PointerRand(%t) -> %s", tc, mem)
		}
	}
}

func memory(data interface{}) []byte {
	v := reflect.ValueOf(data).Elem()
	ptr := unsafe.Pointer(uintptr(v.Addr().UnsafePointer()))
	size := v.Type().Size()

	return (*[1 << 30]byte)(ptr)[:size:size]
}
