package chatGroup

import (
	"pair/service/chat/rpc/chat"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"strconv"

	"pair/service/chat/api/internal/svc"
	"pair/service/chat/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatHistoryLogic {
	return ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatHistoryLogic) ChatHistory() (resp *types.ChatHistoryRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, _ := l.svcCtx.ChatRPC.ChatHistoryList(l.ctx, &chat.ChatHistoryReq{
		Id: uid,
	})

	user := make([]*types.User, len(r.User))

	for i, u := range r.User {
		item := types.User{}
		copier.Copy(&item, &u)
		item.Id = strconv.FormatInt(u.Id, 10)
		user[i] = &item
	}

	return &types.ChatHistoryRsp{User: user}, nil
}
