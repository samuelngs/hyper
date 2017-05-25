// Code generated by protoc-gen-go. DO NOT EDIT.
// source: packet.proto

package sync

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Packet struct {
	ID        string `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Signature bool   `protobuf:"varint,10,opt,name=Signature" json:"Signature,omitempty"`
	Action    int32  `protobuf:"varint,20,opt,name=Action" json:"Action,omitempty"`
	Namespace string `protobuf:"bytes,30,opt,name=Namespace" json:"Namespace,omitempty"`
	Channel   string `protobuf:"bytes,40,opt,name=Channel" json:"Channel,omitempty"`
	Call      string `protobuf:"bytes,50,opt,name=Call" json:"Call,omitempty"`
	Message   []byte `protobuf:"bytes,60,opt,name=Message,proto3" json:"Message,omitempty"`
	Error     string `protobuf:"bytes,70,opt,name=Error" json:"Error,omitempty"`
}

func (m *Packet) Reset()                    { *m = Packet{} }
func (m *Packet) String() string            { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()               {}
func (*Packet) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Packet) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Packet) GetSignature() bool {
	if m != nil {
		return m.Signature
	}
	return false
}

func (m *Packet) GetAction() int32 {
	if m != nil {
		return m.Action
	}
	return 0
}

func (m *Packet) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *Packet) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *Packet) GetCall() string {
	if m != nil {
		return m.Call
	}
	return ""
}

func (m *Packet) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Packet) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Packet)(nil), "sync.Packet")
}

func init() { proto.RegisterFile("packet.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x48, 0x4c, 0xce,
	0x4e, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0xae, 0xcc, 0x4b, 0x56, 0x3a,
	0xc5, 0xc8, 0xc5, 0x16, 0x00, 0x16, 0x16, 0xe2, 0xe3, 0x62, 0xf2, 0x74, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x62, 0xf2, 0x74, 0x11, 0x92, 0xe1, 0xe2, 0x0c, 0xce, 0x4c, 0xcf, 0x4b, 0x2c,
	0x29, 0x2d, 0x4a, 0x95, 0xe0, 0x52, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x08, 0x08, 0x89, 0x71, 0xb1,
	0x39, 0x26, 0x97, 0x64, 0xe6, 0xe7, 0x49, 0x88, 0x28, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x79, 0x20,
	0x5d, 0x7e, 0x89, 0xb9, 0xa9, 0xc5, 0x05, 0x89, 0xc9, 0xa9, 0x12, 0x72, 0x60, 0xc3, 0x10, 0x02,
	0x42, 0x12, 0x5c, 0xec, 0xce, 0x19, 0x89, 0x79, 0x79, 0xa9, 0x39, 0x12, 0x1a, 0x60, 0x39, 0x18,
	0x57, 0x48, 0x88, 0x8b, 0xc5, 0x39, 0x31, 0x27, 0x47, 0xc2, 0x08, 0x2c, 0x0c, 0x66, 0x83, 0x54,
	0xfb, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x4a, 0xd8, 0x28, 0x30, 0x6a, 0xf0, 0x04, 0xc1, 0xb8,
	0x42, 0x22, 0x5c, 0xac, 0xae, 0x45, 0x45, 0xf9, 0x45, 0x12, 0x6e, 0x60, 0xe5, 0x10, 0x4e, 0x12,
	0x1b, 0xd8, 0x67, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0xec, 0x6d, 0x53, 0xe9, 0x00,
	0x00, 0x00,
}