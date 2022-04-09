package chatGroup

import (
	"pair/service/chat/rpc/chat"
	"context"
	"encoding/json"
	"strconv"

	"pair/service/chat/api/internal/svc"
	"pair/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatMessageSaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatMessageSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatMessageSaveLogic {
	return ChatMessageSaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatMessageSaveLogic) ChatMessageSave(req *types.ChatMessageSaveReq) (resp *types.Rsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	toUid, _ := strconv.ParseInt(req.ToUid, 10, 64)

	_, err = l.svcCtx.ChatRPC.MessageSave(l.ctx, &chat.MsgSaveReq{
		Uid:     uid,
		ToUid:   toUid,
		Message: req.Message,
		Status:  req.Status,
	})

	if err != nil {
		return nil, err
	}

	return &types.Rsp{Code: 0}, nil
}
