/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/apimachinery/pkg/api/resource/generated.proto
// DO NOT EDIT!

/*
	Package resource is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/apimachinery/pkg/api/resource/generated.proto

	It has these top-level messages:
		Quantity
*/
package resource

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *Quantity) Reset()                    { *m = Quantity{} }
func (*Quantity) ProtoMessage()               {}
func (*Quantity) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func init() {
	proto.RegisterType((*Quantity)(nil), "k8s.io.kubernetes.pkg.api.resource.Quantity")
}

var fileDescriptorGenerated = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xa1, 0x4e, 0x03, 0x41,
	0x10, 0x86, 0x67, 0x0d, 0x29, 0x27, 0x1b, 0x42, 0x48, 0xc5, 0x5e, 0x53, 0x45, 0x48, 0xd8, 0x09,
	0xa8, 0x06, 0xc9, 0x1b, 0x80, 0xc4, 0xdd, 0x95, 0x61, 0x99, 0x1c, 0xec, 0x6e, 0x76, 0x67, 0x05,
	0xae, 0x12, 0x59, 0x89, 0xec, 0xbd, 0x4d, 0x65, 0x25, 0x02, 0xc1, 0x1d, 0x2f, 0x42, 0x72, 0xa5,
	0x21, 0x21, 0xb8, 0xf9, 0xc4, 0x37, 0xf9, 0xfe, 0xe2, 0xb2, 0x99, 0x27, 0xc3, 0x1e, 0x9b, 0x5c,
	0x53, 0x74, 0x24, 0x94, 0x30, 0x34, 0x16, 0xab, 0xc0, 0x18, 0x29, 0xf9, 0x1c, 0x17, 0x84, 0x96,
	0x1c, 0xc5, 0x4a, 0xe8, 0xde, 0x84, 0xe8, 0xc5, 0x8f, 0x67, 0x3b, 0xc7, 0xfc, 0x3a, 0x26, 0x34,
	0xd6, 0x54, 0x81, 0xcd, 0xde, 0x99, 0x9c, 0x5b, 0x96, 0xc7, 0x5c, 0x9b, 0x85, 0x7f, 0x46, 0xeb,
	0xad, 0xc7, 0x41, 0xad, 0xf3, 0xc3, 0x40, 0x03, 0x0c, 0xd7, 0xee, 0xe5, 0xe4, 0xe2, 0xff, 0x8c,
	0x2c, 0xfc, 0x84, 0xec, 0x24, 0x49, 0xfc, 0x5b, 0x31, 0x9b, 0x17, 0xa3, 0x9b, 0x5c, 0x39, 0x61,
	0x79, 0x19, 0x1f, 0x17, 0x07, 0x49, 0x22, 0x3b, 0x7b, 0xa2, 0xa6, 0xea, 0xf4, 0xf0, 0xf6, 0x87,
	0xae, 0x8e, 0xde, 0xd6, 0x25, 0xbc, 0xb6, 0x25, 0xac, 0xda, 0x12, 0xd6, 0x6d, 0x09, 0xcb, 0x8f,
	0x29, 0x5c, 0x9f, 0x6d, 0x3a, 0x0d, 0xdb, 0x4e, 0xc3, 0x7b, 0xa7, 0x61, 0xd9, 0x6b, 0xb5, 0xe9,
	0xb5, 0xda, 0xf6, 0x5a, 0x7d, 0xf6, 0x5a, 0xad, 0xbe, 0x34, 0xdc, 0x8d, 0xf6, 0x3b, 0xbe, 0x03,
	0x00, 0x00, 0xff, 0xff, 0x90, 0x1c, 0x7f, 0xff, 0x20, 0x01, 0x00, 0x00,
}
