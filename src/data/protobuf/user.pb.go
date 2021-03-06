// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package entity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type User struct {
	Id               *int64    `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name             *string   `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Gender           *int32    `protobuf:"varint,3,req,name=gender" json:"gender,omitempty"`
	Desc             *string   `protobuf:"bytes,4,opt,name=desc" json:"desc,omitempty"`
	Degree           *float32  `protobuf:"fixed32,5,opt,name=degree" json:"degree,omitempty"`
	Flags            []int32   `protobuf:"varint,6,rep,name=flags" json:"flags,omitempty"`
	Platform         *PLATFORM `protobuf:"varint,7,req,name=platform,enum=entity.PLATFORM" json:"platform,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor4, []int{0}
}

func (m *User) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetGender() int32 {
	if m != nil && m.Gender != nil {
		return *m.Gender
	}
	return 0
}

func (m *User) GetDesc() string {
	if m != nil && m.Desc != nil {
		return *m.Desc
	}
	return ""
}

func (m *User) GetDegree() float32 {
	if m != nil && m.Degree != nil {
		return *m.Degree
	}
	return 0
}

func (m *User) GetFlags() []int32 {
	if m != nil {
		return m.Flags
	}
	return nil
}

func (m *User) GetPlatform() PLATFORM {
	if m != nil && m.Platform != nil {
		return *m.Platform
	}
	return PLATFORM_PHONE
}

func init() {
	proto.RegisterType((*User)(nil), "entity.User")
}

func init() {
	proto.RegisterFile("user.proto", fileDescriptor4)
}

var fileDescriptor4 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0xc6, 0x71, 0x72, 0xda, 0x54, 0x3d, 0x60, 0x91, 0x20, 0x12, 0x9c, 0x82, 0x53, 0x06, 0xe9,
	0xe0, 0x1b, 0xb8, 0x38, 0x29, 0x4a, 0xd0, 0x07, 0x28, 0xcd, 0x69, 0x08, 0xb4, 0x49, 0x49, 0xe2,
	0xe0, 0x6b, 0xf9, 0x84, 0x97, 0xdc, 0x96, 0xbb, 0x7d, 0xff, 0xc3, 0x6f, 0x38, 0x88, 0xbf, 0x99,
	0xd2, 0xb0, 0xa5, 0x58, 0xa2, 0xe8, 0x28, 0x14, 0x5f, 0xfe, 0x1e, 0x6f, 0xa7, 0x18, 0x66, 0xef,
	0xf2, 0x7e, 0x7e, 0xfa, 0x67, 0xd8, 0xfe, 0x64, 0x4a, 0xa2, 0x47, 0xf0, 0x56, 0x32, 0x05, 0xba,
	0x31, 0xe0, 0xad, 0x10, 0xd8, 0x86, 0x71, 0x25, 0x09, 0x0a, 0xf4, 0x8d, 0x39, 0x6f, 0xf1, 0x80,
	0x9d, 0xa3, 0x60, 0x29, 0xc9, 0x46, 0x81, 0xe6, 0xe6, 0xa8, 0x6a, 0x2d, 0xe5, 0x49, 0xb6, 0x8a,
	0x55, 0x5b, 0x77, 0xb5, 0x96, 0x5c, 0x22, 0x92, 0x5c, 0x31, 0x0d, 0xe6, 0x28, 0x71, 0x8f, 0x7c,
	0x5e, 0x46, 0x97, 0x65, 0xa7, 0x1a, 0xcd, 0xcd, 0x1e, 0xe2, 0x19, 0xaf, 0xb7, 0x65, 0x2c, 0x73,
	0x4c, 0xab, 0xbc, 0x52, 0xa0, 0xfb, 0x97, 0xbb, 0x61, 0x7f, 0x78, 0xf8, 0x7a, 0x7f, 0xfd, 0x7e,
	0xfb, 0x34, 0x1f, 0xe6, 0x22, 0x4e, 0x01, 0x00, 0x00, 0xff, 0xff, 0x87, 0x64, 0xac, 0xeb, 0xd8,
	0x00, 0x00, 0x00,
}
