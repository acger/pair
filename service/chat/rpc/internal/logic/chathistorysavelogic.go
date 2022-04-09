package logic

import (
	"pair/common/aerror"
	"pair/service/chat/model"
	"context"
	"gorm.io/gorm"

	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistorySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatHistorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistorySaveLogic {
	return &ChatHistorySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatHistorySaveLogic) ChatHistorySave(in *pb.CHSaveReq) (*pb.Rsp, error) {
	db := l.svcCtx.DB
	err := db.Transaction(func(tx *gorm.DB) error {
		var fromHistory model.ChatHistories
		tx.Model(&model.ChatHistories{}).Where("uid = ?", in.Uid).Where("to_uid = ?", in.ToUid).Find(&fromHistory)
		if fromHistory.Id == 0 {
			tx.Model(&model.ChatHistories{}).Create(&model.ChatHistories{
				Uid:   in.Uid,
				ToUid: in.ToUid,
			})
		}

		var toHistory model.ChatHistories
		tx.Model(&model.ChatHistories{}).Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid).Find(&toHistory)
		if toHistory.Id == 0 {
			tx.Model(&model.ChatHistories{}).Create(&model.ChatHistories{
				Uid:   in.ToUid,
				ToUid: in.Uid,
			})
		}

		return nil
	})

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	return &pb.Rsp{Code: 0}, nil
}
