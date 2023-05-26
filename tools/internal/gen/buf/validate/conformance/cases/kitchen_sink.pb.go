// Copyright 2023 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: buf/validate/conformance/cases/kitchen_sink.proto

package cases

import (
	_ "github.com/bufbuild/protovalidate/tools/internal/gen/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ComplexTestEnum int32

const (
	ComplexTestEnum_COMPLEX_TEST_ENUM_UNSPECIFIED ComplexTestEnum = 0
	ComplexTestEnum_COMPLEX_TEST_ENUM_ONE         ComplexTestEnum = 1
	ComplexTestEnum_COMPLEX_TEST_ENUM_TWO         ComplexTestEnum = 2
)

// Enum value maps for ComplexTestEnum.
var (
	ComplexTestEnum_name = map[int32]string{
		0: "COMPLEX_TEST_ENUM_UNSPECIFIED",
		1: "COMPLEX_TEST_ENUM_ONE",
		2: "COMPLEX_TEST_ENUM_TWO",
	}
	ComplexTestEnum_value = map[string]int32{
		"COMPLEX_TEST_ENUM_UNSPECIFIED": 0,
		"COMPLEX_TEST_ENUM_ONE":         1,
		"COMPLEX_TEST_ENUM_TWO":         2,
	}
)

func (x ComplexTestEnum) Enum() *ComplexTestEnum {
	p := new(ComplexTestEnum)
	*p = x
	return p
}

func (x ComplexTestEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ComplexTestEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes[0].Descriptor()
}

func (ComplexTestEnum) Type() protoreflect.EnumType {
	return &file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes[0]
}

func (x ComplexTestEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ComplexTestEnum.Descriptor instead.
func (ComplexTestEnum) EnumDescriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescGZIP(), []int{0}
}

type ComplexTestMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Const      string                   `protobuf:"bytes,1,opt,name=const,proto3" json:"const,omitempty"`
	Nested     *ComplexTestMsg          `protobuf:"bytes,2,opt,name=nested,proto3" json:"nested,omitempty"`
	IntConst   int32                    `protobuf:"varint,3,opt,name=int_const,json=intConst,proto3" json:"int_const,omitempty"`
	BoolConst  bool                     `protobuf:"varint,4,opt,name=bool_const,json=boolConst,proto3" json:"bool_const,omitempty"`
	FloatVal   *wrapperspb.FloatValue   `protobuf:"bytes,5,opt,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	DurVal     *durationpb.Duration     `protobuf:"bytes,6,opt,name=dur_val,json=durVal,proto3" json:"dur_val,omitempty"`
	TsVal      *timestamppb.Timestamp   `protobuf:"bytes,7,opt,name=ts_val,json=tsVal,proto3" json:"ts_val,omitempty"`
	Another    *ComplexTestMsg          `protobuf:"bytes,8,opt,name=another,proto3" json:"another,omitempty"`
	FloatConst float32                  `protobuf:"fixed32,9,opt,name=float_const,json=floatConst,proto3" json:"float_const,omitempty"`
	DoubleIn   float64                  `protobuf:"fixed64,10,opt,name=double_in,json=doubleIn,proto3" json:"double_in,omitempty"`
	EnumConst  ComplexTestEnum          `protobuf:"varint,11,opt,name=enum_const,json=enumConst,proto3,enum=buf.validate.conformance.cases.ComplexTestEnum" json:"enum_const,omitempty"`
	AnyVal     *anypb.Any               `protobuf:"bytes,12,opt,name=any_val,json=anyVal,proto3" json:"any_val,omitempty"`
	RepTsVal   []*timestamppb.Timestamp `protobuf:"bytes,13,rep,name=rep_ts_val,json=repTsVal,proto3" json:"rep_ts_val,omitempty"`
	MapVal     map[int32]string         `protobuf:"bytes,14,rep,name=map_val,json=mapVal,proto3" json:"map_val,omitempty" protobuf_key:"zigzag32,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	BytesVal   []byte                   `protobuf:"bytes,15,opt,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
	// Types that are assignable to O:
	//
	//	*ComplexTestMsg_X
	//	*ComplexTestMsg_Y
	O isComplexTestMsg_O `protobuf_oneof:"o"`
}

func (x *ComplexTestMsg) Reset() {
	*x = ComplexTestMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComplexTestMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComplexTestMsg) ProtoMessage() {}

func (x *ComplexTestMsg) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComplexTestMsg.ProtoReflect.Descriptor instead.
func (*ComplexTestMsg) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescGZIP(), []int{0}
}

func (x *ComplexTestMsg) GetConst() string {
	if x != nil {
		return x.Const
	}
	return ""
}

func (x *ComplexTestMsg) GetNested() *ComplexTestMsg {
	if x != nil {
		return x.Nested
	}
	return nil
}

func (x *ComplexTestMsg) GetIntConst() int32 {
	if x != nil {
		return x.IntConst
	}
	return 0
}

func (x *ComplexTestMsg) GetBoolConst() bool {
	if x != nil {
		return x.BoolConst
	}
	return false
}

func (x *ComplexTestMsg) GetFloatVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *ComplexTestMsg) GetDurVal() *durationpb.Duration {
	if x != nil {
		return x.DurVal
	}
	return nil
}

func (x *ComplexTestMsg) GetTsVal() *timestamppb.Timestamp {
	if x != nil {
		return x.TsVal
	}
	return nil
}

func (x *ComplexTestMsg) GetAnother() *ComplexTestMsg {
	if x != nil {
		return x.Another
	}
	return nil
}

func (x *ComplexTestMsg) GetFloatConst() float32 {
	if x != nil {
		return x.FloatConst
	}
	return 0
}

func (x *ComplexTestMsg) GetDoubleIn() float64 {
	if x != nil {
		return x.DoubleIn
	}
	return 0
}

func (x *ComplexTestMsg) GetEnumConst() ComplexTestEnum {
	if x != nil {
		return x.EnumConst
	}
	return ComplexTestEnum_COMPLEX_TEST_ENUM_UNSPECIFIED
}

func (x *ComplexTestMsg) GetAnyVal() *anypb.Any {
	if x != nil {
		return x.AnyVal
	}
	return nil
}

func (x *ComplexTestMsg) GetRepTsVal() []*timestamppb.Timestamp {
	if x != nil {
		return x.RepTsVal
	}
	return nil
}

func (x *ComplexTestMsg) GetMapVal() map[int32]string {
	if x != nil {
		return x.MapVal
	}
	return nil
}

func (x *ComplexTestMsg) GetBytesVal() []byte {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

func (m *ComplexTestMsg) GetO() isComplexTestMsg_O {
	if m != nil {
		return m.O
	}
	return nil
}

func (x *ComplexTestMsg) GetX() string {
	if x, ok := x.GetO().(*ComplexTestMsg_X); ok {
		return x.X
	}
	return ""
}

func (x *ComplexTestMsg) GetY() int32 {
	if x, ok := x.GetO().(*ComplexTestMsg_Y); ok {
		return x.Y
	}
	return 0
}

type isComplexTestMsg_O interface {
	isComplexTestMsg_O()
}

type ComplexTestMsg_X struct {
	X string `protobuf:"bytes,16,opt,name=x,proto3,oneof"`
}

type ComplexTestMsg_Y struct {
	Y int32 `protobuf:"varint,17,opt,name=y,proto3,oneof"`
}

func (*ComplexTestMsg_X) isComplexTestMsg_O() {}

func (*ComplexTestMsg_Y) isComplexTestMsg_O() {}

type KitchenSinkMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *ComplexTestMsg `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *KitchenSinkMessage) Reset() {
	*x = KitchenSinkMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KitchenSinkMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KitchenSinkMessage) ProtoMessage() {}

func (x *KitchenSinkMessage) ProtoReflect() protoreflect.Message {
	mi := &file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KitchenSinkMessage.ProtoReflect.Descriptor instead.
func (*KitchenSinkMessage) Descriptor() ([]byte, []int) {
	return file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescGZIP(), []int{1}
}

func (x *KitchenSinkMessage) GetVal() *ComplexTestMsg {
	if x != nil {
		return x.Val
	}
	return nil
}

var File_buf_validate_conformance_cases_kitchen_sink_proto protoreflect.FileDescriptor

var file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc = []byte{
	0x0a, 0x31, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73,
	0x2f, 0x6b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x5f, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61,
	0x73, 0x65, 0x73, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xca, 0x08, 0x0a,
	0x0e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x12,
	0x22, 0x0a, 0x05, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xfa, 0xf7, 0x18, 0x08, 0x72, 0x06, 0x0a, 0x04, 0x61, 0x62, 0x63, 0x64, 0x52, 0x05, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74,
	0x4d, 0x73, 0x67, 0x52, 0x06, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x25, 0x0a, 0x09, 0x69,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x08,
	0xfa, 0xf7, 0x18, 0x04, 0x1a, 0x02, 0x08, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x12, 0x27, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x6a, 0x02, 0x08, 0x00,
	0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x45, 0x0a, 0x09, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x0b, 0xfa, 0xf7, 0x18,
	0x07, 0x0a, 0x05, 0x25, 0x00, 0x00, 0x00, 0x00, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56,
	0x61, 0x6c, 0x12, 0x42, 0x0a, 0x07, 0x64, 0x75, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0e,
	0xfa, 0xf7, 0x18, 0x0a, 0xc8, 0x01, 0x01, 0xaa, 0x01, 0x04, 0x1a, 0x02, 0x08, 0x11, 0x52, 0x06,
	0x64, 0x75, 0x72, 0x56, 0x61, 0x6c, 0x12, 0x3e, 0x0a, 0x06, 0x74, 0x73, 0x5f, 0x76, 0x61, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x42, 0x0b, 0xfa, 0xf7, 0x18, 0x07, 0xb2, 0x01, 0x04, 0x2a, 0x02, 0x08, 0x07, 0x52,
	0x05, 0x74, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x48, 0x0a, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65,
	0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78,
	0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x07, 0x61, 0x6e, 0x6f, 0x74, 0x68, 0x65, 0x72,
	0x12, 0x2c, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0b, 0xfa, 0xf7, 0x18, 0x07, 0x0a, 0x05, 0x15, 0x00, 0x00,
	0x00, 0x41, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x12, 0x35,
	0x0a, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x01, 0x42, 0x18, 0xfa, 0xf7, 0x18, 0x14, 0x12, 0x12, 0x32, 0x10, 0xb4, 0xc8, 0x76, 0xbe, 0x9f,
	0x8c, 0x7c, 0x40, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x5e, 0x40, 0x52, 0x08, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x49, 0x6e, 0x12, 0x59, 0x0a, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x6f,
	0x6e, 0x73, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x62, 0x75, 0x66, 0x2e,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x42, 0x09, 0xfa, 0xf7, 0x18, 0x05,
	0x82, 0x01, 0x02, 0x08, 0x02, 0x52, 0x09, 0x65, 0x6e, 0x75, 0x6d, 0x43, 0x6f, 0x6e, 0x73, 0x74,
	0x12, 0x64, 0x0a, 0x07, 0x61, 0x6e, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x42, 0x35, 0xfa, 0xf7, 0x18, 0x31, 0xa2, 0x01, 0x2e,
	0x12, 0x2c, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06,
	0x61, 0x6e, 0x79, 0x56, 0x61, 0x6c, 0x12, 0x4c, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x5f, 0x74, 0x73,
	0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x12, 0xfa, 0xf7, 0x18, 0x0e, 0x92, 0x01, 0x0b, 0x22,
	0x09, 0xb2, 0x01, 0x06, 0x32, 0x04, 0x10, 0xc0, 0x84, 0x3d, 0x52, 0x08, 0x72, 0x65, 0x70, 0x54,
	0x73, 0x56, 0x61, 0x6c, 0x12, 0x62, 0x0a, 0x07, 0x6d, 0x61, 0x70, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x73, 0x67, 0x2e, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x42, 0x0d, 0xfa, 0xf7, 0x18, 0x09, 0x9a, 0x01, 0x06, 0x22, 0x04, 0x3a, 0x02, 0x10, 0x00,
	0x52, 0x06, 0x6d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x12, 0x27, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x0a, 0xfa, 0xf7, 0x18,
	0x06, 0x7a, 0x04, 0x0a, 0x02, 0x00, 0x99, 0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61,
	0x6c, 0x12, 0x0e, 0x0a, 0x01, 0x78, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x01,
	0x78, 0x12, 0x0e, 0x0a, 0x01, 0x79, 0x18, 0x11, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x01,
	0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4d, 0x61, 0x70, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x01,
	0x6f, 0x12, 0x06, 0xfa, 0xf7, 0x18, 0x02, 0x08, 0x01, 0x22, 0x56, 0x0a, 0x12, 0x4b, 0x69, 0x74,
	0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x40, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x62,
	0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x61, 0x73, 0x65, 0x73, 0x2e, 0x43, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x73, 0x67, 0x52, 0x03, 0x76, 0x61,
	0x6c, 0x2a, 0x6a, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x78, 0x54, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c,
	0x45, 0x58, 0x5f, 0x54, 0x45, 0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x4f, 0x4e, 0x45,
	0x10, 0x01, 0x12, 0x19, 0x0a, 0x15, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x58, 0x5f, 0x54, 0x45,
	0x53, 0x54, 0x5f, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x42, 0xa7, 0x02,
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x61, 0x73, 0x65, 0x73, 0x42, 0x10, 0x4b, 0x69, 0x74, 0x63, 0x68, 0x65, 0x6e, 0x53, 0x69, 0x6e,
	0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x61, 0x73, 0x65, 0x73, 0xa2, 0x02, 0x04,
	0x42, 0x56, 0x43, 0x43, 0xaa, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x43, 0x61, 0x73, 0x65, 0x73, 0xca, 0x02, 0x1e, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0xe2, 0x02, 0x2a, 0x42, 0x75, 0x66, 0x5c, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63,
	0x65, 0x5c, 0x43, 0x61, 0x73, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x21, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x3a, 0x3a, 0x43, 0x61, 0x73, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescOnce sync.Once
	file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescData = file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc
)

func file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescGZIP() []byte {
	file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescOnce.Do(func() {
		file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescData)
	})
	return file_buf_validate_conformance_cases_kitchen_sink_proto_rawDescData
}

var file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes = []interface{}{
	(ComplexTestEnum)(0),          // 0: buf.validate.conformance.cases.ComplexTestEnum
	(*ComplexTestMsg)(nil),        // 1: buf.validate.conformance.cases.ComplexTestMsg
	(*KitchenSinkMessage)(nil),    // 2: buf.validate.conformance.cases.KitchenSinkMessage
	nil,                           // 3: buf.validate.conformance.cases.ComplexTestMsg.MapValEntry
	(*wrapperspb.FloatValue)(nil), // 4: google.protobuf.FloatValue
	(*durationpb.Duration)(nil),   // 5: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 7: google.protobuf.Any
}
var file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs = []int32{
	1,  // 0: buf.validate.conformance.cases.ComplexTestMsg.nested:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	4,  // 1: buf.validate.conformance.cases.ComplexTestMsg.float_val:type_name -> google.protobuf.FloatValue
	5,  // 2: buf.validate.conformance.cases.ComplexTestMsg.dur_val:type_name -> google.protobuf.Duration
	6,  // 3: buf.validate.conformance.cases.ComplexTestMsg.ts_val:type_name -> google.protobuf.Timestamp
	1,  // 4: buf.validate.conformance.cases.ComplexTestMsg.another:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	0,  // 5: buf.validate.conformance.cases.ComplexTestMsg.enum_const:type_name -> buf.validate.conformance.cases.ComplexTestEnum
	7,  // 6: buf.validate.conformance.cases.ComplexTestMsg.any_val:type_name -> google.protobuf.Any
	6,  // 7: buf.validate.conformance.cases.ComplexTestMsg.rep_ts_val:type_name -> google.protobuf.Timestamp
	3,  // 8: buf.validate.conformance.cases.ComplexTestMsg.map_val:type_name -> buf.validate.conformance.cases.ComplexTestMsg.MapValEntry
	1,  // 9: buf.validate.conformance.cases.KitchenSinkMessage.val:type_name -> buf.validate.conformance.cases.ComplexTestMsg
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_buf_validate_conformance_cases_kitchen_sink_proto_init() }
func file_buf_validate_conformance_cases_kitchen_sink_proto_init() {
	if File_buf_validate_conformance_cases_kitchen_sink_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComplexTestMsg); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KitchenSinkMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ComplexTestMsg_X)(nil),
		(*ComplexTestMsg_Y)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes,
		DependencyIndexes: file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs,
		EnumInfos:         file_buf_validate_conformance_cases_kitchen_sink_proto_enumTypes,
		MessageInfos:      file_buf_validate_conformance_cases_kitchen_sink_proto_msgTypes,
	}.Build()
	File_buf_validate_conformance_cases_kitchen_sink_proto = out.File
	file_buf_validate_conformance_cases_kitchen_sink_proto_rawDesc = nil
	file_buf_validate_conformance_cases_kitchen_sink_proto_goTypes = nil
	file_buf_validate_conformance_cases_kitchen_sink_proto_depIdxs = nil
}