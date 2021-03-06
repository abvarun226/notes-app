// Code generated by protoc-gen-go. DO NOT EDIT.
// source: note.proto

package note

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

type NotesRequest struct {
	UserId               string      `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Sort                 *SortOption `protobuf:"bytes,2,opt,name=sort,proto3" json:"sort,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *NotesRequest) Reset()         { *m = NotesRequest{} }
func (m *NotesRequest) String() string { return proto.CompactTextString(m) }
func (*NotesRequest) ProtoMessage()    {}
func (*NotesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{0}
}

func (m *NotesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotesRequest.Unmarshal(m, b)
}
func (m *NotesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotesRequest.Marshal(b, m, deterministic)
}
func (m *NotesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotesRequest.Merge(m, src)
}
func (m *NotesRequest) XXX_Size() int {
	return xxx_messageInfo_NotesRequest.Size(m)
}
func (m *NotesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NotesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NotesRequest proto.InternalMessageInfo

func (m *NotesRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NotesRequest) GetSort() *SortOption {
	if m != nil {
		return m.Sort
	}
	return nil
}

type NotesResponse struct {
	Notes                []string `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotesResponse) Reset()         { *m = NotesResponse{} }
func (m *NotesResponse) String() string { return proto.CompactTextString(m) }
func (*NotesResponse) ProtoMessage()    {}
func (*NotesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{1}
}

func (m *NotesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotesResponse.Unmarshal(m, b)
}
func (m *NotesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotesResponse.Marshal(b, m, deterministic)
}
func (m *NotesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotesResponse.Merge(m, src)
}
func (m *NotesResponse) XXX_Size() int {
	return xxx_messageInfo_NotesResponse.Size(m)
}
func (m *NotesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NotesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NotesResponse proto.InternalMessageInfo

func (m *NotesResponse) GetNotes() []string {
	if m != nil {
		return m.Notes
	}
	return nil
}

type SortOption struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
	Order                string   `protobuf:"bytes,2,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SortOption) Reset()         { *m = SortOption{} }
func (m *SortOption) String() string { return proto.CompactTextString(m) }
func (*SortOption) ProtoMessage()    {}
func (*SortOption) Descriptor() ([]byte, []int) {
	return fileDescriptor_640dafe07df50d4e, []int{2}
}

func (m *SortOption) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortOption.Unmarshal(m, b)
}
func (m *SortOption) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortOption.Marshal(b, m, deterministic)
}
func (m *SortOption) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortOption.Merge(m, src)
}
func (m *SortOption) XXX_Size() int {
	return xxx_messageInfo_SortOption.Size(m)
}
func (m *SortOption) XXX_DiscardUnknown() {
	xxx_messageInfo_SortOption.DiscardUnknown(m)
}

var xxx_messageInfo_SortOption proto.InternalMessageInfo

func (m *SortOption) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *SortOption) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func init() {
	proto.RegisterType((*NotesRequest)(nil), "note.NotesRequest")
	proto.RegisterType((*NotesResponse)(nil), "note.NotesResponse")
	proto.RegisterType((*SortOption)(nil), "note.SortOption")
}

func init() { proto.RegisterFile("note.proto", fileDescriptor_640dafe07df50d4e) }

var fileDescriptor_640dafe07df50d4e = []byte{
	// 203 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xcf, 0x6a, 0xc3, 0x30,
	0x0c, 0x87, 0xe7, 0x2d, 0xcb, 0xb0, 0xb6, 0xc1, 0xf0, 0x06, 0x0b, 0x3b, 0x05, 0xb3, 0x41, 0x4e,
	0x39, 0x64, 0x30, 0x76, 0xda, 0x7d, 0xd0, 0x3f, 0xe0, 0x3e, 0x40, 0xa1, 0x44, 0x05, 0x43, 0x89,
	0x52, 0x4b, 0x79, 0xff, 0x62, 0xbb, 0xa5, 0xed, 0xcd, 0x9f, 0x2c, 0xe9, 0xd3, 0x0f, 0x60, 0x20,
	0xc1, 0x76, 0x0c, 0x24, 0x64, 0x8a, 0xf8, 0xb6, 0x73, 0x78, 0x5a, 0x90, 0x20, 0x3b, 0xdc, 0x4f,
	0xc8, 0x62, 0xde, 0xe1, 0x61, 0x62, 0x0c, 0x6b, 0xdf, 0x57, 0xaa, 0x56, 0x8d, 0x76, 0x65, 0xc4,
	0xff, 0xde, 0x7c, 0x42, 0xc1, 0x14, 0xa4, 0xba, 0xad, 0x55, 0xf3, 0xd8, 0xbd, 0xb4, 0x69, 0xd3,
	0x8a, 0x82, 0x2c, 0x47, 0xf1, 0x34, 0xb8, 0xf4, 0x6b, 0xbf, 0xe0, 0xf9, 0xb8, 0x8e, 0x47, 0x1a,
	0x18, 0xcd, 0x1b, 0xdc, 0xc7, 0x4e, 0xae, 0x54, 0x7d, 0xd7, 0x68, 0x97, 0xc1, 0xfe, 0x02, 0x9c,
	0x47, 0x63, 0xcf, 0xd6, 0xe3, 0xee, 0x64, 0xcc, 0x10, 0xab, 0x14, 0x7a, 0x0c, 0xc9, 0xa8, 0x5d,
	0x86, 0xee, 0x0f, 0x8a, 0x28, 0x30, 0x3f, 0xa0, 0x67, 0x9e, 0x25, 0xc9, 0x8c, 0xc9, 0xd7, 0x5c,
	0x06, 0xf9, 0x78, 0xbd, 0xaa, 0xe5, 0x6b, 0xec, 0xcd, 0xa6, 0x4c, 0xe1, 0xbf, 0x0f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xb0, 0x82, 0x2d, 0x6b, 0x0a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NoteClient is the client API for Note service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NoteClient interface {
	ListNotes(ctx context.Context, in *NotesRequest, opts ...grpc.CallOption) (*NotesResponse, error)
}

type noteClient struct {
	cc *grpc.ClientConn
}

func NewNoteClient(cc *grpc.ClientConn) NoteClient {
	return &noteClient{cc}
}

func (c *noteClient) ListNotes(ctx context.Context, in *NotesRequest, opts ...grpc.CallOption) (*NotesResponse, error) {
	out := new(NotesResponse)
	err := c.cc.Invoke(ctx, "/note.Note/ListNotes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NoteServer is the server API for Note service.
type NoteServer interface {
	ListNotes(context.Context, *NotesRequest) (*NotesResponse, error)
}

// UnimplementedNoteServer can be embedded to have forward compatible implementations.
type UnimplementedNoteServer struct {
}

func (*UnimplementedNoteServer) ListNotes(ctx context.Context, req *NotesRequest) (*NotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNotes not implemented")
}

func RegisterNoteServer(s *grpc.Server, srv NoteServer) {
	s.RegisterService(&_Note_serviceDesc, srv)
}

func _Note_ListNotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NoteServer).ListNotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/note.Note/ListNotes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NoteServer).ListNotes(ctx, req.(*NotesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Note_serviceDesc = grpc.ServiceDesc{
	ServiceName: "note.Note",
	HandlerType: (*NoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListNotes",
			Handler:    _Note_ListNotes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "note.proto",
}
