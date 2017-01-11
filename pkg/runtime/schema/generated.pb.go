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
// source: k8s.io/apimachinery/pkg/runtime/schema/generated.proto
// DO NOT EDIT!

/*
	Package schema is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/apimachinery/pkg/runtime/schema/generated.proto

	It has these top-level messages:
*/
package schema

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

var fileDescriptorGenerated = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0xc9, 0xb6, 0x28, 0xd6,
	0xcb, 0xcc, 0xd7, 0xcf, 0x2e, 0x4d, 0x4a, 0x2d, 0xca, 0x4b, 0x2d, 0x49, 0x2d, 0xd6, 0x2f, 0xc8,
	0x4e, 0xd7, 0x2f, 0x2a, 0xcd, 0x2b, 0xc9, 0xcc, 0x4d, 0xd5, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d,
	0xd4, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x52, 0x81, 0xe8, 0xd2, 0x43, 0xe8, 0xd2, 0x2b, 0xc8, 0x4e, 0xd7, 0x83, 0xea, 0xd2, 0x83,
	0xe8, 0x92, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0x6b, 0x4e, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0x62, 0xa8,
	0x94, 0x21, 0x76, 0xa7, 0x94, 0x96, 0x64, 0xe6, 0xe8, 0x67, 0xe6, 0x95, 0x14, 0x97, 0x14, 0xa1,
	0xbb, 0xc3, 0x49, 0xe3, 0xc4, 0x43, 0x39, 0x86, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63,
	0x68, 0x78, 0x24, 0xc7, 0x78, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9,
	0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x06, 0x71, 0x0b, 0x20, 0x00, 0x00, 0xff, 0xff, 0x8d,
	0x74, 0x75, 0xa7, 0xe8, 0x00, 0x00, 0x00,
}
