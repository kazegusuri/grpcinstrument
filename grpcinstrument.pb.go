// Code generated by protoc-gen-go.
// source: grpcinstrument/grpcinstrument.proto
// DO NOT EDIT!

/*
Package grpcinstrument is a generated protocol buffer package.

It is generated from these files:
	grpcinstrument/grpcinstrument.proto

It has these top-level messages:
	Call
	Input
	Output
	Error
*/
package grpcinstrument

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Call struct {
	Service  string                    `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Method   string                    `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	Input    *Input                    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Output   *Output                   `protobuf:"bytes,4,opt,name=output" json:"output,omitempty"`
	Error    *Error                    `protobuf:"bytes,5,opt,name=error" json:"error,omitempty"`
	Duration *google_protobuf.Duration `protobuf:"bytes,6,opt,name=duration" json:"duration,omitempty"`
}

func (m *Call) Reset()                    { *m = Call{} }
func (m *Call) String() string            { return proto.CompactTextString(m) }
func (*Call) ProtoMessage()               {}
func (*Call) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Call) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Call) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *Call) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *Call) GetDuration() *google_protobuf.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

type Input struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Input) Reset()                    { *m = Input{} }
func (m *Input) String() string            { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()               {}
func (*Input) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Output struct {
	Type string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
}

func (m *Output) Reset()                    { *m = Output{} }
func (m *Output) String() string            { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()               {}
func (*Output) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Error struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*Call)(nil), "grpcinstrument.Call")
	proto.RegisterType((*Input)(nil), "grpcinstrument.Input")
	proto.RegisterType((*Output)(nil), "grpcinstrument.Output")
	proto.RegisterType((*Error)(nil), "grpcinstrument.Error")
}

var fileDescriptor0 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xcd, 0x4a, 0xc4, 0x40,
	0x10, 0x84, 0x89, 0x26, 0xa3, 0xb6, 0xb2, 0xc2, 0xc0, 0x2e, 0xa3, 0x07, 0x91, 0x28, 0x22, 0x08,
	0x13, 0xd0, 0x47, 0x50, 0x0f, 0x9e, 0x7c, 0x86, 0xec, 0xee, 0x18, 0x03, 0x49, 0x3a, 0xf4, 0xf4,
	0x08, 0xbe, 0xa1, 0x8f, 0xe5, 0xce, 0xcf, 0x1e, 0x12, 0xf6, 0x58, 0x5d, 0x1f, 0x55, 0x5d, 0x70,
	0xd7, 0xd0, 0xb8, 0x69, 0x07, 0xcb, 0xe4, 0x7a, 0x33, 0x70, 0x35, 0x95, 0x7a, 0x24, 0x64, 0x94,
	0x8b, 0xe9, 0xf5, 0xfa, 0xa6, 0x41, 0x6c, 0x3a, 0x53, 0x05, 0x77, 0xed, 0xbe, 0xaa, 0xad, 0xa3,
	0x9a, 0x5b, 0x1c, 0x22, 0x5f, 0xfe, 0x65, 0x90, 0xbf, 0xd6, 0x5d, 0x27, 0x2f, 0xe1, 0xc4, 0x1a,
	0xfa, 0x69, 0x37, 0x46, 0x65, 0xb7, 0xd9, 0xe3, 0x99, 0x5c, 0x80, 0xe8, 0x0d, 0x7f, 0xe3, 0x56,
	0x1d, 0x05, 0x7d, 0x0f, 0x45, 0x3b, 0x8c, 0x8e, 0xd5, 0xf1, 0x4e, 0x9e, 0x3f, 0x2f, 0xf5, 0xac,
	0xff, 0xc3, 0x9b, 0xf2, 0x01, 0x04, 0x3a, 0xf6, 0x58, 0x1e, 0xb0, 0xd5, 0x1c, 0xfb, 0x0c, 0xae,
	0x4f, 0x33, 0x44, 0x48, 0xaa, 0x38, 0x9c, 0xf6, 0xee, 0x4d, 0xf9, 0x04, 0xa7, 0xfb, 0x7f, 0x95,
	0x08, 0xe0, 0x95, 0x8e, 0x83, 0xf4, 0x7e, 0x90, 0x7e, 0x4b, 0x40, 0xb9, 0x84, 0x22, 0xfe, 0x70,
	0x01, 0x39, 0xff, 0x8e, 0x69, 0x47, 0xb9, 0x02, 0x91, 0x3a, 0xa7, 0x77, 0x05, 0x45, 0x2c, 0xd9,
	0x2d, 0xef, 0x8d, 0xb5, 0x75, 0x93, 0x9c, 0xb5, 0x08, 0xd9, 0x2f, 0xff, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x5d, 0x61, 0xf0, 0x71, 0x01, 0x00, 0x00,
}
