// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.8.0
// source: pkg/config/po/re.proto

package po

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Rq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 当前机器号
	DeviceId string `protobuf:"bytes,1,opt,name=deviceId,proto3" json:"deviceId,omitempty"`
	// 当前公网IP
	Ip string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	// 当前代理端口号
	Port int64 `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	// 拨号时间戳
	DialTime int64 `protobuf:"varint,4,opt,name=dialTime,proto3" json:"dialTime,omitempty"`
}

func (x *Rq) Reset() {
	*x = Rq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_config_po_re_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rq) ProtoMessage() {}

func (x *Rq) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_po_re_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rq.ProtoReflect.Descriptor instead.
func (*Rq) Descriptor() ([]byte, []int) {
	return file_pkg_config_po_re_proto_rawDescGZIP(), []int{0}
}

func (x *Rq) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Rq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Rq) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Rq) GetDialTime() int64 {
	if x != nil {
		return x.DialTime
	}
	return 0
}

type Rp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// -------------- 标准参数
	// 服务器是否处理成功
	Code int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	// 是否需要验证 0需要 1不需要
	NoAuth int64 `protobuf:"varint,2,opt,name=noAuth,proto3" json:"noAuth,omitempty"`
	// 授权/限流等配置
	ConfigData *Config `protobuf:"bytes,3,opt,name=configData,proto3" json:"configData,omitempty"`
	// -------------- 标准代理参数
	// IP切换时间间隔（秒）
	IpInterval int64 `protobuf:"varint,4,opt,name=ipInterval,proto3" json:"ipInterval,omitempty"`
	// 当日IP是否允许重复 0允许 1不允许
	IpRepeat int64 `protobuf:"varint,5,opt,name=ipRepeat,proto3" json:"ipRepeat,omitempty"`
}

func (x *Rp) Reset() {
	*x = Rp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_config_po_re_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rp) ProtoMessage() {}

func (x *Rp) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_config_po_re_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rp.ProtoReflect.Descriptor instead.
func (*Rp) Descriptor() ([]byte, []int) {
	return file_pkg_config_po_re_proto_rawDescGZIP(), []int{1}
}

func (x *Rp) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Rp) GetNoAuth() int64 {
	if x != nil {
		return x.NoAuth
	}
	return 0
}

func (x *Rp) GetConfigData() *Config {
	if x != nil {
		return x.ConfigData
	}
	return nil
}

func (x *Rp) GetIpInterval() int64 {
	if x != nil {
		return x.IpInterval
	}
	return 0
}

func (x *Rp) GetIpRepeat() int64 {
	if x != nil {
		return x.IpRepeat
	}
	return 0
}

var File_pkg_config_po_re_proto protoreflect.FileDescriptor

var file_pkg_config_po_re_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x6f, 0x2f,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2f, 0x70, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a, 0x02, 0x52, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x95, 0x01, 0x0a, 0x02, 0x52, 0x70, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x6e, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x70, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x70, 0x52, 0x65, 0x70, 0x65, 0x61, 0x74, 0x42, 0x34,
	0x0a, 0x1b, 0x6f, 0x72, 0x67, 0x2e, 0x6a, 0x65, 0x65, 0x63, 0x67, 0x2e, 0x6d, 0x6f, 0x64, 0x75,
	0x6c, 0x65, 0x73, 0x2e, 0x76, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x06, 0x52,
	0x65, 0x50, 0x6f, 0x6a, 0x6f, 0x5a, 0x0d, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x70, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_config_po_re_proto_rawDescOnce sync.Once
	file_pkg_config_po_re_proto_rawDescData = file_pkg_config_po_re_proto_rawDesc
)

func file_pkg_config_po_re_proto_rawDescGZIP() []byte {
	file_pkg_config_po_re_proto_rawDescOnce.Do(func() {
		file_pkg_config_po_re_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_config_po_re_proto_rawDescData)
	})
	return file_pkg_config_po_re_proto_rawDescData
}

var file_pkg_config_po_re_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_config_po_re_proto_goTypes = []interface{}{
	(*Rq)(nil),     // 0: Rq
	(*Rp)(nil),     // 1: Rp
	(*Config)(nil), // 2: Config
}
var file_pkg_config_po_re_proto_depIdxs = []int32{
	2, // 0: Rp.configData:type_name -> Config
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_config_po_re_proto_init() }
func file_pkg_config_po_re_proto_init() {
	if File_pkg_config_po_re_proto != nil {
		return
	}
	file_pkg_config_po_config_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_config_po_re_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_config_po_re_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_config_po_re_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_config_po_re_proto_goTypes,
		DependencyIndexes: file_pkg_config_po_re_proto_depIdxs,
		MessageInfos:      file_pkg_config_po_re_proto_msgTypes,
	}.Build()
	File_pkg_config_po_re_proto = out.File
	file_pkg_config_po_re_proto_rawDesc = nil
	file_pkg_config_po_re_proto_goTypes = nil
	file_pkg_config_po_re_proto_depIdxs = nil
}
