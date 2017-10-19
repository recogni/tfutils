////////////////////////////////////////////////////////////////////////////////

package tfutils

////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
)

////////////////////////////////////////////////////////////////////////////////

const (
	crc32Polynomial = crc32.Castagnoli
	crc32MaskDelta  = uint32(0xa282ead8)
)

var (
	crc = crc32.MakeTable(crc32Polynomial)
)

func CrcMask(v uint32) uint32 {
	return ((v >> 15) | (v << 17)) + crc32MaskDelta
}

func CrcUnmask(v uint32) uint32 {
	r := v - crc32MaskDelta
	return ((r >> 17) | (r << 15))
}

func MaskedCrc(bs []byte, n int64) uint32 {
	return CrcMask(crc32.Checksum(bs[:n], crc))
}

////////////////////////////////////////////////////////////////////////////////

type recordEntry struct {
	length    uint64
	lengthCrc uint32
	data      []byte
	dataCrc   uint32
}

func newEntry(data []byte) (*recordEntry, error) {
	length := uint64(len(data))
	if length == 0 {
		return nil, errors.New("data array is empty")
	}

	bs := make([]byte, 8)
	binary.LittleEndian.PutUint64(bs[0:], length)
	return &recordEntry{
		length:    length,
		lengthCrc: MaskedCrc(bs, 8),
		data:      data,
		dataCrc:   MaskedCrc(data, int64(length)),
	}, nil
}

func (re *recordEntry) Marshal() ([]byte, error) {
	b := bytes.NewBuffer(make([]byte, 0, re.length+8+4+4))
	binary.Write(b, binary.LittleEndian, re.length)
	binary.Write(b, binary.LittleEndian, re.lengthCrc)
	b.Write(re.data[:])
	binary.Write(b, binary.LittleEndian, re.dataCrc)
	return b.Bytes(), nil
}

////////////////////////////////////////////////////////////////////////////////

type CompressionType int

const (
	CompressionTypeNone CompressionType = iota
	CompressionTypeZlib
)

////////////////////////////////////////////////////////////////////////////////
