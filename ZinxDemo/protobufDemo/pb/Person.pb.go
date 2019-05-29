// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Person.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//定义一个protobuf协议
type Person struct {
	Name   string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age    int32          `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Emails []string       `protobuf:"bytes,3,rep,name=emails,proto3" json:"emails,omitempty"`
	Phones []*PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*Person_School
	//	*Person_Score
	Data                 isPerson_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Person) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *Person) GetEmails() []string {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *Person) GetPhones() []*PhoneNumber {
	if m != nil {
		return m.Phones
	}
	return nil
}

type isPerson_Data interface {
	isPerson_Data()
}

type Person_School struct {
	School string `protobuf:"bytes,5,opt,name=school,proto3,oneof"`
}

type Person_Score struct {
	Score uint32 `protobuf:"varint,6,opt,name=score,proto3,oneof"`
}

func (*Person_School) isPerson_Data() {}

func (*Person_Score) isPerson_Data() {}

func (m *Person) GetData() isPerson_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Person) GetSchool() string {
	if x, ok := m.GetData().(*Person_School); ok {
		return x.School
	}
	return ""
}

func (m *Person) GetScore() uint32 {
	if x, ok := m.GetData().(*Person_Score); ok {
		return x.Score
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Person) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Person_OneofMarshaler, _Person_OneofUnmarshaler, _Person_OneofSizer, []interface{}{
		(*Person_School)(nil),
		(*Person_Score)(nil),
	}
}

func _Person_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Person)
	// Data
	switch x := m.Data.(type) {
	case *Person_School:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.School)
	case *Person_Score:
		b.EncodeVarint(6<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.Score))
	case nil:
	default:
		return fmt.Errorf("Person.Data has unexpected type %T", x)
	}
	return nil
}

func _Person_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Person)
	switch tag {
	case 5: // Data.school
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &Person_School{x}
		return true, err
	case 6: // Data.score
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &Person_Score{uint32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _Person_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Person)
	// Data
	switch x := m.Data.(type) {
	case *Person_School:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.School)))
		n += len(x.School)
	case *Person_Score:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.Score))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//一个protobuf协议的消息
type PhoneNumber struct {
	Number               string   `protobuf:"bytes,1,opt,name=Number,proto3" json:"Number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_841ab6396175eaf3, []int{1}
}

func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (m *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(m, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

func init() {
	proto.RegisterType((*Person)(nil), "pb.Person")
	proto.RegisterType((*PhoneNumber)(nil), "pb.PhoneNumber")
}

func init() { proto.RegisterFile("Person.proto", fileDescriptor_841ab6396175eaf3) }

var fileDescriptor_841ab6396175eaf3 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x41, 0x6a, 0x85, 0x30,
	0x10, 0x86, 0x8d, 0xd1, 0x80, 0x63, 0x4b, 0xcb, 0x2c, 0x24, 0xcb, 0x20, 0x94, 0x66, 0xe5, 0xa2,
	0xbd, 0x41, 0xe9, 0xc2, 0x55, 0x91, 0xdc, 0x20, 0x91, 0x50, 0x0b, 0x6a, 0x42, 0x62, 0x0f, 0xf5,
	0x6e, 0xf9, 0x88, 0x66, 0xf1, 0x76, 0xdf, 0xf7, 0x0f, 0xcc, 0xfc, 0x03, 0x4f, 0x93, 0x0d, 0xd1,
	0xed, 0x83, 0x0f, 0xee, 0x70, 0x58, 0x7a, 0xd3, 0xdf, 0x08, 0xb0, 0x2b, 0x44, 0x84, 0x6a, 0xd7,
	0x9b, 0xe5, 0x44, 0x10, 0xd9, 0xa8, 0x93, 0xf1, 0x15, 0xa8, 0xfe, 0xb5, 0xbc, 0x14, 0x44, 0xd6,
	0x2a, 0x21, 0x76, 0xc0, 0xec, 0xa6, 0xff, 0xd6, 0xc8, 0xa9, 0xa0, 0xb2, 0x51, 0xd9, 0xf0, 0x1d,
	0x98, 0x5f, 0xdc, 0x6e, 0x23, 0xaf, 0x04, 0x95, 0xed, 0xc7, 0xcb, 0xe0, 0xcd, 0x30, 0xa5, 0xe4,
	0xe7, 0x7f, 0x33, 0x36, 0xa8, 0x3c, 0x46, 0x0e, 0x2c, 0xce, 0x8b, 0x73, 0x2b, 0xaf, 0xd3, 0xa1,
	0xb1, 0x50, 0xd9, 0xb1, 0x83, 0x3a, 0xce, 0x2e, 0x58, 0xce, 0x04, 0x91, 0xcf, 0x63, 0xa1, 0x2e,
	0xfd, 0x62, 0x50, 0x7d, 0xeb, 0x43, 0xf7, 0x6f, 0xd0, 0x3e, 0x2c, 0x4c, 0x4d, 0x2e, 0xca, 0x8d,
	0xb3, 0x19, 0x76, 0x7e, 0xf7, 0x79, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x61, 0xf6, 0x4a, 0xed,
	0x00, 0x00, 0x00,
}
