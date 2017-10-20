////////////////////////////////////////////////////////////////////////////////

package tfutils

////////////////////////////////////////////////////////////////////////////////

import (
	"os"
)

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
	f       *os.File
	dstfile string
	options *RecordWriterOptions
}

// NewWriter returns a new instance of a tfrecrod writer.
func NewWriter(path string, options *RecordWriterOptions) (*RecordWriter, error) {
	// Try to open the file, if this does not work the writer should fail.
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return nil, err
	}

	return &RecordWriter{
		f:       f,
		dstfile: path,
		options: options,
	}, nil
}

func (rw *RecordWriter) WriteRecord(data []byte) error {
	e, err := newEntry(data)
	if err != nil {
		return err
	}

	bs, err := e.Marshal()
	if err != nil {
		return err
	}

	_, err = rw.f.Write(bs)
	return err
}

func (rw *RecordWriter) Close() error {
	// Nothing to do since we open and close the file handle on each write.
	// We might have some extra stuff to handle here in the ZLIB compression
	// is enabled.
	if rw.f != nil {
		rw.f.Close()
	}
	return nil
}

func (rw *RecordWriter) Flush() error {
	// Nothing to do yet, once ZLIB compression is supported we will have to
	// flush that stream to file as needed.
	return nil
}

////////////////////////////////////////////////////////////////////////////////
