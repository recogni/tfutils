package tfutils

import (
	"fmt"
	"reflect"

	"github.com/golang/protobuf/proto"

	// Import the "example" and "feature" protocol buffers
	tf "tensorflow/core/example"
)

// GetFeatureFromGoType returns a tensorflow Feature instance for the
// underlying golang type.  Note that not all primitive types are massaged
// into one of the feature types here (yet).
func GetFeatureFromGoType(x interface{}) (*tf.Feature, error) {
	tstr := reflect.TypeOf(x).String()
	switch tstr {
	case "int":
		return &tf.Feature{Kind: &tf.Feature_Int64List{Int64List: &tf.Int64List{[]int64{int64(x.(int))}}}}, nil
	case "string":
		return &tf.Feature{Kind: &tf.Feature_BytesList{BytesList: &tf.BytesList{[][]byte{[]byte(x.(string))}}}}, nil
	case "[]uint8":
		return &tf.Feature{Kind: &tf.Feature_BytesList{BytesList: &tf.BytesList{[][]byte{[]byte(x.([]uint8))}}}}, nil
	case "[]int64":
		return &tf.Feature{Kind: &tf.Feature_Int64List{Int64List: &tf.Int64List{x.([]int64)}}}, nil
	case "[][]uint8":
		return &tf.Feature{Kind: &tf.Feature_BytesList{BytesList: &tf.BytesList{x.([][]byte)}}}, nil
	case "[]float32":
		return &tf.Feature{Kind: &tf.Feature_FloatList{FloatList: &tf.FloatList{x.([]float32)}}}, nil
	}
	return nil, fmt.Errorf("unknown base type %s", tstr)
}

// GetFeaturesFromMap returns a tensorflow Features instance based on
// the generic map of string->value being passed in.
func GetFeaturesFromMap(m map[string]interface{}) (*tf.Features, error) {
	fs := map[string]*tf.Feature{}
	for k, v := range m {
		f, err := GetFeatureFromGoType(v)
		if err != nil {
			return nil, err
		}
		fs[k] = f
	}
	return &tf.Features{Feature: fs}, nil
}

// GetTFRecordStringForFeatures returns a serialized version of an "Example"
// protobuffer for writing using a TF RecordWriter.
func GetTFRecordStringForFeatures(fs *tf.Features) ([]byte, error) {
	ex := &tf.Example{
		Features: fs,
	}
	return proto.Marshal(ex)
}

func GetFeatureMapFromTFRecord(data []byte) (*tf.Features, error) {
	ex := tf.Example{}
	if err := proto.Unmarshal(data, &ex); err != nil {
		return nil, err
	}
	return ex.Features, nil
}
