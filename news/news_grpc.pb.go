// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: news/news.proto

package newsv1

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
	News_GetArticlesByFilters_FullMethodName    = "/news.News/GetArticlesByFilters"
	News_GetArticlesBySources_FullMethodName    = "/news.News/GetArticlesBySources"
	News_GetArticlesByCategories_FullMethodName = "/news.News/GetArticlesByCategories"
	News_GetArticlesByUuids_FullMethodName      = "/news.News/GetArticlesByUuids"
	News_GetArticlesByQuery_FullMethodName      = "/news.News/GetArticlesByQuery"
	News_GetArticle_FullMethodName              = "/news.News/GetArticle"
	News_GetArticleContent_FullMethodName       = "/news.News/GetArticleContent"
	News_GetArticleSummary_FullMethodName       = "/news.News/GetArticleSummary"
	News_CheckExistArticle_FullMethodName       = "/news.News/CheckExistArticle"
	News_GetFilters_FullMethodName              = "/news.News/GetFilters"
	News_GetFilter_FullMethodName               = "/news.News/GetFilter"
	News_CreateFilter_FullMethodName            = "/news.News/CreateFilter"
	News_UpdateFilterMeta_FullMethodName        = "/news.News/UpdateFilterMeta"
	News_UpdateFilterBody_FullMethodName        = "/news.News/UpdateFilterBody"
	News_RemoveFilter_FullMethodName            = "/news.News/RemoveFilter"
	News_CopyFilter_FullMethodName              = "/news.News/CopyFilter"
	News_GetCategories_FullMethodName           = "/news.News/GetCategories"
	News_GetSourcesByUuids_FullMethodName       = "/news.News/GetSourcesByUuids"
	News_GetSourcesByQuery_FullMethodName       = "/news.News/GetSourcesByQuery"
)

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsClient interface {
	GetArticlesByFilters(ctx context.Context, in *GetArticlesByFiltersParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticlesBySources(ctx context.Context, in *GetArticlesBySourcesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticlesByCategories(ctx context.Context, in *GetArticlesByCategoriesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticlesByUuids(ctx context.Context, in *GetArticlesByUuidsParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticlesByQuery(ctx context.Context, in *GetArticlesByQueryParams, opts ...grpc.CallOption) (*GetArticlesResponse, error)
	GetArticle(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*Article, error)
	GetArticleContent(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleContentResponse, error)
	GetArticleSummary(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleSummaryResponse, error)
	CheckExistArticle(ctx context.Context, in *CheckExistsArticleParams, opts ...grpc.CallOption) (*common.BoolStatus, error)
	GetFilters(ctx context.Context, in *GetFiltersParams, opts ...grpc.CallOption) (*GetFiltersResponse, error)
	GetFilter(ctx context.Context, in *GetFilterParams, opts ...grpc.CallOption) (*Filter, error)
	CreateFilter(ctx context.Context, in *CreateFilterParams, opts ...grpc.CallOption) (*Filter, error)
	UpdateFilterMeta(ctx context.Context, in *UpdateFilterMetaParams, opts ...grpc.CallOption) (*common.BoolStatus, error)
	UpdateFilterBody(ctx context.Context, in *UpdateFilterBodyParams, opts ...grpc.CallOption) (*common.BoolStatus, error)
	RemoveFilter(ctx context.Context, in *RemoveFilterParams, opts ...grpc.CallOption) (*common.BoolStatus, error)
	CopyFilter(ctx context.Context, in *CopyFilterParams, opts ...grpc.CallOption) (*Filter, error)
	GetCategories(ctx context.Context, in *GetCategoriesParams, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetSourcesByUuids(ctx context.Context, in *GetSourcesParams, opts ...grpc.CallOption) (*GetSourcesResponse, error)
	GetSourcesByQuery(ctx context.Context, in *SearchSourcesParams, opts ...grpc.CallOption) (*SearchSourcesResponse, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) GetArticlesByFilters(ctx context.Context, in *GetArticlesByFiltersParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticlesByFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticlesBySources(ctx context.Context, in *GetArticlesBySourcesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticlesBySources_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticlesByCategories(ctx context.Context, in *GetArticlesByCategoriesParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticlesByCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticlesByUuids(ctx context.Context, in *GetArticlesByUuidsParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticlesByUuids_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticlesByQuery(ctx context.Context, in *GetArticlesByQueryParams, opts ...grpc.CallOption) (*GetArticlesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticlesResponse)
	err := c.cc.Invoke(ctx, News_GetArticlesByQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticle(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*Article, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Article)
	err := c.cc.Invoke(ctx, News_GetArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticleContent(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleContentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleContentResponse)
	err := c.cc.Invoke(ctx, News_GetArticleContent_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetArticleSummary(ctx context.Context, in *GetArticleParams, opts ...grpc.CallOption) (*GetArticleSummaryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetArticleSummaryResponse)
	err := c.cc.Invoke(ctx, News_GetArticleSummary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) CheckExistArticle(ctx context.Context, in *CheckExistsArticleParams, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, News_CheckExistArticle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetFilters(ctx context.Context, in *GetFiltersParams, opts ...grpc.CallOption) (*GetFiltersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFiltersResponse)
	err := c.cc.Invoke(ctx, News_GetFilters_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetFilter(ctx context.Context, in *GetFilterParams, opts ...grpc.CallOption) (*Filter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Filter)
	err := c.cc.Invoke(ctx, News_GetFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) CreateFilter(ctx context.Context, in *CreateFilterParams, opts ...grpc.CallOption) (*Filter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Filter)
	err := c.cc.Invoke(ctx, News_CreateFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) UpdateFilterMeta(ctx context.Context, in *UpdateFilterMetaParams, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, News_UpdateFilterMeta_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) UpdateFilterBody(ctx context.Context, in *UpdateFilterBodyParams, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, News_UpdateFilterBody_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) RemoveFilter(ctx context.Context, in *RemoveFilterParams, opts ...grpc.CallOption) (*common.BoolStatus, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(common.BoolStatus)
	err := c.cc.Invoke(ctx, News_RemoveFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) CopyFilter(ctx context.Context, in *CopyFilterParams, opts ...grpc.CallOption) (*Filter, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Filter)
	err := c.cc.Invoke(ctx, News_CopyFilter_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetCategories(ctx context.Context, in *GetCategoriesParams, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, News_GetCategories_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetSourcesByUuids(ctx context.Context, in *GetSourcesParams, opts ...grpc.CallOption) (*GetSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSourcesResponse)
	err := c.cc.Invoke(ctx, News_GetSourcesByUuids_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) GetSourcesByQuery(ctx context.Context, in *SearchSourcesParams, opts ...grpc.CallOption) (*SearchSourcesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchSourcesResponse)
	err := c.cc.Invoke(ctx, News_GetSourcesByQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServer is the server API for News service.
// All implementations must embed UnimplementedNewsServer
// for forward compatibility.
type NewsServer interface {
	GetArticlesByFilters(context.Context, *GetArticlesByFiltersParams) (*GetArticlesResponse, error)
	GetArticlesBySources(context.Context, *GetArticlesBySourcesParams) (*GetArticlesResponse, error)
	GetArticlesByCategories(context.Context, *GetArticlesByCategoriesParams) (*GetArticlesResponse, error)
	GetArticlesByUuids(context.Context, *GetArticlesByUuidsParams) (*GetArticlesResponse, error)
	GetArticlesByQuery(context.Context, *GetArticlesByQueryParams) (*GetArticlesResponse, error)
	GetArticle(context.Context, *GetArticleParams) (*Article, error)
	GetArticleContent(context.Context, *GetArticleParams) (*GetArticleContentResponse, error)
	GetArticleSummary(context.Context, *GetArticleParams) (*GetArticleSummaryResponse, error)
	CheckExistArticle(context.Context, *CheckExistsArticleParams) (*common.BoolStatus, error)
	GetFilters(context.Context, *GetFiltersParams) (*GetFiltersResponse, error)
	GetFilter(context.Context, *GetFilterParams) (*Filter, error)
	CreateFilter(context.Context, *CreateFilterParams) (*Filter, error)
	UpdateFilterMeta(context.Context, *UpdateFilterMetaParams) (*common.BoolStatus, error)
	UpdateFilterBody(context.Context, *UpdateFilterBodyParams) (*common.BoolStatus, error)
	RemoveFilter(context.Context, *RemoveFilterParams) (*common.BoolStatus, error)
	CopyFilter(context.Context, *CopyFilterParams) (*Filter, error)
	GetCategories(context.Context, *GetCategoriesParams) (*GetCategoriesResponse, error)
	GetSourcesByUuids(context.Context, *GetSourcesParams) (*GetSourcesResponse, error)
	GetSourcesByQuery(context.Context, *SearchSourcesParams) (*SearchSourcesResponse, error)
	mustEmbedUnimplementedNewsServer()
}

// UnimplementedNewsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNewsServer struct{}

func (UnimplementedNewsServer) GetArticlesByFilters(context.Context, *GetArticlesByFiltersParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByFilters not implemented")
}
func (UnimplementedNewsServer) GetArticlesBySources(context.Context, *GetArticlesBySourcesParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesBySources not implemented")
}
func (UnimplementedNewsServer) GetArticlesByCategories(context.Context, *GetArticlesByCategoriesParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByCategories not implemented")
}
func (UnimplementedNewsServer) GetArticlesByUuids(context.Context, *GetArticlesByUuidsParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByUuids not implemented")
}
func (UnimplementedNewsServer) GetArticlesByQuery(context.Context, *GetArticlesByQueryParams) (*GetArticlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticlesByQuery not implemented")
}
func (UnimplementedNewsServer) GetArticle(context.Context, *GetArticleParams) (*Article, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticle not implemented")
}
func (UnimplementedNewsServer) GetArticleContent(context.Context, *GetArticleParams) (*GetArticleContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleContent not implemented")
}
func (UnimplementedNewsServer) GetArticleSummary(context.Context, *GetArticleParams) (*GetArticleSummaryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArticleSummary not implemented")
}
func (UnimplementedNewsServer) CheckExistArticle(context.Context, *CheckExistsArticleParams) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistArticle not implemented")
}
func (UnimplementedNewsServer) GetFilters(context.Context, *GetFiltersParams) (*GetFiltersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilters not implemented")
}
func (UnimplementedNewsServer) GetFilter(context.Context, *GetFilterParams) (*Filter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilter not implemented")
}
func (UnimplementedNewsServer) CreateFilter(context.Context, *CreateFilterParams) (*Filter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFilter not implemented")
}
func (UnimplementedNewsServer) UpdateFilterMeta(context.Context, *UpdateFilterMetaParams) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilterMeta not implemented")
}
func (UnimplementedNewsServer) UpdateFilterBody(context.Context, *UpdateFilterBodyParams) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFilterBody not implemented")
}
func (UnimplementedNewsServer) RemoveFilter(context.Context, *RemoveFilterParams) (*common.BoolStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFilter not implemented")
}
func (UnimplementedNewsServer) CopyFilter(context.Context, *CopyFilterParams) (*Filter, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyFilter not implemented")
}
func (UnimplementedNewsServer) GetCategories(context.Context, *GetCategoriesParams) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedNewsServer) GetSourcesByUuids(context.Context, *GetSourcesParams) (*GetSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourcesByUuids not implemented")
}
func (UnimplementedNewsServer) GetSourcesByQuery(context.Context, *SearchSourcesParams) (*SearchSourcesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSourcesByQuery not implemented")
}
func (UnimplementedNewsServer) mustEmbedUnimplementedNewsServer() {}
func (UnimplementedNewsServer) testEmbeddedByValue()              {}

// UnsafeNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServer will
// result in compilation errors.
type UnsafeNewsServer interface {
	mustEmbedUnimplementedNewsServer()
}

func RegisterNewsServer(s grpc.ServiceRegistrar, srv NewsServer) {
	// If the following call pancis, it indicates UnimplementedNewsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&News_ServiceDesc, srv)
}

func _News_GetArticlesByFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByFiltersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticlesByFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticlesByFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticlesByFilters(ctx, req.(*GetArticlesByFiltersParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticlesBySources_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesBySourcesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticlesBySources(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticlesBySources_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticlesBySources(ctx, req.(*GetArticlesBySourcesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticlesByCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByCategoriesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticlesByCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticlesByCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticlesByCategories(ctx, req.(*GetArticlesByCategoriesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticlesByUuids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByUuidsParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticlesByUuids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticlesByUuids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticlesByUuids(ctx, req.(*GetArticlesByUuidsParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticlesByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticlesByQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticlesByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticlesByQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticlesByQuery(ctx, req.(*GetArticlesByQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticle(ctx, req.(*GetArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticleContent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticleContent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticleContent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticleContent(ctx, req.(*GetArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetArticleSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetArticleSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetArticleSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetArticleSummary(ctx, req.(*GetArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_CheckExistArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExistsArticleParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).CheckExistArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_CheckExistArticle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).CheckExistArticle(ctx, req.(*CheckExistsArticleParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetFilters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFiltersParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetFilters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetFilters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetFilters(ctx, req.(*GetFiltersParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetFilter(ctx, req.(*GetFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_CreateFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).CreateFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_CreateFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).CreateFilter(ctx, req.(*CreateFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_UpdateFilterMeta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilterMetaParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).UpdateFilterMeta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_UpdateFilterMeta_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).UpdateFilterMeta(ctx, req.(*UpdateFilterMetaParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_UpdateFilterBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFilterBodyParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).UpdateFilterBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_UpdateFilterBody_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).UpdateFilterBody(ctx, req.(*UpdateFilterBodyParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_RemoveFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).RemoveFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_RemoveFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).RemoveFilter(ctx, req.(*RemoveFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_CopyFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CopyFilterParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).CopyFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_CopyFilter_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).CopyFilter(ctx, req.(*CopyFilterParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetCategories(ctx, req.(*GetCategoriesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetSourcesByUuids_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourcesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetSourcesByUuids(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetSourcesByUuids_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetSourcesByUuids(ctx, req.(*GetSourcesParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_GetSourcesByQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchSourcesParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).GetSourcesByQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_GetSourcesByQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).GetSourcesByQuery(ctx, req.(*SearchSourcesParams))
	}
	return interceptor(ctx, in, info, handler)
}

// News_ServiceDesc is the grpc.ServiceDesc for News service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var News_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "news.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArticlesByFilters",
			Handler:    _News_GetArticlesByFilters_Handler,
		},
		{
			MethodName: "GetArticlesBySources",
			Handler:    _News_GetArticlesBySources_Handler,
		},
		{
			MethodName: "GetArticlesByCategories",
			Handler:    _News_GetArticlesByCategories_Handler,
		},
		{
			MethodName: "GetArticlesByUuids",
			Handler:    _News_GetArticlesByUuids_Handler,
		},
		{
			MethodName: "GetArticlesByQuery",
			Handler:    _News_GetArticlesByQuery_Handler,
		},
		{
			MethodName: "GetArticle",
			Handler:    _News_GetArticle_Handler,
		},
		{
			MethodName: "GetArticleContent",
			Handler:    _News_GetArticleContent_Handler,
		},
		{
			MethodName: "GetArticleSummary",
			Handler:    _News_GetArticleSummary_Handler,
		},
		{
			MethodName: "CheckExistArticle",
			Handler:    _News_CheckExistArticle_Handler,
		},
		{
			MethodName: "GetFilters",
			Handler:    _News_GetFilters_Handler,
		},
		{
			MethodName: "GetFilter",
			Handler:    _News_GetFilter_Handler,
		},
		{
			MethodName: "CreateFilter",
			Handler:    _News_CreateFilter_Handler,
		},
		{
			MethodName: "UpdateFilterMeta",
			Handler:    _News_UpdateFilterMeta_Handler,
		},
		{
			MethodName: "UpdateFilterBody",
			Handler:    _News_UpdateFilterBody_Handler,
		},
		{
			MethodName: "RemoveFilter",
			Handler:    _News_RemoveFilter_Handler,
		},
		{
			MethodName: "CopyFilter",
			Handler:    _News_CopyFilter_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _News_GetCategories_Handler,
		},
		{
			MethodName: "GetSourcesByUuids",
			Handler:    _News_GetSourcesByUuids_Handler,
		},
		{
			MethodName: "GetSourcesByQuery",
			Handler:    _News_GetSourcesByQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news/news.proto",
}
