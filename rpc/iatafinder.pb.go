// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iatafinder.proto

package iatafinder

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

type EmptyRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyRequest) Reset()         { *m = EmptyRequest{} }
func (m *EmptyRequest) String() string { return proto.CompactTextString(m) }
func (*EmptyRequest) ProtoMessage()    {}
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{0}
}

func (m *EmptyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyRequest.Unmarshal(m, b)
}
func (m *EmptyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyRequest.Marshal(b, m, deterministic)
}
func (m *EmptyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyRequest.Merge(m, src)
}
func (m *EmptyRequest) XXX_Size() int {
	return xxx_messageInfo_EmptyRequest.Size(m)
}
func (m *EmptyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyRequest proto.InternalMessageInfo

type Airport struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Iata                 string   `protobuf:"bytes,5,opt,name=iata,proto3" json:"iata,omitempty"`
	Icao                 string   `protobuf:"bytes,6,opt,name=icao,proto3" json:"icao,omitempty"`
	Latitude             float64  `protobuf:"fixed64,7,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude            float64  `protobuf:"fixed64,8,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Altitude             float64  `protobuf:"fixed64,9,opt,name=altitude,proto3" json:"altitude,omitempty"`
	Timezone             string   `protobuf:"bytes,10,opt,name=timezone,proto3" json:"timezone,omitempty"`
	DaylightSavingsTime  string   `protobuf:"bytes,11,opt,name=daylight_savings_time,json=daylightSavingsTime,proto3" json:"daylight_savings_time,omitempty"`
	Tz                   string   `protobuf:"bytes,12,opt,name=tz,proto3" json:"tz,omitempty"`
	Type                 string   `protobuf:"bytes,13,opt,name=type,proto3" json:"type,omitempty"`
	Source               string   `protobuf:"bytes,14,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Airport) Reset()         { *m = Airport{} }
func (m *Airport) String() string { return proto.CompactTextString(m) }
func (*Airport) ProtoMessage()    {}
func (*Airport) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{1}
}

func (m *Airport) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Airport.Unmarshal(m, b)
}
func (m *Airport) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Airport.Marshal(b, m, deterministic)
}
func (m *Airport) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airport.Merge(m, src)
}
func (m *Airport) XXX_Size() int {
	return xxx_messageInfo_Airport.Size(m)
}
func (m *Airport) XXX_DiscardUnknown() {
	xxx_messageInfo_Airport.DiscardUnknown(m)
}

var xxx_messageInfo_Airport proto.InternalMessageInfo

func (m *Airport) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Airport) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Airport) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *Airport) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Airport) GetIata() string {
	if m != nil {
		return m.Iata
	}
	return ""
}

func (m *Airport) GetIcao() string {
	if m != nil {
		return m.Icao
	}
	return ""
}

func (m *Airport) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Airport) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Airport) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Airport) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *Airport) GetDaylightSavingsTime() string {
	if m != nil {
		return m.DaylightSavingsTime
	}
	return ""
}

func (m *Airport) GetTz() string {
	if m != nil {
		return m.Tz
	}
	return ""
}

func (m *Airport) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Airport) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type Airports struct {
	Airport              []*Airport `protobuf:"bytes,1,rep,name=airport,proto3" json:"airport,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Airports) Reset()         { *m = Airports{} }
func (m *Airports) String() string { return proto.CompactTextString(m) }
func (*Airports) ProtoMessage()    {}
func (*Airports) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{2}
}

func (m *Airports) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Airports.Unmarshal(m, b)
}
func (m *Airports) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Airports.Marshal(b, m, deterministic)
}
func (m *Airports) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Airports.Merge(m, src)
}
func (m *Airports) XXX_Size() int {
	return xxx_messageInfo_Airports.Size(m)
}
func (m *Airports) XXX_DiscardUnknown() {
	xxx_messageInfo_Airports.DiscardUnknown(m)
}

var xxx_messageInfo_Airports proto.InternalMessageInfo

func (m *Airports) GetAirport() []*Airport {
	if m != nil {
		return m.Airport
	}
	return nil
}

type IATA struct {
	Iata                 string   `protobuf:"bytes,1,opt,name=iata,proto3" json:"iata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IATA) Reset()         { *m = IATA{} }
func (m *IATA) String() string { return proto.CompactTextString(m) }
func (*IATA) ProtoMessage()    {}
func (*IATA) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{3}
}

func (m *IATA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IATA.Unmarshal(m, b)
}
func (m *IATA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IATA.Marshal(b, m, deterministic)
}
func (m *IATA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IATA.Merge(m, src)
}
func (m *IATA) XXX_Size() int {
	return xxx_messageInfo_IATA.Size(m)
}
func (m *IATA) XXX_DiscardUnknown() {
	xxx_messageInfo_IATA.DiscardUnknown(m)
}

var xxx_messageInfo_IATA proto.InternalMessageInfo

func (m *IATA) GetIata() string {
	if m != nil {
		return m.Iata
	}
	return ""
}

type ICAO struct {
	Icao                 string   `protobuf:"bytes,1,opt,name=icao,proto3" json:"icao,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ICAO) Reset()         { *m = ICAO{} }
func (m *ICAO) String() string { return proto.CompactTextString(m) }
func (*ICAO) ProtoMessage()    {}
func (*ICAO) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{4}
}

func (m *ICAO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ICAO.Unmarshal(m, b)
}
func (m *ICAO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ICAO.Marshal(b, m, deterministic)
}
func (m *ICAO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ICAO.Merge(m, src)
}
func (m *ICAO) XXX_Size() int {
	return xxx_messageInfo_ICAO.Size(m)
}
func (m *ICAO) XXX_DiscardUnknown() {
	xxx_messageInfo_ICAO.DiscardUnknown(m)
}

var xxx_messageInfo_ICAO proto.InternalMessageInfo

func (m *ICAO) GetIcao() string {
	if m != nil {
		return m.Icao
	}
	return ""
}

type AirportDescriptor struct {
	Descriptor_          string   `protobuf:"bytes,1,opt,name=descriptor,proto3" json:"descriptor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AirportDescriptor) Reset()         { *m = AirportDescriptor{} }
func (m *AirportDescriptor) String() string { return proto.CompactTextString(m) }
func (*AirportDescriptor) ProtoMessage()    {}
func (*AirportDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff5d7323532d9adb, []int{5}
}

func (m *AirportDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AirportDescriptor.Unmarshal(m, b)
}
func (m *AirportDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AirportDescriptor.Marshal(b, m, deterministic)
}
func (m *AirportDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirportDescriptor.Merge(m, src)
}
func (m *AirportDescriptor) XXX_Size() int {
	return xxx_messageInfo_AirportDescriptor.Size(m)
}
func (m *AirportDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_AirportDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_AirportDescriptor proto.InternalMessageInfo

func (m *AirportDescriptor) GetDescriptor_() string {
	if m != nil {
		return m.Descriptor_
	}
	return ""
}

func init() {
	proto.RegisterType((*EmptyRequest)(nil), "iatafinder.EmptyRequest")
	proto.RegisterType((*Airport)(nil), "iatafinder.Airport")
	proto.RegisterType((*Airports)(nil), "iatafinder.Airports")
	proto.RegisterType((*IATA)(nil), "iatafinder.IATA")
	proto.RegisterType((*ICAO)(nil), "iatafinder.ICAO")
	proto.RegisterType((*AirportDescriptor)(nil), "iatafinder.AirportDescriptor")
}

func init() { proto.RegisterFile("iatafinder.proto", fileDescriptor_ff5d7323532d9adb) }

var fileDescriptor_ff5d7323532d9adb = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0xc6, 0xd9, 0x6e, 0xd3, 0x4e, 0x97, 0x68, 0xf1, 0x02, 0xb2, 0x22, 0x40, 0x51, 0x4e, 0xb9,
	0xb0, 0x87, 0xee, 0x69, 0x4f, 0x28, 0x2a, 0x08, 0x71, 0x5a, 0x29, 0xf4, 0x5e, 0x99, 0xc4, 0x14,
	0x4b, 0x49, 0x1c, 0x9c, 0x09, 0x52, 0xfa, 0x04, 0xbc, 0x28, 0xef, 0x81, 0xec, 0xfc, 0x22, 0xda,
	0xdb, 0xcc, 0xf7, 0x33, 0xf6, 0x7c, 0x71, 0xe0, 0x56, 0x72, 0xe4, 0xdf, 0x65, 0x99, 0x09, 0x7d,
	0x5f, 0x69, 0x85, 0x8a, 0xc2, 0x84, 0x84, 0x1e, 0xdc, 0x7c, 0x2a, 0x2a, 0x6c, 0x13, 0xf1, 0xb3,
	0x11, 0x35, 0x86, 0x7f, 0x1c, 0x70, 0x63, 0xa9, 0x2b, 0xa5, 0x91, 0x7a, 0xe0, 0xc8, 0x8c, 0x91,
	0x80, 0x44, 0xd7, 0x89, 0x23, 0x33, 0x4a, 0x61, 0x51, 0xf2, 0x42, 0x30, 0x27, 0x20, 0xd1, 0x3a,
	0xb1, 0xb5, 0xc1, 0x52, 0x89, 0x2d, 0xbb, 0xea, 0x30, 0x53, 0x53, 0x06, 0x6e, 0xaa, 0x9a, 0x12,
	0x75, 0xcb, 0x16, 0x16, 0x1e, 0x5a, 0xa3, 0x36, 0x67, 0xb3, 0xeb, 0x4e, 0x6d, 0x6a, 0x8b, 0xa5,
	0x5c, 0xb1, 0x65, 0x8f, 0xa5, 0x5c, 0x51, 0x1f, 0x56, 0x39, 0x47, 0x89, 0x4d, 0x26, 0x98, 0x1b,
	0x90, 0x88, 0x24, 0x63, 0x4f, 0xdf, 0xc0, 0x3a, 0x57, 0xe5, 0xb1, 0x23, 0x57, 0x96, 0x9c, 0x00,
	0xe3, 0xe4, 0x79, 0xef, 0x5c, 0x77, 0xce, 0xa1, 0x37, 0x1c, 0xca, 0x42, 0x9c, 0x54, 0x29, 0x18,
	0xd8, 0xd3, 0xc6, 0x9e, 0x6e, 0xe1, 0x55, 0xc6, 0xdb, 0x5c, 0x1e, 0x7f, 0xe0, 0xa1, 0xe6, 0xbf,
	0x64, 0x79, 0xac, 0x0f, 0x86, 0x64, 0x1b, 0x2b, 0xbc, 0x1b, 0xc8, 0xaf, 0x1d, 0xb7, 0x97, 0x85,
	0x30, 0xf9, 0xe0, 0x89, 0xdd, 0x58, 0x81, 0x83, 0x27, 0xb3, 0x09, 0xb6, 0x95, 0x60, 0xcf, 0xbb,
	0x4d, 0x4c, 0x4d, 0x5f, 0xc3, 0xb2, 0x56, 0x8d, 0x4e, 0x05, 0xf3, 0x2c, 0xda, 0x77, 0xe1, 0x23,
	0xac, 0xfa, 0x98, 0x6b, 0xfa, 0x1e, 0x5c, 0xde, 0xd5, 0x8c, 0x04, 0x57, 0xd1, 0x66, 0x7b, 0x77,
	0x3f, 0xfb, 0x66, 0xbd, 0x2c, 0x19, 0x34, 0xa1, 0x0f, 0x8b, 0x2f, 0xf1, 0x3e, 0x1e, 0xc3, 0x24,
	0x53, 0x98, 0x96, 0xdb, 0xc5, 0x4f, 0x63, 0xa8, 0x64, 0x0a, 0x35, 0x7c, 0x80, 0x17, 0xfd, 0xac,
	0x8f, 0xa2, 0x4e, 0xb5, 0xac, 0x50, 0x69, 0xfa, 0x0e, 0x20, 0x1b, 0xbb, 0x5e, 0x3e, 0x43, 0xb6,
	0xbf, 0x1d, 0x98, 0x3d, 0x17, 0xfa, 0x01, 0x36, 0x9f, 0x05, 0x8e, 0x37, 0x67, 0xf3, 0x8b, 0xce,
	0xdf, 0x91, 0xff, 0xf2, 0xcc, 0x0a, 0x75, 0xf8, 0x8c, 0xee, 0x00, 0xa6, 0x01, 0xf4, 0xed, 0x19,
	0xd5, 0x74, 0xb9, 0x8b, 0x43, 0x1e, 0xc1, 0x9b, 0x86, 0xd8, 0x2c, 0x6e, 0xe7, 0x4a, 0x83, 0xf8,
	0xe7, 0x32, 0xfc, 0xcf, 0x6a, 0xa2, 0xfa, 0xd7, 0xba, 0x8b, 0x9f, 0x2e, 0x58, 0xbf, 0x2d, 0xed,
	0xdf, 0xf3, 0xf0, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xa8, 0x05, 0x24, 0x51, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IatafinderClient is the client API for Iatafinder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IatafinderClient interface {
	GetAirports(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Airports, error)
	GetAirport(ctx context.Context, in *AirportDescriptor, opts ...grpc.CallOption) (*Airports, error)
	GetAirportIATA(ctx context.Context, in *IATA, opts ...grpc.CallOption) (*Airport, error)
	GetAirportICAO(ctx context.Context, in *ICAO, opts ...grpc.CallOption) (*Airport, error)
}

type iatafinderClient struct {
	cc *grpc.ClientConn
}

func NewIatafinderClient(cc *grpc.ClientConn) IatafinderClient {
	return &iatafinderClient{cc}
}

func (c *iatafinderClient) GetAirports(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*Airports, error) {
	out := new(Airports)
	err := c.cc.Invoke(ctx, "/iatafinder.iatafinder/GetAirports", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iatafinderClient) GetAirport(ctx context.Context, in *AirportDescriptor, opts ...grpc.CallOption) (*Airports, error) {
	out := new(Airports)
	err := c.cc.Invoke(ctx, "/iatafinder.iatafinder/GetAirport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iatafinderClient) GetAirportIATA(ctx context.Context, in *IATA, opts ...grpc.CallOption) (*Airport, error) {
	out := new(Airport)
	err := c.cc.Invoke(ctx, "/iatafinder.iatafinder/GetAirportIATA", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iatafinderClient) GetAirportICAO(ctx context.Context, in *ICAO, opts ...grpc.CallOption) (*Airport, error) {
	out := new(Airport)
	err := c.cc.Invoke(ctx, "/iatafinder.iatafinder/GetAirportICAO", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IatafinderServer is the server API for Iatafinder service.
type IatafinderServer interface {
	GetAirports(context.Context, *EmptyRequest) (*Airports, error)
	GetAirport(context.Context, *AirportDescriptor) (*Airports, error)
	GetAirportIATA(context.Context, *IATA) (*Airport, error)
	GetAirportICAO(context.Context, *ICAO) (*Airport, error)
}

// UnimplementedIatafinderServer can be embedded to have forward compatible implementations.
type UnimplementedIatafinderServer struct {
}

func (*UnimplementedIatafinderServer) GetAirports(ctx context.Context, req *EmptyRequest) (*Airports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAirports not implemented")
}
func (*UnimplementedIatafinderServer) GetAirport(ctx context.Context, req *AirportDescriptor) (*Airports, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAirport not implemented")
}
func (*UnimplementedIatafinderServer) GetAirportIATA(ctx context.Context, req *IATA) (*Airport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAirportIATA not implemented")
}
func (*UnimplementedIatafinderServer) GetAirportICAO(ctx context.Context, req *ICAO) (*Airport, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAirportICAO not implemented")
}

func RegisterIatafinderServer(s *grpc.Server, srv IatafinderServer) {
	s.RegisterService(&_Iatafinder_serviceDesc, srv)
}

func _Iatafinder_GetAirports_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IatafinderServer).GetAirports(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iatafinder.iatafinder/GetAirports",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IatafinderServer).GetAirports(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iatafinder_GetAirport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirportDescriptor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IatafinderServer).GetAirport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iatafinder.iatafinder/GetAirport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IatafinderServer).GetAirport(ctx, req.(*AirportDescriptor))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iatafinder_GetAirportIATA_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IATA)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IatafinderServer).GetAirportIATA(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iatafinder.iatafinder/GetAirportIATA",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IatafinderServer).GetAirportIATA(ctx, req.(*IATA))
	}
	return interceptor(ctx, in, info, handler)
}

func _Iatafinder_GetAirportICAO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ICAO)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IatafinderServer).GetAirportICAO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iatafinder.iatafinder/GetAirportICAO",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IatafinderServer).GetAirportICAO(ctx, req.(*ICAO))
	}
	return interceptor(ctx, in, info, handler)
}

var _Iatafinder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "iatafinder.iatafinder",
	HandlerType: (*IatafinderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAirports",
			Handler:    _Iatafinder_GetAirports_Handler,
		},
		{
			MethodName: "GetAirport",
			Handler:    _Iatafinder_GetAirport_Handler,
		},
		{
			MethodName: "GetAirportIATA",
			Handler:    _Iatafinder_GetAirportIATA_Handler,
		},
		{
			MethodName: "GetAirportICAO",
			Handler:    _Iatafinder_GetAirportICAO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iatafinder.proto",
}