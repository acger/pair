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
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoReq) (*pb.UserInfoRsp, error) {
	var err error
	var u *model.Users

	if len(in.Account) > 0 {
		u, err = l.svcCtx.UserModel.FindOneByAccount(l.ctx, in.Account)
	} else {
		u, err = l.svcCtx.UserModel.FindOne(l.ctx, int64(in.Id))
	}

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	userDetail := user.UserDetail{}
	copier.Copy(&userDetail, u)

	return &pb.UserInfoRsp{
		Code: 0,
		User: &userDetail,
	}, nil
}
