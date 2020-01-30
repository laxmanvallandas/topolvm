// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lvmd/proto/lvmd.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

// Represents a logical volume.
type LogicalVolume struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	DevMajor             uint32   `protobuf:"varint,3,opt,name=dev_major,json=devMajor,proto3" json:"dev_major,omitempty"`
	DevMinor             uint32   `protobuf:"varint,4,opt,name=dev_minor,json=devMinor,proto3" json:"dev_minor,omitempty"`
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogicalVolume) Reset()         { *m = LogicalVolume{} }
func (m *LogicalVolume) String() string { return proto.CompactTextString(m) }
func (*LogicalVolume) ProtoMessage()    {}
func (*LogicalVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{1}
}

func (m *LogicalVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalVolume.Unmarshal(m, b)
}
func (m *LogicalVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalVolume.Marshal(b, m, deterministic)
}
func (m *LogicalVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalVolume.Merge(m, src)
}
func (m *LogicalVolume) XXX_Size() int {
	return xxx_messageInfo_LogicalVolume.Size(m)
}
func (m *LogicalVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalVolume.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalVolume proto.InternalMessageInfo

func (m *LogicalVolume) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LogicalVolume) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

func (m *LogicalVolume) GetDevMajor() uint32 {
	if m != nil {
		return m.DevMajor
	}
	return 0
}

func (m *LogicalVolume) GetDevMinor() uint32 {
	if m != nil {
		return m.DevMinor
	}
	return 0
}

func (m *LogicalVolume) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Represents the input for CreateLV.
type CreateLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateLVRequest) Reset()         { *m = CreateLVRequest{} }
func (m *CreateLVRequest) String() string { return proto.CompactTextString(m) }
func (*CreateLVRequest) ProtoMessage()    {}
func (*CreateLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{2}
}

func (m *CreateLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLVRequest.Unmarshal(m, b)
}
func (m *CreateLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLVRequest.Marshal(b, m, deterministic)
}
func (m *CreateLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLVRequest.Merge(m, src)
}
func (m *CreateLVRequest) XXX_Size() int {
	return xxx_messageInfo_CreateLVRequest.Size(m)
}
func (m *CreateLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLVRequest proto.InternalMessageInfo

func (m *CreateLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateLVRequest) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

func (m *CreateLVRequest) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Represents the response of CreateLV.
type CreateLVResponse struct {
	Volume               *LogicalVolume `protobuf:"bytes,1,opt,name=volume,proto3" json:"volume,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateLVResponse) Reset()         { *m = CreateLVResponse{} }
func (m *CreateLVResponse) String() string { return proto.CompactTextString(m) }
func (*CreateLVResponse) ProtoMessage()    {}
func (*CreateLVResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{3}
}

func (m *CreateLVResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateLVResponse.Unmarshal(m, b)
}
func (m *CreateLVResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateLVResponse.Marshal(b, m, deterministic)
}
func (m *CreateLVResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateLVResponse.Merge(m, src)
}
func (m *CreateLVResponse) XXX_Size() int {
	return xxx_messageInfo_CreateLVResponse.Size(m)
}
func (m *CreateLVResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateLVResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateLVResponse proto.InternalMessageInfo

func (m *CreateLVResponse) GetVolume() *LogicalVolume {
	if m != nil {
		return m.Volume
	}
	return nil
}

// Represents the input for RemoveLV.
type RemoveLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveLVRequest) Reset()         { *m = RemoveLVRequest{} }
func (m *RemoveLVRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveLVRequest) ProtoMessage()    {}
func (*RemoveLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{4}
}

func (m *RemoveLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveLVRequest.Unmarshal(m, b)
}
func (m *RemoveLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveLVRequest.Marshal(b, m, deterministic)
}
func (m *RemoveLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveLVRequest.Merge(m, src)
}
func (m *RemoveLVRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveLVRequest.Size(m)
}
func (m *RemoveLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveLVRequest proto.InternalMessageInfo

func (m *RemoveLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Represents the input for ResizeLV.
//
// The volume must already exist.
// The volume size will be set to exactly "size_gb".
type ResizeLVRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SizeGb               uint64   `protobuf:"varint,2,opt,name=size_gb,json=sizeGb,proto3" json:"size_gb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResizeLVRequest) Reset()         { *m = ResizeLVRequest{} }
func (m *ResizeLVRequest) String() string { return proto.CompactTextString(m) }
func (*ResizeLVRequest) ProtoMessage()    {}
func (*ResizeLVRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{5}
}

func (m *ResizeLVRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResizeLVRequest.Unmarshal(m, b)
}
func (m *ResizeLVRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResizeLVRequest.Marshal(b, m, deterministic)
}
func (m *ResizeLVRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResizeLVRequest.Merge(m, src)
}
func (m *ResizeLVRequest) XXX_Size() int {
	return xxx_messageInfo_ResizeLVRequest.Size(m)
}
func (m *ResizeLVRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResizeLVRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResizeLVRequest proto.InternalMessageInfo

func (m *ResizeLVRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ResizeLVRequest) GetSizeGb() uint64 {
	if m != nil {
		return m.SizeGb
	}
	return 0
}

// Represents the response of GetLVList.
type GetLVListResponse struct {
	Volumes              []*LogicalVolume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetLVListResponse) Reset()         { *m = GetLVListResponse{} }
func (m *GetLVListResponse) String() string { return proto.CompactTextString(m) }
func (*GetLVListResponse) ProtoMessage()    {}
func (*GetLVListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{6}
}

func (m *GetLVListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetLVListResponse.Unmarshal(m, b)
}
func (m *GetLVListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetLVListResponse.Marshal(b, m, deterministic)
}
func (m *GetLVListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetLVListResponse.Merge(m, src)
}
func (m *GetLVListResponse) XXX_Size() int {
	return xxx_messageInfo_GetLVListResponse.Size(m)
}
func (m *GetLVListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetLVListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetLVListResponse proto.InternalMessageInfo

func (m *GetLVListResponse) GetVolumes() []*LogicalVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

// Represents the response of GetFreeBytes.
type GetFreeBytesResponse struct {
	FreeBytes            uint64   `protobuf:"varint,1,opt,name=free_bytes,json=freeBytes,proto3" json:"free_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFreeBytesResponse) Reset()         { *m = GetFreeBytesResponse{} }
func (m *GetFreeBytesResponse) String() string { return proto.CompactTextString(m) }
func (*GetFreeBytesResponse) ProtoMessage()    {}
func (*GetFreeBytesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{7}
}

func (m *GetFreeBytesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFreeBytesResponse.Unmarshal(m, b)
}
func (m *GetFreeBytesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFreeBytesResponse.Marshal(b, m, deterministic)
}
func (m *GetFreeBytesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFreeBytesResponse.Merge(m, src)
}
func (m *GetFreeBytesResponse) XXX_Size() int {
	return xxx_messageInfo_GetFreeBytesResponse.Size(m)
}
func (m *GetFreeBytesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFreeBytesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFreeBytesResponse proto.InternalMessageInfo

func (m *GetFreeBytesResponse) GetFreeBytes() uint64 {
	if m != nil {
		return m.FreeBytes
	}
	return 0
}

// Represents the stream output from Watch.
type WatchResponse struct {
	FreeBytes            uint64   `protobuf:"varint,1,opt,name=free_bytes,json=freeBytes,proto3" json:"free_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb3c510e545f3bbd, []int{8}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetFreeBytes() uint64 {
	if m != nil {
		return m.FreeBytes
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "proto.Empty")
	proto.RegisterType((*LogicalVolume)(nil), "proto.LogicalVolume")
	proto.RegisterType((*CreateLVRequest)(nil), "proto.CreateLVRequest")
	proto.RegisterType((*CreateLVResponse)(nil), "proto.CreateLVResponse")
	proto.RegisterType((*RemoveLVRequest)(nil), "proto.RemoveLVRequest")
	proto.RegisterType((*ResizeLVRequest)(nil), "proto.ResizeLVRequest")
	proto.RegisterType((*GetLVListResponse)(nil), "proto.GetLVListResponse")
	proto.RegisterType((*GetFreeBytesResponse)(nil), "proto.GetFreeBytesResponse")
	proto.RegisterType((*WatchResponse)(nil), "proto.WatchResponse")
}

func init() { proto.RegisterFile("lvmd/proto/lvmd.proto", fileDescriptor_cb3c510e545f3bbd) }

var fileDescriptor_cb3c510e545f3bbd = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0x19, 0x93, 0xb4, 0xcd, 0x71, 0xcb, 0xea, 0xb0, 0xba, 0xa1, 0x8b, 0x10, 0x06, 0x84,
	0x5e, 0x68, 0x76, 0xe9, 0xe2, 0x85, 0x08, 0x22, 0x2e, 0x9a, 0x9b, 0x78, 0x33, 0x42, 0xbc, 0x2c,
	0x49, 0x7b, 0xb6, 0x46, 0x9a, 0x4e, 0xcd, 0xcc, 0x06, 0xea, 0x23, 0xf8, 0x28, 0xe2, 0x43, 0xca,
	0x4c, 0xfe, 0xb4, 0x49, 0x51, 0xd4, 0xab, 0x9e, 0x73, 0xe6, 0xfb, 0xce, 0x7c, 0xfd, 0x4d, 0xe0,
	0xd1, 0xba, 0xcc, 0x97, 0x97, 0xdb, 0x42, 0x28, 0x71, 0xa9, 0xcb, 0xc0, 0x94, 0xd4, 0x31, 0x3f,
	0x6c, 0x08, 0xce, 0xbb, 0x7c, 0xab, 0x76, 0xec, 0x3b, 0x81, 0x71, 0x24, 0x56, 0xd9, 0x22, 0x59,
	0xc7, 0x62, 0x7d, 0x97, 0x23, 0xa5, 0x60, 0x6f, 0x92, 0x1c, 0x3d, 0xe2, 0x93, 0xa9, 0xcb, 0x4d,
	0x4d, 0xcf, 0x61, 0x28, 0xb3, 0x6f, 0x38, 0x5f, 0xa5, 0xde, 0x3d, 0x9f, 0x4c, 0x6d, 0x3e, 0xd0,
	0x6d, 0x98, 0xd2, 0x0b, 0x70, 0x97, 0x58, 0xce, 0xf3, 0xe4, 0x8b, 0x28, 0x3c, 0xcb, 0x27, 0xd3,
	0x31, 0x1f, 0x2d, 0xb1, 0xfc, 0xa0, 0xfb, 0xf6, 0x30, 0xdb, 0x88, 0xc2, 0xb3, 0xf7, 0x87, 0xba,
	0xd7, 0xd7, 0xa8, 0x64, 0x25, 0x3d, 0xc7, 0xb7, 0xf4, 0x35, 0xba, 0x66, 0x1c, 0x4e, 0x6f, 0x0a,
	0x4c, 0x14, 0x46, 0x31, 0xc7, 0xaf, 0x77, 0x28, 0xd5, 0xbf, 0xa5, 0x69, 0x76, 0x5a, 0x07, 0x3b,
	0xdf, 0xc0, 0x83, 0xfd, 0x4e, 0xb9, 0x15, 0x1b, 0x89, 0xf4, 0x19, 0x0c, 0x4a, 0xf3, 0x67, 0xcd,
	0xda, 0xfb, 0xb3, 0xb3, 0x0a, 0x4e, 0xd0, 0x01, 0xc1, 0x6b, 0x0d, 0x7b, 0x0a, 0xa7, 0x1c, 0x73,
	0x51, 0xfe, 0x39, 0x15, 0x7b, 0xad, 0x65, 0x3a, 0xc8, 0xff, 0x85, 0x67, 0x37, 0xf0, 0x30, 0x44,
	0x15, 0xc5, 0x51, 0x26, 0x55, 0x9b, 0x34, 0x80, 0x61, 0x95, 0x42, 0x7a, 0xc4, 0xb7, 0x7e, 0x1b,
	0xb5, 0x11, 0xb1, 0x17, 0x70, 0x16, 0xa2, 0x7a, 0x5f, 0x20, 0xbe, 0xdd, 0x29, 0x94, 0xed, 0x9e,
	0x27, 0x00, 0xb7, 0x05, 0xe2, 0x3c, 0xd5, 0x53, 0x93, 0xc7, 0xe6, 0xee, 0x6d, 0x23, 0x63, 0x01,
	0x8c, 0x3f, 0x25, 0x6a, 0xf1, 0xf9, 0x2f, 0xf5, 0xb3, 0x9f, 0x04, 0xdc, 0x28, 0xfe, 0x88, 0x45,
	0x99, 0x2d, 0x90, 0xbe, 0x82, 0x51, 0x83, 0x98, 0x3e, 0xae, 0xf3, 0xf5, 0xde, 0x71, 0x72, 0x7e,
	0x34, 0xaf, 0x6f, 0xba, 0x82, 0x51, 0x43, 0xb7, 0x35, 0xf7, 0x70, 0x4f, 0x4e, 0xea, 0xb9, 0xf9,
	0x64, 0x2b, 0x47, 0x05, 0xfa, 0xc0, 0xd1, 0x21, 0xdf, 0x75, 0xcc, 0x7e, 0x10, 0x70, 0xe3, 0xb0,
	0x89, 0x7b, 0x0d, 0x6e, 0x0b, 0x9a, 0x76, 0x84, 0x13, 0xaf, 0xee, 0x8e, 0x1f, 0xe2, 0x25, 0x9c,
	0x1c, 0x82, 0xed, 0xf9, 0x2e, 0xf6, 0xbe, 0x63, 0xf6, 0xcf, 0xc1, 0x31, 0x70, 0x7b, 0x9e, 0xe6,
	0x25, 0x3b, 0xe0, 0xaf, 0x48, 0x3a, 0x30, 0xe3, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0xae, 0x9b, 0x5b, 0xc1, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LVServiceClient is the client API for LVService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LVServiceClient interface {
	// Create a logical volume.
	CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVResponse, error)
	// Remove a logical volume.
	RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*Empty, error)
	// Resize a logical volume.
	ResizeLV(ctx context.Context, in *ResizeLVRequest, opts ...grpc.CallOption) (*Empty, error)
}

type lVServiceClient struct {
	cc *grpc.ClientConn
}

func NewLVServiceClient(cc *grpc.ClientConn) LVServiceClient {
	return &lVServiceClient{cc}
}

func (c *lVServiceClient) CreateLV(ctx context.Context, in *CreateLVRequest, opts ...grpc.CallOption) (*CreateLVResponse, error) {
	out := new(CreateLVResponse)
	err := c.cc.Invoke(ctx, "/proto.LVService/CreateLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVServiceClient) RemoveLV(ctx context.Context, in *RemoveLVRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LVService/RemoveLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lVServiceClient) ResizeLV(ctx context.Context, in *ResizeLVRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/proto.LVService/ResizeLV", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LVServiceServer is the server API for LVService service.
type LVServiceServer interface {
	// Create a logical volume.
	CreateLV(context.Context, *CreateLVRequest) (*CreateLVResponse, error)
	// Remove a logical volume.
	RemoveLV(context.Context, *RemoveLVRequest) (*Empty, error)
	// Resize a logical volume.
	ResizeLV(context.Context, *ResizeLVRequest) (*Empty, error)
}

// UnimplementedLVServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLVServiceServer struct {
}

func (*UnimplementedLVServiceServer) CreateLV(ctx context.Context, req *CreateLVRequest) (*CreateLVResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLV not implemented")
}
func (*UnimplementedLVServiceServer) RemoveLV(ctx context.Context, req *RemoveLVRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveLV not implemented")
}
func (*UnimplementedLVServiceServer) ResizeLV(ctx context.Context, req *ResizeLVRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResizeLV not implemented")
}

func RegisterLVServiceServer(s *grpc.Server, srv LVServiceServer) {
	s.RegisterService(&_LVService_serviceDesc, srv)
}

func _LVService_CreateLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).CreateLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/CreateLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).CreateLV(ctx, req.(*CreateLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVService_RemoveLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).RemoveLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/RemoveLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).RemoveLV(ctx, req.(*RemoveLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LVService_ResizeLV_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResizeLVRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LVServiceServer).ResizeLV(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.LVService/ResizeLV",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LVServiceServer).ResizeLV(ctx, req.(*ResizeLVRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LVService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.LVService",
	HandlerType: (*LVServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLV",
			Handler:    _LVService_CreateLV_Handler,
		},
		{
			MethodName: "RemoveLV",
			Handler:    _LVService_RemoveLV_Handler,
		},
		{
			MethodName: "ResizeLV",
			Handler:    _LVService_ResizeLV_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lvmd/proto/lvmd.proto",
}

// VGServiceClient is the client API for VGService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VGServiceClient interface {
	// Get the list of logical volumes in the volume group.
	GetLVList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetLVListResponse, error)
	// Get the free space of the volume group in bytes.
	GetFreeBytes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFreeBytesResponse, error)
	// Stream the volume group metrics.
	Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (VGService_WatchClient, error)
}

type vGServiceClient struct {
	cc *grpc.ClientConn
}

func NewVGServiceClient(cc *grpc.ClientConn) VGServiceClient {
	return &vGServiceClient{cc}
}

func (c *vGServiceClient) GetLVList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetLVListResponse, error) {
	out := new(GetLVListResponse)
	err := c.cc.Invoke(ctx, "/proto.VGService/GetLVList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vGServiceClient) GetFreeBytes(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*GetFreeBytesResponse, error) {
	out := new(GetFreeBytesResponse)
	err := c.cc.Invoke(ctx, "/proto.VGService/GetFreeBytes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vGServiceClient) Watch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (VGService_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_VGService_serviceDesc.Streams[0], "/proto.VGService/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &vGServiceWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type VGService_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type vGServiceWatchClient struct {
	grpc.ClientStream
}

func (x *vGServiceWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VGServiceServer is the server API for VGService service.
type VGServiceServer interface {
	// Get the list of logical volumes in the volume group.
	GetLVList(context.Context, *Empty) (*GetLVListResponse, error)
	// Get the free space of the volume group in bytes.
	GetFreeBytes(context.Context, *Empty) (*GetFreeBytesResponse, error)
	// Stream the volume group metrics.
	Watch(*Empty, VGService_WatchServer) error
}

// UnimplementedVGServiceServer can be embedded to have forward compatible implementations.
type UnimplementedVGServiceServer struct {
}

func (*UnimplementedVGServiceServer) GetLVList(ctx context.Context, req *Empty) (*GetLVListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLVList not implemented")
}
func (*UnimplementedVGServiceServer) GetFreeBytes(ctx context.Context, req *Empty) (*GetFreeBytesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFreeBytes not implemented")
}
func (*UnimplementedVGServiceServer) Watch(req *Empty, srv VGService_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterVGServiceServer(s *grpc.Server, srv VGServiceServer) {
	s.RegisterService(&_VGService_serviceDesc, srv)
}

func _VGService_GetLVList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VGServiceServer).GetLVList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VGService/GetLVList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VGServiceServer).GetLVList(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VGService_GetFreeBytes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VGServiceServer).GetFreeBytes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.VGService/GetFreeBytes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VGServiceServer).GetFreeBytes(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VGService_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VGServiceServer).Watch(m, &vGServiceWatchServer{stream})
}

type VGService_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type vGServiceWatchServer struct {
	grpc.ServerStream
}

func (x *vGServiceWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _VGService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.VGService",
	HandlerType: (*VGServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLVList",
			Handler:    _VGService_GetLVList_Handler,
		},
		{
			MethodName: "GetFreeBytes",
			Handler:    _VGService_GetFreeBytes_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _VGService_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lvmd/proto/lvmd.proto",
}
