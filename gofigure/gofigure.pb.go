// Code generated by protoc-gen-go.
// source: gofigure.proto
// DO NOT EDIT!

/*
Package gofigure is a generated protocol buffer package.

It is generated from these files:
	gofigure.proto

It has these top-level messages:
	ConfigVersion
	ProtoConfig
	YamlConfig
	Config
	CheckConfigRequest
	CheckConfigResponse
	NewConfigRequest
	NewConfigResponse
	GetConfigRequest
	GetConfigResponse
	UpdateConfigRequest
	UpdateConfigResponse
*/
package gofigure

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/struct"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Status int32

const (
	Status_SUCCESS             Status = 0
	Status_ERR_INVALID_SERVICE Status = 1
	Status_ERR_INVALID_CONFIG  Status = 2
	Status_ERR_UNSPECIFIED     Status = 3
)

var Status_name = map[int32]string{
	0: "SUCCESS",
	1: "ERR_INVALID_SERVICE",
	2: "ERR_INVALID_CONFIG",
	3: "ERR_UNSPECIFIED",
}
var Status_value = map[string]int32{
	"SUCCESS":             0,
	"ERR_INVALID_SERVICE": 1,
	"ERR_INVALID_CONFIG":  2,
	"ERR_UNSPECIFIED":     3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}
func (Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ConfigVersion struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *ConfigVersion) Reset()                    { *m = ConfigVersion{} }
func (m *ConfigVersion) String() string            { return proto.CompactTextString(m) }
func (*ConfigVersion) ProtoMessage()               {}
func (*ConfigVersion) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ProtoConfig struct {
	Data *google_protobuf.Struct `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *ProtoConfig) Reset()                    { *m = ProtoConfig{} }
func (m *ProtoConfig) String() string            { return proto.CompactTextString(m) }
func (*ProtoConfig) ProtoMessage()               {}
func (*ProtoConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ProtoConfig) GetData() *google_protobuf.Struct {
	if m != nil {
		return m.Data
	}
	return nil
}

type YamlConfig struct {
	RawData []byte `protobuf:"bytes,1,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty"`
}

func (m *YamlConfig) Reset()                    { *m = YamlConfig{} }
func (m *YamlConfig) String() string            { return proto.CompactTextString(m) }
func (*YamlConfig) ProtoMessage()               {}
func (*YamlConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Map of Values that represents configuration data for a service
// version specifies version bucket wherein configuration is stored
type Config struct {
	Version *ConfigVersion `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	// Types that are valid to be assigned to AConfig:
	//	*Config_Proto
	//	*Config_Yaml
	AConfig isConfig_AConfig `protobuf_oneof:"a_config"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isConfig_AConfig interface {
	isConfig_AConfig()
}

type Config_Proto struct {
	Proto *ProtoConfig `protobuf:"bytes,2,opt,name=proto,oneof"`
}
type Config_Yaml struct {
	Yaml *YamlConfig `protobuf:"bytes,3,opt,name=yaml,oneof"`
}

func (*Config_Proto) isConfig_AConfig() {}
func (*Config_Yaml) isConfig_AConfig()  {}

func (m *Config) GetAConfig() isConfig_AConfig {
	if m != nil {
		return m.AConfig
	}
	return nil
}

func (m *Config) GetVersion() *ConfigVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

func (m *Config) GetProto() *ProtoConfig {
	if x, ok := m.GetAConfig().(*Config_Proto); ok {
		return x.Proto
	}
	return nil
}

func (m *Config) GetYaml() *YamlConfig {
	if x, ok := m.GetAConfig().(*Config_Yaml); ok {
		return x.Yaml
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Config) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Config_OneofMarshaler, _Config_OneofUnmarshaler, _Config_OneofSizer, []interface{}{
		(*Config_Proto)(nil),
		(*Config_Yaml)(nil),
	}
}

func _Config_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Config)
	// a_config
	switch x := m.AConfig.(type) {
	case *Config_Proto:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Proto); err != nil {
			return err
		}
	case *Config_Yaml:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Yaml); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Config.AConfig has unexpected type %T", x)
	}
	return nil
}

func _Config_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Config)
	switch tag {
	case 2: // a_config.proto
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ProtoConfig)
		err := b.DecodeMessage(msg)
		m.AConfig = &Config_Proto{msg}
		return true, err
	case 3: // a_config.yaml
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(YamlConfig)
		err := b.DecodeMessage(msg)
		m.AConfig = &Config_Yaml{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Config_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Config)
	// a_config
	switch x := m.AConfig.(type) {
	case *Config_Proto:
		s := proto.Size(x.Proto)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Config_Yaml:
		s := proto.Size(x.Yaml)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// GRPC request to check service configuration
type CheckConfigRequest struct {
	ServiceName string         `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Version     *ConfigVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *CheckConfigRequest) Reset()                    { *m = CheckConfigRequest{} }
func (m *CheckConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigRequest) ProtoMessage()               {}
func (*CheckConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CheckConfigRequest) GetVersion() *ConfigVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

// Used by the client to check its configuration status
type CheckConfigResponse struct {
	Status  Status         `protobuf:"varint,1,opt,name=status,enum=gofigure.Status" json:"status,omitempty"`
	Version *ConfigVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *CheckConfigResponse) Reset()                    { *m = CheckConfigResponse{} }
func (m *CheckConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigResponse) ProtoMessage()               {}
func (*CheckConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckConfigResponse) GetVersion() *ConfigVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

// Store a new configuration on behalf of a service
type NewConfigRequest struct {
	ServiceName   string  `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Configuration *Config `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *NewConfigRequest) Reset()                    { *m = NewConfigRequest{} }
func (m *NewConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*NewConfigRequest) ProtoMessage()               {}
func (*NewConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *NewConfigRequest) GetConfiguration() *Config {
	if m != nil {
		return m.Configuration
	}
	return nil
}

// Response for new service request
type NewConfigResponse struct {
	Status  Status         `protobuf:"varint,1,opt,name=status,enum=gofigure.Status" json:"status,omitempty"`
	Version *ConfigVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *NewConfigResponse) Reset()                    { *m = NewConfigResponse{} }
func (m *NewConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*NewConfigResponse) ProtoMessage()               {}
func (*NewConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *NewConfigResponse) GetVersion() *ConfigVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

// Returns config given service name and optional version
type GetConfigRequest struct {
	ServiceName string         `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Version     *ConfigVersion `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *GetConfigRequest) Reset()                    { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()               {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetConfigRequest) GetVersion() *ConfigVersion {
	if m != nil {
		return m.Version
	}
	return nil
}

// Response for get service request
type GetConfigResponse struct {
	Status        Status  `protobuf:"varint,1,opt,name=status,enum=gofigure.Status" json:"status,omitempty"`
	Configuration *Config `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *GetConfigResponse) Reset()                    { *m = GetConfigResponse{} }
func (m *GetConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*GetConfigResponse) ProtoMessage()               {}
func (*GetConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetConfigResponse) GetConfiguration() *Config {
	if m != nil {
		return m.Configuration
	}
	return nil
}

// Store new service configuration in service
type UpdateConfigRequest struct {
	ServiceName   string  `protobuf:"bytes,1,opt,name=service_name,json=serviceName" json:"service_name,omitempty"`
	Configuration *Config `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *UpdateConfigRequest) Reset()                    { *m = UpdateConfigRequest{} }
func (m *UpdateConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigRequest) ProtoMessage()               {}
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *UpdateConfigRequest) GetConfiguration() *Config {
	if m != nil {
		return m.Configuration
	}
	return nil
}

//  Response for update config request
type UpdateConfigResponse struct {
	Status        Status  `protobuf:"varint,1,opt,name=status,enum=gofigure.Status" json:"status,omitempty"`
	Configuration *Config `protobuf:"bytes,2,opt,name=configuration" json:"configuration,omitempty"`
}

func (m *UpdateConfigResponse) Reset()                    { *m = UpdateConfigResponse{} }
func (m *UpdateConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateConfigResponse) ProtoMessage()               {}
func (*UpdateConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *UpdateConfigResponse) GetConfiguration() *Config {
	if m != nil {
		return m.Configuration
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfigVersion)(nil), "gofigure.ConfigVersion")
	proto.RegisterType((*ProtoConfig)(nil), "gofigure.ProtoConfig")
	proto.RegisterType((*YamlConfig)(nil), "gofigure.YamlConfig")
	proto.RegisterType((*Config)(nil), "gofigure.Config")
	proto.RegisterType((*CheckConfigRequest)(nil), "gofigure.CheckConfigRequest")
	proto.RegisterType((*CheckConfigResponse)(nil), "gofigure.CheckConfigResponse")
	proto.RegisterType((*NewConfigRequest)(nil), "gofigure.NewConfigRequest")
	proto.RegisterType((*NewConfigResponse)(nil), "gofigure.NewConfigResponse")
	proto.RegisterType((*GetConfigRequest)(nil), "gofigure.GetConfigRequest")
	proto.RegisterType((*GetConfigResponse)(nil), "gofigure.GetConfigResponse")
	proto.RegisterType((*UpdateConfigRequest)(nil), "gofigure.UpdateConfigRequest")
	proto.RegisterType((*UpdateConfigResponse)(nil), "gofigure.UpdateConfigResponse")
	proto.RegisterEnum("gofigure.Status", Status_name, Status_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for GoFigurator service

type GoFiguratorClient interface {
	NewConfig(ctx context.Context, in *NewConfigRequest, opts ...grpc.CallOption) (*NewConfigResponse, error)
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigResponse, error)
}

type goFiguratorClient struct {
	cc *grpc.ClientConn
}

func NewGoFiguratorClient(cc *grpc.ClientConn) GoFiguratorClient {
	return &goFiguratorClient{cc}
}

func (c *goFiguratorClient) NewConfig(ctx context.Context, in *NewConfigRequest, opts ...grpc.CallOption) (*NewConfigResponse, error) {
	out := new(NewConfigResponse)
	err := grpc.Invoke(ctx, "/gofigure.GoFigurator/NewConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goFiguratorClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := grpc.Invoke(ctx, "/gofigure.GoFigurator/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goFiguratorClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := grpc.Invoke(ctx, "/gofigure.GoFigurator/UpdateConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goFiguratorClient) CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigResponse, error) {
	out := new(CheckConfigResponse)
	err := grpc.Invoke(ctx, "/gofigure.GoFigurator/CheckConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GoFigurator service

type GoFiguratorServer interface {
	NewConfig(context.Context, *NewConfigRequest) (*NewConfigResponse, error)
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	CheckConfig(context.Context, *CheckConfigRequest) (*CheckConfigResponse, error)
}

func RegisterGoFiguratorServer(s *grpc.Server, srv GoFiguratorServer) {
	s.RegisterService(&_GoFigurator_serviceDesc, srv)
}

func _GoFigurator_NewConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoFiguratorServer).NewConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofigure.GoFigurator/NewConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoFiguratorServer).NewConfig(ctx, req.(*NewConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoFigurator_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoFiguratorServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofigure.GoFigurator/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoFiguratorServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoFigurator_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoFiguratorServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofigure.GoFigurator/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoFiguratorServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GoFigurator_CheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoFiguratorServer).CheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gofigure.GoFigurator/CheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoFiguratorServer).CheckConfig(ctx, req.(*CheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GoFigurator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gofigure.GoFigurator",
	HandlerType: (*GoFiguratorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewConfig",
			Handler:    _GoFigurator_NewConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _GoFigurator_GetConfig_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _GoFigurator_UpdateConfig_Handler,
		},
		{
			MethodName: "CheckConfig",
			Handler:    _GoFigurator_CheckConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 575 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xda, 0x4e,
	0x10, 0x0d, 0x90, 0x1f, 0x7f, 0xc6, 0x84, 0x9f, 0xb3, 0xa4, 0x0d, 0xa5, 0x49, 0xd5, 0xfa, 0xd2,
	0x28, 0x55, 0x41, 0xa5, 0x6a, 0x0f, 0x95, 0x7a, 0x68, 0x8d, 0xa1, 0x96, 0x22, 0x12, 0xad, 0x05,
	0x52, 0xd5, 0x03, 0x32, 0x64, 0x63, 0xdc, 0x62, 0xd6, 0xb5, 0xd7, 0x49, 0xf3, 0x75, 0xfa, 0x39,
	0xfa, 0xe1, 0xba, 0xac, 0x0d, 0x5e, 0x48, 0x22, 0x35, 0x39, 0x70, 0x02, 0xcf, 0xbc, 0x79, 0x6f,
	0xe6, 0xed, 0xec, 0x42, 0xc5, 0xa1, 0x17, 0xae, 0x13, 0x05, 0xa4, 0xe1, 0x07, 0x94, 0x51, 0x54,
	0x5c, 0x7c, 0xd7, 0xdf, 0x39, 0x2e, 0x9b, 0x44, 0xa3, 0xc6, 0x98, 0x7a, 0x4d, 0x87, 0x4e, 0xed,
	0x99, 0xd3, 0x14, 0x90, 0x51, 0x74, 0xd1, 0xf4, 0xd9, 0xb5, 0x4f, 0xc2, 0x66, 0xc8, 0x82, 0x68,
	0xcc, 0x92, 0x9f, 0x98, 0x40, 0xfb, 0x08, 0x3b, 0x3a, 0x9d, 0x71, 0x8e, 0x01, 0x09, 0x42, 0x97,
	0xce, 0x50, 0x05, 0xb2, 0xee, 0x79, 0x2d, 0xf3, 0x3c, 0x73, 0x54, 0xc2, 0xfc, 0x1f, 0x3a, 0x80,
	0x12, 0x73, 0x3d, 0x12, 0x32, 0xdb, 0xf3, 0x6b, 0x59, 0x1e, 0xce, 0xe1, 0x34, 0xa0, 0x7d, 0x00,
	0xe5, 0x6c, 0xce, 0x13, 0x73, 0xa0, 0x57, 0xb0, 0x7d, 0x6e, 0x33, 0x5b, 0x94, 0x2b, 0xad, 0xfd,
	0x86, 0x43, 0xa9, 0x33, 0x4d, 0x7a, 0xe5, 0x8d, 0x34, 0x2c, 0x21, 0x8d, 0x05, 0x48, 0x7b, 0x09,
	0xf0, 0xd5, 0xf6, 0xa6, 0x49, 0xe9, 0x13, 0x28, 0x06, 0xf6, 0xd5, 0x70, 0x59, 0x5e, 0xc6, 0x05,
	0xfe, 0xdd, 0x9e, 0x03, 0x7f, 0x67, 0x20, 0x9f, 0xa0, 0xde, 0x40, 0xe1, 0x32, 0x6e, 0x54, 0xd2,
	0x48, 0x1c, 0x59, 0x99, 0x03, 0x2f, 0x70, 0xe8, 0x35, 0xfc, 0x27, 0xf4, 0x45, 0xf3, 0x4a, 0xeb,
	0x51, 0x5a, 0x20, 0x75, 0xfe, 0x65, 0x0b, 0xc7, 0x28, 0x74, 0x0c, 0xdb, 0xd7, 0xbc, 0xab, 0x5a,
	0x4e, 0xa0, 0xf7, 0x52, 0x74, 0xda, 0x2b, 0x07, 0x0b, 0xcc, 0x67, 0x80, 0xa2, 0x3d, 0x1c, 0x8b,
	0x98, 0xf6, 0x1d, 0x90, 0x3e, 0x21, 0xe3, 0x1f, 0x31, 0x04, 0x93, 0x9f, 0x11, 0xb7, 0x08, 0xbd,
	0x80, 0x72, 0x48, 0x82, 0x4b, 0x77, 0x4c, 0x86, 0x33, 0xdb, 0x23, 0x89, 0xaf, 0x4a, 0x12, 0xeb,
	0xf1, 0x90, 0x3c, 0x52, 0xf6, 0xdf, 0x46, 0xd2, 0x02, 0xa8, 0xae, 0x68, 0x85, 0x3e, 0x9d, 0x85,
	0x04, 0x1d, 0x41, 0x9e, 0x9f, 0x0a, 0x8b, 0x42, 0x21, 0x53, 0x69, 0xa9, 0x29, 0x91, 0x25, 0xe2,
	0x38, 0xc9, 0x3f, 0x44, 0xd3, 0x03, 0xb5, 0x47, 0xae, 0xee, 0x3d, 0xdd, 0x7b, 0xd8, 0x89, 0x0d,
	0x8a, 0x02, 0x9b, 0xa5, 0x7a, 0xea, 0xba, 0x1e, 0x5e, 0x85, 0x69, 0x3e, 0xec, 0x4a, 0x72, 0x9b,
	0x18, 0x70, 0x02, 0x6a, 0x97, 0xb0, 0x4d, 0x1c, 0x5f, 0x04, 0xbb, 0x92, 0xd2, 0xbd, 0x67, 0x7b,
	0xb8, 0xa5, 0xd5, 0xbe, 0xcf, 0xef, 0x17, 0xd9, 0xd8, 0x21, 0xfe, 0x82, 0xbd, 0x55, 0xc5, 0x4d,
	0xcd, 0x7a, 0xfc, 0x0d, 0xf2, 0x31, 0x13, 0x52, 0xa0, 0x60, 0xf5, 0x75, 0xdd, 0xb0, 0x2c, 0x75,
	0x0b, 0xed, 0x43, 0xd5, 0xc0, 0x78, 0x68, 0xf6, 0x06, 0x9f, 0x4e, 0xcc, 0xf6, 0xd0, 0x32, 0xf0,
	0xc0, 0xd4, 0x0d, 0x35, 0x83, 0x1e, 0x03, 0x92, 0x13, 0xfa, 0x69, 0xaf, 0x63, 0x76, 0xd5, 0x2c,
	0xaa, 0xc2, 0xff, 0xf3, 0x78, 0xbf, 0x67, 0x9d, 0x19, 0xba, 0xd9, 0x31, 0x8d, 0xb6, 0x9a, 0x6b,
	0xfd, 0xc9, 0x82, 0xd2, 0xa5, 0x9d, 0x58, 0x8d, 0x06, 0xa8, 0x03, 0xa5, 0xe5, 0xae, 0xa2, 0x7a,
	0xda, 0xda, 0xfa, 0x7d, 0xa9, 0x3f, 0xbd, 0x35, 0x17, 0x9b, 0xa2, 0x6d, 0xcd, 0x79, 0x96, 0x7b,
	0x21, 0xf3, 0xac, 0xaf, 0xa5, 0xcc, 0x73, 0x63, 0x91, 0x38, 0xcf, 0x29, 0x94, 0x65, 0xdb, 0xd1,
	0x61, 0x0a, 0xbf, 0x65, 0x01, 0xea, 0xcf, 0xee, 0x4a, 0x2f, 0x09, 0x4f, 0x40, 0x91, 0xde, 0x1b,
	0x74, 0x20, 0xb9, 0x7f, 0xe3, 0xc9, 0xab, 0x1f, 0xde, 0x91, 0x5d, 0xb0, 0x8d, 0xf2, 0xe2, 0xa1,
	0x7d, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x6d, 0x19, 0xfe, 0xcc, 0x06, 0x00, 0x00,
}
