// Code generated by protoc-gen-go.
// source: user.proto
// DO NOT EDIT!

package models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Uuid             *string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Email            *string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
	HashedPassword   *string `protobuf:"bytes,3,opt,name=hashed_password,json=hashedPassword" json:"hashed_password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *User) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil && m.Email != nil {
		return *m.Email
	}
	return ""
}

func (m *User) GetHashedPassword() string {
	if m != nil && m.HashedPassword != nil {
		return *m.HashedPassword
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "press.models.User")
}
func (m *User) SetUuid(v *string) {
	m.Uuid = v
}

func (m *User) SetEmail(v *string) {
	m.Email = v
}

func (m *User) SetHashedPassword(v *string) {
	m.HashedPassword = v
}

func init() { proto.RegisterFile("user.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x29, 0x28, 0x4a, 0x2d, 0x2e, 0xd6, 0xcb, 0xcd,
	0x4f, 0x49, 0xcd, 0x29, 0x56, 0x8a, 0xe4, 0x62, 0x09, 0x2d, 0x4e, 0x2d, 0x12, 0x12, 0xe2, 0x62,
	0x29, 0x2d, 0xcd, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8,
	0x58, 0x53, 0x73, 0x13, 0x33, 0x73, 0x24, 0x98, 0xc0, 0x82, 0x10, 0x8e, 0x90, 0x3a, 0x17, 0x7f,
	0x46, 0x62, 0x71, 0x46, 0x6a, 0x4a, 0x7c, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51, 0x8a, 0x04,
	0x33, 0x58, 0x9e, 0x0f, 0x22, 0x1c, 0x00, 0x15, 0x05, 0x04, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xca,
	0x07, 0x92, 0x75, 0x00, 0x00, 0x00,
}
