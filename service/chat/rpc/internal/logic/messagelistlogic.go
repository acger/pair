package logic

import (
	"pair/common/aerror"
	"pair/service/chat/model"
	"pair/service/chat/rpc/chat"
	"context"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageListLogic {
	return &MessageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MessageListLogic) MessageList(in *pb.MsgListReq) (*pb.MsgListRsp, error) {
	var result *pb.MsgListRsp

	err := l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		//更新已读状态
		tx.Model(&model.Chats{}).Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid).Update("status", true)

		var list []*model.Chats

		tx = tx.Where(
			tx.Where("uid = ?", in.Uid).Where("to_uid = ?", in.ToUid),
		).Or(
			tx.Where("uid = ?", in.ToUid).Where("to_uid = ?", in.Uid),
		)

		/* 聊天记录分页尚有问题，暂时屏蔽
		if in.PageSize == 0 {
			in.PageSize = 30
		}

		if in.Page == 0 {
			in.Page = 1
		}

		offset := (in.Page - 1) * in.PageSize
		tx.Limit(int(in.PageSize)).Offset(int(offset)).Order("id desc").Find(&list)
		*/

		tx.Order("create_time").Find(&list)

		var total int64
		tx.Model(model.Chats{}).Count(&total)

		result = &pb.MsgListRsp{
			Code:     0,
			Total:    total,
			PageSize: in.PageSize,
			Page:     in.Page,
			Msg:      make([]*pb.ChatMessage, len(list)),
		}

		for i, item := range list {
			tmp := chat.ChatMessage{}
			copier.Copy(&tmp, &item)
			result.Msg[i] = &tmp
		}

		return nil
	})

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	return result, nil
}
