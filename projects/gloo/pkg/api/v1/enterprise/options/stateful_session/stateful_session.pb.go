// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/stateful_session/stateful_session.proto

package stateful_session

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This configures the Envoy [Stateful Session](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/stateful_session_filter) filter for a listener
type StatefulSession struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to SessionState:
	//
	//	*StatefulSession_CookieBased
	//	*StatefulSession_HeaderBased
	SessionState isStatefulSession_SessionState `protobuf_oneof:"SessionState"`
	// If set to True, the HTTP request must be routed to the requested destination. If the requested destination is not available, Envoy returns 503. Defaults to False.
	Strict        bool `protobuf:"varint,3,opt,name=strict,proto3" json:"strict,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StatefulSession) Reset() {
	*x = StatefulSession{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StatefulSession) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatefulSession) ProtoMessage() {}

func (x *StatefulSession) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatefulSession.ProtoReflect.Descriptor instead.
func (*StatefulSession) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{0}
}

func (x *StatefulSession) GetSessionState() isStatefulSession_SessionState {
	if x != nil {
		return x.SessionState
	}
	return nil
}

func (x *StatefulSession) GetCookieBased() *CookieBasedSessionState {
	if x != nil {
		if x, ok := x.SessionState.(*StatefulSession_CookieBased); ok {
			return x.CookieBased
		}
	}
	return nil
}

func (x *StatefulSession) GetHeaderBased() *HeaderBasedSessionState {
	if x != nil {
		if x, ok := x.SessionState.(*StatefulSession_HeaderBased); ok {
			return x.HeaderBased
		}
	}
	return nil
}

func (x *StatefulSession) GetStrict() bool {
	if x != nil {
		return x.Strict
	}
	return false
}

type isStatefulSession_SessionState interface {
	isStatefulSession_SessionState()
}

type StatefulSession_CookieBased struct {
	// Configure a cookie based session state - https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate
	// Exactly one of `cookie_based` or `header_based` must be set.
	CookieBased *CookieBasedSessionState `protobuf:"bytes,1,opt,name=cookie_based,json=cookieBased,proto3,oneof"`
}

type StatefulSession_HeaderBased struct {
	// Configure a header based session state - https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate
	// Exactly one of `cookie_based` or `header_based` must be set.
	HeaderBased *HeaderBasedSessionState `protobuf:"bytes,2,opt,name=header_based,json=headerBased,proto3,oneof"`
}

func (*StatefulSession_CookieBased) isStatefulSession_SessionState() {}

func (*StatefulSession_HeaderBased) isStatefulSession_SessionState() {}

// Configuration for the [cookie-based session state](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/cookie/v3/cookie.proto#envoy-v3-api-msg-extensions-http-stateful-session-cookie-v3-cookiebasedsessionstate) filter
type CookieBasedSessionState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required, the cookie configuration used to track session state.
	Cookie        *CookieBasedSessionState_Cookie `protobuf:"bytes,1,opt,name=cookie,proto3" json:"cookie,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CookieBasedSessionState) Reset() {
	*x = CookieBasedSessionState{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CookieBasedSessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieBasedSessionState) ProtoMessage() {}

func (x *CookieBasedSessionState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieBasedSessionState.ProtoReflect.Descriptor instead.
func (*CookieBasedSessionState) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{1}
}

func (x *CookieBasedSessionState) GetCookie() *CookieBasedSessionState_Cookie {
	if x != nil {
		return x.Cookie
	}
	return nil
}

// Configuration for the [header-based session state](https://www.envoyproxy.io/docs/envoy/latest/api-v3/extensions/http/stateful_session/header/v3/header.proto#extension-envoy-http-stateful-session-header) filter
type HeaderBasedSessionState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required, the header used to track session state.
	HeaderName    string `protobuf:"bytes,1,opt,name=header_name,json=headerName,proto3" json:"header_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HeaderBasedSessionState) Reset() {
	*x = HeaderBasedSessionState{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeaderBasedSessionState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeaderBasedSessionState) ProtoMessage() {}

func (x *HeaderBasedSessionState) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeaderBasedSessionState.ProtoReflect.Descriptor instead.
func (*HeaderBasedSessionState) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{2}
}

func (x *HeaderBasedSessionState) GetHeaderName() string {
	if x != nil {
		return x.HeaderName
	}
	return ""
}

type CookieBasedSessionState_Cookie struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required, the name that will be used to obtain cookie value from downstream HTTP request or generate new cookie for downstream.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Path of cookie. This will be used to set the path of a new cookie when it is generated. If no path is specified here, no path will be set for the cookie.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// Duration of cookie. This will be used to set the expiry time of a new cookie when it is generated. Set this to 0s to use a session cookie and disable cookie expiration.
	Ttl           *durationpb.Duration `protobuf:"bytes,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CookieBasedSessionState_Cookie) Reset() {
	*x = CookieBasedSessionState_Cookie{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CookieBasedSessionState_Cookie) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CookieBasedSessionState_Cookie) ProtoMessage() {}

func (x *CookieBasedSessionState_Cookie) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CookieBasedSessionState_Cookie.ProtoReflect.Descriptor instead.
func (*CookieBasedSessionState_Cookie) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CookieBasedSessionState_Cookie) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CookieBasedSessionState_Cookie) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CookieBasedSessionState_Cookie) GetTtl() *durationpb.Duration {
	if x != nil {
		return x.Ttl
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto protoreflect.FileDescriptor

const file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc = "" +
	"\n" +
	"ggithub.com/solo-io/gloo/projects/gloo/api/v1/enterprise/options/stateful_session/stateful_session.proto\x12%stateful_session.options.gloo.solo.io\x1a\x1egoogle/protobuf/duration.proto\x1a\x17validate/validate.proto\"\x88\x02\n" +
	"\x0fStatefulSession\x12c\n" +
	"\fcookie_based\x18\x01 \x01(\v2>.stateful_session.options.gloo.solo.io.CookieBasedSessionStateH\x00R\vcookieBased\x12c\n" +
	"\fheader_based\x18\x02 \x01(\v2>.stateful_session.options.gloo.solo.io.HeaderBasedSessionStateH\x00R\vheaderBased\x12\x16\n" +
	"\x06strict\x18\x03 \x01(\bR\x06strictB\x13\n" +
	"\fSessionState\x12\x03\xf8B\x01\"\xf7\x01\n" +
	"\x17CookieBasedSessionState\x12g\n" +
	"\x06cookie\x18\x01 \x01(\v2E.stateful_session.options.gloo.solo.io.CookieBasedSessionState.CookieB\b\xfaB\x05\x8a\x01\x02\x10\x01R\x06cookie\x1as\n" +
	"\x06Cookie\x12\x1c\n" +
	"\x04name\x18\x01 \x01(\tB\b\xfaB\x05\x8a\x01\x02\x10\x01R\x04name\x12\x12\n" +
	"\x04path\x18\x02 \x01(\tR\x04path\x127\n" +
	"\x03ttl\x18\x03 \x01(\v2\x19.google.protobuf.DurationB\n" +
	"\xfaB\a\xaa\x01\x04\b\x012\x00R\x03ttl\"D\n" +
	"\x17HeaderBasedSessionState\x12)\n" +
	"\vheader_name\x18\x01 \x01(\tB\b\xfaB\x05\x8a\x01\x02\x10\x01R\n" +
	"headerNameBVZTgithub.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/stateful_sessionb\x06proto3"

var (
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData []byte
)

func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc), len(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc)))
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes = []any{
	(*StatefulSession)(nil),                // 0: stateful_session.options.gloo.solo.io.StatefulSession
	(*CookieBasedSessionState)(nil),        // 1: stateful_session.options.gloo.solo.io.CookieBasedSessionState
	(*HeaderBasedSessionState)(nil),        // 2: stateful_session.options.gloo.solo.io.HeaderBasedSessionState
	(*CookieBasedSessionState_Cookie)(nil), // 3: stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie
	(*durationpb.Duration)(nil),            // 4: google.protobuf.Duration
}
var file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs = []int32{
	1, // 0: stateful_session.options.gloo.solo.io.StatefulSession.cookie_based:type_name -> stateful_session.options.gloo.solo.io.CookieBasedSessionState
	2, // 1: stateful_session.options.gloo.solo.io.StatefulSession.header_based:type_name -> stateful_session.options.gloo.solo.io.HeaderBasedSessionState
	3, // 2: stateful_session.options.gloo.solo.io.CookieBasedSessionState.cookie:type_name -> stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie
	4, // 3: stateful_session.options.gloo.solo.io.CookieBasedSessionState.Cookie.ttl:type_name -> google.protobuf.Duration
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto != nil {
		return
	}
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes[0].OneofWrappers = []any{
		(*StatefulSession_CookieBased)(nil),
		(*StatefulSession_HeaderBased)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc), len(file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_v1_enterprise_options_stateful_session_stateful_session_proto_depIdxs = nil
}
