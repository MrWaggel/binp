package binp

import "testing"

var valInt = int(123456789)
var valFloat64 = float64(-1.5)
var valBool = false
var valByte = byte(64)
var valUint32 = uint32(999)

var targetInt int
var targetFloat64 float64
var targetBool bool
var targetByte byte
var targetUint32 uint32

func BenchmarkPackInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pack(valInt)
	}
}

func BenchmarkPackToInt(b *testing.B) {
	var i int
	vs := []interface{}{valInt}
	bs := make([]byte, 9)
	b.ResetTimer()
	for ; i < b.N; i++ {
		PackTo(bs, vs)
	}
}

func BenchmarkPackMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pack(valInt, valFloat64, valBool, valByte, valUint32)
	}
}

func BenchmarkPacToMany(b *testing.B) {
	var i int
	vs := []interface{}{valInt, valFloat64, valBool, valByte, valUint32}
	bs := make([]byte, 26)
	b.ResetTimer()
	for ; i < b.N; i++ {
		PackTo(bs, vs)
	}
}

func BenchmarkPackFixedInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackFixed(9, valInt)
	}
}
func BenchmarkPackFixedMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackFixed(26, valInt, valFloat64, valBool, valByte, valUint32)
	}
}

func BenchmarkPackNetworkInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackNetwork(valInt)
	}
}

func BenchmarkPackNetworkMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackNetwork(valInt, valFloat64, valBool, valByte, valUint32)
	}
}

func BenchmarkPackNetworkFixedInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackNetworkFixed(9, valInt)
	}
}

func BenchmarkPackNetworkFixedMany(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PackNetworkFixed(26, valInt, valFloat64, valBool, valByte, valUint32)
	}
}

func BenchmarkReadInt(b *testing.B) {
	d, _ := Pack(valInt)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(d, &targetInt)
	}
}

func BenchmarkReadMany(b *testing.B) {
	d, _ := Pack(valInt, valFloat64, valBool, valByte, valUint32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Read(d, &targetInt, &targetFloat64, &targetBool, &targetByte, &targetUint32)
	}
}

func BenchmarkSizeFunctionInt(b *testing.B) {
	v := []interface{}{valInt}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Size(v)
	}
}

func BenchmarkSizeFunctionMany(b *testing.B) {
	v := []interface{}{valInt, valFloat64, valBool, valByte, valUint32}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Size(v)
	}
}
