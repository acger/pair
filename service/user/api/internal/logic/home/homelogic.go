package home

import (
	"context"

	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) HomeLogic {
	return HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home() (resp *types.Rsp, err error) {
	return &types.Rsp{Code: 0, Message: "a4"}, nil
}
