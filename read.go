package binp

import (
	"encoding/binary"
	"math"
)

func Read(b []byte, receivers ...interface{}) error {
	var offset, receiveIndex, sliceLen, stringLen, i, numBytes int
	numBytes = len(b)
	for offset < numBytes {
		offset++
		switch b[offset-1] {
		case Int:
			*receivers[receiveIndex].(*int) = int(binary.BigEndian.Uint64(b[offset : offset+8]))
			offset += 8
		case IntSlice:
			// Get size of slice array
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]int) = make([]int, sliceLen, sliceLen)
			// Allocate underlying slice
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]int))[i] = int(binary.BigEndian.Uint64(b[offset : offset+8]))
				offset += 8
			}
		case Int64:
			*receivers[receiveIndex].(*int64) = int64(binary.BigEndian.Uint64(b[offset : offset+8]))
			offset += 8
		case Int64Slice:
			// Get size of slice array
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]int64) = make([]int64, sliceLen, sliceLen)
			// Allocate underlying slice
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]int64))[i] = int64(binary.BigEndian.Uint64(b[offset : offset+8]))
				offset += 8
			}
		case Int32:
			*receivers[receiveIndex].(*int32) = int32(binary.BigEndian.Uint32(b[offset : offset+4]))
			offset += 4
		case Int32Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]int32) = make([]int32, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]int32))[i] = int32(binary.BigEndian.Uint32(b[offset : offset+4]))
				offset += 4
			}
		case Int16:
			*receivers[receiveIndex].(*int16) = int16(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
		case Int16Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]int16) = make([]int16, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]int16))[i] = int16(binary.BigEndian.Uint16(b[offset : offset+2]))
				offset += 2
			}
		case Int8:
			*receivers[receiveIndex].(*int8) = int8(b[offset])
			offset += 1
		case Int8Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]int8) = make([]int8, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]int8))[i] = int8(b[offset])
				offset++
			}
		case Uint:
			*receivers[receiveIndex].(*uint) = uint(binary.BigEndian.Uint64(b[offset : offset+8]))
			offset += 8
		case UintSlice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]uint) = make([]uint, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]uint))[i] = uint(binary.BigEndian.Uint64(b[offset : offset+8]))
				offset += 8
			}
		case Uint64:
			*receivers[receiveIndex].(*uint64) = binary.BigEndian.Uint64(b[offset : offset+8])
			offset += 8
		case Uint64Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]uint64) = make([]uint64, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]uint64))[i] = binary.BigEndian.Uint64(b[offset : offset+8])
				offset += 8
			}
		case Uint32:
			*receivers[receiveIndex].(*uint32) = binary.BigEndian.Uint32(b[offset : offset+4])
			offset += 4
		case Uint32Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]uint32) = make([]uint32, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]uint32))[i] = binary.BigEndian.Uint32(b[offset : offset+4])
				offset += 4
			}
		case Uint16:
			*receivers[receiveIndex].(*uint16) = binary.BigEndian.Uint16(b[offset : offset+2])
			offset += 2
		case Uint16Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]uint16) = make([]uint16, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]uint16))[i] = binary.BigEndian.Uint16(b[offset : offset+2])
				offset += 2
			}
		case Byte:
			*receivers[receiveIndex].(*byte) = b[offset]
			offset++
		case ByteSlice:
			sliceLen = int(binary.BigEndian.Uint32(b[offset : offset+4]))
			offset += 4
			*receivers[receiveIndex].(*[]byte) = make([]byte, sliceLen, sliceLen)
			copy(*receivers[receiveIndex].(*[]byte), b[offset:offset+sliceLen])
			offset += sliceLen
		case Float64:
			*receivers[receiveIndex].(*float64) = math.Float64frombits(binary.BigEndian.Uint64(b[offset : offset+8]))
			offset += 8
		case Float64Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]float64) = make([]float64, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]float64))[i] = math.Float64frombits(binary.BigEndian.Uint64(b[offset : offset+8]))
				offset += 8
			}
		case Float32:
			*receivers[receiveIndex].(*float32) = math.Float32frombits(binary.BigEndian.Uint32(b[offset : offset+4]))
			offset += 4
		case Float32Slice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]float32) = make([]float32, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				(*receivers[receiveIndex].(*[]float32))[i] = math.Float32frombits(binary.BigEndian.Uint32(b[offset : offset+4]))
				offset += 4
			}
		case BoolTrue:
			*receivers[receiveIndex].(*bool) = true
		case BoolFalse:
			*receivers[receiveIndex].(*bool) = false
		case BoolSlice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]bool) = make([]bool, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				if b[offset] == BoolTrue {
					(*receivers[receiveIndex].(*[]bool))[i] = true
				} else {
					(*receivers[receiveIndex].(*[]bool))[i] = false
				}
				offset++
			}
		case String:
			sliceLen = int(binary.BigEndian.Uint32(b[offset : offset+4]))
			offset += 4
			*receivers[receiveIndex].(*string) = (string)(b[offset : offset+sliceLen])
			offset += sliceLen
		case StringSlice:
			sliceLen = int(binary.BigEndian.Uint16(b[offset : offset+2]))
			offset += 2
			*receivers[receiveIndex].(*[]string) = make([]string, sliceLen, sliceLen)
			i = 0
			for ; i < sliceLen; i++ {
				stringLen = int(binary.BigEndian.Uint32(b[offset : offset+4]))
				offset += 4
				(*receivers[receiveIndex].(*[]string))[i] = (string)(b[offset : offset+stringLen])
				offset += stringLen
			}
		default:
			return ErrExpectedTypeIdentifier
		}

		receiveIndex++
	}
	return nil
}
