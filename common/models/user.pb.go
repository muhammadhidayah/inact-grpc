// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type Person struct {
	PersonNo             int64    `protobuf:"varint,1,opt,name=person_no,json=personNo,proto3" json:"person_no,omitempty"`
	PersonId             string   `protobuf:"bytes,2,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	PersonName           string   `protobuf:"bytes,3,opt,name=person_name,json=personName,proto3" json:"person_name,omitempty"`
	PersonPosition       string   `protobuf:"bytes,4,opt,name=person_position,json=personPosition,proto3" json:"person_position,omitempty"`
	PersonMail           string   `protobuf:"bytes,5,opt,name=person_mail,json=personMail,proto3" json:"person_mail,omitempty"`
	PersonTelephone      string   `protobuf:"bytes,6,opt,name=person_telephone,json=personTelephone,proto3" json:"person_telephone,omitempty"`
	PersonPassword       string   `protobuf:"bytes,7,opt,name=person_password,json=personPassword,proto3" json:"person_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Person) Reset()         { *m = Person{} }
func (m *Person) String() string { return proto.CompactTextString(m) }
func (*Person) ProtoMessage()    {}
func (*Person) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *Person) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Person.Unmarshal(m, b)
}
func (m *Person) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Person.Marshal(b, m, deterministic)
}
func (m *Person) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Person.Merge(m, src)
}
func (m *Person) XXX_Size() int {
	return xxx_messageInfo_Person.Size(m)
}
func (m *Person) XXX_DiscardUnknown() {
	xxx_messageInfo_Person.DiscardUnknown(m)
}

var xxx_messageInfo_Person proto.InternalMessageInfo

func (m *Person) GetPersonNo() int64 {
	if m != nil {
		return m.PersonNo
	}
	return 0
}

func (m *Person) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

func (m *Person) GetPersonName() string {
	if m != nil {
		return m.PersonName
	}
	return ""
}

func (m *Person) GetPersonPosition() string {
	if m != nil {
		return m.PersonPosition
	}
	return ""
}

func (m *Person) GetPersonMail() string {
	if m != nil {
		return m.PersonMail
	}
	return ""
}

func (m *Person) GetPersonTelephone() string {
	if m != nil {
		return m.PersonTelephone
	}
	return ""
}

func (m *Person) GetPersonPassword() string {
	if m != nil {
		return m.PersonPassword
	}
	return ""
}

type PersonNo struct {
	PersonNo             int64    `protobuf:"varint,1,opt,name=person_no,json=personNo,proto3" json:"person_no,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonNo) Reset()         { *m = PersonNo{} }
func (m *PersonNo) String() string { return proto.CompactTextString(m) }
func (*PersonNo) ProtoMessage()    {}
func (*PersonNo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *PersonNo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonNo.Unmarshal(m, b)
}
func (m *PersonNo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonNo.Marshal(b, m, deterministic)
}
func (m *PersonNo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonNo.Merge(m, src)
}
func (m *PersonNo) XXX_Size() int {
	return xxx_messageInfo_PersonNo.Size(m)
}
func (m *PersonNo) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonNo.DiscardUnknown(m)
}

var xxx_messageInfo_PersonNo proto.InternalMessageInfo

func (m *PersonNo) GetPersonNo() int64 {
	if m != nil {
		return m.PersonNo
	}
	return 0
}

type PersonID struct {
	PersonId             string   `protobuf:"bytes,1,opt,name=person_id,json=personId,proto3" json:"person_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersonID) Reset()         { *m = PersonID{} }
func (m *PersonID) String() string { return proto.CompactTextString(m) }
func (*PersonID) ProtoMessage()    {}
func (*PersonID) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *PersonID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonID.Unmarshal(m, b)
}
func (m *PersonID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonID.Marshal(b, m, deterministic)
}
func (m *PersonID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonID.Merge(m, src)
}
func (m *PersonID) XXX_Size() int {
	return xxx_messageInfo_PersonID.Size(m)
}
func (m *PersonID) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonID.DiscardUnknown(m)
}

var xxx_messageInfo_PersonID proto.InternalMessageInfo

func (m *PersonID) GetPersonId() string {
	if m != nil {
		return m.PersonId
	}
	return ""
}

type PersonList struct {
	List                 []*Person `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PersonList) Reset()         { *m = PersonList{} }
func (m *PersonList) String() string { return proto.CompactTextString(m) }
func (*PersonList) ProtoMessage()    {}
func (*PersonList) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *PersonList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersonList.Unmarshal(m, b)
}
func (m *PersonList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersonList.Marshal(b, m, deterministic)
}
func (m *PersonList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersonList.Merge(m, src)
}
func (m *PersonList) XXX_Size() int {
	return xxx_messageInfo_PersonList.Size(m)
}
func (m *PersonList) XXX_DiscardUnknown() {
	xxx_messageInfo_PersonList.DiscardUnknown(m)
}

var xxx_messageInfo_PersonList proto.InternalMessageInfo

func (m *PersonList) GetList() []*Person {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Person)(nil), "Person")
	proto.RegisterType((*PersonNo)(nil), "PersonNo")
	proto.RegisterType((*PersonID)(nil), "PersonID")
	proto.RegisterType((*PersonList)(nil), "PersonList")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0x5d, 0x4b, 0x32, 0x41,
	0x14, 0x76, 0xd5, 0x57, 0xdd, 0xa3, 0x6f, 0xc5, 0x5c, 0xc4, 0xb0, 0x5e, 0x24, 0x0b, 0xa1, 0xde,
	0x8c, 0xa4, 0xbf, 0xa0, 0x30, 0x42, 0xa8, 0x10, 0xa9, 0xeb, 0x58, 0xd9, 0x93, 0x0d, 0xec, 0xee,
	0x2c, 0x3b, 0x23, 0xe1, 0x65, 0x3f, 0x3c, 0x08, 0xe7, 0x23, 0x77, 0x11, 0xa1, 0xbb, 0x99, 0xe7,
	0x63, 0x9e, 0x33, 0xe7, 0x01, 0xd8, 0x4a, 0x2c, 0x58, 0x5e, 0x08, 0x25, 0x82, 0xfe, 0x46, 0x88,
	0x4d, 0x82, 0x13, 0x7d, 0x5b, 0x6f, 0xdf, 0x27, 0x98, 0xe6, 0x6a, 0x67, 0xc8, 0xf0, 0xab, 0x0e,
	0xad, 0x25, 0x16, 0x52, 0x64, 0xa4, 0x0f, 0x7e, 0xae, 0x4f, 0x6f, 0x99, 0xa0, 0xde, 0xc0, 0x1b,
	0x35, 0x56, 0x1d, 0x03, 0x3c, 0x8b, 0x12, 0xc9, 0x63, 0x5a, 0x1f, 0x78, 0x23, 0xdf, 0x91, 0x8b,
	0x98, 0x5c, 0x41, 0xd7, 0x39, 0xa3, 0x14, 0x69, 0x43, 0xd3, 0x60, 0xbd, 0x51, 0x8a, 0x64, 0x08,
	0xe7, 0x56, 0x90, 0x0b, 0xc9, 0x15, 0x17, 0x19, 0x6d, 0x6a, 0xd1, 0x99, 0x81, 0x97, 0x16, 0x2d,
	0xbd, 0x94, 0x46, 0x3c, 0xa1, 0xff, 0xca, 0x2f, 0x3d, 0x45, 0x3c, 0x21, 0x63, 0xb8, 0xb0, 0x02,
	0x85, 0x09, 0xe6, 0x1f, 0x22, 0x43, 0xda, 0xd2, 0x2a, 0x9b, 0xf0, 0xe2, 0xe0, 0x72, 0x68, 0x24,
	0xe5, 0xa7, 0x28, 0x62, 0xda, 0xae, 0x84, 0x5a, 0x34, 0x1c, 0x42, 0x67, 0x79, 0xfc, 0xcf, 0xe3,
	0x25, 0x1c, 0x84, 0x8b, 0x79, 0x75, 0x21, 0x5e, 0x75, 0x21, 0xe1, 0x18, 0xc0, 0x08, 0x1f, 0xb9,
	0x54, 0xa4, 0x0f, 0xcd, 0x84, 0x4b, 0x45, 0xbd, 0x41, 0x63, 0xd4, 0x9d, 0xb6, 0x99, 0xa1, 0x56,
	0x1a, 0x9c, 0x7e, 0x7b, 0xd0, 0x36, 0x80, 0x24, 0xd7, 0xf0, 0xff, 0x01, 0x95, 0xb9, 0xdd, 0xed,
	0x16, 0x73, 0xe2, 0x33, 0x97, 0x17, 0x38, 0x5b, 0x58, 0x23, 0x37, 0xd0, 0x7b, 0xcd, 0xe3, 0x48,
	0xa1, 0x2d, 0xce, 0x51, 0xc1, 0x25, 0x33, 0x55, 0x33, 0x57, 0x35, 0xbb, 0xdf, 0x57, 0x1d, 0xd6,
	0x08, 0x03, 0xff, 0x36, 0x8e, 0xff, 0xae, 0x9f, 0x01, 0xec, 0x47, 0xb7, 0x86, 0x13, 0xba, 0xa0,
	0xcb, 0x0e, 0xbf, 0xd4, 0xa6, 0xde, 0x1c, 0x13, 0xfc, 0x9d, 0xab, 0x34, 0xfd, 0xc9, 0xa4, 0x75,
	0x4b, 0x23, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x18, 0x91, 0x92, 0x24, 0xb2, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PersonsClient is the client API for Persons service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PersonsClient interface {
	GetPersonByID(ctx context.Context, in *PersonID, opts ...grpc.CallOption) (*Person, error)
	UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*empty.Empty, error)
	AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*empty.Empty, error)
	ListPerson(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PersonList, error)
	DeletePerson(ctx context.Context, in *PersonID, opts ...grpc.CallOption) (*empty.Empty, error)
}

type personsClient struct {
	cc *grpc.ClientConn
}

func NewPersonsClient(cc *grpc.ClientConn) PersonsClient {
	return &personsClient{cc}
}

func (c *personsClient) GetPersonByID(ctx context.Context, in *PersonID, opts ...grpc.CallOption) (*Person, error) {
	out := new(Person)
	err := c.cc.Invoke(ctx, "/Persons/GetPersonByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) UpdatePerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Persons/UpdatePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) AddPerson(ctx context.Context, in *Person, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Persons/AddPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) ListPerson(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PersonList, error) {
	out := new(PersonList)
	err := c.cc.Invoke(ctx, "/Persons/ListPerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *personsClient) DeletePerson(ctx context.Context, in *PersonID, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/Persons/DeletePerson", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PersonsServer is the server API for Persons service.
type PersonsServer interface {
	GetPersonByID(context.Context, *PersonID) (*Person, error)
	UpdatePerson(context.Context, *Person) (*empty.Empty, error)
	AddPerson(context.Context, *Person) (*empty.Empty, error)
	ListPerson(context.Context, *empty.Empty) (*PersonList, error)
	DeletePerson(context.Context, *PersonID) (*empty.Empty, error)
}

// UnimplementedPersonsServer can be embedded to have forward compatible implementations.
type UnimplementedPersonsServer struct {
}

func (*UnimplementedPersonsServer) GetPersonByID(ctx context.Context, req *PersonID) (*Person, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPersonByID not implemented")
}
func (*UnimplementedPersonsServer) UpdatePerson(ctx context.Context, req *Person) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePerson not implemented")
}
func (*UnimplementedPersonsServer) AddPerson(ctx context.Context, req *Person) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPerson not implemented")
}
func (*UnimplementedPersonsServer) ListPerson(ctx context.Context, req *empty.Empty) (*PersonList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPerson not implemented")
}
func (*UnimplementedPersonsServer) DeletePerson(ctx context.Context, req *PersonID) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePerson not implemented")
}

func RegisterPersonsServer(s *grpc.Server, srv PersonsServer) {
	s.RegisterService(&_Persons_serviceDesc, srv)
}

func _Persons_GetPersonByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).GetPersonByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Persons/GetPersonByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).GetPersonByID(ctx, req.(*PersonID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_UpdatePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).UpdatePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Persons/UpdatePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).UpdatePerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_AddPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Person)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).AddPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Persons/AddPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).AddPerson(ctx, req.(*Person))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_ListPerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).ListPerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Persons/ListPerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).ListPerson(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Persons_DeletePerson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PersonsServer).DeletePerson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Persons/DeletePerson",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PersonsServer).DeletePerson(ctx, req.(*PersonID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Persons_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Persons",
	HandlerType: (*PersonsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPersonByID",
			Handler:    _Persons_GetPersonByID_Handler,
		},
		{
			MethodName: "UpdatePerson",
			Handler:    _Persons_UpdatePerson_Handler,
		},
		{
			MethodName: "AddPerson",
			Handler:    _Persons_AddPerson_Handler,
		},
		{
			MethodName: "ListPerson",
			Handler:    _Persons_ListPerson_Handler,
		},
		{
			MethodName: "DeletePerson",
			Handler:    _Persons_DeletePerson_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
