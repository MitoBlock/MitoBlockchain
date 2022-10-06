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

type MsgCreateDiscountToken struct {
	Creator           string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Timestamp         string `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	ActivityName      string `protobuf:"bytes,3,opt,name=activityName,proto3" json:"activityName,omitempty"`
	Score             string `protobuf:"bytes,4,opt,name=score,proto3" json:"score,omitempty"`
	Message           string `protobuf:"bytes,5,opt,name=message,proto3" json:"message,omitempty"`
	DiscountValue     string `protobuf:"bytes,6,opt,name=discountValue,proto3" json:"discountValue,omitempty"`
	EligibleCompanies string `protobuf:"bytes,7,opt,name=eligibleCompanies,proto3" json:"eligibleCompanies,omitempty"`
	ItemType          string `protobuf:"bytes,8,opt,name=itemType,proto3" json:"itemType,omitempty"`
	ExpiryDate        string `protobuf:"bytes,9,opt,name=expiryDate,proto3" json:"expiryDate,omitempty"`
}

func (m *MsgCreateDiscountToken) Reset()         { *m = MsgCreateDiscountToken{} }
func (m *MsgCreateDiscountToken) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDiscountToken) ProtoMessage()    {}
func (*MsgCreateDiscountToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccdad15e0fb52345, []int{0}
}
func (m *MsgCreateDiscountToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDiscountToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDiscountToken.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDiscountToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDiscountToken.Merge(m, src)
}
func (m *MsgCreateDiscountToken) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDiscountToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDiscountToken.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDiscountToken proto.InternalMessageInfo

func (m *MsgCreateDiscountToken) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetActivityName() string {
	if m != nil {
		return m.ActivityName
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetScore() string {
	if m != nil {
		return m.Score
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetDiscountValue() string {
	if m != nil {
		return m.DiscountValue
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetEligibleCompanies() string {
	if m != nil {
		return m.EligibleCompanies
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetItemType() string {
	if m != nil {
		return m.ItemType
	}
	return ""
}

func (m *MsgCreateDiscountToken) GetExpiryDate() string {
	if m != nil {
		return m.ExpiryDate
	}
	return ""
}

type MsgCreateDiscountTokenResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreateDiscountTokenResponse) Reset()         { *m = MsgCreateDiscountTokenResponse{} }
func (m *MsgCreateDiscountTokenResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateDiscountTokenResponse) ProtoMessage()    {}
func (*MsgCreateDiscountTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccdad15e0fb52345, []int{1}
}
func (m *MsgCreateDiscountTokenResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateDiscountTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateDiscountTokenResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateDiscountTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateDiscountTokenResponse.Merge(m, src)
}
func (m *MsgCreateDiscountTokenResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateDiscountTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateDiscountTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateDiscountTokenResponse proto.InternalMessageInfo

func (m *MsgCreateDiscountTokenResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateDiscountToken)(nil), "mitoblock.mitoblockchain.mitoblockchain.MsgCreateDiscountToken")
	proto.RegisterType((*MsgCreateDiscountTokenResponse)(nil), "mitoblock.mitoblockchain.mitoblockchain.MsgCreateDiscountTokenResponse")
}

func init() { proto.RegisterFile("mitoblockchain/tx.proto", fileDescriptor_ccdad15e0fb52345) }

var fileDescriptor_ccdad15e0fb52345 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xbf, 0x4a, 0x2b, 0x41,
	0x14, 0xc6, 0xb3, 0x9b, 0xff, 0x87, 0x7b, 0x2f, 0xdc, 0x51, 0x74, 0x08, 0x32, 0x48, 0x10, 0xb4,
	0x90, 0x8d, 0x68, 0x67, 0x23, 0x98, 0x80, 0x55, 0x2c, 0x42, 0xb0, 0xb0, 0x9b, 0xdd, 0x1c, 0x36,
	0x43, 0x76, 0x76, 0x96, 0x9d, 0x89, 0x64, 0xdf, 0xc2, 0x37, 0xb0, 0xf4, 0x05, 0x7c, 0x08, 0xcb,
	0x94, 0x96, 0x92, 0xbc, 0x88, 0x64, 0xf2, 0x47, 0x13, 0x53, 0x08, 0x96, 0xdf, 0xf7, 0xfb, 0x98,
	0x73, 0xf8, 0xe6, 0xc0, 0xbe, 0x14, 0x46, 0xf9, 0x91, 0x0a, 0x06, 0x41, 0x9f, 0x8b, 0xb8, 0x61,
	0x46, 0x5e, 0x92, 0x2a, 0xa3, 0xc8, 0xf1, 0x0a, 0x78, 0xeb, 0x91, 0x0d, 0x59, 0x7f, 0x71, 0x61,
	0xaf, 0xad, 0xc3, 0x66, 0x8a, 0xdc, 0x60, 0x4b, 0xe8, 0x40, 0x0d, 0x63, 0xd3, 0x55, 0x03, 0x8c,
	0x09, 0x85, 0x72, 0x30, 0xb3, 0x55, 0x4a, 0x9d, 0x43, 0xe7, 0xa4, 0xda, 0x59, 0x4a, 0x72, 0x00,
	0x55, 0x23, 0x24, 0x6a, 0xc3, 0x65, 0x42, 0x5d, 0xcb, 0x3e, 0x0d, 0x52, 0x87, 0x3f, 0x3c, 0x30,
	0xe2, 0x41, 0x98, 0xec, 0x96, 0x4b, 0xa4, 0x79, 0x1b, 0x58, 0xf3, 0xc8, 0x2e, 0x14, 0x75, 0xa0,
	0x52, 0xa4, 0x05, 0x0b, 0xe7, 0x62, 0x36, 0x51, 0xa2, 0xd6, 0x3c, 0x44, 0x5a, 0x9c, 0x4f, 0x5c,
	0x48, 0x72, 0x04, 0x7f, 0x7b, 0x8b, 0xe5, 0xee, 0x78, 0x34, 0x44, 0x5a, 0xb2, 0x7c, 0xdd, 0x24,
	0xa7, 0xf0, 0x1f, 0x23, 0x11, 0x0a, 0x3f, 0xc2, 0xa6, 0x92, 0x09, 0x8f, 0x05, 0x6a, 0x5a, 0xb6,
	0xc9, 0xef, 0x80, 0xd4, 0xa0, 0x22, 0x0c, 0xca, 0x6e, 0x96, 0x20, 0xad, 0xd8, 0xd0, 0x4a, 0x13,
	0x06, 0x80, 0xa3, 0x44, 0xa4, 0x59, 0x8b, 0x1b, 0xa4, 0x55, 0x4b, 0xbf, 0x38, 0xf5, 0x33, 0x60,
	0xdb, 0x5b, 0xeb, 0xa0, 0x4e, 0x54, 0xac, 0x91, 0xfc, 0x03, 0x57, 0xf4, 0x6c, 0x71, 0x85, 0x8e,
	0x2b, 0x7a, 0xe7, 0xcf, 0x0e, 0xe4, 0xdb, 0x3a, 0x24, 0x4f, 0x0e, 0xec, 0x6c, 0x6b, 0xfb, 0xca,
	0xfb, 0xe1, 0x97, 0x79, 0xdb, 0x07, 0xd7, 0x6e, 0x7e, 0xf9, 0xc0, 0x72, 0xf3, 0xeb, 0xee, 0xeb,
	0x84, 0x39, 0xe3, 0x09, 0x73, 0xde, 0x27, 0xcc, 0x79, 0x9c, 0xb2, 0xdc, 0x78, 0xca, 0x72, 0x6f,
	0x53, 0x96, 0xbb, 0xbf, 0x0c, 0x85, 0xe9, 0x0f, 0x7d, 0x2f, 0x50, 0xb2, 0xb1, 0x7a, 0xb3, 0xb1,
	0x71, 0x83, 0xa3, 0x4d, 0xc3, 0x64, 0x09, 0x6a, 0xbf, 0x64, 0x0f, 0xf3, 0xe2, 0x23, 0x00, 0x00,
	0xff, 0xff, 0x40, 0x2d, 0xd6, 0xdf, 0xb3, 0x02, 0x00, 0x00,
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
	CreateDiscountToken(ctx context.Context, in *MsgCreateDiscountToken, opts ...grpc.CallOption) (*MsgCreateDiscountTokenResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDiscountToken(ctx context.Context, in *MsgCreateDiscountToken, opts ...grpc.CallOption) (*MsgCreateDiscountTokenResponse, error) {
	out := new(MsgCreateDiscountTokenResponse)
	err := c.cc.Invoke(ctx, "/mitoblock.mitoblockchain.mitoblockchain.Msg/CreateDiscountToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateDiscountToken(context.Context, *MsgCreateDiscountToken) (*MsgCreateDiscountTokenResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateDiscountToken(ctx context.Context, req *MsgCreateDiscountToken) (*MsgCreateDiscountTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDiscountToken not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateDiscountToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDiscountToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDiscountToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mitoblock.mitoblockchain.mitoblockchain.Msg/CreateDiscountToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDiscountToken(ctx, req.(*MsgCreateDiscountToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mitoblock.mitoblockchain.mitoblockchain.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDiscountToken",
			Handler:    _Msg_CreateDiscountToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mitoblockchain/tx.proto",
}

func (m *MsgCreateDiscountToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDiscountToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDiscountToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExpiryDate) > 0 {
		i -= len(m.ExpiryDate)
		copy(dAtA[i:], m.ExpiryDate)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ExpiryDate)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ItemType) > 0 {
		i -= len(m.ItemType)
		copy(dAtA[i:], m.ItemType)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ItemType)))
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
	if len(m.DiscountValue) > 0 {
		i -= len(m.DiscountValue)
		copy(dAtA[i:], m.DiscountValue)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DiscountValue)))
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

func (m *MsgCreateDiscountTokenResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateDiscountTokenResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateDiscountTokenResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
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
func (m *MsgCreateDiscountToken) Size() (n int) {
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
	l = len(m.DiscountValue)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EligibleCompanies)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ItemType)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ExpiryDate)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateDiscountTokenResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateDiscountToken) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateDiscountToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDiscountToken: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field DiscountValue", wireType)
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
			m.DiscountValue = string(dAtA[iNdEx:postIndex])
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
				return fmt.Errorf("proto: wrong wireType = %d for field ItemType", wireType)
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
			m.ItemType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
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
func (m *MsgCreateDiscountTokenResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateDiscountTokenResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateDiscountTokenResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
