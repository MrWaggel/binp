package binp

import (
	"encoding/binary"
	"errors"
)

var ErrUnsupportedType = errors.New("unsupported type")
var ErrExpectedTypeIdentifier = errors.New("expected type identifier")

// Pack packs all the given primitive datatype values into a slice of bytes.
func Pack(values ...interface{}) ([]byte, error) {
	size, err := Size(values)
	if err != nil {
		return nil, err
	}

	// Allocate slice
	b := make([]byte, size)

	// Write size to first four byes
	err = write(b, values)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PackFixed packs all the given values into a slice of bytes, this is useful
// if it is known that all the primitive data types are of fixed length, e.g. no
// strings, or []bytes. Has a slight performance boost as it skips counting the
// size of all values.
//
// Note: For every primitive data type that has to packed, the size has to be
// incremented with 1, that extra byte is used as the datatype identifier, except
// for boolean values. See Size().
//
// Example: Packing an int64 and a bool would take up 10bytes.
//
// PackFixed(10, int64(123), true)
//
// 1 byte that holds the identifier for the int64, 8 bytes that holds the int64,
// 1 byte that holds the bool.
func PackFixed(size int, values ...interface{}) (b []byte, err error) {
	b = make([]byte, size)
	err = write(b, values)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// PackTo packs all the given values into the given b ([]byte).
//
// Warning: make sure you allocated enough capacity to the []byte
// or runtime panics will occur. See the functions Size() and PackFixed()
// to calculate the needed capacity.
func PackTo(b []byte, values ...interface{}) error {
	return write(b, values)
}

// PackNetwork packs all the given values into a slice of bytes, whereas the
// first 4 bytes contain the length of all the values.
func PackNetwork(values ...interface{}) ([]byte, error) {
	var sizeTotal int
	size, err := Size(values)
	if err != nil {
		return nil, err
	}
	sizeTotal = size + 4 // First four bytes are the length of the rest

	// Allocate slice
	b := make([]byte, sizeTotal)

	// Write size to first four byes
	binary.BigEndian.PutUint32(b[0:4], uint32(size))
	err = write(b[4:], values)
	if err != nil {
		return nil, err
	}
	return b, nil
}

// PackNetworkFixed packs all the given values into a slice of bytes.
//
// See PackFixed and PackNetwork
func PackNetworkFixed(size int, values ...interface{}) (b []byte, err error) {
	b = make([]byte, size+4)

	binary.BigEndian.PutUint32(b[0:4], uint32(size))
	err = write(b[4:], values)
	if err != nil {
		return nil, err
	}
	return nil, err
}

// Gilles Van Vlasselaer https://mrwaggel.be
