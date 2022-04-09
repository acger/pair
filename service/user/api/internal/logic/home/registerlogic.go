package home

import (
	"pair/common/aauth"
	"pair/common/aerror"
	"pair/service/user/rpc/user/pb"
	"context"
	"github.com/go-playground/validator/v10"
	"strconv"
	"time"

	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRsp, err error) {
	validate := validator.New()
	err = validate.Struct(req)
	if err != nil {
		return nil, aerror.Err(aerror.ErrCodeParamError)
	}

	r, err := l.svcCtx.UserRPC.UserAdd(l.ctx, &pb.UserAddReq{
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
	})

	if err != nil {
		return nil, err
	}

	u, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &pb.UserInfoReq{Id: r.Uid})

	if err != nil {
		return nil, err
	}

	now := time.Now().Unix()
	auth := l.svcCtx.Config.Auth
	jwt, _ := aauth.GetJwtToken(auth.AccessSecret, now, auth.AccessExpire, u.User.Id)

	return &types.RegisterRsp{
		Code: 0,
		User: types.User{
			Id:      strconv.FormatInt(u.User.Id, 10),
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
			Account: u.User.Account,
		},
		Token: jwt,
	}, nil
}
