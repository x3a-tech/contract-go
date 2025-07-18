// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: social/social.proto

package socialv1

import (
	context "context"
	common "github.com/x3a-tech/contract-go/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Socials_UpdateComment_FullMethodName        = "/social.Socials/UpdateComment"
	Socials_RemoveComment_FullMethodName        = "/social.Socials/RemoveComment"
	Socials_SetCommentReaction_FullMethodName   = "/social.Socials/SetCommentReaction"
	Socials_GetArticlesReactions_FullMethodName = "/social.Socials/GetArticlesReactions"
	Socials_GetCommentsByArticle_FullMethodName = "/social.Socials/GetCommentsByArticle"
	Socials_CreateArticleComment_FullMethodName = "/social.Socials/CreateArticleComment"
	Socials_SetArticleReaction_FullMethodName   = "/social.Socials/SetArticleReaction"
	Socials_GetCommentsByEntity_FullMethodName  = "/social.Socials/GetCommentsByEntity"
	Socials_CreateEntityComment_FullMethodName  = "/social.Socials/CreateEntityComment"
	Socials_GetCommentsByAccount_FullMethodName = "/social.Socials/GetCommentsByAccount"
)

// SocialsClient is the client API for Socials service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SocialsClient interface {
	UpdateComment(ctx context.Context, in *UpdateCommentParamsInner, opts ...grpc.CallOption) (*Comment, error)
	RemoveComment(ctx context.Context, in *RemoveCommentParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error)
	SetCommentReaction(ctx context.Context, in *SetCommentReactionParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error)
	GetArticlesReactions(ctx context.Context, in *GetArticlesReactionsParamsInner, opts ...grpc.CallOption) (*GetArticlesReactionsResponse, error)
	GetCommentsByArticle(ctx context.Context, in *GetCommentsByArticleParams, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	CreateArticleComment(ctx context.Context, in *CreateCommentArticleParamsInner, opts ...grpc.CallOption) (*Comment, error)
	SetArticleReaction(ctx context.Context, in *SetArticleReactionParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error)
	GetCommentsByEntity(ctx context.Context, in *GetCommentsByEntityParams, opts ...grpc.CallOption) (*GetCommentsResponse, error)
	CreateEntityComment(ctx context.Context, in *CreateCommentEntityParamsInner, opts ...grpc.CallOption) (*Comment, error)
	GetCommentsByAccount(ctx context.Context, in *GetCommentsByAccountParamsInner, opts ...grpc.CallOption) (*GetCommentsResponse, error)
}

type socialsClient struct {
	cc grpc.ClientConnInterface
}

func NewSocialsClient(cc grpc.ClientConnInterface) SocialsClient {
	return &socialsClient{cc}
}

func (c *socialsClient) UpdateComment(ctx context.Context, in *UpdateCommentParamsInner, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, Socials_UpdateComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) RemoveComment(ctx context.Context, in *RemoveCommentParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, Socials_RemoveComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) SetCommentReaction(ctx context.Context, in *SetCommentReactionParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, Socials_SetCommentReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) GetArticlesReactions(ctx context.Context, in *GetArticlesReactionsParamsInner, opts ...grpc.CallOption) (*GetArticlesReactionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesReactionsResponse)
	err := c.cc.Invoke(ctx, Socials_GetArticlesReactions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) GetCommentsByArticle(ctx context.Context, in *GetCommentsByArticleParams, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, Socials_GetCommentsByArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) CreateArticleComment(ctx context.Context, in *CreateCommentArticleParamsInner, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, Socials_CreateArticleComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) SetArticleReaction(ctx context.Context, in *SetArticleReactionParamsInner, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, Socials_SetArticleReaction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) GetCommentsByEntity(ctx context.Context, in *GetCommentsByEntityParams, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, Socials_GetCommentsByEntity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) CreateEntityComment(ctx context.Context, in *CreateCommentEntityParamsInner, opts ...grpc.CallOption) (*Comment, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Comment)
	err := c.cc.Invoke(ctx, Socials_CreateEntityComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *socialsClient) GetCommentsByAccount(ctx context.Context, in *GetCommentsByAccountParamsInner, opts ...grpc.CallOption) (*GetCommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCommentsResponse)
	err := c.cc.Invoke(ctx, Socials_GetCommentsByAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SocialsServer is the server API for Socials service.
// All implementations must embed UnimplementedSocialsServer
// for forward compatibility.
type SocialsServer interface {
	UpdateComment(context.Context, *UpdateCommentParamsInner) (*Comment, error)
	RemoveComment(context.Context, *RemoveCommentParamsInner) (*common.BoolStatus, error)
	SetCommentReaction(context.Context, *SetCommentReactionParamsInner) (*common.BoolStatus, error)
	GetArticlesReactions(context.Context, *GetArticlesReactionsParamsInner) (*GetArticlesReactionsResponse, error)
	GetCommentsByArticle(context.Context, *GetCommentsByArticleParams) (*GetCommentsResponse, error)
	CreateArticleComment(context.Context, *CreateCommentArticleParamsInner) (*Comment, error)
	SetArticleReaction(context.Context, *SetArticleReactionParamsInner) (*common.BoolStatus, error)
	GetCommentsByEntity(context.Context, *GetCommentsByEntityParams) (*GetCommentsResponse, error)
	CreateEntityComment(context.Context, *CreateCommentEntityParamsInner) (*Comment, error)
	GetCommentsByAccount(context.Context, *GetCommentsByAccountParamsInner) (*GetCommentsResponse, error)
	mustEmbedUnimplementedSocialsServer()
}

// UnimplementedSocialsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSocialsServer struct{}

func (UnimplementedSocialsServer) UpdateComment(context.Context, *UpdateCommentParamsInner) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedSocialsServer) RemoveComment(context.Context, *RemoveCommentParamsInner) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveComment not implemented")
}
func (UnimplementedSocialsServer) SetCommentReaction(context.Context, *SetCommentReactionParamsInner) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCommentReaction not implemented")
}
func (UnimplementedSocialsServer) GetArticlesReactions(context.Context, *GetArticlesReactionsParamsInner) (*GetArticlesReactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesReactions not implemented")
}
func (UnimplementedSocialsServer) GetCommentsByArticle(context.Context, *GetCommentsByArticleParams) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByArticle not implemented")
}
func (UnimplementedSocialsServer) CreateArticleComment(context.Context, *CreateCommentArticleParamsInner) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArticleComment not implemented")
}
func (UnimplementedSocialsServer) SetArticleReaction(context.Context, *SetArticleReactionParamsInner) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetArticleReaction not implemented")
}
func (UnimplementedSocialsServer) GetCommentsByEntity(context.Context, *GetCommentsByEntityParams) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByEntity not implemented")
}
func (UnimplementedSocialsServer) CreateEntityComment(context.Context, *CreateCommentEntityParamsInner) (*Comment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEntityComment not implemented")
}
func (UnimplementedSocialsServer) GetCommentsByAccount(context.Context, *GetCommentsByAccountParamsInner) (*GetCommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommentsByAccount not implemented")
}
func (UnimplementedSocialsServer) mustEmbedUnimplementedSocialsServer() {}
func (UnimplementedSocialsServer) testEmbeddedByValue()                 {}

// UnsafeSocialsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SocialsServer will
// result in compilation errors.
type UnsafeSocialsServer interface {
	mustEmbedUnimplementedSocialsServer()
}

func RegisterSocialsServer(s grpc.ServiceRegistrar, srv SocialsServer) {
	// If the following call pancis, it indicates UnimplementedSocialsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Socials_ServiceDesc, srv)
}

func _Socials_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_UpdateComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).UpdateComment(ctx, req.(*UpdateCommentParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_RemoveComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCommentParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).RemoveComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_RemoveComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).RemoveComment(ctx, req.(*RemoveCommentParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_SetCommentReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCommentReactionParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).SetCommentReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_SetCommentReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).SetCommentReaction(ctx, req.(*SetCommentReactionParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_GetArticlesReactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesReactionsParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).GetArticlesReactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_GetArticlesReactions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).GetArticlesReactions(ctx, req.(*GetArticlesReactionsParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_GetCommentsByArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).GetCommentsByArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_GetCommentsByArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).GetCommentsByArticle(ctx, req.(*GetCommentsByArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_CreateArticleComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentArticleParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).CreateArticleComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_CreateArticleComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).CreateArticleComment(ctx, req.(*CreateCommentArticleParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_SetArticleReaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetArticleReactionParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).SetArticleReaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_SetArticleReaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).SetArticleReaction(ctx, req.(*SetArticleReactionParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_GetCommentsByEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByEntityParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).GetCommentsByEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_GetCommentsByEntity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).GetCommentsByEntity(ctx, req.(*GetCommentsByEntityParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_CreateEntityComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentEntityParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).CreateEntityComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_CreateEntityComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).CreateEntityComment(ctx, req.(*CreateCommentEntityParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

func _Socials_GetCommentsByAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCommentsByAccountParamsInner)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SocialsServer).GetCommentsByAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Socials_GetCommentsByAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SocialsServer).GetCommentsByAccount(ctx, req.(*GetCommentsByAccountParamsInner))
	}
	return interceptor(ctx, in, info, handler)
}

// Socials_ServiceDesc is the grpc.ServiceDesc for Socials service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Socials_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "social.Socials",
	HandlerType: (*SocialsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateComment",
			Handler:    _Socials_UpdateComment_Handler,
		},
		{
			MethodName: "RemoveComment",
			Handler:    _Socials_RemoveComment_Handler,
		},
		{
			MethodName: "SetCommentReaction",
			Handler:    _Socials_SetCommentReaction_Handler,
		},
		{
			MethodName: "GetArticlesReactions",
			Handler:    _Socials_GetArticlesReactions_Handler,
		},
		{
			MethodName: "GetCommentsByArticle",
			Handler:    _Socials_GetCommentsByArticle_Handler,
		},
		{
			MethodName: "CreateArticleComment",
			Handler:    _Socials_CreateArticleComment_Handler,
		},
		{
			MethodName: "SetArticleReaction",
			Handler:    _Socials_SetArticleReaction_Handler,
		},
		{
			MethodName: "GetCommentsByEntity",
			Handler:    _Socials_GetCommentsByEntity_Handler,
		},
		{
			MethodName: "CreateEntityComment",
			Handler:    _Socials_CreateEntityComment_Handler,
		},
		{
			MethodName: "GetCommentsByAccount",
			Handler:    _Socials_GetCommentsByAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "social/social.proto",
}
