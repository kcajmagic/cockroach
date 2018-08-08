// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/closedts/ctpb/entry.proto

/*
	Package ctpb is a generated protocol buffer package.

	It is generated from these files:
		storage/closedts/ctpb/entry.proto

	It has these top-level messages:
		Entry
		Reaction
*/
package ctpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"

import context "context"
import grpc "google.golang.org/grpc"

import sortkeys "github.com/gogo/protobuf/sortkeys"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// An Entry is a closed timestamp update. It consists of a closed timestamp
// (i.e. a timestamp at or below which the origin node guarantees no more new
// writes are going to be permitted), an associated epoch in which the origin
// node promises it was live (for the closed timestamp), a map of minimum lease
// applied indexes (which have to be caught up to before being allowed to use
// the closed timestamp) as well as an indicator of whether this update supplies
// a full initial state or an increment to be merged into a previous state. In
// practice, the first Entry received for each epoch is full, while the remainder
// are incremental. An incremental update represents the implicit promise that
// the state accumulated since the last full Entry is the true full state.
type Entry struct {
	Epoch           Epoch                                                        `protobuf:"varint,1,opt,name=epoch,proto3,casttype=Epoch" json:"epoch,omitempty"`
	ClosedTimestamp cockroach_util_hlc.Timestamp                                 `protobuf:"bytes,2,opt,name=closed_timestamp,json=closedTimestamp" json:"closed_timestamp"`
	MLAI            map[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID]LAI `protobuf:"bytes,3,rep,name=mlai,castkey=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID,castvalue=LAI" json:"mlai,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Full is true if the emitter promises that any future write to any range
	// mentioned in this Entry will be reflected in a subsequent Entry before any
	// stale follower reads are possible. For example, if range 1 is assigned an
	// MLAI of 12 in this Entry and isn't mentioned in the five subsequent
	// entries, the recipient may behave as if the MLAI of 12 were repeated across
	// all of these entries.
	//
	// In practice, a Full message is received when a stream of Entries is first
	// established (or the Epoch changes), and all other updates are incremental
	// (i.e. not Full).
	Full bool `protobuf:"varint,4,opt,name=full,proto3" json:"full,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptorEntry, []int{0} }

// Reactions flow in the direction opposite to Entries and request for ranges to
// be included in the next Entry. Under rare circumstances, ranges may be omitted
// from closed timestamp updates, and so serving follower reads from them would
// fail. The Reaction mechanism serves to explicitly request the missing information
// when that happens.
type Reaction struct {
	Requested []github_com_cockroachdb_cockroach_pkg_roachpb.RangeID `protobuf:"varint,1,rep,packed,name=Requested,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"Requested,omitempty"`
}

func (m *Reaction) Reset()                    { *m = Reaction{} }
func (*Reaction) ProtoMessage()               {}
func (*Reaction) Descriptor() ([]byte, []int) { return fileDescriptorEntry, []int{1} }

func init() {
	proto.RegisterType((*Entry)(nil), "cockroach.storage.ctupdate.Entry")
	proto.RegisterType((*Reaction)(nil), "cockroach.storage.ctupdate.Reaction")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ClosedTimestamp service

type ClosedTimestampClient interface {
	Get(ctx context.Context, opts ...grpc.CallOption) (ClosedTimestamp_GetClient, error)
}

type closedTimestampClient struct {
	cc *grpc.ClientConn
}

func NewClosedTimestampClient(cc *grpc.ClientConn) ClosedTimestampClient {
	return &closedTimestampClient{cc}
}

func (c *closedTimestampClient) Get(ctx context.Context, opts ...grpc.CallOption) (ClosedTimestamp_GetClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ClosedTimestamp_serviceDesc.Streams[0], c.cc, "/cockroach.storage.ctupdate.ClosedTimestamp/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &closedTimestampGetClient{stream}
	return x, nil
}

type ClosedTimestamp_GetClient interface {
	Send(*Reaction) error
	Recv() (*Entry, error)
	grpc.ClientStream
}

type closedTimestampGetClient struct {
	grpc.ClientStream
}

func (x *closedTimestampGetClient) Send(m *Reaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *closedTimestampGetClient) Recv() (*Entry, error) {
	m := new(Entry)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ClosedTimestamp service

type ClosedTimestampServer interface {
	Get(ClosedTimestamp_GetServer) error
}

func RegisterClosedTimestampServer(s *grpc.Server, srv ClosedTimestampServer) {
	s.RegisterService(&_ClosedTimestamp_serviceDesc, srv)
}

func _ClosedTimestamp_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClosedTimestampServer).Get(&closedTimestampGetServer{stream})
}

type ClosedTimestamp_GetServer interface {
	Send(*Entry) error
	Recv() (*Reaction, error)
	grpc.ServerStream
}

type closedTimestampGetServer struct {
	grpc.ServerStream
}

func (x *closedTimestampGetServer) Send(m *Entry) error {
	return x.ServerStream.SendMsg(m)
}

func (x *closedTimestampGetServer) Recv() (*Reaction, error) {
	m := new(Reaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClosedTimestamp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.storage.ctupdate.ClosedTimestamp",
	HandlerType: (*ClosedTimestampServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _ClosedTimestamp_Get_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "storage/closedts/ctpb/entry.proto",
}

func (m *Entry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Entry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Epoch != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintEntry(dAtA, i, uint64(m.Epoch))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintEntry(dAtA, i, uint64(m.ClosedTimestamp.Size()))
	n1, err := m.ClosedTimestamp.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	if len(m.MLAI) > 0 {
		keysForMLAI := make([]int32, 0, len(m.MLAI))
		for k := range m.MLAI {
			keysForMLAI = append(keysForMLAI, int32(k))
		}
		sortkeys.Int32s(keysForMLAI)
		for _, k := range keysForMLAI {
			dAtA[i] = 0x1a
			i++
			v := m.MLAI[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(k)]
			mapSize := 1 + sovEntry(uint64(k)) + 1 + sovEntry(uint64(v))
			i = encodeVarintEntry(dAtA, i, uint64(mapSize))
			dAtA[i] = 0x8
			i++
			i = encodeVarintEntry(dAtA, i, uint64(k))
			dAtA[i] = 0x10
			i++
			i = encodeVarintEntry(dAtA, i, uint64(v))
		}
	}
	if m.Full {
		dAtA[i] = 0x20
		i++
		if m.Full {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	return i, nil
}

func (m *Reaction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Reaction) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Requested) > 0 {
		dAtA3 := make([]byte, len(m.Requested)*10)
		var j2 int
		for _, num1 := range m.Requested {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA3[j2] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j2++
			}
			dAtA3[j2] = uint8(num)
			j2++
		}
		dAtA[i] = 0xa
		i++
		i = encodeVarintEntry(dAtA, i, uint64(j2))
		i += copy(dAtA[i:], dAtA3[:j2])
	}
	return i, nil
}

func encodeVarintEntry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Entry) Size() (n int) {
	var l int
	_ = l
	if m.Epoch != 0 {
		n += 1 + sovEntry(uint64(m.Epoch))
	}
	l = m.ClosedTimestamp.Size()
	n += 1 + l + sovEntry(uint64(l))
	if len(m.MLAI) > 0 {
		for k, v := range m.MLAI {
			_ = k
			_ = v
			mapEntrySize := 1 + sovEntry(uint64(k)) + 1 + sovEntry(uint64(v))
			n += mapEntrySize + 1 + sovEntry(uint64(mapEntrySize))
		}
	}
	if m.Full {
		n += 2
	}
	return n
}

func (m *Reaction) Size() (n int) {
	var l int
	_ = l
	if len(m.Requested) > 0 {
		l = 0
		for _, e := range m.Requested {
			l += sovEntry(uint64(e))
		}
		n += 1 + sovEntry(uint64(l)) + l
	}
	return n
}

func sovEntry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEntry(x uint64) (n int) {
	return sovEntry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Entry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Entry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Entry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= (Epoch(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClosedTimestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClosedTimestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MLAI", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthEntry
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.MLAI == nil {
				m.MLAI = make(map[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID]LAI)
			}
			var mapkey int32
			var mapvalue int64
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapkey |= (int32(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else if fieldNum == 2 {
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapvalue |= (int64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEntry(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthEntry
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.MLAI[github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(mapkey)] = ((LAI)(mapvalue))
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Full", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Full = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntry
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
func (m *Reaction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEntry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Reaction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Reaction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Requested = append(m.Requested, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthEntry
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v github_com_cockroachdb_cockroach_pkg_roachpb.RangeID
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEntry
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Requested = append(m.Requested, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Requested", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEntry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEntry
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
func skipEntry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEntry
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
					return 0, ErrIntOverflowEntry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowEntry
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthEntry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEntry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipEntry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthEntry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEntry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("storage/closedts/ctpb/entry.proto", fileDescriptorEntry) }

var fileDescriptorEntry = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x4f, 0x6b, 0xd4, 0x40,
	0x1c, 0xcd, 0x6c, 0x12, 0xe9, 0x4e, 0x0f, 0x2d, 0x43, 0x0f, 0x21, 0x68, 0x92, 0x2e, 0x1e, 0x02,
	0xc2, 0x8c, 0xac, 0x82, 0xa5, 0xb7, 0x46, 0x4b, 0x59, 0x68, 0x3d, 0x0c, 0xc5, 0x83, 0x17, 0x99,
	0xcc, 0x8e, 0x49, 0xd8, 0xd9, 0x4c, 0xdc, 0x4c, 0x84, 0x5e, 0x3d, 0x89, 0x27, 0x8f, 0x1e, 0xfd,
	0x38, 0x7b, 0xf4, 0x24, 0x9e, 0xb6, 0x1a, 0xbf, 0x45, 0x4f, 0x92, 0x49, 0x77, 0x17, 0x04, 0x15,
	0xbc, 0xbd, 0xfc, 0xfe, 0xbc, 0xf7, 0xf2, 0xe6, 0x07, 0x0f, 0x6b, 0xad, 0x16, 0x2c, 0x13, 0x84,
	0x4b, 0x55, 0x8b, 0xa9, 0xae, 0x09, 0xd7, 0x55, 0x4a, 0x44, 0xa9, 0x17, 0x57, 0xb8, 0x5a, 0x28,
	0xad, 0x90, 0xcf, 0x15, 0x9f, 0x2d, 0x14, 0xe3, 0x39, 0xbe, 0x1d, 0xc6, 0x5c, 0x37, 0xd5, 0x94,
	0x69, 0xe1, 0x1f, 0x64, 0x2a, 0x53, 0x66, 0x8c, 0x74, 0xa8, 0xdf, 0xf0, 0xef, 0x66, 0x4a, 0x65,
	0x52, 0x10, 0x56, 0x15, 0x84, 0x95, 0xa5, 0xd2, 0x4c, 0x17, 0xaa, 0xac, 0x6f, 0xbb, 0x5e, 0xa3,
	0x0b, 0x49, 0x72, 0xc9, 0x89, 0x2e, 0xe6, 0xa2, 0xd6, 0x6c, 0x5e, 0xf5, 0x9d, 0xd1, 0xd7, 0x01,
	0x74, 0x4f, 0x3b, 0x65, 0x14, 0x42, 0x57, 0x54, 0x8a, 0xe7, 0x1e, 0x88, 0x40, 0x6c, 0x27, 0xc3,
	0x9b, 0x55, 0xe8, 0x9e, 0x76, 0x05, 0xda, 0xd7, 0xd1, 0x73, 0xb8, 0xdf, 0x3b, 0x7e, 0xb5, 0x21,
	0xf1, 0x06, 0x11, 0x88, 0x77, 0xc7, 0xf7, 0xf0, 0xd6, 0x6f, 0xa7, 0x84, 0x73, 0xc9, 0xf1, 0xe5,
	0x7a, 0x28, 0x71, 0x96, 0xab, 0xd0, 0xa2, 0x7b, 0xfd, 0xf2, 0xa6, 0x8c, 0xde, 0x03, 0xe8, 0xcc,
	0x25, 0x2b, 0x3c, 0x3b, 0xb2, 0xe3, 0xdd, 0xf1, 0x03, 0xfc, 0xe7, 0x9f, 0xc6, 0xc6, 0x22, 0xbe,
	0x90, 0xac, 0x30, 0x28, 0x39, 0x6b, 0x57, 0xa1, 0x73, 0x71, 0x7e, 0x32, 0x79, 0x77, 0x1d, 0x3e,
	0xce, 0x0a, 0x9d, 0x37, 0x29, 0xe6, 0x6a, 0x4e, 0x36, 0x14, 0xd3, 0x74, 0x8b, 0x49, 0x35, 0xcb,
	0x88, 0x41, 0x55, 0x8a, 0x29, 0x2b, 0x33, 0x31, 0x79, 0xf6, 0xe1, 0x3a, 0xb4, 0xcf, 0x4f, 0x26,
	0xd4, 0x38, 0x40, 0x08, 0x3a, 0xaf, 0x1b, 0x29, 0x3d, 0x27, 0x02, 0xf1, 0x0e, 0x35, 0xd8, 0x7f,
	0x02, 0x87, 0x1b, 0x3d, 0xb4, 0x0f, 0xed, 0x99, 0xb8, 0x32, 0xd1, 0xb8, 0xb4, 0x83, 0xe8, 0x00,
	0xba, 0x6f, 0x99, 0x6c, 0x84, 0x89, 0xc0, 0xa6, 0xfd, 0xc7, 0xf1, 0xe0, 0x08, 0x1c, 0x3b, 0x9f,
	0x3e, 0x87, 0xd6, 0x28, 0x87, 0x3b, 0x54, 0x30, 0xde, 0xbd, 0x02, 0x7a, 0x01, 0x87, 0x54, 0xbc,
	0x69, 0x44, 0xad, 0xc5, 0xd4, 0x03, 0x91, 0x1d, 0xbb, 0xc9, 0xd1, 0xcd, 0xea, 0xff, 0x8c, 0xd3,
	0x2d, 0x55, 0xaf, 0x34, 0xce, 0xe0, 0xde, 0xd3, 0xdf, 0xa2, 0xbd, 0x84, 0xf6, 0x99, 0xd0, 0xe8,
	0xfe, 0xdf, 0x22, 0x5d, 0xbb, 0xf3, 0x0f, 0xff, 0x19, 0xfc, 0xc8, 0x8a, 0xc1, 0x43, 0x90, 0x04,
	0xcb, 0x1f, 0x81, 0xb5, 0x6c, 0x03, 0xf0, 0xa5, 0x0d, 0xc0, 0xb7, 0x36, 0x00, 0xdf, 0xdb, 0x00,
	0x7c, 0xfc, 0x19, 0x58, 0x2f, 0x9d, 0xee, 0x82, 0xd3, 0x3b, 0xe6, 0xa4, 0x1e, 0xfd, 0x0a, 0x00,
	0x00, 0xff, 0xff, 0x30, 0xde, 0x98, 0x28, 0xe1, 0x02, 0x00, 0x00,
}