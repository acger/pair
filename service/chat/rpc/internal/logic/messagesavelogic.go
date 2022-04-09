package logic

import (
	"pair/common/aerror"
	"pair/service/chat/model"
	"context"
	"github.com/jinzhu/copier"

	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageSaveLogic {
	return &MessageSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageSaveLogic) MessageSave(in *pb.MsgSaveReq) (*pb.Rsp, error) {
	msg := model.Chats{}
	copier.Copy(&msg, in)
	tx := l.svcCtx.DB.Create(&msg)

	if tx.Error != nil {
		return nil, aerror.ErrLog(tx.Error, in)
	}

	return &pb.Rsp{Code: 0}, nil
}
