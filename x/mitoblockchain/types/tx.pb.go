// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mitoblockchain/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateMembershipToken struct {
	Creator            string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Timestamp          string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ActivityName       string `protobuf:"bytes,3,opt,name=activityName,proto3" json:"activityName,omitempty"`
	Score              string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	Message            string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	MembershipDuration string `protobuf:"bytes,6,opt,name=membershipDuration,proto3" json:"membershipDuration,omitempty"`
	EligibleCompanies  string `protobuf:"bytes,7,opt,name=eligibleCompanies,proto3" json:"eligibleCompanies,omitempty"`
	ExpiryDate         string `protobuf:"bytes,8,opt,name=expiryDate,proto3" json:"expiryDate,omitempty"`
}

func (m *MsgCreateMembershipToken) Reset()         { *m = MsgCreateMembershipToken{} }
func (m *MsgCreateMembershipToken) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMembershipToken) ProtoMessage()    {}
func (*MsgCreateMembershipToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccdad15e0fb52345, []int{0}
}
func (m *MsgCreateMembershipToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMembershipToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMembershipToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMembershipToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMembershipToken.Merge(m, src)
}
func (m *MsgCreateMembershipToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMembershipToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMembershipToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMembershipToken proto.InternalMessageInfo

func (m *MsgCreateMembershipToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetMembershipDuration() string {
	if m != nil {
		return m.MembershipDuration
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetEligibleCompanies() string {
	if m != nil {
		return m.EligibleCompanies
	}
	return ""
}

func (m *MsgCreateMembershipToken) GetExpiryDate() string {
	if m != nil {
		return m.ExpiryDate
	}
	return ""
}

type MsgCreateMembershipTokenResponse struct {
}

func (m *MsgCreateMembershipTokenResponse) Reset()         { *m = MsgCreateMembershipTokenResponse{} }
func (m *MsgCreateMembershipTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateMembershipTokenResponse) ProtoMessage()    {}
func (*MsgCreateMembershipTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccdad15e0fb52345, []int{1}
}
func (m *MsgCreateMembershipTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateMembershipTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateMembershipTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateMembershipTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateMembershipTokenResponse.Merge(m, src)
}
func (m *MsgCreateMembershipTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateMembershipTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateMembershipTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateMembershipTokenResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateMembershipToken)(nil), "mitoblock.mitoblockchain.mitoblockchain.MsgCreateMembershipToken")
	proto.RegisterType((*MsgCreateMembershipTokenResponse)(nil), "mitoblock.mitoblockchain.mitoblockchain.MsgCreateMembershipTokenResponse")
}

func init() { proto.RegisterFile("mitoblockchain/tx.proto", fileDescriptor_ccdad15e0fb52345) }

var fileDescriptor_ccdad15e0fb52345 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xb1, 0x4e, 0x7a, 0x31,
	0x14, 0xc6, 0x29, 0xfc, 0x81, 0x3f, 0x27, 0x2e, 0x36, 0x1a, 0x1b, 0x63, 0x1a, 0x72, 0x17, 0x1d,
	0xcc, 0x25, 0xd1, 0xcd, 0x4d, 0x61, 0x71, 0xc0, 0x81, 0x30, 0xb9, 0xf5, 0xde, 0x9c, 0x5c, 0x1a,
	0xe8, 0x6d, 0xd3, 0x16, 0x03, 0x6f, 0xe1, 0x43, 0xe8, 0xe8, 0x7b, 0x38, 0x32, 0x3a, 0x1a, 0x78,
	0x11, 0x43, 0x11, 0x14, 0x84, 0xc4, 0xc4, 0xf1, 0xfb, 0x7e, 0x5f, 0xfa, 0x9d, 0x9c, 0x1e, 0x38,
	0x52, 0xd2, 0xeb, 0x64, 0xa0, 0xd3, 0x7e, 0xda, 0x13, 0x32, 0x6f, 0xf8, 0x51, 0x6c, 0xac, 0xf6,
	0x9a, 0x9e, 0xae, 0x40, 0xbc, 0x1e, 0xd9, 0x90, 0xd1, 0x53, 0x11, 0x58, 0xdb, 0x65, 0x4d, 0x8b,
	0xc2, 0x63, 0x1b, 0x55, 0x82, 0xd6, 0xf5, 0xa4, 0xe9, 0xea, 0x3e, 0xe6, 0x94, 0x41, 0x35, 0x9d,
	0x03, 0x6d, 0x19, 0xa9, 0x93, 0xb3, 0x5a, 0x67, 0x29, 0xe9, 0x09, 0xd4, 0xbc, 0x54, 0xe8, 0xbc,
	0x50, 0x86, 0x15, 0x03, 0xfb, 0x32, 0x68, 0x04, 0x7b, 0x22, 0xf5, 0xf2, 0x41, 0xfa, 0xf1, 0x9d,
	0x50, 0xc8, 0x4a, 0x21, 0xb0, 0xe6, 0xd1, 0x03, 0x28, 0xbb, 0x54, 0x5b, 0x64, 0xff, 0x02, 0x5c,
	0x88, 0x79, 0xa3, 0x42, 0xe7, 0x44, 0x86, 0xac, 0xbc, 0x68, 0xfc, 0x94, 0x34, 0x06, 0xaa, 0x56,
	0xe3, 0xb5, 0x86, 0x56, 0x78, 0xa9, 0x73, 0x56, 0x09, 0xa1, 0x2d, 0x84, 0x9e, 0xc3, 0x3e, 0x0e,
	0x64, 0x26, 0x93, 0x01, 0x36, 0xb5, 0x32, 0x22, 0x97, 0xe8, 0x58, 0x35, 0xc4, 0x7f, 0x02, 0xca,
	0x01, 0x70, 0x64, 0xa4, 0x1d, 0xb7, 0x84, 0x47, 0xf6, 0x3f, 0xc4, 0xbe, 0x39, 0x51, 0x04, 0xf5,
	0x5d, 0x5b, 0xea, 0xa0, 0x33, 0x3a, 0x77, 0x78, 0xf1, 0x42, 0xa0, 0xd4, 0x76, 0x19, 0x7d, 0x26,
	0x70, 0xb8, 0x7d, 0x9f, 0xd7, 0xf1, 0x2f, 0xbf, 0x25, 0xde, 0x55, 0x76, 0x7c, 0xfb, 0xe7, 0x27,
	0x96, 0xf3, 0xde, 0x74, 0x5f, 0xa7, 0x9c, 0x4c, 0xa6, 0x9c, 0xbc, 0x4f, 0x39, 0x79, 0x9c, 0xf1,
	0xc2, 0x64, 0xc6, 0x0b, 0x6f, 0x33, 0x5e, 0xb8, 0xbf, 0xca, 0xa4, 0xef, 0x0d, 0x93, 0x38, 0xd5,
	0xaa, 0xb1, 0x7a, 0xb5, 0xb1, 0x71, 0x6b, 0xa3, 0x4d, 0xc3, 0x8f, 0x0d, 0xba, 0xa4, 0x12, 0x0e,
	0xf0, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x86, 0x48, 0x11, 0x4a, 0x9b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateMembershipToken(ctx context.Context, in *MsgCreateMembershipToken, opts ...grpc.CallOption) (*MsgCreateMembershipTokenResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateMembershipToken(ctx context.Context, in *MsgCreateMembershipToken, opts ...grpc.CallOption) (*MsgCreateMembershipTokenResponse, error) {
	out := new(MsgCreateMembershipTokenResponse)
	err := c.cc.Invoke(ctx, "/mitoblock.mitoblockchain.mitoblockchain.Msg/CreateMembershipToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateMembershipToken(context.Context, *MsgCreateMembershipToken) (*MsgCreateMembershipTokenResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateMembershipToken(ctx context.Context, req *MsgCreateMembershipToken) (*MsgCreateMembershipTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMembershipToken not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateMembershipToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMembershipToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMembershipToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mitoblock.mitoblockchain.mitoblockchain.Msg/CreateMembershipToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMembershipToken(ctx, req.(*MsgCreateMembershipToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mitoblock.mitoblockchain.mitoblockchain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMembershipToken",
			Handler:    _Msg_CreateMembershipToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitoblockchain/tx.proto",
}

func (m *MsgCreateMembershipToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMembershipToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMembershipToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExpiryDate) > 0 {
		i -= len(m.ExpiryDate)
		copy(dAtA[i:], m.ExpiryDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ExpiryDate)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.EligibleCompanies) > 0 {
		i -= len(m.EligibleCompanies)
		copy(dAtA[i:], m.EligibleCompanies)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EligibleCompanies)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.MembershipDuration) > 0 {
		i -= len(m.MembershipDuration)
		copy(dAtA[i:], m.MembershipDuration)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MembershipDuration)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Message) > 0 {
		i -= len(m.Message)
		copy(dAtA[i:], m.Message)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Message)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Score) > 0 {
		i -= len(m.Score)
		copy(dAtA[i:], m.Score)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Score)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ActivityName) > 0 {
		i -= len(m.ActivityName)
		copy(dAtA[i:], m.ActivityName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ActivityName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Timestamp) > 0 {
		i -= len(m.Timestamp)
		copy(dAtA[i:], m.Timestamp)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Timestamp)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateMembershipTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateMembershipTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateMembershipTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateMembershipToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Timestamp)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ActivityName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Score)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MembershipDuration)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EligibleCompanies)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ExpiryDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateMembershipTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateMembershipToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateMembershipToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMembershipToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Timestamp = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ActivityName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ActivityName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Score = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MembershipDuration", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MembershipDuration = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EligibleCompanies", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EligibleCompanies = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiryDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpiryDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCreateMembershipTokenResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCreateMembershipTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateMembershipTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
