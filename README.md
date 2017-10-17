# tfutils

tensorflow utils for golang

A collection of `tensorflow` helpers useful in `golang`.

## Motivation

[tensorflow's](https://github.com/tensorflow/tensorflow) `TFRecord` format can be read and written to using the native [RecordWriter](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/lib/io/record_writer.cc) and [RecordReader](https://github.com/tensorflow/tensorflow/blob/master/tensorflow/core/lib/io/record_reader.cc).  While this is exposed to your go environment via tensorflow, this is not handy for building simple translation tools which might need to read or operate on images to generate `tfrecords` -- enter `tfutils`.

## What you get (so far ...)

A reader that implements the `go` equivalent of the `RecordReader`, and a similar writer.

The `vendor` directory contains the generated go code for the protobuf, which is done using the `protoc` tool.  The `.proto` files required to generate the same, are found in the `tensorflow` project and are not checked into this repo.

## Record format:

| Field | Type | Description |
|:---------:|:------:|:------------------------------------------------------------------------------:|
| length | uint64 | The length of the `Example` tfrecord |
| lengthCrc | uint32 | 32-bit CRC of the 8-byte length |
| data | []byte | `length`-sized slice of bytes - this is the serialized version of the tfrecord |
| dataCrc | uint32 | 32-bit CRC of the data |

## Caveats

The `protobuf` files needed to interpret the `Example` are bundled in this "library" and so is the generated `go` code that allows the serialization and de-serialization of the data.  This is done to decrease the complexity of building this tool-set, but might get out of sync with the latest and greatest reader / writer logic.  However, this is not something that should change often (if at all) and most `protobuf` implementations promise backwards compatibility which will ignore old fields that it cannot recognize.

## Usage

See tests defined in `reader_test.go` and `writer_test.go`.

## Generating the protocol buffer stuff

```
protoc -I=$TENSORFLOW_DIR --go_out=$GOPATH/src/ $TENSORFLOW_DIR/tensorflow/core/example/*.proto
```
