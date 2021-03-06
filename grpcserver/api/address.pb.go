// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/address.proto

/*
Package envoy_api_v2 is a generated protocol buffer package.

It is generated from these files:
	api/address.proto

It has these top-level messages:
	Pipe
	SocketAddress
	BindConfig
	Address
	CidrRange
*/
package envoy_api_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SocketAddress_Protocol int32

const (
	SocketAddress_TCP SocketAddress_Protocol = 0
	// [#not-implemented-hide:]
	SocketAddress_UDP SocketAddress_Protocol = 1
)

var SocketAddress_Protocol_name = map[int32]string{
	0: "TCP",
	1: "UDP",
}
var SocketAddress_Protocol_value = map[string]int32{
	"TCP": 0,
	"UDP": 1,
}

func (x SocketAddress_Protocol) String() string {
	return proto.EnumName(SocketAddress_Protocol_name, int32(x))
}
func (SocketAddress_Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type Pipe struct {
	// Unix Domain Socket path.
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
}

func (m *Pipe) Reset()                    { *m = Pipe{} }
func (m *Pipe) String() string            { return proto.CompactTextString(m) }
func (*Pipe) ProtoMessage()               {}
func (*Pipe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Pipe) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type SocketAddress struct {
	Protocol SocketAddress_Protocol `protobuf:"varint,1,opt,name=protocol,enum=envoy.api.v2.SocketAddress_Protocol" json:"protocol,omitempty"`
	// The address for this socket. :ref:`Listeners <config_listeners>` will bind
	// to the address or outbound connections will be made. An empty address is
	// not allowed, specify ``0.0.0.0`` or ``::`` to bind any. It's still possible to
	// distinguish on an address via the prefix/suffix matching in
	// FilterChainMatch after connection. For :ref:`clusters
	// <config_cluster_manager_cluster>`, an address may be either an IP or
	// hostname to be resolved via DNS. If it is a hostname, :ref:`resolver_name
	// <envoy_api_field_SocketAddress.resolver_name>` should be set unless default
	// (i.e. DNS) resolution is expected.
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	// Types that are valid to be assigned to PortSpecifier:
	//	*SocketAddress_PortValue
	//	*SocketAddress_NamedPort
	PortSpecifier isSocketAddress_PortSpecifier `protobuf_oneof:"port_specifier"`
	// The name of the resolver. This must have been registered with Envoy. If this is
	// empty, a context dependent default applies. If address is a hostname this
	// should be set for resolution other than DNS. If the address is a concrete
	// IP address, no resolution will occur.
	ResolverName string `protobuf:"bytes,5,opt,name=resolver_name,json=resolverName" json:"resolver_name,omitempty"`
	// When binding to an IPv6 address above, this enables `IPv4 compatibity
	// <https://tools.ietf.org/html/rfc3493#page-11>`_. Binding to ``::`` will
	// allow both IPv4 and IPv6 connections, with peer IPv4 addresses mapped into
	// IPv6 space as ``::FFFF:<IPv4-address>``.
	Ipv4Compat bool `protobuf:"varint,6,opt,name=ipv4_compat,json=ipv4Compat" json:"ipv4_compat,omitempty"`
}

func (m *SocketAddress) Reset()                    { *m = SocketAddress{} }
func (m *SocketAddress) String() string            { return proto.CompactTextString(m) }
func (*SocketAddress) ProtoMessage()               {}
func (*SocketAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type isSocketAddress_PortSpecifier interface {
	isSocketAddress_PortSpecifier()
}

type SocketAddress_PortValue struct {
	PortValue uint32 `protobuf:"varint,3,opt,name=port_value,json=portValue,oneof"`
}
type SocketAddress_NamedPort struct {
	NamedPort string `protobuf:"bytes,4,opt,name=named_port,json=namedPort,oneof"`
}

func (*SocketAddress_PortValue) isSocketAddress_PortSpecifier() {}
func (*SocketAddress_NamedPort) isSocketAddress_PortSpecifier() {}

func (m *SocketAddress) GetPortSpecifier() isSocketAddress_PortSpecifier {
	if m != nil {
		return m.PortSpecifier
	}
	return nil
}

func (m *SocketAddress) GetProtocol() SocketAddress_Protocol {
	if m != nil {
		return m.Protocol
	}
	return SocketAddress_TCP
}

func (m *SocketAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *SocketAddress) GetPortValue() uint32 {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_PortValue); ok {
		return x.PortValue
	}
	return 0
}

func (m *SocketAddress) GetNamedPort() string {
	if x, ok := m.GetPortSpecifier().(*SocketAddress_NamedPort); ok {
		return x.NamedPort
	}
	return ""
}

func (m *SocketAddress) GetResolverName() string {
	if m != nil {
		return m.ResolverName
	}
	return ""
}

func (m *SocketAddress) GetIpv4Compat() bool {
	if m != nil {
		return m.Ipv4Compat
	}
	return false
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SocketAddress) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SocketAddress_OneofMarshaler, _SocketAddress_OneofUnmarshaler, _SocketAddress_OneofSizer, []interface{}{
		(*SocketAddress_PortValue)(nil),
		(*SocketAddress_NamedPort)(nil),
	}
}

func _SocketAddress_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		b.EncodeVarint(3<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.NamedPort)
	case nil:
	default:
		return fmt.Errorf("SocketAddress.PortSpecifier has unexpected type %T", x)
	}
	return nil
}

func _SocketAddress_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SocketAddress)
	switch tag {
	case 3: // port_specifier.port_value
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.PortSpecifier = &SocketAddress_PortValue{uint32(x)}
		return true, err
	case 4: // port_specifier.named_port
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PortSpecifier = &SocketAddress_NamedPort{x}
		return true, err
	default:
		return false, nil
	}
}

func _SocketAddress_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SocketAddress)
	// port_specifier
	switch x := m.PortSpecifier.(type) {
	case *SocketAddress_PortValue:
		n += proto.SizeVarint(3<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.PortValue))
	case *SocketAddress_NamedPort:
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(len(x.NamedPort)))
		n += len(x.NamedPort)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type BindConfig struct {
	// The address to bind to when creating a socket.
	SourceAddress *SocketAddress `protobuf:"bytes,1,opt,name=source_address,json=sourceAddress" json:"source_address,omitempty"`
}

func (m *BindConfig) Reset()                    { *m = BindConfig{} }
func (m *BindConfig) String() string            { return proto.CompactTextString(m) }
func (*BindConfig) ProtoMessage()               {}
func (*BindConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BindConfig) GetSourceAddress() *SocketAddress {
	if m != nil {
		return m.SourceAddress
	}
	return nil
}

// Addresses specify either a logical or physical address and port, which are
// used to tell Envoy where to bind/listen, connect to upstream and find
// management servers.
type Address struct {
	// Types that are valid to be assigned to Address:
	//	*Address_SocketAddress
	//	*Address_Pipe
	Address isAddress_Address `protobuf_oneof:"address"`
}

func (m *Address) Reset()                    { *m = Address{} }
func (m *Address) String() string            { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()               {}
func (*Address) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type isAddress_Address interface {
	isAddress_Address()
}

type Address_SocketAddress struct {
	SocketAddress *SocketAddress `protobuf:"bytes,1,opt,name=socket_address,json=socketAddress,oneof"`
}
type Address_Pipe struct {
	Pipe *Pipe `protobuf:"bytes,2,opt,name=pipe,oneof"`
}

func (*Address_SocketAddress) isAddress_Address() {}
func (*Address_Pipe) isAddress_Address()          {}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Address) GetSocketAddress() *SocketAddress {
	if x, ok := m.GetAddress().(*Address_SocketAddress); ok {
		return x.SocketAddress
	}
	return nil
}

func (m *Address) GetPipe() *Pipe {
	if x, ok := m.GetAddress().(*Address_Pipe); ok {
		return x.Pipe
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Address) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Address_OneofMarshaler, _Address_OneofUnmarshaler, _Address_OneofSizer, []interface{}{
		(*Address_SocketAddress)(nil),
		(*Address_Pipe)(nil),
	}
}

func _Address_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.SocketAddress); err != nil {
			return err
		}
	case *Address_Pipe:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Pipe); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Address.Address has unexpected type %T", x)
	}
	return nil
}

func _Address_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Address)
	switch tag {
	case 1: // address.socket_address
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(SocketAddress)
		err := b.DecodeMessage(msg)
		m.Address = &Address_SocketAddress{msg}
		return true, err
	case 2: // address.pipe
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Pipe)
		err := b.DecodeMessage(msg)
		m.Address = &Address_Pipe{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Address_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Address)
	// address
	switch x := m.Address.(type) {
	case *Address_SocketAddress:
		s := proto.Size(x.SocketAddress)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Address_Pipe:
		s := proto.Size(x.Pipe)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// CidrRange specifies an IP Address and a prefix length to construct
// the subnet mask for a `CIDR <https://tools.ietf.org/html/rfc4632>`_ range.
type CidrRange struct {
	// IPv4 or IPv6 address, e.g. ``192.0.0.0`` or ``2001:db8::``.
	AddressPrefix string `protobuf:"bytes,1,opt,name=address_prefix,json=addressPrefix" json:"address_prefix,omitempty"`
	// Length of prefix, e.g. 0, 32.
	PrefixLen *google_protobuf.UInt32Value `protobuf:"bytes,2,opt,name=prefix_len,json=prefixLen" json:"prefix_len,omitempty"`
}

func (m *CidrRange) Reset()                    { *m = CidrRange{} }
func (m *CidrRange) String() string            { return proto.CompactTextString(m) }
func (*CidrRange) ProtoMessage()               {}
func (*CidrRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CidrRange) GetAddressPrefix() string {
	if m != nil {
		return m.AddressPrefix
	}
	return ""
}

func (m *CidrRange) GetPrefixLen() *google_protobuf.UInt32Value {
	if m != nil {
		return m.PrefixLen
	}
	return nil
}

func init() {
	proto.RegisterType((*Pipe)(nil), "envoy.api.v2.Pipe")
	proto.RegisterType((*SocketAddress)(nil), "envoy.api.v2.SocketAddress")
	proto.RegisterType((*BindConfig)(nil), "envoy.api.v2.BindConfig")
	proto.RegisterType((*Address)(nil), "envoy.api.v2.Address")
	proto.RegisterType((*CidrRange)(nil), "envoy.api.v2.CidrRange")
	proto.RegisterEnum("envoy.api.v2.SocketAddress_Protocol", SocketAddress_Protocol_name, SocketAddress_Protocol_value)
}

func init() { proto.RegisterFile("api/address.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 505 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x6b, 0xd4, 0x40,
	0x18, 0xde, 0xd9, 0x8f, 0xee, 0xe6, 0xdd, 0x26, 0xac, 0x83, 0xd0, 0x50, 0xab, 0x5d, 0x52, 0x85,
	0xc5, 0x43, 0x22, 0xa9, 0x7f, 0xc0, 0x6c, 0xc1, 0x15, 0x8a, 0x84, 0x68, 0xbd, 0xc6, 0x69, 0x32,
	0x1b, 0x07, 0xb3, 0x99, 0x61, 0x92, 0x8d, 0x7a, 0x13, 0x0f, 0x1e, 0xf4, 0x0f, 0x89, 0xa7, 0x7a,
	0xf3, 0x57, 0x78, 0xef, 0xbf, 0x90, 0x99, 0x24, 0xa5, 0x55, 0x10, 0xbc, 0xbd, 0x3c, 0xcf, 0xf3,
	0x3e, 0xbc, 0x1f, 0x0f, 0xdc, 0x22, 0x82, 0x79, 0x24, 0x4d, 0x25, 0x2d, 0x4b, 0x57, 0x48, 0x5e,
	0x71, 0xbc, 0x4b, 0x8b, 0x9a, 0x7f, 0x70, 0x89, 0x60, 0x6e, 0xed, 0xef, 0xdf, 0xcb, 0x38, 0xcf,
	0x72, 0xea, 0x69, 0xee, 0x7c, 0xbb, 0xf6, 0xde, 0x49, 0x22, 0x04, 0x95, 0xad, 0x7a, 0x7f, 0xaf,
	0x26, 0x39, 0x4b, 0x49, 0x45, 0xbd, 0xae, 0x68, 0x89, 0xdb, 0x19, 0xcf, 0xb8, 0x2e, 0x3d, 0x55,
	0x35, 0xa8, 0xf3, 0x00, 0x86, 0x21, 0x13, 0x14, 0xdf, 0x85, 0xa1, 0x20, 0xd5, 0x1b, 0x1b, 0xcd,
	0xd1, 0xc2, 0x08, 0x8c, 0xef, 0x97, 0x17, 0x83, 0xa1, 0xec, 0xcf, 0x51, 0xa4, 0x61, 0xe7, 0x47,
	0x1f, 0xcc, 0x17, 0x3c, 0x79, 0x4b, 0xab, 0x27, 0xcd, 0x6c, 0xf8, 0x14, 0x26, 0xda, 0x21, 0xe1,
	0xb9, 0x6e, 0xb2, 0xfc, 0xfb, 0xee, 0xf5, 0x41, 0xdd, 0x1b, 0x72, 0x37, 0x6c, 0xb5, 0x01, 0x28,
	0xeb, 0xd1, 0x27, 0xd4, 0x9f, 0xa1, 0xe8, 0xca, 0x01, 0x1f, 0xc1, 0xb8, 0x5d, 0xda, 0xee, 0xff,
	0x39, 0x41, 0xc7, 0xe0, 0x43, 0x00, 0xc1, 0x65, 0x15, 0xd7, 0x24, 0xdf, 0x52, 0x7b, 0x30, 0x47,
	0x0b, 0x73, 0xd5, 0x8b, 0x0c, 0x85, 0xbd, 0x52, 0x90, 0x12, 0x14, 0x64, 0x43, 0xd3, 0x58, 0x41,
	0xf6, 0x50, 0x19, 0x29, 0x81, 0xc6, 0x42, 0x2e, 0x2b, 0x7c, 0x04, 0xa6, 0xa4, 0x25, 0xcf, 0x6b,
	0x2a, 0x63, 0x85, 0xda, 0x23, 0xa5, 0x89, 0x76, 0x3b, 0xf0, 0x39, 0xd9, 0x28, 0x97, 0x29, 0x13,
	0xf5, 0xe3, 0x38, 0xe1, 0x1b, 0x41, 0x2a, 0x7b, 0x67, 0x8e, 0x16, 0x93, 0x08, 0x14, 0xb4, 0xd4,
	0x88, 0x73, 0x00, 0x93, 0x6e, 0x1d, 0x3c, 0x86, 0xc1, 0xcb, 0x65, 0x38, 0xeb, 0xa9, 0xe2, 0xec,
	0x24, 0x9c, 0xa1, 0x60, 0x0f, 0x2c, 0x3d, 0x65, 0x29, 0x68, 0xc2, 0xd6, 0x8c, 0x4a, 0x3c, 0xfa,
	0x76, 0x79, 0x31, 0x40, 0xce, 0x6b, 0x80, 0x80, 0x15, 0xe9, 0x92, 0x17, 0x6b, 0x96, 0xe1, 0x08,
	0xac, 0x92, 0x6f, 0x65, 0x42, 0xe3, 0x6e, 0x71, 0x75, 0xc5, 0xa9, 0x7f, 0xe7, 0x1f, 0x57, 0x0c,
	0xac, 0x9f, 0xbf, 0x0e, 0x7b, 0xfa, 0x80, 0x5f, 0xf4, 0x01, 0xcd, 0xc6, 0xa2, 0xa5, 0x9d, 0xaf,
	0x08, 0xc6, 0xdd, 0x7f, 0x4e, 0x94, 0xbf, 0xea, 0xfd, 0x0f, 0xff, 0x55, 0x4f, 0x39, 0x5e, 0xff,
	0xf2, 0x02, 0x86, 0x82, 0x09, 0xaa, 0x9f, 0x32, 0xf5, 0xf1, 0xcd, 0x5e, 0x15, 0x9c, 0x55, 0x2f,
	0xd2, 0x8a, 0x60, 0x76, 0xf5, 0xc1, 0x6e, 0xdf, 0xcf, 0x08, 0x8c, 0x25, 0x4b, 0x65, 0x44, 0x8a,
	0x8c, 0xe2, 0x47, 0x60, 0xb5, 0x7c, 0x2c, 0x24, 0x5d, 0xb3, 0xf7, 0x7f, 0x47, 0xcd, 0x6c, 0x05,
	0xa1, 0xe6, 0xf1, 0x53, 0x80, 0x46, 0x19, 0xe7, 0xb4, 0x68, 0x27, 0x38, 0x70, 0x9b, 0xf8, 0xbb,
	0x5d, 0xfc, 0xdd, 0xb3, 0x67, 0x45, 0x75, 0xec, 0xeb, 0xff, 0xb7, 0xd9, 0x7a, 0x38, 0xb0, 0x3f,
	0xa2, 0xc8, 0x68, 0x7a, 0x4f, 0x69, 0x71, 0xbe, 0xa3, 0xc5, 0xc7, 0xbf, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x1e, 0x4b, 0x2b, 0xa8, 0x5c, 0x03, 0x00, 0x00,
}
