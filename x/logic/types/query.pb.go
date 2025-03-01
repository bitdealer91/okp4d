// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: logic/v1beta/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryServiceParamsRequest is request type for the QueryService/Params RPC method.
type QueryServiceParamsRequest struct {
}

func (m *QueryServiceParamsRequest) Reset()         { *m = QueryServiceParamsRequest{} }
func (m *QueryServiceParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryServiceParamsRequest) ProtoMessage()    {}
func (*QueryServiceParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{0}
}
func (m *QueryServiceParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceParamsRequest.Merge(m, src)
}
func (m *QueryServiceParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceParamsRequest proto.InternalMessageInfo

// QueryServiceParamsResponse is response type for the QueryService/Params RPC method.
type QueryServiceParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryServiceParamsResponse) Reset()         { *m = QueryServiceParamsResponse{} }
func (m *QueryServiceParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryServiceParamsResponse) ProtoMessage()    {}
func (*QueryServiceParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a2e6b98fc71dc0f, []int{1}
}
func (m *QueryServiceParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryServiceParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryServiceParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryServiceParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryServiceParamsResponse.Merge(m, src)
}
func (m *QueryServiceParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryServiceParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryServiceParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryServiceParamsResponse proto.InternalMessageInfo

func (m *QueryServiceParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func init() {
	proto.RegisterType((*QueryServiceParamsRequest)(nil), "logic.v1beta.QueryServiceParamsRequest")
	proto.RegisterType((*QueryServiceParamsResponse)(nil), "logic.v1beta.QueryServiceParamsResponse")
}

func init() { proto.RegisterFile("logic/v1beta/query.proto", fileDescriptor_2a2e6b98fc71dc0f) }

var fileDescriptor_2a2e6b98fc71dc0f = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xc9, 0x4f, 0xcf,
	0x4c, 0xd6, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x01, 0xcb, 0xe8, 0x41, 0x64, 0xa4, 0x44, 0xd2, 0xf3, 0xd3,
	0xf3, 0xc1, 0x12, 0xfa, 0x20, 0x16, 0x44, 0x8d, 0x94, 0x4c, 0x7a, 0x7e, 0x7e, 0x7a, 0x4e, 0xaa,
	0x7e, 0x62, 0x41, 0xa6, 0x7e, 0x62, 0x5e, 0x5e, 0x7e, 0x49, 0x62, 0x49, 0x66, 0x7e, 0x5e, 0x31,
	0x54, 0x56, 0x12, 0xc5, 0xec, 0x82, 0xc4, 0xa2, 0xc4, 0x5c, 0xa8, 0x94, 0x92, 0x34, 0x97, 0x64,
	0x20, 0xc8, 0xae, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0xd4, 0x00, 0xb0, 0x5c, 0x50, 0x6a, 0x61,
	0x69, 0x6a, 0x71, 0x89, 0x52, 0x00, 0x97, 0x14, 0x36, 0xc9, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54,
	0x21, 0x23, 0x2e, 0x36, 0x88, 0x51, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0xc8,
	0x0e, 0xd5, 0x83, 0xa8, 0x76, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0xaa, 0xd2, 0xa8, 0x97,
	0x91, 0x8b, 0x07, 0xd9, 0x48, 0xa1, 0x5a, 0x2e, 0x36, 0x88, 0x42, 0x21, 0x75, 0x54, 0xed, 0x38,
	0x5d, 0x25, 0xa5, 0x41, 0x58, 0x21, 0xc4, 0x85, 0x4a, 0x0a, 0x4d, 0x97, 0x9f, 0x4c, 0x66, 0x92,
	0x12, 0x92, 0xd0, 0xcf, 0xcf, 0x2e, 0x30, 0x01, 0x13, 0x29, 0xfa, 0x90, 0xb0, 0x80, 0xb8, 0xc7,
	0xc9, 0xf6, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x94, 0xd3, 0x33, 0x4b,
	0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0x91, 0x75, 0x57, 0x40, 0xf5, 0x97, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x03, 0xd1, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x93, 0x7d, 0xb2, 0x7c,
	0xbd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryServiceParamsRequest, opts ...grpc.CallOption) (*QueryServiceParamsResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) Params(ctx context.Context, in *QueryServiceParamsRequest, opts ...grpc.CallOption) (*QueryServiceParamsResponse, error) {
	out := new(QueryServiceParamsResponse)
	err := c.cc.Invoke(ctx, "/logic.v1beta.QueryService/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryServiceParamsRequest) (*QueryServiceParamsResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) Params(ctx context.Context, req *QueryServiceParamsRequest) (*QueryServiceParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryServiceParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/logic.v1beta.QueryService/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).Params(ctx, req.(*QueryServiceParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "logic.v1beta.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _QueryService_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "logic/v1beta/query.proto",
}

func (m *QueryServiceParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryServiceParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryServiceParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryServiceParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryServiceParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryServiceParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryServiceParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryServiceParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryServiceParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryServiceParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryServiceParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
