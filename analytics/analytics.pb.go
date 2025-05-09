// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: analytics/analytics.proto

package analyticsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ArticleAction int32

const (
	ArticleAction_UNKNOWN ArticleAction = 0
	ArticleAction_VIEW    ArticleAction = 1
	ArticleAction_CLICK   ArticleAction = 2
	ArticleAction_SHARE   ArticleAction = 3
	ArticleAction_COMMENT ArticleAction = 4
)

// Enum value maps for ArticleAction.
var (
	ArticleAction_name = map[int32]string{
		0: "UNKNOWN",
		1: "VIEW",
		2: "CLICK",
		3: "SHARE",
		4: "COMMENT",
	}
	ArticleAction_value = map[string]int32{
		"UNKNOWN": 0,
		"VIEW":    1,
		"CLICK":   2,
		"SHARE":   3,
		"COMMENT": 4,
	}
)

func (x ArticleAction) Enum() *ArticleAction {
	p := new(ArticleAction)
	*p = x
	return p
}

func (x ArticleAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ArticleAction) Descriptor() protoreflect.EnumDescriptor {
	return file_analytics_analytics_proto_enumTypes[0].Descriptor()
}

func (ArticleAction) Type() protoreflect.EnumType {
	return &file_analytics_analytics_proto_enumTypes[0]
}

func (x ArticleAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArticleAction.Descriptor instead.
func (ArticleAction) EnumDescriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

type GetListParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Limit         int32                  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset        int32                  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	StartedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	EndedAt       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=endedAt,proto3" json:"endedAt,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetListParams) Reset() {
	*x = GetListParams{}
	mi := &file_analytics_analytics_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetListParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListParams) ProtoMessage() {}

func (x *GetListParams) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListParams.ProtoReflect.Descriptor instead.
func (*GetListParams) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{0}
}

func (x *GetListParams) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetListParams) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetListParams) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *GetListParams) GetEndedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedAt
	}
	return nil
}

type IncrementArticleParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ArticleUuid   []byte                 `protobuf:"bytes,1,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	UserUuid      []byte                 `protobuf:"bytes,2,opt,name=user_uuid,json=userUuid,proto3" json:"user_uuid,omitempty"`
	Action        ArticleAction          `protobuf:"varint,3,opt,name=action,proto3,enum=analytics.ArticleAction" json:"action,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IncrementArticleParams) Reset() {
	*x = IncrementArticleParams{}
	mi := &file_analytics_analytics_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IncrementArticleParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncrementArticleParams) ProtoMessage() {}

func (x *IncrementArticleParams) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncrementArticleParams.ProtoReflect.Descriptor instead.
func (*IncrementArticleParams) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *IncrementArticleParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

func (x *IncrementArticleParams) GetUserUuid() []byte {
	if x != nil {
		return x.UserUuid
	}
	return nil
}

func (x *IncrementArticleParams) GetAction() ArticleAction {
	if x != nil {
		return x.Action
	}
	return ArticleAction_UNKNOWN
}

func (x *IncrementArticleParams) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type GetArticleCountsParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ArticleUuid   []byte                 `protobuf:"bytes,1,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetArticleCountsParams) Reset() {
	*x = GetArticleCountsParams{}
	mi := &file_analytics_analytics_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetArticleCountsParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetArticleCountsParams) ProtoMessage() {}

func (x *GetArticleCountsParams) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetArticleCountsParams.ProtoReflect.Descriptor instead.
func (*GetArticleCountsParams) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *GetArticleCountsParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

type ArticleCountsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Views         int32                  `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	Likes         int32                  `protobuf:"varint,2,opt,name=likes,proto3" json:"likes,omitempty"`
	Comments      int32                  `protobuf:"varint,3,opt,name=comments,proto3" json:"comments,omitempty"`
	Shares        int32                  `protobuf:"varint,4,opt,name=shares,proto3" json:"shares,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArticleCountsResponse) Reset() {
	*x = ArticleCountsResponse{}
	mi := &file_analytics_analytics_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArticleCountsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArticleCountsResponse) ProtoMessage() {}

func (x *ArticleCountsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_analytics_analytics_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArticleCountsResponse.ProtoReflect.Descriptor instead.
func (*ArticleCountsResponse) Descriptor() ([]byte, []int) {
	return file_analytics_analytics_proto_rawDescGZIP(), []int{3}
}

func (x *ArticleCountsResponse) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *ArticleCountsResponse) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *ArticleCountsResponse) GetComments() int32 {
	if x != nil {
		return x.Comments
	}
	return 0
}

func (x *ArticleCountsResponse) GetShares() int32 {
	if x != nil {
		return x.Shares
	}
	return 0
}

var File_analytics_analytics_proto protoreflect.FileDescriptor

var file_analytics_analytics_proto_rawDesc = string([]byte{
	0x0a, 0x19, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2f, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x34, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x41, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x16, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x75, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x55, 0x75, 0x69, 0x64,
	0x12, 0x30, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x3b, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c,
	0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x22, 0x77, 0x0a, 0x15, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x73, 0x2a, 0x49, 0x0a, 0x0d, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x56, 0x49, 0x45, 0x57, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x43, 0x4c,
	0x49, 0x43, 0x4b, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x48, 0x41, 0x52, 0x45, 0x10, 0x03,
	0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x04, 0x32, 0xb3, 0x01,
	0x0a, 0x09, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x4d, 0x0a, 0x10, 0x49,
	0x6e, 0x63, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x12,
	0x21, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x57, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x21,
	0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x1a, 0x20, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x78, 0x33, 0x61, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x3b, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_analytics_analytics_proto_rawDescOnce sync.Once
	file_analytics_analytics_proto_rawDescData []byte
)

func file_analytics_analytics_proto_rawDescGZIP() []byte {
	file_analytics_analytics_proto_rawDescOnce.Do(func() {
		file_analytics_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_analytics_analytics_proto_rawDesc), len(file_analytics_analytics_proto_rawDesc)))
	})
	return file_analytics_analytics_proto_rawDescData
}

var file_analytics_analytics_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_analytics_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_analytics_analytics_proto_goTypes = []any{
	(ArticleAction)(0),             // 0: analytics.ArticleAction
	(*GetListParams)(nil),          // 1: analytics.GetListParams
	(*IncrementArticleParams)(nil), // 2: analytics.IncrementArticleParams
	(*GetArticleCountsParams)(nil), // 3: analytics.GetArticleCountsParams
	(*ArticleCountsResponse)(nil),  // 4: analytics.ArticleCountsResponse
	(*timestamppb.Timestamp)(nil),  // 5: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),          // 6: google.protobuf.Empty
}
var file_analytics_analytics_proto_depIdxs = []int32{
	5, // 0: analytics.GetListParams.started_at:type_name -> google.protobuf.Timestamp
	5, // 1: analytics.GetListParams.endedAt:type_name -> google.protobuf.Timestamp
	0, // 2: analytics.IncrementArticleParams.action:type_name -> analytics.ArticleAction
	5, // 3: analytics.IncrementArticleParams.timestamp:type_name -> google.protobuf.Timestamp
	2, // 4: analytics.Analytics.IncrementArticle:input_type -> analytics.IncrementArticleParams
	3, // 5: analytics.Analytics.GetArticleCounts:input_type -> analytics.GetArticleCountsParams
	6, // 6: analytics.Analytics.IncrementArticle:output_type -> google.protobuf.Empty
	4, // 7: analytics.Analytics.GetArticleCounts:output_type -> analytics.ArticleCountsResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_analytics_analytics_proto_init() }
func file_analytics_analytics_proto_init() {
	if File_analytics_analytics_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_analytics_analytics_proto_rawDesc), len(file_analytics_analytics_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_analytics_analytics_proto_goTypes,
		DependencyIndexes: file_analytics_analytics_proto_depIdxs,
		EnumInfos:         file_analytics_analytics_proto_enumTypes,
		MessageInfos:      file_analytics_analytics_proto_msgTypes,
	}.Build()
	File_analytics_analytics_proto = out.File
	file_analytics_analytics_proto_goTypes = nil
	file_analytics_analytics_proto_depIdxs = nil
}
