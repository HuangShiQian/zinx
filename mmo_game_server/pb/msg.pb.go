// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

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

//返回给玩家上线的ID信息
type SyncPid struct {
	Pid                  int32    `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPid) Reset()         { *m = SyncPid{} }
func (m *SyncPid) String() string { return proto.CompactTextString(m) }
func (*SyncPid) ProtoMessage()    {}
func (*SyncPid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *SyncPid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPid.Unmarshal(m, b)
}
func (m *SyncPid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPid.Marshal(b, m, deterministic)
}
func (m *SyncPid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPid.Merge(m, src)
}
func (m *SyncPid) XXX_Size() int {
	return xxx_messageInfo_SyncPid.Size(m)
}
func (m *SyncPid) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPid.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPid proto.InternalMessageInfo

func (m *SyncPid) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//返回给上线玩家初始的坐标
type BroadCast struct {
	Pid int32 `protobuf:"varint,1,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Tp  int32 `protobuf:"varint,2,opt,name=Tp,proto3" json:"Tp,omitempty"`
	// Types that are valid to be assigned to Data:
	//	*BroadCast_Content
	//	*BroadCast_P
	//	*BroadCast_ActionData
	Data                 isBroadCast_Data `protobuf_oneof:"Data"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *BroadCast) Reset()         { *m = BroadCast{} }
func (m *BroadCast) String() string { return proto.CompactTextString(m) }
func (*BroadCast) ProtoMessage()    {}
func (*BroadCast) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{1}
}

func (m *BroadCast) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BroadCast.Unmarshal(m, b)
}
func (m *BroadCast) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BroadCast.Marshal(b, m, deterministic)
}
func (m *BroadCast) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BroadCast.Merge(m, src)
}
func (m *BroadCast) XXX_Size() int {
	return xxx_messageInfo_BroadCast.Size(m)
}
func (m *BroadCast) XXX_DiscardUnknown() {
	xxx_messageInfo_BroadCast.DiscardUnknown(m)
}

var xxx_messageInfo_BroadCast proto.InternalMessageInfo

func (m *BroadCast) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *BroadCast) GetTp() int32 {
	if m != nil {
		return m.Tp
	}
	return 0
}

type isBroadCast_Data interface {
	isBroadCast_Data()
}

type BroadCast_Content struct {
	Content string `protobuf:"bytes,3,opt,name=Content,proto3,oneof"`
}

type BroadCast_P struct {
	P *Position `protobuf:"bytes,4,opt,name=P,proto3,oneof"`
}

type BroadCast_ActionData struct {
	ActionData int32 `protobuf:"varint,5,opt,name=ActionData,proto3,oneof"`
}

func (*BroadCast_Content) isBroadCast_Data() {}

func (*BroadCast_P) isBroadCast_Data() {}

func (*BroadCast_ActionData) isBroadCast_Data() {}

func (m *BroadCast) GetData() isBroadCast_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroadCast) GetContent() string {
	if x, ok := m.GetData().(*BroadCast_Content); ok {
		return x.Content
	}
	return ""
}

func (m *BroadCast) GetP() *Position {
	if x, ok := m.GetData().(*BroadCast_P); ok {
		return x.P
	}
	return nil
}

func (m *BroadCast) GetActionData() int32 {
	if x, ok := m.GetData().(*BroadCast_ActionData); ok {
		return x.ActionData
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*BroadCast) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _BroadCast_OneofMarshaler, _BroadCast_OneofUnmarshaler, _BroadCast_OneofSizer, []interface{}{
		(*BroadCast_Content)(nil),
		(*BroadCast_P)(nil),
		(*BroadCast_ActionData)(nil),
	}
}

func _BroadCast_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*BroadCast)
	// Data
	switch x := m.Data.(type) {
	case *BroadCast_Content:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Content)
	case *BroadCast_P:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.P); err != nil {
			return err
		}
	case *BroadCast_ActionData:
		b.EncodeVarint(5<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.ActionData))
	case nil:
	default:
		return fmt.Errorf("BroadCast.Data has unexpected type %T", x)
	}
	return nil
}

func _BroadCast_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*BroadCast)
	switch tag {
	case 3: // Data.Content
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &BroadCast_Content{x}
		return true, err
	case 4: // Data.P
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Position)
		err := b.DecodeMessage(msg)
		m.Data = &BroadCast_P{msg}
		return true, err
	case 5: // Data.ActionData
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.Data = &BroadCast_ActionData{int32(x)}
		return true, err
	default:
		return false, nil
	}
}

func _BroadCast_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*BroadCast)
	// Data
	switch x := m.Data.(type) {
	case *BroadCast_Content:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Content)))
		n += len(x.Content)
	case *BroadCast_P:
		s := proto.Size(x.P)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *BroadCast_ActionData:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(x.ActionData))
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//位置信息
type Position struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=X,proto3" json:"X,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=Y,proto3" json:"Y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=Z,proto3" json:"Z,omitempty"`
	V                    float32  `protobuf:"fixed32,4,opt,name=V,proto3" json:"V,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}
func (*Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{2}
}

func (m *Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Position.Unmarshal(m, b)
}
func (m *Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Position.Marshal(b, m, deterministic)
}
func (m *Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Position.Merge(m, src)
}
func (m *Position) XXX_Size() int {
	return xxx_messageInfo_Position.Size(m)
}
func (m *Position) XXX_DiscardUnknown() {
	xxx_messageInfo_Position.DiscardUnknown(m)
}

var xxx_messageInfo_Position proto.InternalMessageInfo

func (m *Position) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Position) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *Position) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

func (m *Position) GetV() float32 {
	if m != nil {
		return m.V
	}
	return 0
}

func init() {
	proto.RegisterType((*SyncPid)(nil), "pb.SyncPid")
	proto.RegisterType((*BroadCast)(nil), "pb.BroadCast")
	proto.RegisterType((*Position)(nil), "pb.Position")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x8f, 0xbf, 0x4e, 0xc4, 0x30,
	0x0c, 0x87, 0x6b, 0xdf, 0x3f, 0x6a, 0x4e, 0x08, 0x65, 0x8a, 0x80, 0x21, 0xea, 0xd4, 0xa9, 0x03,
	0x3c, 0x01, 0x3d, 0x86, 0x8e, 0x91, 0x39, 0x9d, 0xee, 0x6e, 0x4b, 0x5b, 0x84, 0x32, 0x90, 0x44,
	0x6d, 0x16, 0x1e, 0x83, 0x37, 0x46, 0x89, 0x54, 0x09, 0x89, 0x29, 0xf9, 0x7e, 0x96, 0xed, 0xcf,
	0x54, 0x7e, 0xcd, 0x9f, 0x4d, 0x98, 0x7c, 0xf4, 0x02, 0x43, 0x5f, 0x3d, 0xd2, 0xee, 0xfd, 0xdb,
	0x0d, 0xda, 0x8e, 0xe2, 0x9e, 0x56, 0xda, 0x8e, 0x12, 0x14, 0xd4, 0x1b, 0x4e, 0xdf, 0xea, 0x07,
	0xa8, 0x6c, 0x27, 0x6f, 0xc6, 0x83, 0x99, 0xe3, 0xff, 0xba, 0xb8, 0x23, 0x3c, 0x06, 0x89, 0x39,
	0xc0, 0x63, 0x10, 0x0f, 0xb4, 0x3b, 0x78, 0x17, 0x3f, 0x5c, 0x94, 0x2b, 0x05, 0x75, 0xd9, 0x15,
	0xbc, 0x04, 0xe2, 0x89, 0x40, 0xcb, 0xb5, 0x82, 0xfa, 0xf6, 0x79, 0xdf, 0x84, 0xbe, 0xd1, 0x7e,
	0xb6, 0xd1, 0x7a, 0xd7, 0x15, 0x0c, 0x5a, 0x28, 0xa2, 0xd7, 0x21, 0xe1, 0x9b, 0x89, 0x46, 0x6e,
	0xd2, 0xc4, 0xae, 0xe0, 0x3f, 0x59, 0xbb, 0xa5, 0x75, 0x7a, 0xab, 0x96, 0x6e, 0x96, 0x56, 0xb1,
	0x27, 0x38, 0x67, 0x1f, 0x64, 0x38, 0x27, 0xba, 0x64, 0x19, 0x64, 0xb8, 0x24, 0xba, 0x66, 0x0b,
	0x64, 0xb8, 0x26, 0x3a, 0xe5, 0xed, 0xc8, 0x70, 0xea, 0xb7, 0xf9, 0xfe, 0x97, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x36, 0xa6, 0xa6, 0x1a, 0x0c, 0x01, 0x00, 0x00,
}