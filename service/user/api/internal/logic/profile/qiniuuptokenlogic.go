package profile

import (
	"context"
	"fmt"

	"pair/service/user/api/internal/svc"
	"pair/service/user/api/internal/types"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"github.com/zeromicro/go-zero/core/logx"
)

type QiniuUpTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQiniuUpTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) QiniuUpTokenLogic {
	return QiniuUpTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QiniuUpTokenLogic) QiniuUpToken(req *types.QiniuUpReq) (resp *types.QiniuUpRsp, err error) {
	config := l.svcCtx.Config.Qiniu
	var scope string

	if req.Name != "" {
		scope = fmt.Sprintf("%s:%s", config.Bucket, req.Name)
	} else {
		scope = config.Bucket
	}

	putPolicy := storage.PutPolicy{
		Scope: scope,
	}

	mac := qbox.NewMac(config.AK, config.SK)
	upToken := putPolicy.UploadToken(mac)

	return &types.QiniuUpRsp{Code: 0, Token: upToken}, nil
}
