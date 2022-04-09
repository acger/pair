package profile

import (
	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"
	"pair/service/user/rpc/user"
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProfileLogic {
	return ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile() (resp *types.ProfileRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	u, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoReq{Id: uid})

	if err != nil {
		return nil, err
	}

	return &types.ProfileRsp{
		Code: 0,
		User: types.User{
			Id:      strconv.FormatInt(u.User.Id, 10),
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
			Account: u.User.Account,
		},
	}, nil
}
