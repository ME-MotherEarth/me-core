// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mect.proto

package mect

import (
	bytes "bytes"
	fmt "fmt"
	github_com_ME_MotherEarth_me_core_data "github.com/ME-MotherEarth/me-core/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// MECToken holds the data for a MotherEarth standard digital token transaction
type MECToken struct {
	Type          uint32        `protobuf:"varint,1,opt,name=Type,proto3" json:"Type"`
	Value         *math_big.Int `protobuf:"bytes,2,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/ME-MotherEarth/me-core/data.BigIntCaster" json:"Value"`
	Properties    []byte        `protobuf:"bytes,3,opt,name=Properties,proto3" json:"Properties"`
	TokenMetaData *MetaData     `protobuf:"bytes,4,opt,name=TokenMetaData,proto3" json:"MetaData"`
	Reserved      []byte        `protobuf:"bytes,5,opt,name=Reserved,proto3" json:"Reserved"`
}

func (m *MECToken) Reset()      { *m = MECToken{} }
func (*MECToken) ProtoMessage() {}
func (*MECToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebfb413144f0c81, []int{0}
}
func (m *MECToken) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MECToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MECToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MECToken.Merge(m, src)
}
func (m *MECToken) XXX_Size() int {
	return m.Size()
}
func (m *MECToken) XXX_DiscardUnknown() {
	xxx_messageInfo_MECToken.DiscardUnknown(m)
}

var xxx_messageInfo_MECToken proto.InternalMessageInfo

func (m *MECToken) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *MECToken) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *MECToken) GetProperties() []byte {
	if m != nil {
		return m.Properties
	}
	return nil
}

func (m *MECToken) GetTokenMetaData() *MetaData {
	if m != nil {
		return m.TokenMetaData
	}
	return nil
}

func (m *MECToken) GetReserved() []byte {
	if m != nil {
		return m.Reserved
	}
	return nil
}

// MECTRoles holds the roles for a given token and the given address
type MECTRoles struct {
	Roles [][]byte `protobuf:"bytes,1,rep,name=Roles,proto3" json:"roles"`
}

func (m *MECTRoles) Reset()      { *m = MECTRoles{} }
func (*MECTRoles) ProtoMessage() {}
func (*MECTRoles) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebfb413144f0c81, []int{1}
}
func (m *MECTRoles) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MECTRoles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MECTRoles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MECTRoles.Merge(m, src)
}
func (m *MECTRoles) XXX_Size() int {
	return m.Size()
}
func (m *MECTRoles) XXX_DiscardUnknown() {
	xxx_messageInfo_MECTRoles.DiscardUnknown(m)
}

var xxx_messageInfo_MECTRoles proto.InternalMessageInfo

func (m *MECTRoles) GetRoles() [][]byte {
	if m != nil {
		return m.Roles
	}
	return nil
}

// MetaData hold the metadata structure for the MECT token
type MetaData struct {
	Nonce      uint64   `protobuf:"varint,1,opt,name=Nonce,proto3" json:"Nonce"`
	Name       []byte   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name"`
	Creator    []byte   `protobuf:"bytes,3,opt,name=Creator,proto3" json:"Creator"`
	Royalties  uint32   `protobuf:"varint,4,opt,name=Royalties,proto3" json:"Royalties"`
	Hash       []byte   `protobuf:"bytes,5,opt,name=Hash,proto3" json:"Hash"`
	URIs       [][]byte `protobuf:"bytes,6,rep,name=URIs,proto3" json:"URIs"`
	Attributes []byte   `protobuf:"bytes,7,opt,name=Attributes,proto3" json:"Attributes"`
}

func (m *MetaData) Reset()      { *m = MetaData{} }
func (*MetaData) ProtoMessage() {}
func (*MetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_2ebfb413144f0c81, []int{2}
}
func (m *MetaData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *MetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetaData.Merge(m, src)
}
func (m *MetaData) XXX_Size() int {
	return m.Size()
}
func (m *MetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetaData proto.InternalMessageInfo

func (m *MetaData) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *MetaData) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *MetaData) GetCreator() []byte {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MetaData) GetRoyalties() uint32 {
	if m != nil {
		return m.Royalties
	}
	return 0
}

func (m *MetaData) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *MetaData) GetURIs() [][]byte {
	if m != nil {
		return m.URIs
	}
	return nil
}

func (m *MetaData) GetAttributes() []byte {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func init() {
	proto.RegisterType((*MECToken)(nil), "protoBuiltInFunctions.MECToken")
	proto.RegisterType((*MECTRoles)(nil), "protoBuiltInFunctions.MECTRoles")
	proto.RegisterType((*MetaData)(nil), "protoBuiltInFunctions.MetaData")
}

func init() { proto.RegisterFile("mect.proto", fileDescriptor_2ebfb413144f0c81) }

var fileDescriptor_2ebfb413144f0c81 = []byte{
	// 504 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xce, 0xec, 0xb6, 0xbb, 0xed, 0x6c, 0xeb, 0x21, 0x20, 0x04, 0x91, 0x99, 0x52, 0x10, 0x0a,
	0xda, 0x14, 0xf4, 0x28, 0x88, 0xa6, 0x56, 0xec, 0xa1, 0x8b, 0x8c, 0xab, 0x07, 0x6f, 0x93, 0xec,
	0x98, 0x04, 0x9b, 0x4c, 0x99, 0x4c, 0x84, 0xbd, 0x79, 0xf5, 0xe6, 0xcf, 0x10, 0x7f, 0x83, 0x3f,
	0xc0, 0x63, 0x8f, 0x3d, 0x45, 0x9b, 0x5e, 0x24, 0xa7, 0xfd, 0x09, 0x32, 0x2f, 0x1b, 0x5b, 0xc1,
	0xcb, 0xcb, 0xf7, 0x7d, 0xf3, 0x98, 0xf7, 0xe6, 0xfb, 0x08, 0xc6, 0x89, 0x08, 0xb4, 0xbb, 0x52,
	0x52, 0x4b, 0xfb, 0x36, 0x7c, 0xbc, 0x3c, 0x5e, 0xea, 0x79, 0xfa, 0x22, 0x4f, 0x03, 0x1d, 0xcb,
	0x34, 0xbb, 0x33, 0x0e, 0x63, 0x1d, 0xe5, 0xbe, 0x1b, 0xc8, 0x64, 0x12, 0xca, 0x50, 0x4e, 0xa0,
	0xcd, 0xcf, 0xdf, 0x03, 0x03, 0x02, 0xa8, 0xbe, 0x65, 0xf8, 0xfd, 0x08, 0x77, 0x16, 0xb3, 0xe9,
	0x85, 0xfc, 0x20, 0x52, 0xfb, 0x2e, 0x6e, 0x5d, 0x5c, 0xad, 0x84, 0x83, 0x06, 0x68, 0xd4, 0xf7,
	0x3a, 0x55, 0x41, 0x81, 0x33, 0xa8, 0x76, 0x80, 0xdb, 0x6f, 0xf9, 0x32, 0x17, 0xce, 0xd1, 0x00,
	0x8d, 0x7a, 0xde, 0xa2, 0x2a, 0x68, 0x2d, 0x7c, 0xfb, 0x49, 0x9f, 0x26, 0x5c, 0x47, 0x13, 0x3f,
	0x0e, 0xdd, 0x79, 0xaa, 0x1f, 0x1f, 0xac, 0xb0, 0x98, 0x8d, 0x17, 0x52, 0x47, 0x42, 0xcd, 0xb8,
	0xd2, 0xd1, 0x24, 0x11, 0xe3, 0x40, 0x2a, 0x31, 0xb9, 0xe4, 0x9a, 0xbb, 0x5e, 0x1c, 0xce, 0x53,
	0x3d, 0xe5, 0x99, 0x16, 0x8a, 0xd5, 0x57, 0xd9, 0x2e, 0xc6, 0xaf, 0x94, 0x5c, 0x09, 0xa5, 0x63,
	0x91, 0x39, 0xc7, 0x30, 0xe9, 0x56, 0x55, 0xd0, 0x03, 0x95, 0x1d, 0x60, 0xfb, 0x35, 0xee, 0xc3,
	0xee, 0x0b, 0xa1, 0xf9, 0x73, 0xae, 0xb9, 0xd3, 0x1a, 0xa0, 0xd1, 0xd9, 0x43, 0xea, 0xfe, 0xd7,
	0x1d, 0xb7, 0x69, 0xf3, 0x7a, 0x55, 0x41, 0x3b, 0x0d, 0x63, 0xff, 0xde, 0x61, 0x8f, 0x70, 0x87,
	0x89, 0x4c, 0xa8, 0x8f, 0xe2, 0xd2, 0x69, 0xc3, 0x0a, 0xd0, 0xde, 0x68, 0xec, 0x2f, 0x1a, 0x3e,
	0xc0, 0x5d, 0xe3, 0x1e, 0x93, 0x4b, 0x91, 0xd9, 0x14, 0xb7, 0x01, 0x38, 0x68, 0x70, 0x3c, 0xea,
	0x79, 0x5d, 0x63, 0x90, 0x32, 0x02, 0xab, 0xf5, 0xe1, 0x67, 0x63, 0x76, 0x33, 0x84, 0xe2, 0xf6,
	0xb9, 0x4c, 0x83, 0xda, 0xed, 0x56, 0xdd, 0x0d, 0x02, 0xab, 0x3f, 0x26, 0x8d, 0x73, 0x9e, 0x34,
	0x76, 0x43, 0x1a, 0x86, 0x33, 0xa8, 0xf6, 0x3d, 0x7c, 0x3a, 0x55, 0x82, 0x6b, 0xa9, 0x6e, 0x5c,
	0x3a, 0xab, 0x0a, 0xda, 0x48, 0xac, 0x01, 0xf6, 0x7d, 0xdc, 0x65, 0xf2, 0x8a, 0x2f, 0xc1, 0xce,
	0x16, 0xe4, 0xda, 0xaf, 0x0a, 0xba, 0x17, 0xd9, 0x1e, 0x9a, 0x89, 0x2f, 0x79, 0x16, 0xdd, 0xbc,
	0x19, 0x26, 0x1a, 0xce, 0xa0, 0x9a, 0xd3, 0x37, 0x6c, 0x9e, 0x39, 0x27, 0xf0, 0x3a, 0x38, 0x35,
	0x9c, 0x41, 0x35, 0xc1, 0x3d, 0xd3, 0x5a, 0xc5, 0x7e, 0xae, 0x45, 0xe6, 0x9c, 0xee, 0x83, 0xdb,
	0xab, 0xec, 0x00, 0x7b, 0x4f, 0xd6, 0x5b, 0x62, 0x6d, 0xb6, 0xc4, 0xba, 0xde, 0x12, 0xf4, 0xa9,
	0x24, 0xe8, 0x6b, 0x49, 0xd0, 0x8f, 0x92, 0xa0, 0x75, 0x49, 0xd0, 0xa6, 0x24, 0xe8, 0x57, 0x49,
	0xd0, 0xef, 0x92, 0x58, 0xd7, 0x25, 0x41, 0x5f, 0x76, 0xc4, 0x5a, 0xef, 0x88, 0xb5, 0xd9, 0x11,
	0xeb, 0x5d, 0xcb, 0xfc, 0x04, 0xfe, 0x09, 0x04, 0xfc, 0xe8, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x95, 0xa5, 0xad, 0x8f, 0x13, 0x03, 0x00, 0x00,
}

func (this *MECToken) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MECToken)
	if !ok {
		that2, ok := that.(MECToken)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	{
		__caster := &github_com_ME_MotherEarth_me_core_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.Properties, that1.Properties) {
		return false
	}
	if !this.TokenMetaData.Equal(that1.TokenMetaData) {
		return false
	}
	if !bytes.Equal(this.Reserved, that1.Reserved) {
		return false
	}
	return true
}
func (this *MECTRoles) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MECTRoles)
	if !ok {
		that2, ok := that.(MECTRoles)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Roles) != len(that1.Roles) {
		return false
	}
	for i := range this.Roles {
		if !bytes.Equal(this.Roles[i], that1.Roles[i]) {
			return false
		}
	}
	return true
}
func (this *MetaData) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MetaData)
	if !ok {
		that2, ok := that.(MetaData)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Nonce != that1.Nonce {
		return false
	}
	if !bytes.Equal(this.Name, that1.Name) {
		return false
	}
	if !bytes.Equal(this.Creator, that1.Creator) {
		return false
	}
	if this.Royalties != that1.Royalties {
		return false
	}
	if !bytes.Equal(this.Hash, that1.Hash) {
		return false
	}
	if len(this.URIs) != len(that1.URIs) {
		return false
	}
	for i := range this.URIs {
		if !bytes.Equal(this.URIs[i], that1.URIs[i]) {
			return false
		}
	}
	if !bytes.Equal(this.Attributes, that1.Attributes) {
		return false
	}
	return true
}
func (this *MECToken) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 9)
	s = append(s, "&mect.MECToken{")
	s = append(s, "Type: "+fmt.Sprintf("%#v", this.Type)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "Properties: "+fmt.Sprintf("%#v", this.Properties)+",\n")
	if this.TokenMetaData != nil {
		s = append(s, "TokenMetaData: "+fmt.Sprintf("%#v", this.TokenMetaData)+",\n")
	}
	s = append(s, "Reserved: "+fmt.Sprintf("%#v", this.Reserved)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MECTRoles) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&mect.MECTRoles{")
	s = append(s, "Roles: "+fmt.Sprintf("%#v", this.Roles)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *MetaData) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 11)
	s = append(s, "&mect.MetaData{")
	s = append(s, "Nonce: "+fmt.Sprintf("%#v", this.Nonce)+",\n")
	s = append(s, "Name: "+fmt.Sprintf("%#v", this.Name)+",\n")
	s = append(s, "Creator: "+fmt.Sprintf("%#v", this.Creator)+",\n")
	s = append(s, "Royalties: "+fmt.Sprintf("%#v", this.Royalties)+",\n")
	s = append(s, "Hash: "+fmt.Sprintf("%#v", this.Hash)+",\n")
	s = append(s, "URIs: "+fmt.Sprintf("%#v", this.URIs)+",\n")
	s = append(s, "Attributes: "+fmt.Sprintf("%#v", this.Attributes)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringMect(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *MECToken) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MECToken) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MECToken) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Reserved) > 0 {
		i -= len(m.Reserved)
		copy(dAtA[i:], m.Reserved)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Reserved)))
		i--
		dAtA[i] = 0x2a
	}
	if m.TokenMetaData != nil {
		{
			size, err := m.TokenMetaData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintMect(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Properties) > 0 {
		i -= len(m.Properties)
		copy(dAtA[i:], m.Properties)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Properties)))
		i--
		dAtA[i] = 0x1a
	}
	{
		__caster := &github_com_ME_MotherEarth_me_core_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMect(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.Type != 0 {
		i = encodeVarintMect(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MECTRoles) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MECTRoles) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MECTRoles) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for iNdEx := len(m.Roles) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Roles[iNdEx])
			copy(dAtA[i:], m.Roles[iNdEx])
			i = encodeVarintMect(dAtA, i, uint64(len(m.Roles[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MetaData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MetaData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MetaData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Attributes) > 0 {
		i -= len(m.Attributes)
		copy(dAtA[i:], m.Attributes)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Attributes)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.URIs) > 0 {
		for iNdEx := len(m.URIs) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.URIs[iNdEx])
			copy(dAtA[i:], m.URIs[iNdEx])
			i = encodeVarintMect(dAtA, i, uint64(len(m.URIs[iNdEx])))
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Royalties != 0 {
		i = encodeVarintMect(dAtA, i, uint64(m.Royalties))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMect(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Nonce != 0 {
		i = encodeVarintMect(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMect(dAtA []byte, offset int, v uint64) int {
	offset -= sovMect(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MECToken) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovMect(uint64(m.Type))
	}
	{
		__caster := &github_com_ME_MotherEarth_me_core_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovMect(uint64(l))
	}
	l = len(m.Properties)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	if m.TokenMetaData != nil {
		l = m.TokenMetaData.Size()
		n += 1 + l + sovMect(uint64(l))
	}
	l = len(m.Reserved)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	return n
}

func (m *MECTRoles) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Roles) > 0 {
		for _, b := range m.Roles {
			l = len(b)
			n += 1 + l + sovMect(uint64(l))
		}
	}
	return n
}

func (m *MetaData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Nonce != 0 {
		n += 1 + sovMect(uint64(m.Nonce))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	if m.Royalties != 0 {
		n += 1 + sovMect(uint64(m.Royalties))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	if len(m.URIs) > 0 {
		for _, b := range m.URIs {
			l = len(b)
			n += 1 + l + sovMect(uint64(l))
		}
	}
	l = len(m.Attributes)
	if l > 0 {
		n += 1 + l + sovMect(uint64(l))
	}
	return n
}

func sovMect(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMect(x uint64) (n int) {
	return sovMect(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MECToken) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MECToken{`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Properties:` + fmt.Sprintf("%v", this.Properties) + `,`,
		`TokenMetaData:` + strings.Replace(this.TokenMetaData.String(), "MetaData", "MetaData", 1) + `,`,
		`Reserved:` + fmt.Sprintf("%v", this.Reserved) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MECTRoles) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MECTRoles{`,
		`Roles:` + fmt.Sprintf("%v", this.Roles) + `,`,
		`}`,
	}, "")
	return s
}
func (this *MetaData) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MetaData{`,
		`Nonce:` + fmt.Sprintf("%v", this.Nonce) + `,`,
		`Name:` + fmt.Sprintf("%v", this.Name) + `,`,
		`Creator:` + fmt.Sprintf("%v", this.Creator) + `,`,
		`Royalties:` + fmt.Sprintf("%v", this.Royalties) + `,`,
		`Hash:` + fmt.Sprintf("%v", this.Hash) + `,`,
		`URIs:` + fmt.Sprintf("%v", this.URIs) + `,`,
		`Attributes:` + fmt.Sprintf("%v", this.Attributes) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringMect(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MECToken) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMect
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
			return fmt.Errorf("proto: MECToken: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MECToken: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_ME_MotherEarth_me_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Properties", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Properties = append(m.Properties[:0], dAtA[iNdEx:postIndex]...)
			if m.Properties == nil {
				m.Properties = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenMetaData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
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
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TokenMetaData == nil {
				m.TokenMetaData = &MetaData{}
			}
			if err := m.TokenMetaData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserved", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reserved = append(m.Reserved[:0], dAtA[iNdEx:postIndex]...)
			if m.Reserved == nil {
				m.Reserved = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMect
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
func (m *MECTRoles) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMect
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
			return fmt.Errorf("proto: MECTRoles: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MECTRoles: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Roles", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Roles = append(m.Roles, make([]byte, postIndex-iNdEx))
			copy(m.Roles[len(m.Roles)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMect
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
func (m *MetaData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMect
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
			return fmt.Errorf("proto: MetaData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MetaData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = append(m.Name[:0], dAtA[iNdEx:postIndex]...)
			if m.Name == nil {
				m.Name = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Royalties", wireType)
			}
			m.Royalties = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Royalties |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URIs", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URIs = append(m.URIs, make([]byte, postIndex-iNdEx))
			copy(m.URIs[len(m.URIs)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attributes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMect
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMect
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMect
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attributes = append(m.Attributes[:0], dAtA[iNdEx:postIndex]...)
			if m.Attributes == nil {
				m.Attributes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMect(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMect
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
func skipMect(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMect
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
					return 0, ErrIntOverflowMect
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
					return 0, ErrIntOverflowMect
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
				return 0, ErrInvalidLengthMect
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMect
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMect
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMect        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMect          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMect = fmt.Errorf("proto: unexpected end of group")
)