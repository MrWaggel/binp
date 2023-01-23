package binp

import (
	"encoding/binary"
	"math"
)

func write(b []byte, values []interface{}) error {
	var lenBuf, offset int
	for _, v := range values {
		switch v.(type) {
		case int:
			b[offset] = Int
			offset++
			binary.BigEndian.PutUint64(b[offset:offset+8], uint64(v.(int)))
			offset += 8
		case []int:
			b[offset] = IntSlice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]int))))
			offset += 2
			for _, vv := range v.([]int) {
				binary.BigEndian.PutUint64(b[offset:offset+8], uint64(vv))
				offset += 8
			}
		case int64:
			b[offset] = Int64
			offset++
			binary.BigEndian.PutUint64(b[offset:offset+8], uint64(v.(int64)))
			offset += 8
		case []int64:
			b[offset] = Int64Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]int64))))
			offset += 2
			for _, vv := range v.([]int64) {
				binary.BigEndian.PutUint64(b[offset:offset+8], uint64(vv))
				offset += 8
			}
		case int32:
			b[offset] = Int32
			offset++
			binary.BigEndian.PutUint32(b[offset:offset+4], uint32(v.(int32)))
			offset += 4
		case []int32:
			b[offset] = Int32Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]int32))))
			offset += 2
			for _, vv := range v.([]int32) {
				binary.BigEndian.PutUint32(b[offset:offset+4], uint32(vv))
				offset += 4
			}
		case int16:
			b[offset] = Int16
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(v.(int16)))
			offset += 2
		case []int16:
			b[offset] = Int16Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]int16))))
			offset += 2
			for _, vv := range v.([]int16) {
				binary.BigEndian.PutUint16(b[offset:offset+2], uint16(vv))
				offset += 2
			}
		case int8:
			b[offset] = Int8
			offset++
			b[offset] = byte(v.(int8))
			offset++
		case []int8:
			b[offset] = Int8Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]int8))))
			offset += 2
			for _, vv := range v.([]int8) {
				b[offset] = byte(vv)
				offset++
			}
		case uint:
			b[offset] = Uint
			offset++
			binary.BigEndian.PutUint64(b[offset:offset+8], uint64(v.(uint)))
			offset += 8
		case []uint:
			b[offset] = UintSlice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]uint))))
			offset += 2
			for _, vv := range v.([]uint) {
				binary.BigEndian.PutUint64(b[offset:offset+8], uint64(vv))
				offset += 8
			}
		case uint64:
			b[offset] = Uint64
			offset++
			binary.BigEndian.PutUint64(b[offset:offset+8], v.(uint64))
			offset += 8
		case []uint64:
			b[offset] = Uint64Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]uint64))))
			offset += 2
			for _, vv := range v.([]uint64) {
				binary.BigEndian.PutUint64(b[offset:offset+8], vv)
				offset += 8
			}
		case uint32:
			b[offset] = Uint32
			offset++
			binary.BigEndian.PutUint32(b[offset:offset+4], v.(uint32))
			offset += 4
		case []uint32:
			b[offset] = Uint32Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]uint32))))
			offset += 2
			for _, vv := range v.([]uint32) {
				binary.BigEndian.PutUint32(b[offset:offset+4], vv)
				offset += 4
			}
		case uint16:
			b[offset] = Uint16
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], v.(uint16))
			offset += 2
		case []uint16:
			b[offset] = Uint16Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]uint16))))
			offset += 2
			for _, vv := range v.([]uint16) {
				binary.BigEndian.PutUint16(b[offset:offset+2], vv)
				offset += 2
			}
		case byte:
			b[offset] = Byte
			offset++
			b[offset] = v.(byte)
			offset++
		case []byte:
			b[offset] = ByteSlice
			offset++
			lenBuf = len(v.([]byte))
			binary.BigEndian.PutUint32(b[offset:offset+4], uint32(lenBuf))
			offset += 4
			copy(b[offset:offset+lenBuf], v.([]byte))
			offset += lenBuf
		case bool:
			if v.(bool) {
				b[offset] = BoolTrue
			} else {
				b[offset] = BoolFalse
			}
			offset++
		case []bool:
			b[offset] = BoolSlice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]bool))))
			offset += 2
			for _, vv := range v.([]bool) {
				if vv {
					b[offset] = BoolTrue
				} else {
					b[offset] = BoolFalse
				}
				offset++
			}
		case float64:
			b[offset] = Float64
			offset++
			binary.BigEndian.PutUint64(b[offset:offset+8], math.Float64bits(v.(float64)))
			offset += 8
		case []float64:
			b[offset] = Float64Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]float64))))
			offset += 2
			for _, vv := range v.([]float64) {
				binary.BigEndian.PutUint64(b[offset:offset+8], math.Float64bits(vv))
				offset += 8
			}
		case float32:
			b[offset] = Float32
			offset++
			binary.BigEndian.PutUint32(b[offset:offset+4], math.Float32bits(v.(float32)))
			offset += 4
		case []float32:
			b[offset] = Float32Slice
			offset++
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(len(v.([]float32))))
			offset += 2
			for _, vv := range v.([]float32) {
				binary.BigEndian.PutUint32(b[offset:offset+4], math.Float32bits(vv))
				offset += 4
			}
		case string:
			b[offset] = String
			offset++
			lenBuf = len(v.(string))
			binary.BigEndian.PutUint32(b[offset:offset+4], uint32(lenBuf))
			offset += 4
			copy(b[offset:offset+lenBuf], v.(string))
			offset += lenBuf
		case []string:
			b[offset] = StringSlice
			offset++
			lenBuf = len(v.([]string))
			binary.BigEndian.PutUint16(b[offset:offset+2], uint16(lenBuf))
			offset += 2
			for _, vv := range v.([]string) {
				lenBuf = len(vv)
				binary.BigEndian.PutUint32(b[offset:offset+4], uint32(lenBuf))
				offset += 4
				copy(b[offset:offset+lenBuf], vv)
				offset += lenBuf
			}
		default:
			return ErrUnsupportedType
		}
	}
	return nil
}
