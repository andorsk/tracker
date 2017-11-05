// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

/*
Package user is a generated protocol buffer package.

It is generated from these files:
	user.proto

It has these top-level messages:
	User
*/
package user

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	UserId int64  `protobuf:"varint,1,opt,name=UserId,json=userId" json:"UserId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=Name,json=name" json:"Name,omitempty"`
	Age    int32  `protobuf:"varint,3,opt,name=Age,json=age" json:"Age,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "user.User")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 109 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x5c, 0xb8, 0x58, 0x42, 0x8b,
	0x53, 0x8b, 0x84, 0xc4, 0xb8, 0xd8, 0x40, 0xb4, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x73,
	0x10, 0x5b, 0x29, 0x98, 0x27, 0x24, 0xc4, 0xc5, 0xe2, 0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa4, 0xc0,
	0xa8, 0xc1, 0x19, 0xc4, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0x24, 0xc0, 0xc5, 0xec, 0x98, 0x9e, 0x2a,
	0xc1, 0xac, 0xc0, 0xa8, 0xc1, 0x1a, 0xc4, 0x9c, 0x98, 0x9e, 0x9a, 0xc4, 0x06, 0x36, 0xd2, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x21, 0x4d, 0xd5, 0xd5, 0x60, 0x00, 0x00, 0x00,
}
