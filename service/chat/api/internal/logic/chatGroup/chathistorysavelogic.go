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

type ChatHistorySaveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistorySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistorySaveLogic {
	return ChatHistorySaveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistorySaveLogic) ChatHistorySave(req *types.ChatHistorySaveReq) (resp *types.Rsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	toUid, _ := strconv.ParseInt(req.ToUid, 10, 64)

	_, err = l.svcCtx.ChatRPC.ChatHistorySave(l.ctx, &chat.CHSaveReq{
		Uid:   uid,
		ToUid: toUid,
	})

	if err != nil {
		return nil, err
	}

	return &types.Rsp{Code: 0}, nil
}
