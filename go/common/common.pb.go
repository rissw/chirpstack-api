// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Modulation int32

const (
	// LoRa
	Modulation_LORA Modulation = 0
	// FSK
	Modulation_FSK Modulation = 1
	// LR-FHSS
	Modulation_LR_FHSS Modulation = 2
)

var Modulation_name = map[int32]string{
	0: "LORA",
	1: "FSK",
	2: "LR_FHSS",
}

var Modulation_value = map[string]int32{
	"LORA":    0,
	"FSK":     1,
	"LR_FHSS": 2,
}

func (x Modulation) String() string {
	return proto.EnumName(Modulation_name, int32(x))
}

func (Modulation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type Region int32

const (
	// EU868
	Region_EU868 Region = 0
	// US915
	Region_US915 Region = 2
	// CN779
	Region_CN779 Region = 3
	// EU433
	Region_EU433 Region = 4
	// AU915
	Region_AU915 Region = 5
	// CN470
	Region_CN470 Region = 6
	// AS923
	Region_AS923 Region = 7
	// AS923 with -1.80 MHz frequency offset
	Region_AS923_2 Region = 12
	// AS923 with -6.60 MHz frequency offset
	Region_AS923_3 Region = 13
	// AS923 with -5.90 MHz frequency offset
	Region_AS923_4 Region = 14
	// KR920
	Region_KR920 Region = 8
	// IN865
	Region_IN865 Region = 9
	// RU864
	Region_RU864 Region = 10
	// ISM2400 (LoRaWAN 2.4 GHz)
	Region_ISM2400 Region = 11
	// KZ865
	Region_KZ865 Region = 15
)

var Region_name = map[int32]string{
	0:  "EU868",
	2:  "US915",
	3:  "CN779",
	4:  "EU433",
	5:  "AU915",
	6:  "CN470",
	7:  "AS923",
	12: "AS923_2",
	13: "AS923_3",
	14: "AS923_4",
	8:  "KR920",
	9:  "IN865",
	10: "RU864",
	11: "ISM2400",
	15: "KZ865",
}

var Region_value = map[string]int32{
	"EU868":   0,
	"US915":   2,
	"CN779":   3,
	"EU433":   4,
	"AU915":   5,
	"CN470":   6,
	"AS923":   7,
	"AS923_2": 12,
	"AS923_3": 13,
	"AS923_4": 14,
	"KR920":   8,
	"IN865":   9,
	"RU864":   10,
	"ISM2400": 11,
	"KZ865":   15,
}

func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}

func (Region) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

type MType int32

const (
	// JoinRequest.
	MType_JoinRequest MType = 0
	// JoinAccept.
	MType_JoinAccept MType = 1
	// UnconfirmedDataUp.
	MType_UnconfirmedDataUp MType = 2
	// UnconfirmedDataDown.
	MType_UnconfirmedDataDown MType = 3
	// ConfirmedDataUp.
	MType_ConfirmedDataUp MType = 4
	// ConfirmedDataDown.
	MType_ConfirmedDataDown MType = 5
	// RejoinRequest.
	MType_RejoinRequest MType = 6
	// Proprietary.
	MType_Proprietary MType = 7
)

var MType_name = map[int32]string{
	0: "JoinRequest",
	1: "JoinAccept",
	2: "UnconfirmedDataUp",
	3: "UnconfirmedDataDown",
	4: "ConfirmedDataUp",
	5: "ConfirmedDataDown",
	6: "RejoinRequest",
	7: "Proprietary",
}

var MType_value = map[string]int32{
	"JoinRequest":         0,
	"JoinAccept":          1,
	"UnconfirmedDataUp":   2,
	"UnconfirmedDataDown": 3,
	"ConfirmedDataUp":     4,
	"ConfirmedDataDown":   5,
	"RejoinRequest":       6,
	"Proprietary":         7,
}

func (x MType) String() string {
	return proto.EnumName(MType_name, int32(x))
}

func (MType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

type LocationSource int32

const (
	// Unknown.
	LocationSource_UNKNOWN LocationSource = 0
	// GPS.
	LocationSource_GPS LocationSource = 1
	// Manually configured.
	LocationSource_CONFIG LocationSource = 2
	// Geo resolver (TDOA).
	LocationSource_GEO_RESOLVER_TDOA LocationSource = 3
	// Geo resolver (RSSI).
	LocationSource_GEO_RESOLVER_RSSI LocationSource = 4
	// Geo resolver (GNSS).
	LocationSource_GEO_RESOLVER_GNSS LocationSource = 5
	// Geo resolver (WIFI).
	LocationSource_GEO_RESOLVER_WIFI LocationSource = 6
)

var LocationSource_name = map[int32]string{
	0: "UNKNOWN",
	1: "GPS",
	2: "CONFIG",
	3: "GEO_RESOLVER_TDOA",
	4: "GEO_RESOLVER_RSSI",
	5: "GEO_RESOLVER_GNSS",
	6: "GEO_RESOLVER_WIFI",
}

var LocationSource_value = map[string]int32{
	"UNKNOWN":           0,
	"GPS":               1,
	"CONFIG":            2,
	"GEO_RESOLVER_TDOA": 3,
	"GEO_RESOLVER_RSSI": 4,
	"GEO_RESOLVER_GNSS": 5,
	"GEO_RESOLVER_WIFI": 6,
}

func (x LocationSource) String() string {
	return proto.EnumName(LocationSource_name, int32(x))
}

func (LocationSource) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

type KeyEnvelope struct {
	// KEK label.
	KekLabel string `protobuf:"bytes,1,opt,name=kek_label,json=kekLabel,proto3" json:"kek_label,omitempty"`
	// AES key (when the kek_label is set, this key is encrypted using a key
	// known to the join-server and application-server.
	// For more information please refer to the LoRaWAN Backend Interface
	// 'Key Transport Security' section.
	AesKey               []byte   `protobuf:"bytes,2,opt,name=aes_key,json=aesKey,proto3" json:"aes_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyEnvelope) Reset()         { *m = KeyEnvelope{} }
func (m *KeyEnvelope) String() string { return proto.CompactTextString(m) }
func (*KeyEnvelope) ProtoMessage()    {}
func (*KeyEnvelope) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyEnvelope) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyEnvelope.Unmarshal(m, b)
}
func (m *KeyEnvelope) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyEnvelope.Marshal(b, m, deterministic)
}
func (m *KeyEnvelope) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyEnvelope.Merge(m, src)
}
func (m *KeyEnvelope) XXX_Size() int {
	return xxx_messageInfo_KeyEnvelope.Size(m)
}
func (m *KeyEnvelope) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyEnvelope.DiscardUnknown(m)
}

var xxx_messageInfo_KeyEnvelope proto.InternalMessageInfo

func (m *KeyEnvelope) GetKekLabel() string {
	if m != nil {
		return m.KekLabel
	}
	return ""
}

func (m *KeyEnvelope) GetAesKey() []byte {
	if m != nil {
		return m.AesKey
	}
	return nil
}

type Location struct {
	// Latitude.
	Latitude float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	// Longitude.
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
	// Altitude.
	Altitude float64 `protobuf:"fixed64,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
	// Location source.
	Source LocationSource `protobuf:"varint,4,opt,name=source,proto3,enum=common.LocationSource" json:"source,omitempty"`
	// Accuracy (in meters).
	Accuracy             uint32   `protobuf:"varint,5,opt,name=accuracy,proto3" json:"accuracy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Location) Reset()         { *m = Location{} }
func (m *Location) String() string { return proto.CompactTextString(m) }
func (*Location) ProtoMessage()    {}
func (*Location) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *Location) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Location.Unmarshal(m, b)
}
func (m *Location) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Location.Marshal(b, m, deterministic)
}
func (m *Location) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Location.Merge(m, src)
}
func (m *Location) XXX_Size() int {
	return xxx_messageInfo_Location.Size(m)
}
func (m *Location) XXX_DiscardUnknown() {
	xxx_messageInfo_Location.DiscardUnknown(m)
}

var xxx_messageInfo_Location proto.InternalMessageInfo

func (m *Location) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

func (m *Location) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Location) GetAltitude() float64 {
	if m != nil {
		return m.Altitude
	}
	return 0
}

func (m *Location) GetSource() LocationSource {
	if m != nil {
		return m.Source
	}
	return LocationSource_UNKNOWN
}

func (m *Location) GetAccuracy() uint32 {
	if m != nil {
		return m.Accuracy
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.Modulation", Modulation_name, Modulation_value)
	proto.RegisterEnum("common.Region", Region_name, Region_value)
	proto.RegisterEnum("common.MType", MType_name, MType_value)
	proto.RegisterEnum("common.LocationSource", LocationSource_name, LocationSource_value)
	proto.RegisterType((*KeyEnvelope)(nil), "common.KeyEnvelope")
	proto.RegisterType((*Location)(nil), "common.Location")
}

func init() {
	proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6)
}

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 597 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x53, 0x4d, 0x53, 0x1a, 0x4b,
	0x14, 0xb5, 0xf9, 0x18, 0xe0, 0xa2, 0x78, 0x6d, 0xeb, 0x3d, 0xa9, 0xf7, 0xb2, 0xa0, 0x5c, 0x51,
	0x54, 0x02, 0x04, 0x50, 0x61, 0x89, 0x08, 0x84, 0x80, 0x03, 0xd5, 0xe3, 0xc4, 0x8a, 0x1b, 0xaa,
	0x19, 0x3b, 0x38, 0x61, 0x9c, 0x9e, 0x0c, 0x83, 0x29, 0xfe, 0x44, 0x7e, 0x47, 0x16, 0xd9, 0xe6,
	0xff, 0xa5, 0x7a, 0x06, 0x25, 0xea, 0xaa, 0xcf, 0x39, 0xf7, 0xdc, 0x8f, 0xae, 0xee, 0x0b, 0x87,
	0x96, 0xbc, 0xbf, 0x97, 0x6e, 0x25, 0x3a, 0xca, 0x9e, 0x2f, 0x03, 0x49, 0xb5, 0x88, 0x1d, 0x77,
	0x20, 0x3b, 0x14, 0xeb, 0xae, 0xfb, 0x20, 0x1c, 0xe9, 0x09, 0xfa, 0x3f, 0x64, 0x16, 0x62, 0x31,
	0x75, 0xf8, 0x4c, 0x38, 0x79, 0x52, 0x20, 0xc5, 0x0c, 0x4b, 0x2f, 0xc4, 0x62, 0xa4, 0x38, 0x3d,
	0x82, 0x14, 0x17, 0xcb, 0xe9, 0x42, 0xac, 0xf3, 0xb1, 0x02, 0x29, 0xee, 0x32, 0x8d, 0x8b, 0xe5,
	0x50, 0xac, 0x8f, 0x7f, 0x11, 0x48, 0x8f, 0xa4, 0xc5, 0x03, 0x5b, 0xba, 0xf4, 0x3f, 0x48, 0x3b,
	0x3c, 0xb0, 0x83, 0xd5, 0xad, 0x08, 0x2b, 0x10, 0xf6, 0xc4, 0xe9, 0x1b, 0xc8, 0x38, 0xd2, 0x9d,
	0x47, 0xc1, 0x58, 0x18, 0xdc, 0x0a, 0x2a, 0x93, 0x3b, 0x9b, 0xcc, 0x78, 0x94, 0xf9, 0xc8, 0x69,
	0x19, 0xb4, 0xa5, 0x5c, 0xf9, 0x96, 0xc8, 0x27, 0x0a, 0xa4, 0x98, 0xab, 0xfd, 0x5b, 0xde, 0x5c,
	0xe7, 0xb1, 0xaf, 0x11, 0x46, 0xd9, 0xc6, 0x15, 0xd6, 0xb2, 0xac, 0x95, 0xcf, 0xad, 0x75, 0x3e,
	0x59, 0x20, 0xc5, 0x3d, 0xf6, 0xc4, 0x4b, 0x6f, 0x01, 0x2e, 0xe5, 0xed, 0xca, 0x89, 0xe6, 0x4d,
	0x43, 0x62, 0x34, 0x66, 0x6d, 0xdc, 0xa1, 0x29, 0x88, 0xf7, 0x8c, 0x21, 0x12, 0x9a, 0x85, 0xd4,
	0x88, 0x4d, 0x7b, 0x1f, 0x0c, 0x03, 0x63, 0xa5, 0xdf, 0x04, 0x34, 0x26, 0xe6, 0xca, 0x9a, 0x81,
	0x64, 0xd7, 0x6c, 0x9e, 0x36, 0x71, 0x47, 0x41, 0xd3, 0x68, 0xbd, 0x3f, 0xc1, 0x98, 0x82, 0x1d,
	0xfd, 0xec, 0xac, 0x85, 0xf1, 0xc8, 0xd0, 0xa8, 0xd7, 0x31, 0xa1, 0x60, 0xdb, 0x54, 0x86, 0x64,
	0x64, 0x68, 0x9c, 0x55, 0x51, 0x0b, 0x55, 0xa3, 0x55, 0xab, 0x63, 0x4a, 0x35, 0x09, 0xe1, 0xb4,
	0x86, 0xbb, 0x5b, 0x52, 0xc7, 0xbd, 0x2d, 0x69, 0x60, 0x4e, 0x65, 0x0c, 0x59, 0xab, 0x56, 0xc5,
	0xb4, 0x82, 0x03, 0xbd, 0x79, 0x7a, 0x82, 0x19, 0x05, 0x99, 0xd9, 0x3c, 0x6d, 0x20, 0x28, 0xf7,
	0xc0, 0xb8, 0xac, 0x35, 0xaa, 0x55, 0xcc, 0x86, 0xee, 0x1b, 0x65, 0xd9, 0x2f, 0xfd, 0x24, 0x90,
	0xbc, 0xbc, 0x5a, 0x7b, 0x82, 0xee, 0x43, 0xf6, 0xa3, 0xb4, 0x5d, 0x26, 0xbe, 0xad, 0xc4, 0x32,
	0xc0, 0x1d, 0x9a, 0x03, 0x50, 0x42, 0xdb, 0xb2, 0x84, 0x17, 0x20, 0xa1, 0xff, 0xc0, 0x81, 0xe9,
	0x5a, 0xd2, 0xfd, 0x62, 0xfb, 0xf7, 0xe2, 0xf6, 0x82, 0x07, 0xdc, 0xf4, 0x30, 0x46, 0x8f, 0xe0,
	0xf0, 0x85, 0x7c, 0x21, 0xbf, 0xbb, 0x18, 0xa7, 0x87, 0xb0, 0xdf, 0x79, 0xe1, 0x4e, 0xa8, 0x22,
	0x9d, 0x57, 0xde, 0x24, 0x3d, 0x80, 0x3d, 0x26, 0xbe, 0xfe, 0xd5, 0x5e, 0x53, 0xf3, 0x4c, 0x7c,
	0xe9, 0xf9, 0xb6, 0x08, 0xb8, 0xbf, 0xc6, 0x54, 0xe9, 0x07, 0x81, 0xdc, 0xf3, 0x77, 0x54, 0xb7,
	0x32, 0xf5, 0xa1, 0x3e, 0xbe, 0xd6, 0xa3, 0x87, 0xe9, 0x4f, 0x0c, 0x24, 0x14, 0x40, 0xeb, 0x8c,
	0xf5, 0xde, 0xa0, 0x8f, 0x31, 0xd5, 0xaf, 0xdf, 0x1d, 0x4f, 0x59, 0xd7, 0x18, 0x8f, 0x3e, 0x75,
	0xd9, 0xf4, 0xea, 0x62, 0xdc, 0xc6, 0xf8, 0x2b, 0x99, 0x19, 0xc6, 0x20, 0x9a, 0xee, 0x99, 0xdc,
	0xd7, 0x0d, 0x03, 0x93, 0xaf, 0xe4, 0xeb, 0x41, 0x6f, 0x80, 0xda, 0xf9, 0x67, 0xc8, 0xdb, 0xb2,
	0x6c, 0xdd, 0xd9, 0xbe, 0xb7, 0x0c, 0xb8, 0xb5, 0x28, 0x73, 0xcf, 0xde, 0x7c, 0xb8, 0xf3, 0x6c,
	0x27, 0x3c, 0x27, 0x6a, 0x8d, 0x26, 0xe4, 0xa6, 0x3c, 0xb7, 0x83, 0xbb, 0xd5, 0x4c, 0x45, 0x2b,
	0x33, 0x5f, 0x5a, 0x9c, 0xfb, 0x95, 0x6d, 0xe2, 0x3b, 0xee, 0xd9, 0x95, 0xb9, 0xac, 0x3c, 0xd4,
	0x37, 0xeb, 0x37, 0xd3, 0xc2, 0xfd, 0xab, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xe8, 0x1b,
	0x9c, 0x96, 0x03, 0x00, 0x00,
}
