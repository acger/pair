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

type ChatMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) ChatMessageLogic {
	return ChatMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatMessageLogic) ChatMessage(req *types.ChatMessageReq) (resp *types.ChatMessageRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	toUid, _ := strconv.ParseInt(req.ToUid, 10, 64)

	r, _ := l.svcCtx.ChatRPC.MessageList(l.ctx, &chat.MsgListReq{
		Uid:      uid,
		ToUid:    toUid,
		Page:     req.Page,
		PageSize: req.PageSize,
	})

	msg := make([]*types.ChatMessage, len(r.Msg))

	for i, m := range r.Msg {
		item := types.ChatMessage{}
		copier.Copy(&item, &m)
		item.Id = strconv.FormatInt(m.Id, 10)
		item.Uid = strconv.FormatInt(m.Uid, 10)
		item.ToUid = strconv.FormatInt(m.ToUid, 10)
		msg[i] = &item
	}

	return &types.ChatMessageRsp{Chat: msg}, nil
}
