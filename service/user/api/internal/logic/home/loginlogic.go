package home

import (
	"pair/common/aauth"
	"pair/common/aerror"
	"pair/service/user/rpc/user/pb"
	"context"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRsp, err error) {
	us := l.svcCtx.UserRPC

	u, err := us.UserInfo(l.ctx, &pb.UserInfoReq{
		Account: req.Account,
	})

	if err != nil {
		return nil, err
	}

	pwdErr := bcrypt.CompareHashAndPassword([]byte(u.User.Password), []byte(req.Password))
	if pwdErr != nil {
		return nil, aerror.Err(aerror.ErrCodeUserPasswordIncorrect)
	}

	now := time.Now().Unix()
	auth := l.svcCtx.Config.Auth
	jwt, _ := aauth.GetJwtToken(auth.AccessSecret, now, auth.AccessExpire, int64(u.User.Id))

	return &types.LoginRsp{
		Code: 0,
		User: types.User{
			Id:      strconv.FormatInt(u.User.Id, 10),
			Account: u.User.Account,
			Name:    u.User.Name,
			Avatar:  u.User.Avatar,
		},
		Token: jwt,
	}, nil
}
