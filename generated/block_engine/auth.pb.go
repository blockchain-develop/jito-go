// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: auth.proto

package jito_go

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_RELAYER                Role = 0
	Role_SEARCHER               Role = 1
	Role_VALIDATOR              Role = 2
	Role_SHREDSTREAM_SUBSCRIBER Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "RELAYER",
		1: "SEARCHER",
		2: "VALIDATOR",
		3: "SHREDSTREAM_SUBSCRIBER",
	}
	Role_value = map[string]int32{
		"RELAYER":                0,
		"SEARCHER":               1,
		"VALIDATOR":              2,
		"SHREDSTREAM_SUBSCRIBER": 3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type GenerateAuthChallengeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / Role the client is attempting to generate tokens for.
	Role Role `protobuf:"varint,1,opt,name=role,proto3,enum=auth.Role" json:"role,omitempty"`
	// / Client's 32 byte pubkey.
	Pubkey        []byte `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateAuthChallengeRequest) Reset() {
	*x = GenerateAuthChallengeRequest{}
	mi := &file_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAuthChallengeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthChallengeRequest) ProtoMessage() {}

func (x *GenerateAuthChallengeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthChallengeRequest.ProtoReflect.Descriptor instead.
func (*GenerateAuthChallengeRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateAuthChallengeRequest) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_RELAYER
}

func (x *GenerateAuthChallengeRequest) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

type GenerateAuthChallengeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Challenge     string                 `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateAuthChallengeResponse) Reset() {
	*x = GenerateAuthChallengeResponse{}
	mi := &file_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAuthChallengeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthChallengeResponse) ProtoMessage() {}

func (x *GenerateAuthChallengeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthChallengeResponse.ProtoReflect.Descriptor instead.
func (*GenerateAuthChallengeResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateAuthChallengeResponse) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

type GenerateAuthTokensRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / The pre-signed challenge.
	Challenge string `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	// / The signing keypair's corresponding 32 byte pubkey.
	ClientPubkey []byte `protobuf:"bytes,2,opt,name=client_pubkey,json=clientPubkey,proto3" json:"client_pubkey,omitempty"`
	// / The 64 byte signature of the challenge signed by the client's private key. The private key must correspond to
	// the pubkey passed in the [GenerateAuthChallenge] method. The client is expected to sign the challenge token
	// prepended with their pubkey. For example sign(pubkey, challenge).
	SignedChallenge []byte `protobuf:"bytes,3,opt,name=signed_challenge,json=signedChallenge,proto3" json:"signed_challenge,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GenerateAuthTokensRequest) Reset() {
	*x = GenerateAuthTokensRequest{}
	mi := &file_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAuthTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthTokensRequest) ProtoMessage() {}

func (x *GenerateAuthTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthTokensRequest.ProtoReflect.Descriptor instead.
func (*GenerateAuthTokensRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *GenerateAuthTokensRequest) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *GenerateAuthTokensRequest) GetClientPubkey() []byte {
	if x != nil {
		return x.ClientPubkey
	}
	return nil
}

func (x *GenerateAuthTokensRequest) GetSignedChallenge() []byte {
	if x != nil {
		return x.SignedChallenge
	}
	return nil
}

type Token struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / The token.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// / When the token will expire.
	ExpiresAtUtc  *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=expires_at_utc,json=expiresAtUtc,proto3" json:"expires_at_utc,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Token) Reset() {
	*x = Token{}
	mi := &file_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Token) GetExpiresAtUtc() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAtUtc
	}
	return nil
}

type GenerateAuthTokensResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / The token granting access to resources.
	AccessToken *Token `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	// / The token used to refresh the access_token. This has a longer TTL than the access_token.
	RefreshToken  *Token `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GenerateAuthTokensResponse) Reset() {
	*x = GenerateAuthTokensResponse{}
	mi := &file_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GenerateAuthTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthTokensResponse) ProtoMessage() {}

func (x *GenerateAuthTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthTokensResponse.ProtoReflect.Descriptor instead.
func (*GenerateAuthTokensResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateAuthTokensResponse) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

func (x *GenerateAuthTokensResponse) GetRefreshToken() *Token {
	if x != nil {
		return x.RefreshToken
	}
	return nil
}

type RefreshAccessTokenRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / Non-expired refresh token obtained from the [GenerateAuthTokens] method.
	RefreshToken  string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshAccessTokenRequest) Reset() {
	*x = RefreshAccessTokenRequest{}
	mi := &file_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenRequest) ProtoMessage() {}

func (x *RefreshAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *RefreshAccessTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshAccessTokenResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// / Fresh access_token.
	AccessToken   *Token `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshAccessTokenResponse) Reset() {
	*x = RefreshAccessTokenResponse{}
	mi := &file_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshAccessTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAccessTokenResponse) ProtoMessage() {}

func (x *RefreshAccessTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAccessTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshAccessTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshAccessTokenResponse) GetAccessToken() *Token {
	if x != nil {
		return x.AccessToken
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

const file_auth_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"auth.proto\x12\x04auth\x1a\x1fgoogle/protobuf/timestamp.proto\"V\n" +
	"\x1cGenerateAuthChallengeRequest\x12\x1e\n" +
	"\x04role\x18\x01 \x01(\x0e2\n" +
	".auth.RoleR\x04role\x12\x16\n" +
	"\x06pubkey\x18\x02 \x01(\fR\x06pubkey\"=\n" +
	"\x1dGenerateAuthChallengeResponse\x12\x1c\n" +
	"\tchallenge\x18\x01 \x01(\tR\tchallenge\"\x89\x01\n" +
	"\x19GenerateAuthTokensRequest\x12\x1c\n" +
	"\tchallenge\x18\x01 \x01(\tR\tchallenge\x12#\n" +
	"\rclient_pubkey\x18\x02 \x01(\fR\fclientPubkey\x12)\n" +
	"\x10signed_challenge\x18\x03 \x01(\fR\x0fsignedChallenge\"_\n" +
	"\x05Token\x12\x14\n" +
	"\x05value\x18\x01 \x01(\tR\x05value\x12@\n" +
	"\x0eexpires_at_utc\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\fexpiresAtUtc\"~\n" +
	"\x1aGenerateAuthTokensResponse\x12.\n" +
	"\faccess_token\x18\x01 \x01(\v2\v.auth.TokenR\vaccessToken\x120\n" +
	"\rrefresh_token\x18\x02 \x01(\v2\v.auth.TokenR\frefreshToken\"@\n" +
	"\x19RefreshAccessTokenRequest\x12#\n" +
	"\rrefresh_token\x18\x01 \x01(\tR\frefreshToken\"L\n" +
	"\x1aRefreshAccessTokenResponse\x12.\n" +
	"\faccess_token\x18\x01 \x01(\v2\v.auth.TokenR\vaccessToken*L\n" +
	"\x04Role\x12\v\n" +
	"\aRELAYER\x10\x00\x12\f\n" +
	"\bSEARCHER\x10\x01\x12\r\n" +
	"\tVALIDATOR\x10\x02\x12\x1a\n" +
	"\x16SHREDSTREAM_SUBSCRIBER\x10\x032\xa7\x02\n" +
	"\vAuthService\x12b\n" +
	"\x15GenerateAuthChallenge\x12\".auth.GenerateAuthChallengeRequest\x1a#.auth.GenerateAuthChallengeResponse\"\x00\x12Y\n" +
	"\x12GenerateAuthTokens\x12\x1f.auth.GenerateAuthTokensRequest\x1a .auth.GenerateAuthTokensResponse\"\x00\x12Y\n" +
	"\x12RefreshAccessToken\x12\x1f.auth.RefreshAccessTokenRequest\x1a .auth.RefreshAccessTokenResponse\"\x00b\x06proto3"

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData []byte
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)))
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_auth_proto_goTypes = []any{
	(Role)(0),                             // 0: auth.Role
	(*GenerateAuthChallengeRequest)(nil),  // 1: auth.GenerateAuthChallengeRequest
	(*GenerateAuthChallengeResponse)(nil), // 2: auth.GenerateAuthChallengeResponse
	(*GenerateAuthTokensRequest)(nil),     // 3: auth.GenerateAuthTokensRequest
	(*Token)(nil),                         // 4: auth.Token
	(*GenerateAuthTokensResponse)(nil),    // 5: auth.GenerateAuthTokensResponse
	(*RefreshAccessTokenRequest)(nil),     // 6: auth.RefreshAccessTokenRequest
	(*RefreshAccessTokenResponse)(nil),    // 7: auth.RefreshAccessTokenResponse
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
}
var file_auth_proto_depIdxs = []int32{
	0, // 0: auth.GenerateAuthChallengeRequest.role:type_name -> auth.Role
	8, // 1: auth.Token.expires_at_utc:type_name -> google.protobuf.Timestamp
	4, // 2: auth.GenerateAuthTokensResponse.access_token:type_name -> auth.Token
	4, // 3: auth.GenerateAuthTokensResponse.refresh_token:type_name -> auth.Token
	4, // 4: auth.RefreshAccessTokenResponse.access_token:type_name -> auth.Token
	1, // 5: auth.AuthService.GenerateAuthChallenge:input_type -> auth.GenerateAuthChallengeRequest
	3, // 6: auth.AuthService.GenerateAuthTokens:input_type -> auth.GenerateAuthTokensRequest
	6, // 7: auth.AuthService.RefreshAccessToken:input_type -> auth.RefreshAccessTokenRequest
	2, // 8: auth.AuthService.GenerateAuthChallenge:output_type -> auth.GenerateAuthChallengeResponse
	5, // 9: auth.AuthService.GenerateAuthTokens:output_type -> auth.GenerateAuthTokensResponse
	7, // 10: auth.AuthService.RefreshAccessToken:output_type -> auth.RefreshAccessTokenResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_auth_proto_rawDesc), len(file_auth_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
