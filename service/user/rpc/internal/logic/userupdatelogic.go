package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"pair/common/aerror"
	"pair/service/user/rpc/internal/svc"
	"pair/service/user/rpc/user/pb"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *pb.UserUpdateReq) (*pb.Response, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	if len(in.Password) > 0 {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		in.Password = string(pwd)
	}else{
		in.Password = user.Password
	}

	copier.Copy(user, in)

	if upErr := l.svcCtx.UserModel.Update(l.ctx, user); upErr != nil {
		return nil, aerror.ErrLog(upErr, in)
	}

	return &pb.Response{Code: 0}, nil
}
