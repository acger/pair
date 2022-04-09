package profile

import (
	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"
	"pair/service/user/rpc/user"
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) EditProfileLogic {
	return EditProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditProfileLogic) EditProfile(req *types.EditReq) (resp *types.EditRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	_, err = l.svcCtx.UserRPC.UserUpdate(l.ctx, &user.UserUpdateReq{
		Id:       uid,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Password: req.Password,
	})

	if err != nil {
		return nil, err
	}

	return &types.EditRsp{Code: 0}, nil
}
