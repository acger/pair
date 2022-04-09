package logic

import (
	"pair/common/aerror"
	"pair/service/user/model"
	"pair/service/user/rpc/internal/svc"
	"pair/service/user/rpc/user"
	"pair/service/user/rpc/user/pb"
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *pb.UserListReq) (*pb.UserListRsp, error) {
	list, err := mr.MapReduce(func(source chan<- interface{}) {
		for _, uid := range in.Id {
			source <- uid
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		uid := item.(int64)
		user, err := l.svcCtx.UserModel.FindOne(l.ctx, uid)

		if err != nil {
			cancel(err)
		}

		writer.Write(user)
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {
		var list []*user.UserInfo

		for p := range pipe {
			var item user.UserInfo
			copier.Copy(&item, p.(*model.Users))
			list = append(list, &item)
		}

		writer.Write(list)
	})

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	return &pb.UserListRsp{
		Code: 0,
		User: list.([]*user.UserInfo),
	}, nil
}
