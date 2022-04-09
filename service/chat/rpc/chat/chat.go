// Code generated by goctl. DO NOT EDIT!
// Source: chat.proto

package chat

import (
	"context"

	"pair/service/chat/rpc/chat/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CHSaveReq      = pb.CHSaveReq
	ChatHistoryReq = pb.ChatHistoryReq
	ChatHistoryRsp = pb.ChatHistoryRsp
	ChatMessage    = pb.ChatMessage
	ChatNumberReq  = pb.ChatNumberReq
	ChatNumberRsp  = pb.ChatNumberRsp
	ChatUser       = pb.ChatUser
	MsgListReq     = pb.MsgListReq
	MsgListRsp     = pb.MsgListRsp
	MsgSaveReq     = pb.MsgSaveReq
	Rsp            = pb.Rsp

	Chat interface {
		MessageSave(ctx context.Context, in *MsgSaveReq, opts ...grpc.CallOption) (*Rsp, error)
		MessageList(ctx context.Context, in *MsgListReq, opts ...grpc.CallOption) (*MsgListRsp, error)
		ChatHistoryList(ctx context.Context, in *ChatHistoryReq, opts ...grpc.CallOption) (*ChatHistoryRsp, error)
		ChatNumber(ctx context.Context, in *ChatNumberReq, opts ...grpc.CallOption) (*ChatNumberRsp, error)
		ChatHistorySave(ctx context.Context, in *CHSaveReq, opts ...grpc.CallOption) (*Rsp, error)
	}

	defaultChat struct {
		cli zrpc.Client
	}
)

func NewChat(cli zrpc.Client) Chat {
	return &defaultChat{
		cli: cli,
	}
}

func (m *defaultChat) MessageSave(ctx context.Context, in *MsgSaveReq, opts ...grpc.CallOption) (*Rsp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.MessageSave(ctx, in, opts...)
}

func (m *defaultChat) MessageList(ctx context.Context, in *MsgListReq, opts ...grpc.CallOption) (*MsgListRsp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.MessageList(ctx, in, opts...)
}

func (m *defaultChat) ChatHistoryList(ctx context.Context, in *ChatHistoryReq, opts ...grpc.CallOption) (*ChatHistoryRsp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.ChatHistoryList(ctx, in, opts...)
}

func (m *defaultChat) ChatNumber(ctx context.Context, in *ChatNumberReq, opts ...grpc.CallOption) (*ChatNumberRsp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.ChatNumber(ctx, in, opts...)
}

func (m *defaultChat) ChatHistorySave(ctx context.Context, in *CHSaveReq, opts ...grpc.CallOption) (*Rsp, error) {
	client := pb.NewChatClient(m.cli.Conn())
	return client.ChatHistorySave(ctx, in, opts...)
}