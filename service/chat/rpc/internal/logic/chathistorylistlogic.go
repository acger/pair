package logic

import (
	"pair/service/chat/rpc/chat/pb"
	"pair/service/chat/rpc/internal/svc"
	"pair/service/user/rpc/user"
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ChatHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChatHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryListLogic {
	return &ChatHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChatHistoryListLogic) ChatHistoryList(in *pb.ChatHistoryReq) (*pb.ChatHistoryRsp, error) {
	type ChatHistoryMessage struct {
		Uid    int64
		Status bool
	}

	var chatHistory []*ChatHistoryMessage

	db := l.svcCtx.DB
	query := db.Table("chat_histories a")
	query = query.Joins("left join chats as b on b.uid = a.to_uid and b.to_uid = ?", in.Id)
	query = query.Where("a.uid = ?", in.Id)
	query = query.Group("a.to_uid")
	query = query.Order("status, max(b.create_time) desc")
	query = query.Select("a.to_uid as uid, IFNULL(MIN(b.status), true) as status")

	result := query.Find(&chatHistory)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return &pb.ChatHistoryRsp{}, nil
	}

	var uidList []int64
	for _, e := range chatHistory {
		uidList = append(uidList, e.Uid)
	}

	//获取用户信息
	userListRsp, _ := l.svcCtx.UserRPC.UserList(l.ctx, &user.UserListReq{Id: uidList})
	userMap := make(map[int64]*user.UserInfo)

	for _, u := range userListRsp.User {
		userMap[u.Id] = u
	}

	var chatUserList []*pb.ChatUser

	for _, c := range chatHistory {
		tmp := &pb.ChatUser{
			Id:     c.Uid,
			Status: c.Status,
		}

		copier.Copy(tmp, userMap[c.Uid])
		chatUserList = append(chatUserList, tmp)
	}

	return &pb.ChatHistoryRsp{Code: 0, User: chatUserList}, nil
}
