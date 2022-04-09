package logic

import (
	"pair/common/aerror"
	"pair/service/user/model"
	"context"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"

	"pair/service/user/rpc/internal/svc"
	"pair/service/user/rpc/user/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"math/rand"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *pb.UserAddReq) (*pb.UserAddRsp, error) {
	_, err := l.svcCtx.UserModel.FindOneByAccount(l.ctx, in.Account)

	if err == nil {
		return nil, aerror.Err(aerror.ErrCodeUserAccountExists)
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	rand.Seed(time.Now().UnixNano())

	if in.Avatar == "" {
		randNum := rand.Intn(6)
		randNumStr := strconv.Itoa(randNum)
		in.Avatar = "avatar/" + randNumStr + ".jpg"
	}

	userInfo := &model.Users{
		Account:  in.Account,
		Password: string(pwd),
		Name:     in.Name,
		Avatar:   in.Avatar,
	}

	r, inErr := l.svcCtx.UserModel.Insert(l.ctx, userInfo)

	if inErr != nil {
		return nil, aerror.ErrLog(inErr, in)
	}

	uid, _ := r.LastInsertId()

	return &pb.UserAddRsp{
		Code: 0,
		Uid:  uid,
	}, nil
}
