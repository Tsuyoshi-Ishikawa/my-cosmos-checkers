// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: checkers/stored_game.proto

package types

import (
	fmt "fmt"
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

type StoredGame struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Creator   string `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Game      string `protobuf:"bytes,3,opt,name=game,proto3" json:"game,omitempty"`
	Turn      string `protobuf:"bytes,4,opt,name=turn,proto3" json:"turn,omitempty"`
	Red       string `protobuf:"bytes,5,opt,name=red,proto3" json:"red,omitempty"`
	Black     string `protobuf:"bytes,6,opt,name=black,proto3" json:"black,omitempty"`
	MoveCount uint64 `protobuf:"varint,7,opt,name=moveCount,proto3" json:"moveCount,omitempty"`
	BeforeId  string `protobuf:"bytes,8,opt,name=beforeId,proto3" json:"beforeId,omitempty"`
	AfterId   string `protobuf:"bytes,9,opt,name=afterId,proto3" json:"afterId,omitempty"`
	Deadline  string `protobuf:"bytes,10,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Winner    string `protobuf:"bytes,11,opt,name=winner,proto3" json:"winner,omitempty"`
	Wager     uint64 `protobuf:"varint,12,opt,name=wager,proto3" json:"wager,omitempty"`
	Token     string `protobuf:"bytes,13,opt,name=token,proto3" json:"token,omitempty"`
}

func (m *StoredGame) Reset()         { *m = StoredGame{} }
func (m *StoredGame) String() string { return proto.CompactTextString(m) }
func (*StoredGame) ProtoMessage()    {}
func (*StoredGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_8439c9c90688ff75, []int{0}
}
func (m *StoredGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredGame.Merge(m, src)
}
func (m *StoredGame) XXX_Size() int {
	return m.Size()
}
func (m *StoredGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredGame.DiscardUnknown(m)
}

var xxx_messageInfo_StoredGame proto.InternalMessageInfo

func (m *StoredGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredGame) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *StoredGame) GetGame() string {
	if m != nil {
		return m.Game
	}
	return ""
}

func (m *StoredGame) GetTurn() string {
	if m != nil {
		return m.Turn
	}
	return ""
}

func (m *StoredGame) GetRed() string {
	if m != nil {
		return m.Red
	}
	return ""
}

func (m *StoredGame) GetBlack() string {
	if m != nil {
		return m.Black
	}
	return ""
}

func (m *StoredGame) GetMoveCount() uint64 {
	if m != nil {
		return m.MoveCount
	}
	return 0
}

func (m *StoredGame) GetBeforeId() string {
	if m != nil {
		return m.BeforeId
	}
	return ""
}

func (m *StoredGame) GetAfterId() string {
	if m != nil {
		return m.AfterId
	}
	return ""
}

func (m *StoredGame) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

func (m *StoredGame) GetWinner() string {
	if m != nil {
		return m.Winner
	}
	return ""
}

func (m *StoredGame) GetWager() uint64 {
	if m != nil {
		return m.Wager
	}
	return 0
}

func (m *StoredGame) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*StoredGame)(nil), "alice.checkers.checkers.StoredGame")
}

func init() { proto.RegisterFile("checkers/stored_game.proto", fileDescriptor_8439c9c90688ff75) }

var fileDescriptor_8439c9c90688ff75 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x91, 0xbd, 0x4e, 0x82, 0x31,
	0x14, 0x86, 0x29, 0xff, 0x54, 0x4d, 0x4c, 0x63, 0xf4, 0x84, 0x98, 0x86, 0x38, 0x11, 0x07, 0x18,
	0xbc, 0x03, 0x35, 0x31, 0xac, 0xb8, 0xb9, 0x98, 0x7e, 0xed, 0x01, 0xbe, 0xc0, 0xd7, 0x92, 0x52,
	0x04, 0xef, 0xc2, 0x4b, 0x72, 0x74, 0x64, 0x74, 0x34, 0x70, 0x23, 0xa6, 0xa7, 0xfc, 0x6c, 0xef,
	0xf3, 0xb4, 0x27, 0xe7, 0x4d, 0x0e, 0x6f, 0xeb, 0x09, 0xea, 0x29, 0xfa, 0x45, 0x7f, 0x11, 0x9c,
	0x47, 0xf3, 0x3e, 0x56, 0x05, 0xf6, 0xe6, 0xde, 0x05, 0x27, 0x6e, 0xd4, 0x2c, 0xd7, 0xd8, 0x3b,
	0xfc, 0x38, 0x86, 0xbb, 0xef, 0x32, 0xe7, 0xaf, 0xf4, 0xfd, 0x45, 0x15, 0x28, 0xae, 0x78, 0x2d,
	0xb7, 0x06, 0xd7, 0xc0, 0x3a, 0xac, 0xdb, 0x1a, 0x26, 0x10, 0xc0, 0x1b, 0xda, 0xa3, 0x0a, 0xce,
	0x43, 0x99, 0xfc, 0x01, 0x85, 0xe0, 0xd5, 0xb8, 0x05, 0x2a, 0xa4, 0x29, 0x47, 0x17, 0x96, 0xde,
	0x42, 0x35, 0xb9, 0x98, 0xc5, 0x25, 0xaf, 0x78, 0x34, 0x50, 0x23, 0x15, 0x63, 0xdc, 0x94, 0xcd,
	0x94, 0x9e, 0x42, 0x3d, 0x6d, 0x22, 0x10, 0xb7, 0xbc, 0x55, 0xb8, 0x0f, 0x7c, 0x72, 0x4b, 0x1b,
	0xa0, 0xd1, 0x61, 0xdd, 0xea, 0xf0, 0x24, 0x44, 0x9b, 0x37, 0x33, 0x1c, 0x39, 0x8f, 0x03, 0x03,
	0x4d, 0x1a, 0x3b, 0x72, 0xec, 0xa8, 0x46, 0x01, 0xfd, 0xc0, 0x40, 0x2b, 0x75, 0xdc, 0x63, 0x9c,
	0x32, 0xa8, 0xcc, 0x2c, 0xb7, 0x08, 0x3c, 0x4d, 0x1d, 0x58, 0x5c, 0xf3, 0xfa, 0x2a, 0xb7, 0x16,
	0x3d, 0x9c, 0xd1, 0xcb, 0x9e, 0x62, 0xbb, 0x95, 0x1a, 0xa3, 0x87, 0x73, 0xea, 0x90, 0x20, 0xda,
	0xe0, 0xa6, 0x68, 0xe1, 0x22, 0x75, 0x26, 0x78, 0x7c, 0xfe, 0xd9, 0x4a, 0xb6, 0xd9, 0x4a, 0xf6,
	0xb7, 0x95, 0xec, 0x6b, 0x27, 0x4b, 0x9b, 0x9d, 0x2c, 0xfd, 0xee, 0x64, 0xe9, 0xed, 0x7e, 0x9c,
	0x87, 0xc9, 0x32, 0xeb, 0x69, 0x57, 0xf4, 0xe9, 0x00, 0xfd, 0xe3, 0x89, 0xd6, 0xa7, 0x18, 0x3e,
	0xe7, 0xb8, 0xc8, 0xea, 0x74, 0xa8, 0x87, 0xff, 0x00, 0x00, 0x00, 0xff, 0xff, 0x88, 0x28, 0xd1,
	0xbe, 0xc6, 0x01, 0x00, 0x00,
}

func (m *StoredGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x6a
	}
	if m.Wager != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Wager))
		i--
		dAtA[i] = 0x60
	}
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.AfterId) > 0 {
		i -= len(m.AfterId)
		copy(dAtA[i:], m.AfterId)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.AfterId)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.BeforeId) > 0 {
		i -= len(m.BeforeId)
		copy(dAtA[i:], m.BeforeId)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.BeforeId)))
		i--
		dAtA[i] = 0x42
	}
	if m.MoveCount != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.MoveCount))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Black) > 0 {
		i -= len(m.Black)
		copy(dAtA[i:], m.Black)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Black)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Red) > 0 {
		i -= len(m.Red)
		copy(dAtA[i:], m.Red)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Red)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Turn) > 0 {
		i -= len(m.Turn)
		copy(dAtA[i:], m.Turn)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Turn)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Game) > 0 {
		i -= len(m.Game)
		copy(dAtA[i:], m.Game)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Game)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Game)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Turn)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Red)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Black)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.MoveCount != 0 {
		n += 1 + sovStoredGame(uint64(m.MoveCount))
	}
	l = len(m.BeforeId)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.AfterId)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.Wager != 0 {
		n += 1 + sovStoredGame(uint64(m.Wager))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	return n
}

func sovStoredGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredGame(x uint64) (n int) {
	return sovStoredGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredGame
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
			return fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Game", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Game = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Turn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Red", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Red = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Black", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Black = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MoveCount", wireType)
			}
			m.MoveCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MoveCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BeforeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BeforeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AfterId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AfterId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Wager", wireType)
			}
			m.Wager = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Wager |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStoredGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredGame
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
func skipStoredGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
				return 0, ErrInvalidLengthStoredGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredGame = fmt.Errorf("proto: unexpected end of group")
)
