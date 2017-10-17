// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/attr_value.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value isAttrValue_Value `protobuf_oneof:"value"`
}

func (m *AttrValue) Reset()                    { *m = AttrValue{} }
func (m *AttrValue) String() string            { return proto.CompactTextString(m) }
func (*AttrValue) ProtoMessage()               {}
func (*AttrValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}
type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,oneof"`
}
type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,oneof"`
}
type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,oneof"`
}
type AttrValue_Type struct {
	Type DataType `protobuf:"varint,6,opt,name=type,enum=tensorflow.DataType,oneof"`
}
type AttrValue_Shape struct {
	Shape *TensorShapeProto `protobuf:"bytes,7,opt,name=shape,oneof"`
}
type AttrValue_Tensor struct {
	Tensor *TensorProto `protobuf:"bytes,8,opt,name=tensor,oneof"`
}
type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,oneof"`
}
type AttrValue_Func struct {
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,oneof"`
}
type AttrValue_Placeholder struct {
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,oneof"`
}

func (*AttrValue_S) isAttrValue_Value()           {}
func (*AttrValue_I) isAttrValue_Value()           {}
func (*AttrValue_F) isAttrValue_Value()           {}
func (*AttrValue_B) isAttrValue_Value()           {}
func (*AttrValue_Type) isAttrValue_Value()        {}
func (*AttrValue_Shape) isAttrValue_Value()       {}
func (*AttrValue_Tensor) isAttrValue_Value()      {}
func (*AttrValue_List) isAttrValue_Value()        {}
func (*AttrValue_Func) isAttrValue_Value()        {}
func (*AttrValue_Placeholder) isAttrValue_Value() {}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttrValue) GetS() []byte {
	if x, ok := m.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (m *AttrValue) GetI() int64 {
	if x, ok := m.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (m *AttrValue) GetF() float32 {
	if x, ok := m.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (m *AttrValue) GetB() bool {
	if x, ok := m.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (m *AttrValue) GetType() DataType {
	if x, ok := m.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return DataType_DT_INVALID
}

func (m *AttrValue) GetShape() *TensorShapeProto {
	if x, ok := m.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *AttrValue) GetTensor() *TensorProto {
	if x, ok := m.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (m *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := m.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (m *AttrValue) GetFunc() *NameAttrList {
	if x, ok := m.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (m *AttrValue) GetPlaceholder() string {
	if x, ok := m.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AttrValue) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AttrValue_OneofMarshaler, _AttrValue_OneofUnmarshaler, _AttrValue_OneofSizer, []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
}

func _AttrValue_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeRawBytes(x.S)
	case *AttrValue_I:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.I))
	case *AttrValue_F:
		b.EncodeVarint(4<<3 | proto.WireFixed32)
		b.EncodeFixed32(uint64(math.Float32bits(x.F)))
	case *AttrValue_B:
		t := uint64(0)
		if x.B {
			t = 1
		}
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(t)
	case *AttrValue_Type:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Shape); err != nil {
			return err
		}
	case *AttrValue_Tensor:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tensor); err != nil {
			return err
		}
	case *AttrValue_List:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.List); err != nil {
			return err
		}
	case *AttrValue_Func:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Func); err != nil {
			return err
		}
	case *AttrValue_Placeholder:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Placeholder)
	case nil:
	default:
		return fmt.Errorf("AttrValue.Value has unexpected type %T", x)
	}
	return nil
}

func _AttrValue_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AttrValue)
	switch tag {
	case 2: // value.s
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeRawBytes(true)
		m.Value = &AttrValue_S{x}
		return true, err
	case 3: // value.i
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_I{int64(x)}
		return true, err
	case 4: // value.f
		if wire != proto.WireFixed32 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed32()
		m.Value = &AttrValue_F{math.Float32frombits(uint32(x))}
		return true, err
	case 5: // value.b
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_B{x != 0}
		return true, err
	case 6: // value.type
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Value = &AttrValue_Type{DataType(x)}
		return true, err
	case 7: // value.shape
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorShapeProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Shape{msg}
		return true, err
	case 8: // value.tensor
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TensorProto)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Tensor{msg}
		return true, err
	case 1: // value.list
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AttrValue_ListValue)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_List{msg}
		return true, err
	case 10: // value.func
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NameAttrList)
		err := b.DecodeMessage(msg)
		m.Value = &AttrValue_Func{msg}
		return true, err
	case 9: // value.placeholder
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Value = &AttrValue_Placeholder{x}
		return true, err
	default:
		return false, nil
	}
}

func _AttrValue_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AttrValue)
	// value
	switch x := m.Value.(type) {
	case *AttrValue_S:
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.S)))
		n += len(x.S)
	case *AttrValue_I:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.I))
	case *AttrValue_F:
		n += proto.SizeVarint(4<<3 | proto.WireFixed32)
		n += 4
	case *AttrValue_B:
		n += proto.SizeVarint(5<<3 | proto.WireVarint)
		n += 1
	case *AttrValue_Type:
		n += proto.SizeVarint(6<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.Type))
	case *AttrValue_Shape:
		s := proto.Size(x.Shape)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Tensor:
		s := proto.Size(x.Tensor)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_List:
		s := proto.Size(x.List)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Func:
		s := proto.Size(x.Func)
		n += proto.SizeVarint(10<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AttrValue_Placeholder:
		n += proto.SizeVarint(9<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.Placeholder)))
		n += len(x.Placeholder)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// LINT.IfChange
type AttrValue_ListValue struct {
	S      [][]byte            `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`
	I      []int64             `protobuf:"varint,3,rep,packed,name=i" json:"i,omitempty"`
	F      []float32           `protobuf:"fixed32,4,rep,packed,name=f" json:"f,omitempty"`
	B      []bool              `protobuf:"varint,5,rep,packed,name=b" json:"b,omitempty"`
	Type   []DataType          `protobuf:"varint,6,rep,packed,name=type,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape  []*TensorShapeProto `protobuf:"bytes,7,rep,name=shape" json:"shape,omitempty"`
	Tensor []*TensorProto      `protobuf:"bytes,8,rep,name=tensor" json:"tensor,omitempty"`
	Func   []*NameAttrList     `protobuf:"bytes,9,rep,name=func" json:"func,omitempty"`
}

func (m *AttrValue_ListValue) Reset()                    { *m = AttrValue_ListValue{} }
func (m *AttrValue_ListValue) String() string            { return proto.CompactTextString(m) }
func (*AttrValue_ListValue) ProtoMessage()               {}
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

func (m *AttrValue_ListValue) GetS() [][]byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *AttrValue_ListValue) GetI() []int64 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *AttrValue_ListValue) GetF() []float32 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *AttrValue_ListValue) GetB() []bool {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AttrValue_ListValue) GetType() []DataType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AttrValue_ListValue) GetShape() []*TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *AttrValue_ListValue) GetTensor() []*TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if m != nil {
		return m.Func
	}
	return nil
}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	Name string                `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Attr map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *NameAttrList) Reset()                    { *m = NameAttrList{} }
func (m *NameAttrList) String() string            { return proto.CompactTextString(m) }
func (*NameAttrList) ProtoMessage()               {}
func (*NameAttrList) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *NameAttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameAttrList) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*AttrValue)(nil), "tensorflow.AttrValue")
	proto.RegisterType((*AttrValue_ListValue)(nil), "tensorflow.AttrValue.ListValue")
	proto.RegisterType((*NameAttrList)(nil), "tensorflow.NameAttrList")
}

func init() { proto.RegisterFile("tensorflow/core/framework/attr_value.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 496 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xdf, 0x6a, 0xdb, 0x30,
	0x14, 0x87, 0x23, 0xcb, 0x49, 0x63, 0x25, 0x74, 0x41, 0x6c, 0x4c, 0x98, 0xc1, 0x4c, 0x60, 0x43,
	0x74, 0xc1, 0xd9, 0xbc, 0x3f, 0x8c, 0xdd, 0x2d, 0x6c, 0x90, 0x8b, 0x51, 0x8a, 0x57, 0x76, 0x5b,
	0x94, 0x4c, 0x5e, 0x43, 0x9d, 0xd8, 0xc8, 0xea, 0x4a, 0x9e, 0x60, 0xb7, 0x7b, 0x8e, 0x3d, 0xe1,
	0x2e, 0xc7, 0x39, 0x72, 0x1d, 0xc3, 0x9a, 0xe6, 0x4e, 0xe7, 0xe8, 0xfb, 0x29, 0x27, 0x9f, 0x25,
	0x76, 0x62, 0xf5, 0xa6, 0x2a, 0x4c, 0x96, 0x17, 0x37, 0xd3, 0x65, 0x61, 0xf4, 0x34, 0x33, 0x6a,
	0xad, 0x6f, 0x0a, 0x73, 0x35, 0x55, 0xd6, 0x9a, 0x8b, 0x9f, 0x2a, 0xbf, 0xd6, 0x71, 0x69, 0x0a,
	0x5b, 0x70, 0xb6, 0x63, 0xc3, 0xe7, 0xfb, 0x73, 0x6e, 0xc7, 0x65, 0xc2, 0xc9, 0x21, 0xee, 0xa2,
	0xba, 0x54, 0x65, 0xfd, 0x0b, 0xe1, 0xb3, 0x7b, 0xe8, 0x6d, 0xa9, 0x2b, 0x87, 0x8d, 0x7f, 0x75,
	0x59, 0xf0, 0xd1, 0x5a, 0xf3, 0x0d, 0x86, 0xe3, 0xc7, 0x8c, 0x54, 0xc2, 0x8b, 0x88, 0x1c, 0xce,
	0x3b, 0x29, 0xa9, 0xa0, 0x5e, 0x09, 0x1a, 0x11, 0x49, 0xa1, 0x5e, 0x41, 0x9d, 0x09, 0x3f, 0x22,
	0xd2, 0x83, 0x3a, 0x83, 0x7a, 0x21, 0xba, 0x11, 0x91, 0x7d, 0xa8, 0x17, 0xfc, 0x84, 0xf9, 0x70,
	0xb8, 0xe8, 0x45, 0x44, 0x1e, 0x27, 0x0f, 0xe3, 0xdd, 0x0c, 0xf1, 0x27, 0x65, 0xd5, 0xf9, 0xb6,
	0xd4, 0xf3, 0x4e, 0x8a, 0x0c, 0x7f, 0xc3, 0xba, 0x38, 0xaf, 0x38, 0x8a, 0x88, 0x1c, 0x24, 0x4f,
	0xda, 0xf0, 0x39, 0x2e, 0xbf, 0xc2, 0xf6, 0x19, 0x8c, 0x39, 0xef, 0xa4, 0x0e, 0xe6, 0xaf, 0x58,
	0xcf, 0x71, 0xa2, 0x8f, 0xb1, 0xc7, 0xff, 0xc7, 0x6e, 0x13, 0x35, 0xc8, 0xdf, 0x32, 0x3f, 0x5f,
	0x55, 0x56, 0x10, 0x0c, 0x3c, 0x6d, 0x07, 0x9a, 0x7f, 0x1e, 0x7f, 0x59, 0x55, 0x16, 0x57, 0x30,
	0x1f, 0xe0, 0x3c, 0x66, 0x7e, 0x76, 0xbd, 0x59, 0x0a, 0x86, 0x31, 0xd1, 0x8e, 0x9d, 0xaa, 0xb5,
	0x86, 0x28, 0x84, 0x80, 0x07, 0x8e, 0x8f, 0xd9, 0xa0, 0xcc, 0xd5, 0x52, 0x5f, 0x16, 0xf9, 0x77,
	0x6d, 0x44, 0x10, 0x11, 0x19, 0xcc, 0x3b, 0x69, 0xbb, 0x19, 0xfe, 0xf6, 0x58, 0xd0, 0xfc, 0x12,
	0x1f, 0x3a, 0xdb, 0x54, 0x0e, 0xc1, 0xf5, 0xc8, 0xb9, 0xa6, 0x92, 0xce, 0xbc, 0x11, 0x01, 0xdb,
	0x23, 0x67, 0x9b, 0x4a, 0xcf, 0x75, 0x32, 0xe8, 0x80, 0x6f, 0x2a, 0xfb, 0xae, 0xb3, 0xe0, 0x93,
	0xc6, 0x38, 0xdd, 0x67, 0x1c, 0x51, 0xe7, 0x3c, 0xd9, 0x39, 0xa7, 0x87, 0x9c, 0xdf, 0x1a, 0x9f,
	0xb6, 0x8c, 0xd3, 0x7b, 0x8c, 0x37, 0xbe, 0x27, 0xb5, 0xb8, 0x00, 0xf1, 0xbd, 0xe2, 0x9c, 0xb6,
	0xd9, 0x11, 0xeb, 0xe2, 0xc3, 0x18, 0xff, 0x21, 0x6c, 0xd8, 0xde, 0xe7, 0x9c, 0xf9, 0x1b, 0xb5,
	0xd6, 0xf8, 0xdd, 0x82, 0x14, 0xd7, 0xfc, 0x1d, 0xf3, 0xe1, 0x2d, 0xa1, 0xb5, 0x41, 0x32, 0xde,
	0x77, 0x36, 0x7e, 0xd8, 0xcf, 0x1b, 0x6b, 0xb6, 0x29, 0xf2, 0xe1, 0xa9, 0xbb, 0xe5, 0xd8, 0xe2,
	0x23, 0x46, 0xaf, 0xf4, 0xb6, 0x3e, 0x17, 0x96, 0xfc, 0x45, 0x3d, 0x04, 0xde, 0xfd, 0x41, 0xf2,
	0xe8, 0xce, 0x3b, 0x92, 0x3a, 0xe6, 0x83, 0xf7, 0x9e, 0xcc, 0x5e, 0x32, 0x51, 0x98, 0x1f, 0x6d,
	0xac, 0x79, 0x5e, 0xb3, 0x07, 0x4d, 0x02, 0xbd, 0x54, 0x67, 0xe4, 0x2f, 0x21, 0x8b, 0x1e, 0xbe,
	0xb7, 0xd7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xeb, 0xc8, 0xea, 0x26, 0x04, 0x00, 0x00,
}
