// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto

package v2

import (
	fmt "fmt"
	io "io"
	math "math"

	proto "github.com/gogo/protobuf/proto"
	_ "github.com/lyft/protoc-gen-validate/validate"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// [#protodoc-title: Extensions gRPC Http1 Reverse Bridge]
// gRPC reverse bridge filter configuration
type FilterConfig struct {
	// The content-type to pass to the upstream when the gRPC bridge filter is applied.
	// The filter will also validate that the upstream responds with the same content type.
	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// If true, Envoy will assume that the upstream doesn't understand gRPC frames and
	// strip the gRPC frame from the request, and add it back in to the response. This will
	// hide the gRPC semantics from the upstream, allowing it to receive and respond with a
	// simple binary encoded protobuf.
	WithholdGrpcFrames   bool     `protobuf:"varint,2,opt,name=withhold_grpc_frames,json=withholdGrpcFrames,proto3" json:"withhold_grpc_frames,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterConfig) Reset()         { *m = FilterConfig{} }
func (m *FilterConfig) String() string { return proto.CompactTextString(m) }
func (*FilterConfig) ProtoMessage()    {}
func (*FilterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a86db517160ad0a, []int{0}
}
func (m *FilterConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FilterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FilterConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FilterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterConfig.Merge(m, src)
}
func (m *FilterConfig) XXX_Size() int {
	return m.Size()
}
func (m *FilterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_FilterConfig proto.InternalMessageInfo

func (m *FilterConfig) GetContentType() string {
	if m != nil {
		return m.ContentType
	}
	return ""
}

func (m *FilterConfig) GetWithholdGrpcFrames() bool {
	if m != nil {
		return m.WithholdGrpcFrames
	}
	return false
}

func init() {
	proto.RegisterType((*FilterConfig)(nil), "envoy.extensions.filter.http.grpc_http1_reverse_bridge.v2alpha1.FilterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/grpc_http1_reverse_bridge/v2alpha1/config.proto", fileDescriptor_2a86db517160ad0a)
}

var fileDescriptor_2a86db517160ad0a = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x8f, 0xb1, 0x4a, 0xc3, 0x40,
	0x18, 0xc7, 0xb9, 0x28, 0x62, 0xd3, 0x4e, 0x41, 0xb0, 0x38, 0x84, 0xe0, 0xd4, 0x41, 0xee, 0x6c,
	0x7d, 0x00, 0x21, 0x42, 0x75, 0x11, 0x4a, 0x71, 0x72, 0x09, 0x69, 0xf2, 0x25, 0x39, 0x88, 0x77,
	0xc7, 0x97, 0x8f, 0xb3, 0x79, 0x35, 0x27, 0x47, 0x47, 0x1f, 0x41, 0xb2, 0xf9, 0x16, 0x92, 0x4b,
	0x33, 0x3a, 0x75, 0xfb, 0xc3, 0x8f, 0xfb, 0xdd, 0xf7, 0xf3, 0x9f, 0x40, 0x59, 0xdd, 0x8a, 0x4c,
	0xab, 0x42, 0x96, 0xa2, 0x90, 0x35, 0x01, 0x8a, 0x8a, 0xc8, 0x88, 0x12, 0x4d, 0x96, 0xf4, 0x6b,
	0x99, 0x20, 0x58, 0xc0, 0x06, 0x92, 0x1d, 0xca, 0xbc, 0x04, 0x61, 0x57, 0x69, 0x6d, 0xaa, 0x74,
	0x79, 0x78, 0xc5, 0x0d, 0x6a, 0xd2, 0xc1, 0xbd, 0x33, 0x71, 0xd8, 0x13, 0xa8, 0x46, 0x6a, 0xd5,
	0xf0, 0xc1, 0xc6, 0x7b, 0x07, 0xff, 0xd7, 0xc6, 0x47, 0xdb, 0xd5, 0xa5, 0x4d, 0x6b, 0x99, 0xa7,
	0x04, 0x62, 0x1c, 0x83, 0xf9, 0x5a, 0xf9, 0xb3, 0xb5, 0x53, 0x3d, 0xb8, 0xff, 0x82, 0x1b, 0x7f,
	0x96, 0x69, 0x45, 0xa0, 0x28, 0xa1, 0xd6, 0xc0, 0x9c, 0x45, 0x6c, 0x31, 0x89, 0x27, 0x1f, 0xbf,
	0x9f, 0x27, 0xa7, 0xe8, 0x45, 0x6c, 0x3b, 0x3d, 0xe0, 0x97, 0xd6, 0x40, 0x70, 0xeb, 0x5f, 0xbc,
	0x4b, 0xaa, 0x2a, 0x5d, 0xe7, 0x89, 0xbb, 0xa2, 0xc0, 0xf4, 0x0d, 0x9a, 0xb9, 0x17, 0xb1, 0xc5,
	0xf9, 0x36, 0x18, 0xd9, 0x23, 0x9a, 0x6c, 0xed, 0x48, 0xac, 0xbe, 0xba, 0x90, 0x7d, 0x77, 0x21,
	0xfb, 0xe9, 0x42, 0xe6, 0x3f, 0x4b, 0xcd, 0x5d, 0x9a, 0x41, 0xbd, 0x6f, 0xf9, 0x91, 0x95, 0xf1,
	0x74, 0x88, 0xd8, 0xf4, 0x65, 0x1b, 0xf6, 0xea, 0xd9, 0xd5, 0xee, 0xcc, 0x65, 0xde, 0xfd, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x3b, 0xeb, 0xc5, 0x93, 0x8c, 0x01, 0x00, 0x00,
}

func (m *FilterConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FilterConfig) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ContentType) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintConfig(dAtA, i, uint64(len(m.ContentType)))
		i += copy(dAtA[i:], m.ContentType)
	}
	if m.WithholdGrpcFrames {
		dAtA[i] = 0x10
		i++
		if m.WithholdGrpcFrames {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintConfig(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *FilterConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContentType)
	if l > 0 {
		n += 1 + l + sovConfig(uint64(l))
	}
	if m.WithholdGrpcFrames {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovConfig(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozConfig(x uint64) (n int) {
	return sovConfig(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FilterConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConfig
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
			return fmt.Errorf("proto: FilterConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FilterConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
				return ErrInvalidLengthConfig
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConfig
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field WithholdGrpcFrames", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConfig
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
			m.WithholdGrpcFrames = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipConfig(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthConfig
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConfig(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
					return 0, ErrIntOverflowConfig
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
				return 0, ErrInvalidLengthConfig
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthConfig
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowConfig
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
				next, err := skipConfig(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthConfig
				}
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
	ErrInvalidLengthConfig = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConfig   = fmt.Errorf("proto: integer overflow")
)
