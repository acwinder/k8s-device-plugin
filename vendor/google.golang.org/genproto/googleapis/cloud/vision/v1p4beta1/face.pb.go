// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/vision/v1p4beta1/face.proto

package vision

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Parameters for a celebrity recognition request.
type FaceRecognitionParams struct {
	// The resource names for one or more
	// [CelebritySet][google.cloud.vision.v1p4beta1.CelebritySet]s. A celebrity
	// set is preloaded and can be specified as "builtin/default". If this is
	// specified, the algorithm will try to match the faces detected in the input
	// image to the Celebrities in the CelebritySets.
	CelebritySet         []string `protobuf:"bytes,1,rep,name=celebrity_set,json=celebritySet,proto3" json:"celebrity_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceRecognitionParams) Reset()         { *m = FaceRecognitionParams{} }
func (m *FaceRecognitionParams) String() string { return proto.CompactTextString(m) }
func (*FaceRecognitionParams) ProtoMessage()    {}
func (*FaceRecognitionParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d91499545f74a28, []int{0}
}

func (m *FaceRecognitionParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceRecognitionParams.Unmarshal(m, b)
}
func (m *FaceRecognitionParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceRecognitionParams.Marshal(b, m, deterministic)
}
func (m *FaceRecognitionParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceRecognitionParams.Merge(m, src)
}
func (m *FaceRecognitionParams) XXX_Size() int {
	return xxx_messageInfo_FaceRecognitionParams.Size(m)
}
func (m *FaceRecognitionParams) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceRecognitionParams.DiscardUnknown(m)
}

var xxx_messageInfo_FaceRecognitionParams proto.InternalMessageInfo

func (m *FaceRecognitionParams) GetCelebritySet() []string {
	if m != nil {
		return m.CelebritySet
	}
	return nil
}

// A Celebrity is a group of Faces with an identity.
type Celebrity struct {
	// The resource name of the preloaded Celebrity. Has the format
	// `builtin/{mid}`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Celebrity's display name.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// The Celebrity's description.
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Celebrity) Reset()         { *m = Celebrity{} }
func (m *Celebrity) String() string { return proto.CompactTextString(m) }
func (*Celebrity) ProtoMessage()    {}
func (*Celebrity) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d91499545f74a28, []int{1}
}

func (m *Celebrity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Celebrity.Unmarshal(m, b)
}
func (m *Celebrity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Celebrity.Marshal(b, m, deterministic)
}
func (m *Celebrity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Celebrity.Merge(m, src)
}
func (m *Celebrity) XXX_Size() int {
	return xxx_messageInfo_Celebrity.Size(m)
}
func (m *Celebrity) XXX_DiscardUnknown() {
	xxx_messageInfo_Celebrity.DiscardUnknown(m)
}

var xxx_messageInfo_Celebrity proto.InternalMessageInfo

func (m *Celebrity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Celebrity) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Celebrity) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// Information about a face's identity.
type FaceRecognitionResult struct {
	// The [Celebrity][google.cloud.vision.v1p4beta1.Celebrity] that this face was
	// matched to.
	Celebrity *Celebrity `protobuf:"bytes,1,opt,name=celebrity,proto3" json:"celebrity,omitempty"`
	// Recognition confidence. Range [0, 1].
	Confidence           float32  `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceRecognitionResult) Reset()         { *m = FaceRecognitionResult{} }
func (m *FaceRecognitionResult) String() string { return proto.CompactTextString(m) }
func (*FaceRecognitionResult) ProtoMessage()    {}
func (*FaceRecognitionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d91499545f74a28, []int{2}
}

func (m *FaceRecognitionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceRecognitionResult.Unmarshal(m, b)
}
func (m *FaceRecognitionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceRecognitionResult.Marshal(b, m, deterministic)
}
func (m *FaceRecognitionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceRecognitionResult.Merge(m, src)
}
func (m *FaceRecognitionResult) XXX_Size() int {
	return xxx_messageInfo_FaceRecognitionResult.Size(m)
}
func (m *FaceRecognitionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceRecognitionResult.DiscardUnknown(m)
}

var xxx_messageInfo_FaceRecognitionResult proto.InternalMessageInfo

func (m *FaceRecognitionResult) GetCelebrity() *Celebrity {
	if m != nil {
		return m.Celebrity
	}
	return nil
}

func (m *FaceRecognitionResult) GetConfidence() float32 {
	if m != nil {
		return m.Confidence
	}
	return 0
}

func init() {
	proto.RegisterType((*FaceRecognitionParams)(nil), "google.cloud.vision.v1p4beta1.FaceRecognitionParams")
	proto.RegisterType((*Celebrity)(nil), "google.cloud.vision.v1p4beta1.Celebrity")
	proto.RegisterType((*FaceRecognitionResult)(nil), "google.cloud.vision.v1p4beta1.FaceRecognitionResult")
}

func init() {
	proto.RegisterFile("google/cloud/vision/v1p4beta1/face.proto", fileDescriptor_9d91499545f74a28)
}

var fileDescriptor_9d91499545f74a28 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x6b, 0xeb, 0x30,
	0x14, 0x85, 0x71, 0x12, 0x1e, 0x58, 0xc9, 0x7b, 0x83, 0xe0, 0x41, 0x08, 0xef, 0x95, 0x24, 0x5d,
	0x3c, 0x14, 0x89, 0xb4, 0xdd, 0xda, 0x29, 0x81, 0x74, 0x0b, 0xc1, 0x85, 0x0e, 0x5d, 0x82, 0x22,
	0xdf, 0x08, 0x81, 0xad, 0x6b, 0x2c, 0x25, 0x90, 0xa9, 0x4b, 0x7f, 0x49, 0x7f, 0x65, 0xc7, 0x62,
	0xc9, 0x75, 0x0a, 0xa5, 0xd9, 0xec, 0x73, 0xbf, 0x73, 0xee, 0xe1, 0x8a, 0x24, 0x0a, 0x51, 0xe5,
	0xc0, 0x65, 0x8e, 0xfb, 0x8c, 0x1f, 0xb4, 0xd5, 0x68, 0xf8, 0x61, 0x56, 0xde, 0x6e, 0xc1, 0x89,
	0x19, 0xdf, 0x09, 0x09, 0xac, 0xac, 0xd0, 0x21, 0xfd, 0x1f, 0x48, 0xe6, 0x49, 0x16, 0x48, 0xd6,
	0x92, 0xa3, 0x7f, 0x4d, 0x90, 0x28, 0x35, 0x17, 0xc6, 0xa0, 0x13, 0x4e, 0xa3, 0xb1, 0xc1, 0x3c,
	0xba, 0x3a, 0xbf, 0x46, 0x01, 0x16, 0xe0, 0xaa, 0x63, 0xa0, 0xa7, 0xf7, 0xe4, 0xef, 0x52, 0x48,
	0x48, 0x41, 0xa2, 0x32, 0xba, 0xce, 0x59, 0x8b, 0x4a, 0x14, 0x96, 0x5e, 0x92, 0xdf, 0x12, 0x72,
	0xd8, 0x56, 0xda, 0x1d, 0x37, 0x16, 0xdc, 0x30, 0x1a, 0x77, 0x93, 0x38, 0x1d, 0xb4, 0xe2, 0x23,
	0xb8, 0x69, 0x46, 0xe2, 0xc5, 0xe7, 0x3f, 0xa5, 0xa4, 0x67, 0x44, 0x01, 0xc3, 0x68, 0x1c, 0x25,
	0x71, 0xea, 0xbf, 0xe9, 0x84, 0x0c, 0x32, 0x6d, 0xcb, 0x5c, 0x1c, 0x37, 0x7e, 0xd6, 0xf1, 0xb3,
	0x7e, 0xa3, 0xad, 0x6a, 0x64, 0x4c, 0xfa, 0x19, 0x58, 0x59, 0xe9, 0xb2, 0xde, 0x3e, 0xec, 0x36,
	0xc4, 0x49, 0x9a, 0xbe, 0x7c, 0xeb, 0x98, 0x82, 0xdd, 0xe7, 0x8e, 0x2e, 0x49, 0xdc, 0xd6, 0xf1,
	0x6b, 0xfb, 0xd7, 0x09, 0x3b, 0x7b, 0x3b, 0xd6, 0xd6, 0x4d, 0x4f, 0x56, 0x7a, 0x41, 0x88, 0x44,
	0xb3, 0xd3, 0x19, 0x18, 0x19, 0x3a, 0x76, 0xd2, 0x2f, 0xca, 0xfc, 0x35, 0x22, 0x13, 0x89, 0xc5,
	0xf9, 0xe8, 0xf9, 0x9f, 0x36, 0x7b, 0x5d, 0x9f, 0x76, 0x1d, 0x3d, 0x2f, 0x1a, 0x83, 0xc2, 0x5c,
	0x18, 0xc5, 0xb0, 0x52, 0x5c, 0x81, 0xf1, 0x87, 0xe7, 0x61, 0x24, 0x4a, 0x6d, 0x7f, 0x78, 0xa9,
	0xbb, 0x20, 0xbc, 0x47, 0xd1, 0x5b, 0xa7, 0xf7, 0xb0, 0x78, 0x5a, 0x6d, 0x7f, 0x79, 0xe7, 0xcd,
	0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x94, 0xd4, 0x7f, 0x49, 0x02, 0x00, 0x00,
}