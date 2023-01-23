package binp

import (
	"bytes"
	"errors"
	"testing"
)

var ErrCompared = errors.New("input output comparison mismatched")

func TestTypeInt(t *testing.T) {
	err := genericCall(int(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeInt64(t *testing.T) {
	err := genericCall(int64(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeInt32(t *testing.T) {
	err := genericCall(int32(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeInt16(t *testing.T) {
	err := genericCall(int16(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeInt8(t *testing.T) {
	err := genericCall(int8(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeUInt(t *testing.T) {
	err := genericCall(uint(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeUInt64(t *testing.T) {
	err := genericCall(uint64(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeUInt32(t *testing.T) {
	err := genericCall(uint32(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeUInt16(t *testing.T) {
	err := genericCall(uint16(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeUInt8(t *testing.T) {
	err := genericCall(uint8(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeByte(t *testing.T) {
	err := genericCall(byte(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceInt(t *testing.T) {
	err := genericCall([]int{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceInt64(t *testing.T) {
	err := genericCall([]int64{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceInt32(t *testing.T) {
	err := genericCall([]int32{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceInt16(t *testing.T) {
	err := genericCall([]int16{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceInt8(t *testing.T) {
	err := genericCall([]int8{10, 50, -128})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceUInt(t *testing.T) {
	err := genericCall([]uint{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceUInt64(t *testing.T) {
	err := genericCall([]uint64{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceUInt32(t *testing.T) {
	err := genericCall([]uint32{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceUInt16(t *testing.T) {
	err := genericCall([]uint16{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceUInt8(t *testing.T) {
	err := genericCall([]uint8{10, 50, 200})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeFloat64(t *testing.T) {
	err := genericCall(float64(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeFloat32(t *testing.T) {
	err := genericCall(float32(99))
	if err != nil {
		t.Error(err)
	}
}

func TestTypeIntSliceFloat64(t *testing.T) {
	err := genericCall([]float64{99, 10.3, -5.5})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceFloat32(t *testing.T) {
	err := genericCall([]float32{99, 10.3, -5.5})
	if err != nil {
		t.Error(err)
	}
}

func TestTypeBoolTrue(t *testing.T) {
	err := genericCall(true)
	if err != nil {
		t.Error(err)
	}
}

func TestTypeBoolFalse(t *testing.T) {
	err := genericCall(false)
	if err != nil {
		t.Error(err)
	}
}

func TestTypeString(t *testing.T) {
	err := genericCall("qwerty123")
	if err != nil {
		t.Error(err)
	}
}

func TestTypeSliceString(t *testing.T) {
	err := genericCall([]string{"abc", "xyz", "Special chars 漢字", "123", "azertyuiopmlkjhgfdsqwxcvbn \n \t!"})
	if err != nil {
		t.Error(err)
	}
}

type testData1 struct {
	Bool     bool
	String   string
	Float64  float64
	Bytes    []byte
	Booleans []bool
	Ints16   []int16
	Uint32   uint32
}

func TestMultipleInputs(t *testing.T) {
	var unpacked testData1
	out := testData1{
		Bool:     true,
		String:   "abc-xyz-漢字",
		Float64:  -3.5,
		Bytes:    []byte("Some bytes"),
		Booleans: []bool{true, true, false, true, false, false, false, true},
		Ints16:   []int16{1, 2, 3, 5, 120},
		Uint32:   2500,
	}

	outbytes, err := Pack(out.Bool, out.String, out.Bytes, out.Booleans, out.Ints16, out.Uint32)
	if err != nil {
		t.Error(err)
		return
	}

	err = Read(outbytes, &unpacked.Bool, &unpacked.String, &unpacked.Bytes, &unpacked.Booleans, &unpacked.Ints16, &unpacked.Uint32)
	if err != nil {
		t.Error(err)
		return
	}

	comparebytes, err := Pack(unpacked.Bool, unpacked.String, unpacked.Bytes, unpacked.Booleans, unpacked.Ints16, unpacked.Uint32)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(outbytes, comparebytes) {
		t.Error(ErrCompared)
	}
}

type testData2 struct {
	Int    int
	Int64  int64
	Int32  int32
	Int16  int16
	Int8   int8
	Uint   uint
	Uint64 uint64
	Uint32 uint32
	Uint16 uint16
	Uint8  uint8

	Ints    []int
	Ints64  []int64
	Ints32  []int32
	Ints16  []int16
	Ints8   []int8
	Uints   []uint
	Uints64 []uint64
	Uints32 []uint32
	Uints16 []uint16
	Uints8  []uint8

	Byte  byte
	Bytes []byte

	Float64  float64
	Floats64 []float64
	Float32  float32
	FLoats32 []float32

	BoolTrue  bool
	BoolFalse bool
	Booleans  []bool

	String  string
	Strings []string

	LastString string
	LastBytes  []byte
	LastInt    int
}

func TestMultipleInputsAll(t *testing.T) {
	var unpacked testData2
	out := testData2{
		Int:        500,
		Int64:      89745552,
		Int32:      -8500,
		Int16:      500,
		Int8:       25,
		Uint:       15842156,
		Uint64:     9151582,
		Uint32:     1518,
		Uint16:     35000,
		Uint8:      250,
		Ints:       []int{1, 2, 3, 5, 8, 10},
		Ints64:     []int64{1, 5, 8, 5, 8, 1, -3, 8, 7, 5, 6, 3},
		Ints32:     []int32{12, 55, 6, 77, 268},
		Ints16:     []int16{1, 2, 3, 4, 5, 6},
		Ints8:      []int8{1, 5, 8, 7, 1, 5, 6},
		Uints:      []uint{555, 123, 456, 98},
		Uints64:    []uint64{123, 588, 569, 54, 123, 6},
		Uints32:    []uint32{5489, 78979, 42, 6465, 879, 123},
		Uints16:    []uint16{1, 2, 3, 4, 5, 6, 78, 89},
		Uints8:     []uint8{1, 2, 5, 2, 5, 2, 5, 2, 5, 2, 5},
		Byte:       101,
		Bytes:      []byte{101, 102, 50, 99, 13, 0, 15, 0},
		Float64:    -3.5,
		Floats64:   []float64{1, 2, 3, 0.993, -14.81},
		Float32:    1.23,
		FLoats32:   []float32{1, 2, 3, -14.55, 998},
		BoolTrue:   true,
		BoolFalse:  false,
		Booleans:   []bool{true, false, true, true, false, true, false, false, false, true},
		String:     "abc-xyz-漢字",
		Strings:    []string{"abc-xyz-漢字", "123", "azertyuiop + abc-xyz-漢字 \n \t abc-xyz-漢字"},
		LastString: "Last string",
		LastBytes:  []byte{1, 2, 3, 40, 50},
		LastInt:    -188000,
	}

	outpacked, err := Pack(out.Int, out.Int64, out.Int32, out.Int16, out.Int8, out.Uint, out.Uint64,
		out.Uint32, out.Uint16, out.Uint8, out.Ints, out.Ints64, out.Ints32, out.Ints16, out.Ints8, out.Uints,
		out.Uints64, out.Uints32, out.Uints16, out.Uints8, out.Byte, out.Bytes, out.Float64, out.Floats64,
		out.Float32, out.FLoats32, out.BoolTrue, out.BoolFalse, out.Booleans, out.String, out.Strings,
		out.LastString, out.LastBytes, out.LastInt)
	if err != nil {
		t.Error(err)
		return
	}

	err = Read(outpacked, &unpacked.Int, &unpacked.Int64, &unpacked.Int32, &unpacked.Int16, &unpacked.Int8, &unpacked.Uint, &unpacked.Uint64, &unpacked.Uint32, &unpacked.Uint16, &unpacked.Uint8,
		&unpacked.Ints, &unpacked.Ints64, &unpacked.Ints32, &unpacked.Ints16, &unpacked.Ints8, &unpacked.Uints, &unpacked.Uints64, &unpacked.Uints32, &unpacked.Uints16, &unpacked.Uints8,
		&unpacked.Byte, &unpacked.Bytes, &unpacked.Float64, &unpacked.Floats64, &unpacked.Float32, &unpacked.FLoats32, &unpacked.BoolTrue, &unpacked.BoolFalse, &unpacked.Booleans,
		&unpacked.String, &unpacked.Strings, &unpacked.LastString, &unpacked.LastBytes, &unpacked.LastInt)

	comparebytes, err := Pack(unpacked.Int, unpacked.Int64, unpacked.Int32, unpacked.Int16, unpacked.Int8, unpacked.Uint, unpacked.Uint64, unpacked.Uint32, unpacked.Uint16, unpacked.Uint8, unpacked.Ints, unpacked.Ints64, unpacked.Ints32, unpacked.Ints16, unpacked.Ints8, unpacked.Uints, unpacked.Uints64, unpacked.Uints32, unpacked.Uints16, unpacked.Uints8,
		unpacked.Byte, unpacked.Bytes, unpacked.Float64, unpacked.Floats64, unpacked.Float32, unpacked.FLoats32, unpacked.BoolTrue, unpacked.BoolFalse, unpacked.Booleans,
		unpacked.String, unpacked.Strings, unpacked.LastString, unpacked.LastBytes, unpacked.LastInt)
	if err != nil {
		t.Error(err)
		return
	}

	if !bytes.Equal(outpacked, comparebytes) {
		t.Error(ErrCompared)
		return
	}

	if out.LastString != unpacked.LastString {
		t.Error(ErrCompared)
		return
	}

	if out.String != unpacked.String {
		t.Error(ErrCompared)
		return
	}

	if out.Int32 != unpacked.Int32 {
		t.Error(ErrCompared)
		return
	}

	if out.LastInt != unpacked.LastInt {
		t.Error(ErrCompared)
		return
	}
	return

}

func genericCall[T any](val T) error {
	var final T
	packed, err := Pack(val)
	if err != nil {
		return err
	}

	err = Read(packed, &final)
	if err != nil {
		return err
	}

	packed2, err := Pack(final)
	if err != nil {
		return err
	}

	if !bytes.Equal(packed, packed2) {
		return ErrCompared
	}
	return nil
}
