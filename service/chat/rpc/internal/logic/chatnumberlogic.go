package logic

import (
	"pair/common/aerror"
	"pair/service/chat/model"
	"context"

	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatNumberLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatNumberLogic {
	return &ChatNumberLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatNumberLogic) ChatNumber(in *pb.ChatNumberReq) (*pb.ChatNumberRsp, error) {
	var num int64
	tx := l.svcCtx.DB.Model(&model.ChatHistories{})
	tx.Where("uid = ?", in.Id).Count(&num)

	if tx.Error != nil {
		return nil, aerror.ErrLog(tx.Error, in)
	}

	return &pb.ChatNumberRsp{Code: 0, Number: num}, nil
}
