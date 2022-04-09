package chatGroup

import (
	"pair/service/chat/rpc/chat"
	"context"
	"encoding/json"

	"pair/service/chat/api/internal/svc"
	"pair/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryNumberLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryNumberLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistoryNumberLogic {
	return ChatHistoryNumberLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryNumberLogic) ChatHistoryNumber() (resp *types.ChatNumberRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, _ := l.svcCtx.ChatRPC.ChatNumber(l.ctx, &chat.ChatNumberReq{
		Id: uid,
	})

	return &types.ChatNumberRsp{Number: r.Number}, nil
}
