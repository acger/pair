// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package user

import (
	"context"

	"pair/service/user/rpc/user/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Response      = pb.Response
	UserAddReq    = pb.UserAddReq
	UserAddRsp    = pb.UserAddRsp
	UserDetail    = pb.UserDetail
	UserInfo      = pb.UserInfo
	UserInfoReq   = pb.UserInfoReq
	UserInfoRsp   = pb.UserInfoRsp
	UserListReq   = pb.UserListReq
	UserListRsp   = pb.UserListRsp
	UserUpdateReq = pb.UserUpdateReq

	User interface {
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddRsp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error)
		UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRsp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRsp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultUser) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*Response, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultUser) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultUser) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListRsp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}