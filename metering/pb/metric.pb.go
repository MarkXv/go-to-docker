// Code generated by protoc-gen-gogo.
// source: pb/metric.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type KnownMeter int32

const (
	KnownMeter_CADVISOR   KnownMeter = 0
	KnownMeter_DOCKER     KnownMeter = 1
	KnownMeter_HEAPSTER   KnownMeter = 2
	KnownMeter_PROMETHEUS KnownMeter = 3
)

var KnownMeter_name = map[int32]string{
	0: "CADVISOR",
	1: "DOCKER",
	2: "HEAPSTER",
	3: "PROMETHEUS",
}
var KnownMeter_value = map[string]int32{
	"CADVISOR":   0,
	"DOCKER":     1,
	"HEAPSTER":   2,
	"PROMETHEUS": 3,
}

func (x KnownMeter) String() string {
	return proto.EnumName(KnownMeter_name, int32(x))
}
func (KnownMeter) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

type RequestType int32

const (
	RequestType_CADVISOR_V1_ContainerInfoRequest RequestType = 0
	RequestType_CADVISOR_V2_ContainerInfoRequest RequestType = 1
)

var RequestType_name = map[int32]string{
	0: "CADVISOR_V1_ContainerInfoRequest",
	1: "CADVISOR_V2_ContainerInfoRequest",
}
var RequestType_value = map[string]int32{
	"CADVISOR_V1_ContainerInfoRequest": 0,
	"CADVISOR_V2_ContainerInfoRequest": 1,
}

func (x RequestType) String() string {
	return proto.EnumName(RequestType_name, int32(x))
}
func (RequestType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetric, []int{1} }

type MetricType int32

const (
	MetricType_CADVISOR_V1_CpuUsage MetricType = 0
	MetricType_CADVISOR_V1_CpuCFS   MetricType = 1
)

var MetricType_name = map[int32]string{
	0: "CADVISOR_V1_CpuUsage",
	1: "CADVISOR_V1_CpuCFS",
}
var MetricType_value = map[string]int32{
	"CADVISOR_V1_CpuUsage": 0,
	"CADVISOR_V1_CpuCFS":   1,
}

func (x MetricType) String() string {
	return proto.EnumName(MetricType_name, int32(x))
}
func (MetricType) EnumDescriptor() ([]byte, []int) { return fileDescriptorMetric, []int{2} }

type MetricReqResp struct {
	Meter       KnownMeter  `protobuf:"varint,1,opt,name=meter,proto3,enum=pb.KnownMeter" json:"meter,omitempty"`
	RequestType RequestType `protobuf:"varint,2,opt,name=request_type,json=requestType,proto3,enum=pb.RequestType" json:"request_type,omitempty"`
	RequestInfo []byte      `protobuf:"bytes,3,opt,name=request_info,json=requestInfo,proto3" json:"request_info,omitempty"`
	MetricType  MetricType  `protobuf:"varint,4,opt,name=metric_type,json=metricType,proto3,enum=pb.MetricType" json:"metric_type,omitempty"`
	MetricData  []byte      `protobuf:"bytes,5,opt,name=metric_data,json=metricData,proto3" json:"metric_data,omitempty"`
}

func (m *MetricReqResp) Reset()                    { *m = MetricReqResp{} }
func (m *MetricReqResp) String() string            { return proto.CompactTextString(m) }
func (*MetricReqResp) ProtoMessage()               {}
func (*MetricReqResp) Descriptor() ([]byte, []int) { return fileDescriptorMetric, []int{0} }

func (m *MetricReqResp) GetMeter() KnownMeter {
	if m != nil {
		return m.Meter
	}
	return KnownMeter_CADVISOR
}

func (m *MetricReqResp) GetRequestType() RequestType {
	if m != nil {
		return m.RequestType
	}
	return RequestType_CADVISOR_V1_ContainerInfoRequest
}

func (m *MetricReqResp) GetRequestInfo() []byte {
	if m != nil {
		return m.RequestInfo
	}
	return nil
}

func (m *MetricReqResp) GetMetricType() MetricType {
	if m != nil {
		return m.MetricType
	}
	return MetricType_CADVISOR_V1_CpuUsage
}

func (m *MetricReqResp) GetMetricData() []byte {
	if m != nil {
		return m.MetricData
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricReqResp)(nil), "pb.MetricReqResp")
	proto.RegisterEnum("pb.KnownMeter", KnownMeter_name, KnownMeter_value)
	proto.RegisterEnum("pb.RequestType", RequestType_name, RequestType_value)
	proto.RegisterEnum("pb.MetricType", MetricType_name, MetricType_value)
}
func (m *MetricReqResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetricReqResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Meter != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.Meter))
	}
	if m.RequestType != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.RequestType))
	}
	if len(m.RequestInfo) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.RequestInfo)))
		i += copy(dAtA[i:], m.RequestInfo)
	}
	if m.MetricType != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintMetric(dAtA, i, uint64(m.MetricType))
	}
	if len(m.MetricData) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMetric(dAtA, i, uint64(len(m.MetricData)))
		i += copy(dAtA[i:], m.MetricData)
	}
	return i, nil
}

func encodeFixed64Metric(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Metric(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMetric(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MetricReqResp) Size() (n int) {
	var l int
	_ = l
	if m.Meter != 0 {
		n += 1 + sovMetric(uint64(m.Meter))
	}
	if m.RequestType != 0 {
		n += 1 + sovMetric(uint64(m.RequestType))
	}
	l = len(m.RequestInfo)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	if m.MetricType != 0 {
		n += 1 + sovMetric(uint64(m.MetricType))
	}
	l = len(m.MetricData)
	if l > 0 {
		n += 1 + l + sovMetric(uint64(l))
	}
	return n
}

func sovMetric(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMetric(x uint64) (n int) {
	return sovMetric(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MetricReqResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetric
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
			return fmt.Errorf("proto: MetricReqResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetricReqResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meter", wireType)
			}
			m.Meter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Meter |= (KnownMeter(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestType", wireType)
			}
			m.RequestType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestType |= (RequestType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestInfo", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestInfo = append(m.RequestInfo[:0], dAtA[iNdEx:postIndex]...)
			if m.RequestInfo == nil {
				m.RequestInfo = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricType", wireType)
			}
			m.MetricType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MetricType |= (MetricType(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MetricData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetric
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetric
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MetricData = append(m.MetricData[:0], dAtA[iNdEx:postIndex]...)
			if m.MetricData == nil {
				m.MetricData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetric(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetric
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
func skipMetric(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
					return 0, ErrIntOverflowMetric
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
				return 0, ErrInvalidLengthMetric
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetric
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
				next, err := skipMetric(dAtA[start:])
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
	ErrInvalidLengthMetric = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetric   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/metric.proto", fileDescriptorMetric) }

var fileDescriptorMetric = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x6a, 0xf2, 0x40,
	0x14, 0x85, 0x33, 0xfa, 0x2b, 0x3f, 0x57, 0xab, 0xc3, 0xa5, 0x94, 0xac, 0x52, 0x5b, 0x5c, 0x88,
	0x0b, 0xa5, 0x76, 0x5f, 0xb0, 0x49, 0x8a, 0x22, 0xa2, 0x4c, 0x54, 0xe8, 0x4a, 0x92, 0x76, 0x2c,
	0x2e, 0xcc, 0x8c, 0xe3, 0x48, 0xf1, 0x4d, 0xfa, 0x48, 0x5d, 0xf6, 0x09, 0x4a, 0xb1, 0x2f, 0x52,
	0x92, 0xd8, 0xa4, 0x96, 0xee, 0x92, 0x73, 0xbe, 0x7b, 0x38, 0x87, 0x81, 0xaa, 0x0c, 0xda, 0x2b,
	0xae, 0xd5, 0xf2, 0xa1, 0x25, 0x95, 0xd0, 0x02, 0x73, 0x32, 0xb8, 0x7c, 0x27, 0x70, 0x32, 0x8c,
	0x45, 0xc6, 0xd7, 0x8c, 0x6f, 0x24, 0xd6, 0xa1, 0xb0, 0xe2, 0x9a, 0x2b, 0x93, 0xd4, 0x48, 0xa3,
	0xd2, 0xa9, 0xb4, 0x64, 0xd0, 0x1a, 0x84, 0xe2, 0x39, 0x1c, 0x46, 0x2a, 0x4b, 0x4c, 0xec, 0x40,
	0x59, 0xf1, 0xf5, 0x96, 0x6f, 0xf4, 0x5c, 0xef, 0x24, 0x37, 0x73, 0x31, 0x5c, 0x8d, 0x60, 0x96,
	0xe8, 0x93, 0x9d, 0xe4, 0xac, 0xa4, 0xb2, 0x1f, 0xbc, 0xc8, 0x6e, 0x96, 0xe1, 0x42, 0x98, 0xf9,
	0x1a, 0x69, 0x94, 0x53, 0xa4, 0x1f, 0x2e, 0x04, 0xb6, 0xa1, 0x94, 0x54, 0x4c, 0x52, 0xff, 0x65,
	0x15, 0x92, 0x92, 0x71, 0x28, 0xac, 0xd2, 0x6f, 0x3c, 0x4f, 0x0f, 0x1e, 0x7d, 0xed, 0x9b, 0x85,
	0x38, 0xf2, 0x00, 0x38, 0xbe, 0xf6, 0x9b, 0x0e, 0x40, 0xd6, 0x1e, 0xcb, 0xf0, 0xdf, 0xee, 0x3a,
	0xb3, 0xbe, 0x37, 0x62, 0xd4, 0x40, 0x80, 0xa2, 0x33, 0xb2, 0x07, 0x2e, 0xa3, 0x24, 0x72, 0x7a,
	0x6e, 0x77, 0xec, 0x4d, 0x5c, 0x46, 0x73, 0x58, 0x01, 0x18, 0xb3, 0xd1, 0xd0, 0x9d, 0xf4, 0xdc,
	0xa9, 0x47, 0xf3, 0xcd, 0x7b, 0x28, 0xfd, 0x98, 0x85, 0x75, 0xa8, 0x7d, 0xc7, 0xcc, 0x67, 0x57,
	0x73, 0x5b, 0x84, 0xda, 0x5f, 0x86, 0x5c, 0x45, 0x1b, 0x0e, 0x1c, 0x35, 0x8e, 0xa9, 0xce, 0xdf,
	0x14, 0x69, 0xde, 0x00, 0x64, 0xdb, 0xd0, 0x84, 0xd3, 0xa3, 0x64, 0xb9, 0x9d, 0x6e, 0xfc, 0x27,
	0x4e, 0x0d, 0x3c, 0x03, 0xfc, 0xe5, 0xd8, 0x77, 0x1e, 0x25, 0xb7, 0xf4, 0x75, 0x6f, 0x91, 0xb7,
	0xbd, 0x45, 0x3e, 0xf6, 0x16, 0x79, 0xf9, 0xb4, 0x8c, 0xa0, 0x18, 0x3f, 0xef, 0xf5, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xb3, 0xaa, 0x04, 0x0b, 0xf1, 0x01, 0x00, 0x00,
}