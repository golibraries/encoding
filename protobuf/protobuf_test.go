package protobuf_test

import (
	"testing"

	. "github.com/frankban/quicktest"
	proto "github.com/golibraries/encoding/protobuf"

	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

type Access struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type    int64  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Account []byte `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Enable  bool   `protobuf:"varint,3,opt,name=enable,proto3" json:"enable,omitempty"`
}

func (x *Access) Reset() {
	*x = Access{}
	if protoimpl.UnsafeEnabled {
		mi := &file_external_proto_cli_struct_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_external_proto_cli_struct_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_external_proto_cli_struct_proto_rawDescGZIP(), []int{0}
}

func (x *Access) GetType() int64 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *Access) GetAccount() []byte {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *Access) GetEnable() bool {
	if x != nil {
		return x.Enable
	}
	return false
}

var File_external_proto_cli_struct_proto protoreflect.FileDescriptor

var file_external_proto_cli_struct_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6c, 0x69, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x63, 0x6c, 0x69, 0x22, 0x4e, 0x0a, 0x06, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x73, 0x6c, 0x6f, 0x74, 0x73, 0x2d,
	0x6e, 0x61, 0x6e, 0x6f, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x2f, 0x70, 0x62, 0x2f, 0x63, 0x6c, 0x69, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_external_proto_cli_struct_proto_rawDescOnce sync.Once
	file_external_proto_cli_struct_proto_rawDescData = file_external_proto_cli_struct_proto_rawDesc
)

func file_external_proto_cli_struct_proto_rawDescGZIP() []byte {
	file_external_proto_cli_struct_proto_rawDescOnce.Do(func() {
		file_external_proto_cli_struct_proto_rawDescData = protoimpl.X.CompressGZIP(file_external_proto_cli_struct_proto_rawDescData)
	})
	return file_external_proto_cli_struct_proto_rawDescData
}

var file_external_proto_cli_struct_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_external_proto_cli_struct_proto_goTypes = []interface{}{
	(*Access)(nil), // 0: cli.Access
}
var file_external_proto_cli_struct_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_external_proto_cli_struct_proto_init() }
func file_external_proto_cli_struct_proto_init() {
	if File_external_proto_cli_struct_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_external_proto_cli_struct_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Access); i {
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
			RawDescriptor: file_external_proto_cli_struct_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_external_proto_cli_struct_proto_goTypes,
		DependencyIndexes: file_external_proto_cli_struct_proto_depIdxs,
		MessageInfos:      file_external_proto_cli_struct_proto_msgTypes,
	}.Build()
	File_external_proto_cli_struct_proto = out.File
	file_external_proto_cli_struct_proto_rawDesc = nil
	file_external_proto_cli_struct_proto_goTypes = nil
	file_external_proto_cli_struct_proto_depIdxs = nil
}

func TestProtobuf(t *testing.T) {
	c := New(t)
	c.Run("TestProtobuf", func(c *C) {
		// create a new protobuf message
		msg := &Access{
			Type:    1,
			Account: []byte("test"),
			Enable:  true,
		}
		// encode the message
		data, err := proto.Marshal(msg)
		c.Assert(err, IsNil)
		// decode the message
		msg2 := &Access{}
		err = proto.Unmarshal(data, msg2)
		c.Assert(err, IsNil)
		c.Assert(msg2.Type, Equals, msg.Type)
		c.Assert(msg2.Account, DeepEquals, msg.Account)
		c.Assert(msg2.Enable, Equals, msg.Enable)
	})
}
