// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.24.4
// source: socials/comments.proto

package socialsv1

import (
	common "github.com/x3a-tech/contract-go/common"
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

type CommentState int32

const (
	CommentState_COMMENT_STATE_UNSPECIFIED CommentState = 0
	CommentState_PUBLISHED                 CommentState = 1
	CommentState_BLOCKED                   CommentState = 2
)

// Enum value maps for CommentState.
var (
	CommentState_name = map[int32]string{
		0: "COMMENT_STATE_UNSPECIFIED",
		1: "PUBLISHED",
		2: "BLOCKED",
	}
	CommentState_value = map[string]int32{
		"COMMENT_STATE_UNSPECIFIED": 0,
		"PUBLISHED":                 1,
		"BLOCKED":                   2,
	}
)

func (x CommentState) Enum() *CommentState {
	p := new(CommentState)
	*p = x
	return p
}

func (x CommentState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommentState) Descriptor() protoreflect.EnumDescriptor {
	return file_socials_comments_proto_enumTypes[0].Descriptor()
}

func (CommentState) Type() protoreflect.EnumType {
	return &file_socials_comments_proto_enumTypes[0]
}

func (x CommentState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommentState.Descriptor instead.
func (CommentState) EnumDescriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{0}
}

type Comment struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId      *int64                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	Content       string                 `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	AuthorUuid    []byte                 `protobuf:"bytes,4,opt,name=author_uuid,json=authorUuid,proto3" json:"author_uuid,omitempty"`
	State         CommentState           `protobuf:"varint,5,opt,name=state,proto3,enum=socials.CommentState" json:"state,omitempty"`
	Reactions     map[string]int32       `protobuf:"bytes,6,rep,name=reactions,proto3" json:"reactions,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Reaction      string                 `protobuf:"bytes,7,opt,name=reaction,proto3" json:"reaction,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3,oneof" json:"updated_at,omitempty"`
	DeletedAt     *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=deleted_at,json=deletedAt,proto3,oneof" json:"deleted_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Comment) Reset() {
	*x = Comment{}
	mi := &file_socials_comments_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetAuthorUuid() []byte {
	if x != nil {
		return x.AuthorUuid
	}
	return nil
}

func (x *Comment) GetState() CommentState {
	if x != nil {
		return x.State
	}
	return CommentState_COMMENT_STATE_UNSPECIFIED
}

func (x *Comment) GetReactions() map[string]int32 {
	if x != nil {
		return x.Reactions
	}
	return nil
}

func (x *Comment) GetReaction() string {
	if x != nil {
		return x.Reaction
	}
	return ""
}

func (x *Comment) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Comment) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Comment) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

type GetCommentsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Comments      []*Comment             `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentsResponse) Reset() {
	*x = GetCommentsResponse{}
	mi := &file_socials_comments_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResponse) ProtoMessage() {}

func (x *GetCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsResponse) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{1}
}

func (x *GetCommentsResponse) GetComments() []*Comment {
	if x != nil {
		return x.Comments
	}
	return nil
}

type GetCommentsByArticleParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ArticleUuid   []byte                 `protobuf:"bytes,1,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	ParentId      *int64                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	ListParams    *common.ListParams     `protobuf:"bytes,3,opt,name=listParams,proto3" json:"listParams,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentsByArticleParams) Reset() {
	*x = GetCommentsByArticleParams{}
	mi := &file_socials_comments_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsByArticleParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByArticleParams) ProtoMessage() {}

func (x *GetCommentsByArticleParams) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByArticleParams.ProtoReflect.Descriptor instead.
func (*GetCommentsByArticleParams) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentsByArticleParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

func (x *GetCommentsByArticleParams) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *GetCommentsByArticleParams) GetListParams() *common.ListParams {
	if x != nil {
		return x.ListParams
	}
	return nil
}

type GetCommentsByArticleParamsInner struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	AccountUuid   []byte                      `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3,oneof" json:"account_uuid,omitempty"`
	Data          *GetCommentsByArticleParams `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentsByArticleParamsInner) Reset() {
	*x = GetCommentsByArticleParamsInner{}
	mi := &file_socials_comments_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsByArticleParamsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByArticleParamsInner) ProtoMessage() {}

func (x *GetCommentsByArticleParamsInner) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByArticleParamsInner.ProtoReflect.Descriptor instead.
func (*GetCommentsByArticleParamsInner) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentsByArticleParamsInner) GetAccountUuid() []byte {
	if x != nil {
		return x.AccountUuid
	}
	return nil
}

func (x *GetCommentsByArticleParamsInner) GetData() *GetCommentsByArticleParams {
	if x != nil {
		return x.Data
	}
	return nil
}

type UpsertCommentArticleParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ArticleUuid   []byte                 `protobuf:"bytes,2,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	ParentId      *int64                 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCommentArticleParams) Reset() {
	*x = UpsertCommentArticleParams{}
	mi := &file_socials_comments_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCommentArticleParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCommentArticleParams) ProtoMessage() {}

func (x *UpsertCommentArticleParams) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCommentArticleParams.ProtoReflect.Descriptor instead.
func (*UpsertCommentArticleParams) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertCommentArticleParams) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertCommentArticleParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

func (x *UpsertCommentArticleParams) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

type UpsertCommentArticleParamsInner struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	AccountUuid   []byte                      `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	Data          *UpsertCommentArticleParams `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCommentArticleParamsInner) Reset() {
	*x = UpsertCommentArticleParamsInner{}
	mi := &file_socials_comments_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCommentArticleParamsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCommentArticleParamsInner) ProtoMessage() {}

func (x *UpsertCommentArticleParamsInner) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCommentArticleParamsInner.ProtoReflect.Descriptor instead.
func (*UpsertCommentArticleParamsInner) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertCommentArticleParamsInner) GetAccountUuid() []byte {
	if x != nil {
		return x.AccountUuid
	}
	return nil
}

func (x *UpsertCommentArticleParamsInner) GetData() *UpsertCommentArticleParams {
	if x != nil {
		return x.Data
	}
	return nil
}

type RemoveCommentParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CommentId     int64                  `protobuf:"varint,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCommentParams) Reset() {
	*x = RemoveCommentParams{}
	mi := &file_socials_comments_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCommentParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentParams) ProtoMessage() {}

func (x *RemoveCommentParams) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentParams.ProtoReflect.Descriptor instead.
func (*RemoveCommentParams) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{6}
}

func (x *RemoveCommentParams) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type RemoveCommentParamsInner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccountUuid   []byte                 `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	Data          *RemoveCommentParams   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RemoveCommentParamsInner) Reset() {
	*x = RemoveCommentParamsInner{}
	mi := &file_socials_comments_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RemoveCommentParamsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCommentParamsInner) ProtoMessage() {}

func (x *RemoveCommentParamsInner) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCommentParamsInner.ProtoReflect.Descriptor instead.
func (*RemoveCommentParamsInner) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveCommentParamsInner) GetAccountUuid() []byte {
	if x != nil {
		return x.AccountUuid
	}
	return nil
}

func (x *RemoveCommentParamsInner) GetData() *RemoveCommentParams {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetCommentsByEntityParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ArticleUuid   []byte                 `protobuf:"bytes,1,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	ParentId      *int64                 `protobuf:"varint,2,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	ListParams    *common.ListParams     `protobuf:"bytes,3,opt,name=listParams,proto3" json:"listParams,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCommentsByEntityParams) Reset() {
	*x = GetCommentsByEntityParams{}
	mi := &file_socials_comments_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCommentsByEntityParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsByEntityParams) ProtoMessage() {}

func (x *GetCommentsByEntityParams) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsByEntityParams.ProtoReflect.Descriptor instead.
func (*GetCommentsByEntityParams) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{8}
}

func (x *GetCommentsByEntityParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

func (x *GetCommentsByEntityParams) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

func (x *GetCommentsByEntityParams) GetListParams() *common.ListParams {
	if x != nil {
		return x.ListParams
	}
	return nil
}

type UpsertCommentEntityParams struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Content       string                 `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	ArticleUuid   []byte                 `protobuf:"bytes,2,opt,name=article_uuid,json=articleUuid,proto3" json:"article_uuid,omitempty"`
	ParentId      *int64                 `protobuf:"varint,3,opt,name=parent_id,json=parentId,proto3,oneof" json:"parent_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCommentEntityParams) Reset() {
	*x = UpsertCommentEntityParams{}
	mi := &file_socials_comments_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCommentEntityParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCommentEntityParams) ProtoMessage() {}

func (x *UpsertCommentEntityParams) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCommentEntityParams.ProtoReflect.Descriptor instead.
func (*UpsertCommentEntityParams) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{9}
}

func (x *UpsertCommentEntityParams) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *UpsertCommentEntityParams) GetArticleUuid() []byte {
	if x != nil {
		return x.ArticleUuid
	}
	return nil
}

func (x *UpsertCommentEntityParams) GetParentId() int64 {
	if x != nil && x.ParentId != nil {
		return *x.ParentId
	}
	return 0
}

type UpsertCommentEntityParamsInner struct {
	state         protoimpl.MessageState      `protogen:"open.v1"`
	AccountUuid   []byte                      `protobuf:"bytes,1,opt,name=account_uuid,json=accountUuid,proto3" json:"account_uuid,omitempty"`
	Data          *UpsertCommentArticleParams `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpsertCommentEntityParamsInner) Reset() {
	*x = UpsertCommentEntityParamsInner{}
	mi := &file_socials_comments_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpsertCommentEntityParamsInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertCommentEntityParamsInner) ProtoMessage() {}

func (x *UpsertCommentEntityParamsInner) ProtoReflect() protoreflect.Message {
	mi := &file_socials_comments_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertCommentEntityParamsInner.ProtoReflect.Descriptor instead.
func (*UpsertCommentEntityParamsInner) Descriptor() ([]byte, []int) {
	return file_socials_comments_proto_rawDescGZIP(), []int{10}
}

func (x *UpsertCommentEntityParamsInner) GetAccountUuid() []byte {
	if x != nil {
		return x.AccountUuid
	}
	return nil
}

func (x *UpsertCommentEntityParamsInner) GetData() *UpsertCommentArticleParams {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_socials_comments_proto protoreflect.FileDescriptor

var file_socials_comments_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3,
	0x04, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73,
	0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x01, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x12, 0x3e, 0x0a, 0x0a, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x02, 0x52, 0x09, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x88, 0x01, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x52, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xa3, 0x01, 0x0a, 0x1a, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x41, 0x72, 0x74, 0x69, 0x63,
	0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a,
	0x0a, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22,
	0x93, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e,
	0x6e, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42,
	0x79, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x22, 0x89, 0x01, 0x0a, 0x1a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69,
	0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x22, 0x7d, 0x0a, 0x1f, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49,
	0x6e, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x2e,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x34, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x18, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x6e,
	0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa2, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70,
	0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x32, 0x0a, 0x0a, 0x6c, 0x69,
	0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x88, 0x01, 0x0a,
	0x19, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x75, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x55, 0x75, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x7c, 0x0a, 0x1e, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73, 0x6f, 0x63,
	0x69, 0x61, 0x6c, 0x73, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0x49, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x43, 0x4f, 0x4d, 0x4d, 0x45, 0x4e, 0x54,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02,
	0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78,
	0x33, 0x61, 0x2d, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2d, 0x67, 0x6f, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x73, 0x3b, 0x73, 0x6f, 0x63, 0x69,
	0x61, 0x6c, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_socials_comments_proto_rawDescOnce sync.Once
	file_socials_comments_proto_rawDescData []byte
)

func file_socials_comments_proto_rawDescGZIP() []byte {
	file_socials_comments_proto_rawDescOnce.Do(func() {
		file_socials_comments_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_socials_comments_proto_rawDesc), len(file_socials_comments_proto_rawDesc)))
	})
	return file_socials_comments_proto_rawDescData
}

var file_socials_comments_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_socials_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_socials_comments_proto_goTypes = []any{
	(CommentState)(0),                       // 0: socials.CommentState
	(*Comment)(nil),                         // 1: socials.Comment
	(*GetCommentsResponse)(nil),             // 2: socials.GetCommentsResponse
	(*GetCommentsByArticleParams)(nil),      // 3: socials.GetCommentsByArticleParams
	(*GetCommentsByArticleParamsInner)(nil), // 4: socials.GetCommentsByArticleParamsInner
	(*UpsertCommentArticleParams)(nil),      // 5: socials.UpsertCommentArticleParams
	(*UpsertCommentArticleParamsInner)(nil), // 6: socials.UpsertCommentArticleParamsInner
	(*RemoveCommentParams)(nil),             // 7: socials.RemoveCommentParams
	(*RemoveCommentParamsInner)(nil),        // 8: socials.RemoveCommentParamsInner
	(*GetCommentsByEntityParams)(nil),       // 9: socials.GetCommentsByEntityParams
	(*UpsertCommentEntityParams)(nil),       // 10: socials.UpsertCommentEntityParams
	(*UpsertCommentEntityParamsInner)(nil),  // 11: socials.UpsertCommentEntityParamsInner
	nil,                                     // 12: socials.Comment.ReactionsEntry
	(*timestamppb.Timestamp)(nil),           // 13: google.protobuf.Timestamp
	(*common.ListParams)(nil),               // 14: common.ListParams
}
var file_socials_comments_proto_depIdxs = []int32{
	0,  // 0: socials.Comment.state:type_name -> socials.CommentState
	12, // 1: socials.Comment.reactions:type_name -> socials.Comment.ReactionsEntry
	13, // 2: socials.Comment.created_at:type_name -> google.protobuf.Timestamp
	13, // 3: socials.Comment.updated_at:type_name -> google.protobuf.Timestamp
	13, // 4: socials.Comment.deleted_at:type_name -> google.protobuf.Timestamp
	1,  // 5: socials.GetCommentsResponse.comments:type_name -> socials.Comment
	14, // 6: socials.GetCommentsByArticleParams.listParams:type_name -> common.ListParams
	3,  // 7: socials.GetCommentsByArticleParamsInner.data:type_name -> socials.GetCommentsByArticleParams
	5,  // 8: socials.UpsertCommentArticleParamsInner.data:type_name -> socials.UpsertCommentArticleParams
	7,  // 9: socials.RemoveCommentParamsInner.data:type_name -> socials.RemoveCommentParams
	14, // 10: socials.GetCommentsByEntityParams.listParams:type_name -> common.ListParams
	5,  // 11: socials.UpsertCommentEntityParamsInner.data:type_name -> socials.UpsertCommentArticleParams
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_socials_comments_proto_init() }
func file_socials_comments_proto_init() {
	if File_socials_comments_proto != nil {
		return
	}
	file_socials_comments_proto_msgTypes[0].OneofWrappers = []any{}
	file_socials_comments_proto_msgTypes[2].OneofWrappers = []any{}
	file_socials_comments_proto_msgTypes[3].OneofWrappers = []any{}
	file_socials_comments_proto_msgTypes[4].OneofWrappers = []any{}
	file_socials_comments_proto_msgTypes[8].OneofWrappers = []any{}
	file_socials_comments_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_socials_comments_proto_rawDesc), len(file_socials_comments_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_socials_comments_proto_goTypes,
		DependencyIndexes: file_socials_comments_proto_depIdxs,
		EnumInfos:         file_socials_comments_proto_enumTypes,
		MessageInfos:      file_socials_comments_proto_msgTypes,
	}.Build()
	File_socials_comments_proto = out.File
	file_socials_comments_proto_goTypes = nil
	file_socials_comments_proto_depIdxs = nil
}
