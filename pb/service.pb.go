// Code generated by protoc-gen-gogo.
// source: pb/service.proto
// DO NOT EDIT!

/*
	Package pb is a generated protocol buffer package.

	It is generated from these files:
		pb/service.proto

	It has these top-level messages:
		EchoMessage
		DockerRunData
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"
import moby "github.com/tangfeixiong/go-to-docker/pb/moby"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EchoMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *EchoMessage) Reset()                    { *m = EchoMessage{} }
func (m *EchoMessage) String() string            { return proto.CompactTextString(m) }
func (*EchoMessage) ProtoMessage()               {}
func (*EchoMessage) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{0} }

func (m *EchoMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type DockerRunData struct {
	Config        *moby.Config           `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	HostConfig    *moby.HostConfig       `protobuf:"bytes,2,opt,name=host_config,json=hostConfig" json:"host_config,omitempty"`
	NetworkConfig *moby.NetworkingConfig `protobuf:"bytes,3,opt,name=network_config,json=networkConfig" json:"network_config,omitempty"`
	ContainerName string                 `protobuf:"bytes,4,opt,name=container_name,json=containerName,proto3" json:"container_name,omitempty"`
	StateCode     int32                  `protobuf:"varint,5,opt,name=state_code,json=stateCode,proto3" json:"state_code,omitempty"`
	StateMessage  string                 `protobuf:"bytes,6,opt,name=state_message,json=stateMessage,proto3" json:"state_message,omitempty"`
	ContainerId   string                 `protobuf:"bytes,7,opt,name=container_id,json=containerId,proto3" json:"container_id,omitempty"`
}

func (m *DockerRunData) Reset()                    { *m = DockerRunData{} }
func (m *DockerRunData) String() string            { return proto.CompactTextString(m) }
func (*DockerRunData) ProtoMessage()               {}
func (*DockerRunData) Descriptor() ([]byte, []int) { return fileDescriptorService, []int{1} }

func (m *DockerRunData) GetConfig() *moby.Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *DockerRunData) GetHostConfig() *moby.HostConfig {
	if m != nil {
		return m.HostConfig
	}
	return nil
}

func (m *DockerRunData) GetNetworkConfig() *moby.NetworkingConfig {
	if m != nil {
		return m.NetworkConfig
	}
	return nil
}

func (m *DockerRunData) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *DockerRunData) GetStateCode() int32 {
	if m != nil {
		return m.StateCode
	}
	return 0
}

func (m *DockerRunData) GetStateMessage() string {
	if m != nil {
		return m.StateMessage
	}
	return ""
}

func (m *DockerRunData) GetContainerId() string {
	if m != nil {
		return m.ContainerId
	}
	return ""
}

func init() {
	proto.RegisterType((*EchoMessage)(nil), "pb.EchoMessage")
	proto.RegisterType((*DockerRunData)(nil), "pb.DockerRunData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EchoService service

type EchoServiceClient interface {
	Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error)
	// Like 'docker run' command
	RunContainer(ctx context.Context, in *DockerRunData, opts ...grpc.CallOption) (*DockerRunData, error)
}

type echoServiceClient struct {
	cc *grpc.ClientConn
}

func NewEchoServiceClient(cc *grpc.ClientConn) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) Echo(ctx context.Context, in *EchoMessage, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := grpc.Invoke(ctx, "/pb.EchoService/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoServiceClient) RunContainer(ctx context.Context, in *DockerRunData, opts ...grpc.CallOption) (*DockerRunData, error) {
	out := new(DockerRunData)
	err := grpc.Invoke(ctx, "/pb.EchoService/RunContainer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EchoService service

type EchoServiceServer interface {
	Echo(context.Context, *EchoMessage) (*EchoMessage, error)
	// Like 'docker run' command
	RunContainer(context.Context, *DockerRunData) (*DockerRunData, error)
}

func RegisterEchoServiceServer(s *grpc.Server, srv EchoServiceServer) {
	s.RegisterService(&_EchoService_serviceDesc, srv)
}

func _EchoService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EchoService/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).Echo(ctx, req.(*EchoMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _EchoService_RunContainer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DockerRunData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).RunContainer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.EchoService/RunContainer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).RunContainer(ctx, req.(*DockerRunData))
	}
	return interceptor(ctx, in, info, handler)
}

var _EchoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _EchoService_Echo_Handler,
		},
		{
			MethodName: "RunContainer",
			Handler:    _EchoService_RunContainer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/service.proto",
}

func (m *EchoMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EchoMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	return i, nil
}

func (m *DockerRunData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DockerRunData) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Config != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintService(dAtA, i, uint64(m.Config.Size()))
		n1, err := m.Config.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.HostConfig != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintService(dAtA, i, uint64(m.HostConfig.Size()))
		n2, err := m.HostConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.NetworkConfig != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintService(dAtA, i, uint64(m.NetworkConfig.Size()))
		n3, err := m.NetworkConfig.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.ContainerName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.ContainerName)))
		i += copy(dAtA[i:], m.ContainerName)
	}
	if m.StateCode != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintService(dAtA, i, uint64(m.StateCode))
	}
	if len(m.StateMessage) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.StateMessage)))
		i += copy(dAtA[i:], m.StateMessage)
	}
	if len(m.ContainerId) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintService(dAtA, i, uint64(len(m.ContainerId)))
		i += copy(dAtA[i:], m.ContainerId)
	}
	return i, nil
}

func encodeFixed64Service(dAtA []byte, offset int, v uint64) int {
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
func encodeFixed32Service(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *EchoMessage) Size() (n int) {
	var l int
	_ = l
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *DockerRunData) Size() (n int) {
	var l int
	_ = l
	if m.Config != nil {
		l = m.Config.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.HostConfig != nil {
		l = m.HostConfig.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.NetworkConfig != nil {
		l = m.NetworkConfig.Size()
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.ContainerName)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.StateCode != 0 {
		n += 1 + sovService(uint64(m.StateCode))
	}
	l = len(m.StateMessage)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.ContainerId)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EchoMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: EchoMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EchoMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func (m *DockerRunData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: DockerRunData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DockerRunData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Config == nil {
				m.Config = &moby.Config{}
			}
			if err := m.Config.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.HostConfig == nil {
				m.HostConfig = &moby.HostConfig{}
			}
			if err := m.HostConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NetworkConfig == nil {
				m.NetworkConfig = &moby.NetworkingConfig{}
			}
			if err := m.NetworkConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateCode", wireType)
			}
			m.StateCode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StateCode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContainerId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContainerId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowService
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
				next, err := skipService(dAtA[start:])
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
	ErrInvalidLengthService = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("pb/service.proto", fileDescriptorService) }

var fileDescriptorService = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xcf, 0xae, 0x93, 0x40,
	0x14, 0x87, 0x1d, 0xbc, 0xad, 0xb9, 0x03, 0x54, 0xee, 0x68, 0x0c, 0x12, 0x25, 0x95, 0x6a, 0xd2,
	0xb8, 0x80, 0xb4, 0xee, 0x4c, 0xdc, 0xb4, 0x35, 0xd1, 0x85, 0x35, 0xc1, 0x07, 0x68, 0x06, 0x18,
	0x81, 0xb4, 0xcc, 0x21, 0xcc, 0xb4, 0xc6, 0x18, 0x37, 0xbe, 0x82, 0x1b, 0xd7, 0xbe, 0x83, 0xef,
	0xe0, 0xd2, 0xc4, 0x17, 0x30, 0xd5, 0x07, 0x31, 0xcc, 0x50, 0xfc, 0x73, 0x77, 0xcc, 0x77, 0xbe,
	0xdf, 0x19, 0x38, 0x07, 0xec, 0xd4, 0x49, 0x24, 0x58, 0x73, 0x28, 0x53, 0x16, 0xd6, 0x0d, 0x48,
	0x20, 0x46, 0x9d, 0x78, 0x77, 0x72, 0x80, 0x7c, 0xc7, 0x22, 0x5a, 0x97, 0x11, 0xe5, 0x1c, 0x24,
	0x95, 0x25, 0x70, 0xa1, 0x0d, 0xef, 0xa2, 0x4e, 0xa2, 0x0a, 0x92, 0xb7, 0x6d, 0x59, 0xa3, 0x60,
	0x82, 0xcd, 0xa7, 0x69, 0x01, 0x2f, 0x98, 0x10, 0x34, 0x67, 0xe4, 0x26, 0x1e, 0x1c, 0xe8, 0x6e,
	0xcf, 0x5c, 0x34, 0x46, 0xd3, 0xf3, 0x58, 0x1f, 0x82, 0x2f, 0x06, 0xb6, 0x57, 0x90, 0x6e, 0x59,
	0x13, 0xef, 0xf9, 0x8a, 0x4a, 0x4a, 0xee, 0xe3, 0x61, 0x0a, 0xfc, 0x75, 0x99, 0x2b, 0xd1, 0x9c,
	0x5b, 0x61, 0xdb, 0x37, 0x5c, 0x2a, 0x16, 0x77, 0x35, 0x32, 0xc3, 0x66, 0x01, 0x42, 0x6e, 0x3a,
	0xd5, 0x50, 0xaa, 0xa3, 0xd5, 0x67, 0x20, 0x64, 0xa7, 0xe3, 0xa2, 0x7f, 0x26, 0x4f, 0xf0, 0x88,
	0x33, 0xf9, 0x06, 0x9a, 0xed, 0x29, 0x75, 0x55, 0xa5, 0x6e, 0xe9, 0xd4, 0x5a, 0xd7, 0x4a, 0x9e,
	0x77, 0x59, 0xbb, 0xb3, 0xbb, 0xf8, 0x03, 0x3c, 0x4a, 0x81, 0x4b, 0x5a, 0x72, 0xd6, 0x6c, 0x38,
	0xad, 0x98, 0x7b, 0xa6, 0x3e, 0xc4, 0xee, 0xe9, 0x9a, 0x56, 0x8c, 0xdc, 0xc5, 0x58, 0x48, 0x2a,
	0xd9, 0x26, 0x85, 0x8c, 0xb9, 0x83, 0x31, 0x9a, 0x0e, 0xe2, 0x73, 0x45, 0x96, 0x90, 0x31, 0x32,
	0xc1, 0xb6, 0x2e, 0x57, 0x7a, 0x2c, 0xee, 0x50, 0x35, 0xb1, 0x14, 0x3c, 0x8d, 0xea, 0x1e, 0xb6,
	0xfe, 0x5c, 0x55, 0x66, 0xee, 0x35, 0xe5, 0x98, 0x3d, 0x7b, 0x9e, 0xcd, 0x3f, 0x23, 0x3d, 0xdd,
	0x57, 0x7a, 0x4f, 0x64, 0x81, 0xcf, 0xda, 0x23, 0xb9, 0x1e, 0xd6, 0x49, 0xf8, 0xd7, 0xd8, 0xbd,
	0xff, 0x41, 0xe0, 0x7e, 0xf8, 0xfe, 0xeb, 0xa3, 0x41, 0x88, 0x13, 0x1d, 0x66, 0x11, 0x4b, 0x0b,
	0x88, 0xde, 0xa9, 0x55, 0xbc, 0x27, 0x2f, 0xb1, 0x15, 0xef, 0xf9, 0xf2, 0x74, 0x0b, 0xb9, 0x68,
	0xa3, 0xff, 0x2c, 0xc7, 0xbb, 0x8c, 0x82, 0xdb, 0xaa, 0xdf, 0x8d, 0x60, 0xd4, 0xf6, 0xeb, 0x5f,
	0x51, 0x3c, 0x46, 0x0f, 0x17, 0xce, 0xd7, 0xa3, 0x8f, 0xbe, 0x1d, 0x7d, 0xf4, 0xe3, 0xe8, 0xa3,
	0x4f, 0x3f, 0xfd, 0x2b, 0xc9, 0x50, 0xfd, 0x1a, 0x8f, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x06, 0x39, 0xf0, 0x63, 0x02, 0x00, 0x00,
}
