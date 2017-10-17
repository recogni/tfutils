////////////////////////////////////////////////////////////////////////////////

package tfutils

////////////////////////////////////////////////////////////////////////////////

import (
	"bytes"
	"encoding/binary"
	"errors"
	"hash/crc32"
	"os"

	"fmt"
)

////////////////////////////////////////////////////////////////////////////////

const (
	crc32Polynomial = uint32(0x1EDC6F41)
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

func MaskedCrc(bs []byte) uint32 {
	return CrcMask(crc32.Checksum(bs, crc))
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
		lengthCrc: MaskedCrc(bs),
		data:      data,
		dataCrc:   MaskedCrc(data),
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

type RecordReader struct {
}

////////////////////////////////////////////////////////////////////////////////

// RecordWriterOptions defines the options to open the record writer with.
type RecordWriterOptions struct {
	CompressionType CompressionType
	// TODO: zlibOptions?
}

// RecordWriter implements a writer that appends tfrecord strings to a
// output file.  The `dstfile` is the path to the tfrecord file, and the
// `options` stores a copy of the writer options.
type RecordWriter struct {
	dstfile string
	options *RecordWriterOptions
}

// NewWriter returns a new instance of a tfrecrod writer.
func NewWriter(path string, options *RecordWriterOptions) (*RecordWriter, error) {
	// Try to open the file, if this does not work the writer should fail.
	_, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}

	return &RecordWriter{
		dstfile: path,
		options: options,
	}, nil
}

func (rw *RecordWriter) WriteRecord(data []byte) error {
	e, err := newEntry(data)
	if err != nil {
		return err
	}

	fmt.Printf("RECORD: %#v\n", e)

	bs, err := e.Marshal()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(rw.dstfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bs)
	return err
}

func (rw *RecordWriter) Close() error {
	// Nothing to do since we open and close the file handle on each write.
	// We might have some extra stuff to handle here in the ZLIB compression
	// is enabled.
	return nil
}

func (rw *RecordWriter) Flush() error {
	// Nothing to do yet, once ZLIB compression is supported we will have to
	// flush that stream to file as needed.
	return nil
}

////////////////////////////////////////////////////////////////////////////////
