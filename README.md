# binp
[![Go Reference](https://pkg.go.dev/static/frontend/badge/badge.svg)](https://pkg.go.dev/github.com/MrWaggel/binp)
[![Go Report Card](https://goreportcard.com/badge/github.com/MrWaggel/binp?style=flat-square)](https://goreportcard.com/report/github.com/MrWaggel/binp)

A fast performant bytes packing and unpacking module, to read and write primitive Go types from and to `[]byte`, without using the `unsafe` package.

## Supported types
`bool`, `string`, `int8`, `byte`, `rune`, `int16`, `int32`, `int64`, `int`, `uint8`, `uint16`, `uint32`, `uint64`, `uint`, `float32`, `float64`, `[]bool`, `[]string`, `[]int8`, `[]byte`, `[]rune`, `[]int16`, `[]int32`, `[]int64`, `[]int`, `[]uint8`, `[]uint16`, `[]uint32`, `[]uint64`, `[]uint`, `[]float32`, `[]float64`

## Limitations

- `int` and `uint` are always treated as 64bit, 8 bytes.
- Maximum total size of the input values is 4.2GB.
- Any slice `[]T` can only contain up to 65536 elements.
- Except for `string` and `[]byte` which have a maximum length of 4.2GB.
- `complex64` and `complex128` are not supported.
- `nil` slices are treated as slices with length 0.
- Reading with `Read()` to a targeted `[]T`, will allocate the underlying slice automatically with `make([]T, ...)`.

## Usage

Install 
```
go get github.com/MrWaggel/binp
```

```
import "github.com/MrWaggel/binp"
```

```go
s := "abc"
i := int32(500)
b, _ := binp.Pack(s, i)

fmt.Println(b) 
// [28 0 0 0 3 97 98 99 5 0 0 1 244]

var targetString string
var targetInt int32
_ = binp.Read(b, &targetString, &targetInt)

fmt.Println(targetString, targetInt) 
// abc 500
```

### Over network using the `net.Conn` interface

Sender
```go
var conn net.Conn

s := "abc"
i := int32(500)

b, _ := binp.PackNetwork(s, i)
conn.Write(b)	
```

Receiver
```go
var conn net.Conn

readBufLen := make([]byte, 4)
conn.Read(readBufLen)

readLen := int(binary.BigEndian.Uint32(readBufLen))

b := make([]byte, readLen)
conn.Read(b)

var targetString string
var targetInt int32
binp.Read(b, &targetString, &targetInt)
```

## Benchmarks

```go test -benchmem -bench .```

```
goos: linux
goarch: amd64
pkg: github.com/mrwaggel/binp
cpu: Intel(R) Core(TM) i5-2500K CPU @ 3.30GHz
BenchmarkPackInt-4                      20914693                50.53 ns/op           16 B/op          1 allocs/op
BenchmarkPackToInt-4                    188655639                6.323 ns/op           0 B/op          0 allocs/op
BenchmarkPackMany-4                     11858296               109.9 ns/op            32 B/op          1 allocs/op
BenchmarkPacToMany-4                    188766630                6.361 ns/op           0 B/op          0 allocs/op
BenchmarkPackFixedInt-4                 21688969                47.28 ns/op           16 B/op          1 allocs/op
BenchmarkPackFixedMany-4                14650188                83.23 ns/op           32 B/op          1 allocs/op
BenchmarkPackNetworkInt-4               21861925                53.49 ns/op           16 B/op          1 allocs/op
BenchmarkPackNetworkMany-4              11278044               109.6 ns/op            32 B/op          1 allocs/op
BenchmarkPackNetworkFixedInt-4          25842140                50.84 ns/op           16 B/op          1 allocs/op
BenchmarkPackNetworkFixedMany-4         12346747                83.64 ns/op           32 B/op          1 allocs/op
BenchmarkReadInt-4                      161750701                7.425 ns/op           0 B/op          0 allocs/op
BenchmarkReadMany-4                     51132878                21.19 ns/op            0 B/op          0 allocs/op
BenchmarkSizeFunctionInt-4              173734040                6.893 ns/op           0 B/op          0 allocs/op
BenchmarkSizeFunctionMany-4             60638113                19.76 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/mrwaggel/binp        20.854s
```

One allocation is made per `Pack(...)` since it has to call `make([]byte, ...)`. 

## Tests

```go test -v```

```
=== RUN   TestTypeInt
--- PASS: TestTypeInt (0.00s)
=== RUN   TestTypeInt64
--- PASS: TestTypeInt64 (0.00s)
=== RUN   TestTypeInt32
--- PASS: TestTypeInt32 (0.00s)
=== RUN   TestTypeInt16
--- PASS: TestTypeInt16 (0.00s)
=== RUN   TestTypeInt8
--- PASS: TestTypeInt8 (0.00s)
=== RUN   TestTypeUInt
--- PASS: TestTypeUInt (0.00s)
=== RUN   TestTypeUInt64
--- PASS: TestTypeUInt64 (0.00s)
=== RUN   TestTypeUInt32
--- PASS: TestTypeUInt32 (0.00s)
=== RUN   TestTypeUInt16
--- PASS: TestTypeUInt16 (0.00s)
=== RUN   TestTypeUInt8
--- PASS: TestTypeUInt8 (0.00s)
=== RUN   TestTypeByte
--- PASS: TestTypeByte (0.00s)
=== RUN   TestTypeSliceInt
--- PASS: TestTypeSliceInt (0.00s)
=== RUN   TestTypeSliceInt64
--- PASS: TestTypeSliceInt64 (0.00s)
=== RUN   TestTypeSliceInt32
--- PASS: TestTypeSliceInt32 (0.00s)
=== RUN   TestTypeSliceInt16
--- PASS: TestTypeSliceInt16 (0.00s)
=== RUN   TestTypeSliceInt8
--- PASS: TestTypeSliceInt8 (0.00s)
=== RUN   TestTypeSliceUInt
--- PASS: TestTypeSliceUInt (0.00s)
=== RUN   TestTypeSliceUInt64
--- PASS: TestTypeSliceUInt64 (0.00s)
=== RUN   TestTypeSliceUInt32
--- PASS: TestTypeSliceUInt32 (0.00s)
=== RUN   TestTypeSliceUInt16
--- PASS: TestTypeSliceUInt16 (0.00s)
=== RUN   TestTypeSliceUInt8
--- PASS: TestTypeSliceUInt8 (0.00s)
=== RUN   TestTypeFloat64
--- PASS: TestTypeFloat64 (0.00s)
=== RUN   TestTypeFloat32
--- PASS: TestTypeFloat32 (0.00s)
=== RUN   TestTypeIntSliceFloat64
--- PASS: TestTypeIntSliceFloat64 (0.00s)
=== RUN   TestTypeSliceFloat32
--- PASS: TestTypeSliceFloat32 (0.00s)
=== RUN   TestTypeBoolTrue
--- PASS: TestTypeBoolTrue (0.00s)
=== RUN   TestTypeBoolFalse
--- PASS: TestTypeBoolFalse (0.00s)
=== RUN   TestTypeString
--- PASS: TestTypeString (0.00s)
=== RUN   TestTypeSliceString
--- PASS: TestTypeSliceString (0.00s)
=== RUN   TestMultipleInputs
--- PASS: TestMultipleInputs (0.00s)
=== RUN   TestMultipleInputsAll
--- PASS: TestMultipleInputsAll (0.00s)
PASS
ok      github.com/mrwaggel/binp        0.002s
```

Gilles Van Vlasselaer https://mrwaggel.be