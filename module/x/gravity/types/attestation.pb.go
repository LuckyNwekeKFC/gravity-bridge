// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gravity/v1/attestation.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EventType is the cosmos type of an event from Ethereum
type EventType int32

const (
	EVENT_TYPE_UNSPECIFIED            EventType = 0
	EVENT_TYPE_DEPOSIT                EventType = 1
	EVENT_TYPE_WITHDRAW               EventType = 2
	EVENT_TYPE_COSMOS_ERC20_DEPLOYED  EventType = 3
	EVENT_TYPE_CONTRACT_CALL_EXECUTED EventType = 4
	EVENT_TYPE_SIGNER_SET_UPDATED     EventType = 5
)

var EventType_name = map[int32]string{
	0: "EVENT_TYPE_UNSPECIFIED",
	1: "EVENT_TYPE_DEPOSIT",
	2: "EVENT_TYPE_WITHDRAW",
	3: "EVENT_TYPE_COSMOS_ERC20_DEPLOYED",
	4: "EVENT_TYPE_CONTRACT_CALL_EXECUTED",
	5: "EVENT_TYPE_SIGNER_SET_UPDATED",
}

var EventType_value = map[string]int32{
	"EVENT_TYPE_UNSPECIFIED":            0,
	"EVENT_TYPE_DEPOSIT":                1,
	"EVENT_TYPE_WITHDRAW":               2,
	"EVENT_TYPE_COSMOS_ERC20_DEPLOYED":  3,
	"EVENT_TYPE_CONTRACT_CALL_EXECUTED": 4,
	"EVENT_TYPE_SIGNER_SET_UPDATED":     5,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e3205613bbab7525, []int{0}
}

// EthereumEventVoteRecord is an aggregate of `claims` that eventually becomes `accepted` by
// all orchestrators
type EthereumEventVoteRecord struct {
	// This field stores whether the EthereumEventVoteRecord has had its event applied to the Cosmos state. This happens when
	// enough (usually >2/3s) of the validator power votes that they saw the event on Ethereum.
	// For example, once a DepositClaim has modified the token balance of the account that it was deposited to,
	// this boolean will be set to true.
	Accepted bool `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	// This is an array of the addresses of the validators which have voted that they saw the event on Ethereum.
	Votes []string `protobuf:"bytes,2,rep,name=votes,proto3" json:"votes,omitempty"`
	// This is the Cosmos block height that this event was first submitted by a validator.
	Height uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	// The Ethereum event that this EthereumEventVoteRecord is recording votes for.
	Event *types.Any `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
}

func (m *EthereumEventVoteRecord) Reset()         { *m = EthereumEventVoteRecord{} }
func (m *EthereumEventVoteRecord) String() string { return proto.CompactTextString(m) }
func (*EthereumEventVoteRecord) ProtoMessage()    {}
func (*EthereumEventVoteRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3205613bbab7525, []int{0}
}
func (m *EthereumEventVoteRecord) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumEventVoteRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumEventVoteRecord.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthereumEventVoteRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumEventVoteRecord.Merge(m, src)
}
func (m *EthereumEventVoteRecord) XXX_Size() int {
	return m.Size()
}
func (m *EthereumEventVoteRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumEventVoteRecord.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumEventVoteRecord proto.InternalMessageInfo

func (m *EthereumEventVoteRecord) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *EthereumEventVoteRecord) GetVotes() []string {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *EthereumEventVoteRecord) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *EthereumEventVoteRecord) GetEvent() *types.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

// ERC20Token unique identifier for an Ethereum ERC20 token.
// CONTRACT:
// The contract address on ETH of the token, this could be a Cosmos
// originated token, if so it will be the ERC20 address of the representation
// (note: developers should look up the token symbol using the address on ETH to display for UI)
type ERC20Token struct {
	Contract string                                 `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Amount   github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *ERC20Token) Reset()         { *m = ERC20Token{} }
func (m *ERC20Token) String() string { return proto.CompactTextString(m) }
func (*ERC20Token) ProtoMessage()    {}
func (*ERC20Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_e3205613bbab7525, []int{1}
}
func (m *ERC20Token) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ERC20Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ERC20Token.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ERC20Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20Token.Merge(m, src)
}
func (m *ERC20Token) XXX_Size() int {
	return m.Size()
}
func (m *ERC20Token) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20Token.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20Token proto.InternalMessageInfo

func (m *ERC20Token) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterEnum("gravity.v1.EventType", EventType_name, EventType_value)
	proto.RegisterType((*EthereumEventVoteRecord)(nil), "gravity.v1.EthereumEventVoteRecord")
	proto.RegisterType((*ERC20Token)(nil), "gravity.v1.ERC20Token")
}

func init() { proto.RegisterFile("gravity/v1/attestation.proto", fileDescriptor_e3205613bbab7525) }

var fileDescriptor_e3205613bbab7525 = []byte{
	// 485 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x8e, 0xfb, 0xa5, 0xd5, 0x5c, 0x2a, 0x53, 0x75, 0xa5, 0x82, 0xac, 0x9b, 0x00, 0x55, 0x93,
	0x96, 0xb0, 0x71, 0xe0, 0xdc, 0x25, 0x1e, 0x44, 0x2a, 0x6d, 0x71, 0xdc, 0x8d, 0x71, 0x89, 0xd2,
	0xd4, 0xa4, 0xd5, 0x96, 0x38, 0x4a, 0x9c, 0x88, 0xfc, 0x03, 0x6e, 0xf0, 0x1f, 0xf8, 0x33, 0x3b,
	0xa1, 0x1d, 0x11, 0x87, 0x09, 0xb5, 0x7f, 0x04, 0xe5, 0x63, 0x53, 0x11, 0x27, 0xfb, 0xf9, 0x78,
	0xad, 0xe7, 0x7d, 0xfd, 0xc2, 0xa7, 0x6e, 0x68, 0x27, 0x2b, 0x91, 0xaa, 0xc9, 0xb1, 0x6a, 0x0b,
	0xc1, 0x22, 0x61, 0x8b, 0x15, 0xf7, 0x95, 0x20, 0xe4, 0x82, 0x23, 0x58, 0xaa, 0x4a, 0x72, 0xdc,
	0x6b, 0xbb, 0xdc, 0xe5, 0x39, 0xad, 0x66, 0xb7, 0xc2, 0xd1, 0x7b, 0xe2, 0x72, 0xee, 0x5e, 0x33,
	0x35, 0x47, 0xf3, 0xf8, 0xb3, 0x6a, 0xfb, 0x69, 0x21, 0x1d, 0x7c, 0x03, 0x70, 0x17, 0x8b, 0x25,
	0x0b, 0x59, 0xec, 0xe1, 0x84, 0xf9, 0xe2, 0x9c, 0x0b, 0x46, 0x98, 0xc3, 0xc3, 0x05, 0xea, 0xc1,
	0x1d, 0xdb, 0x71, 0x58, 0x20, 0xd8, 0xa2, 0x0b, 0xfa, 0x60, 0xb0, 0x43, 0x1e, 0x30, 0x6a, 0xc3,
	0x7a, 0xc2, 0x05, 0x8b, 0xba, 0x95, 0x7e, 0x75, 0xd0, 0x24, 0x05, 0x40, 0x1d, 0xd8, 0x58, 0xb2,
	0x95, 0xbb, 0x14, 0xdd, 0x6a, 0x1f, 0x0c, 0x6a, 0xa4, 0x44, 0xe8, 0x10, 0xd6, 0x59, 0xf6, 0x78,
	0xb7, 0xd6, 0x07, 0x83, 0x47, 0x27, 0x6d, 0xa5, 0x08, 0xa4, 0xdc, 0x07, 0x52, 0x86, 0x7e, 0x4a,
	0x0a, 0xcb, 0x41, 0x00, 0x21, 0x26, 0xda, 0xc9, 0x2b, 0xca, 0xaf, 0x98, 0x9f, 0x65, 0x70, 0xb8,
	0x2f, 0x42, 0xdb, 0x11, 0x79, 0x86, 0x26, 0x79, 0xc0, 0xe8, 0x0c, 0x36, 0x6c, 0x8f, 0xc7, 0xbe,
	0xe8, 0x56, 0x32, 0xe5, 0x54, 0xb9, 0xb9, 0xdb, 0x93, 0x7e, 0xdf, 0xed, 0xbd, 0x74, 0x57, 0x62,
	0x19, 0xcf, 0x15, 0x87, 0x7b, 0xaa, 0xc3, 0x23, 0x8f, 0x47, 0xe5, 0x71, 0x14, 0x2d, 0xae, 0x54,
	0x91, 0x06, 0x2c, 0x52, 0x0c, 0x5f, 0x90, 0xb2, 0xfa, 0xf0, 0x27, 0x80, 0xcd, 0xbc, 0x77, 0x9a,
	0x06, 0x0c, 0xf5, 0x60, 0x07, 0x9f, 0xe3, 0x31, 0xb5, 0xe8, 0xe5, 0x14, 0x5b, 0xb3, 0xb1, 0x39,
	0xc5, 0x9a, 0x71, 0x66, 0x60, 0xbd, 0x25, 0xa1, 0x0e, 0x44, 0x5b, 0x9a, 0x8e, 0xa7, 0x13, 0xd3,
	0xa0, 0x2d, 0x80, 0x76, 0xe1, 0xe3, 0x2d, 0xfe, 0xc2, 0xa0, 0xef, 0x74, 0x32, 0xbc, 0x68, 0x55,
	0xd0, 0x73, 0xd8, 0xdf, 0x12, 0xb4, 0x89, 0xf9, 0x7e, 0x62, 0x5a, 0x79, 0x7b, 0x59, 0xf5, 0x68,
	0x72, 0x89, 0xf5, 0x56, 0x15, 0xbd, 0x80, 0xfb, 0xff, 0xb8, 0xc6, 0x94, 0x0c, 0x35, 0x6a, 0x69,
	0xc3, 0xd1, 0xc8, 0xc2, 0x1f, 0xb1, 0x36, 0xa3, 0x58, 0x6f, 0xd5, 0xd0, 0x3e, 0x7c, 0xb6, 0x65,
	0x33, 0x8d, 0xb7, 0x63, 0x4c, 0x2c, 0x13, 0x53, 0x6b, 0x36, 0xd5, 0x87, 0x99, 0xa5, 0xde, 0xab,
	0x7d, 0xfd, 0x21, 0x4b, 0xa7, 0x1f, 0x6e, 0xd6, 0x32, 0xb8, 0x5d, 0xcb, 0xe0, 0xcf, 0x5a, 0x06,
	0xdf, 0x37, 0xb2, 0x74, 0xbb, 0x91, 0xa5, 0x5f, 0x1b, 0x59, 0xfa, 0xf4, 0xe6, 0xff, 0xd1, 0x94,
	0xdb, 0x73, 0x34, 0x0f, 0x57, 0x0b, 0x97, 0xa9, 0x1e, 0x5f, 0xc4, 0xd7, 0x4c, 0xfd, 0x72, 0xcf,
	0x17, 0xf3, 0x9a, 0x37, 0xf2, 0xaf, 0x7a, 0xfd, 0x37, 0x00, 0x00, 0xff, 0xff, 0x27, 0x84, 0x27,
	0x2e, 0x8b, 0x02, 0x00, 0x00,
}

func (m *EthereumEventVoteRecord) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumEventVoteRecord) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthereumEventVoteRecord) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Event != nil {
		{
			size, err := m.Event.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAttestation(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintAttestation(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Votes[iNdEx])
			copy(dAtA[i:], m.Votes[iNdEx])
			i = encodeVarintAttestation(dAtA, i, uint64(len(m.Votes[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Accepted {
		i--
		if m.Accepted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ERC20Token) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ERC20Token) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ERC20Token) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAttestation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintAttestation(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAttestation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAttestation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthereumEventVoteRecord) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Accepted {
		n += 2
	}
	if len(m.Votes) > 0 {
		for _, s := range m.Votes {
			l = len(s)
			n += 1 + l + sovAttestation(uint64(l))
		}
	}
	if m.Height != 0 {
		n += 1 + sovAttestation(uint64(m.Height))
	}
	if m.Event != nil {
		l = m.Event.Size()
		n += 1 + l + sovAttestation(uint64(l))
	}
	return n
}

func (m *ERC20Token) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovAttestation(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAttestation(uint64(l))
	return n
}

func sovAttestation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAttestation(x uint64) (n int) {
	return sovAttestation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthereumEventVoteRecord) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: EthereumEventVoteRecord: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumEventVoteRecord: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accepted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Accepted = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Event", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Event == nil {
				m.Event = &types.Any{}
			}
			if err := m.Event.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func (m *ERC20Token) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAttestation
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
			return fmt.Errorf("proto: ERC20Token: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ERC20Token: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAttestation
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
				return ErrInvalidLengthAttestation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAttestation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAttestation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAttestation
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
func skipAttestation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
					return 0, ErrIntOverflowAttestation
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
				return 0, ErrInvalidLengthAttestation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAttestation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAttestation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAttestation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAttestation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAttestation = fmt.Errorf("proto: unexpected end of group")
)