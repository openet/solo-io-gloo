// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/subset.proto

package v1

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Subset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        map[string]string      `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Subset) Reset() {
	*x = Subset{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subset) ProtoMessage() {}

func (x *Subset) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subset.ProtoReflect.Descriptor instead.
func (*Subset) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescGZIP(), []int{0}
}

func (x *Subset) GetValues() map[string]string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto protoreflect.FileDescriptor

const file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDesc = "" +
	"\n" +
	"9github.com/solo-io/gloo/projects/gloo/api/v1/subset.proto\x12\fgloo.solo.io\x1a\x12extproto/ext.proto\"}\n" +
	"\x06Subset\x128\n" +
	"\x06values\x18\x01 \x03(\v2 .gloo.solo.io.Subset.ValuesEntryR\x06values\x1a9\n" +
	"\vValuesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B>\xb8\xf5\x04\x01\xc0\xf5\x04\x01\xd0\xf5\x04\x01Z0github.com/solo-io/gloo/projects/gloo/pkg/api/v1b\x06proto3"

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescData []byte
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDesc), len(file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDesc)))
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_goTypes = []any{
	(*Subset)(nil), // 0: gloo.solo.io.Subset
	nil,            // 1: gloo.solo.io.Subset.ValuesEntry
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_depIdxs = []int32{
	1, // 0: gloo.solo.io.Subset.values:type_name -> gloo.solo.io.Subset.ValuesEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDesc), len(file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_subset_proto_depIdxs = nil
}
