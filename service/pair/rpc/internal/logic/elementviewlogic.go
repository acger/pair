package logic

import (
	"pair/common/aerror"
	"pair/service/pair/rpc/internal/svc"
	"pair/service/pair/rpc/pair/pb"
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementViewLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementViewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementViewLogic {
	return &ElementViewLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementViewLogic) ElementView(in *pb.EleViewReq) (*pb.EleViewRsp, error) {
	ele, err := l.svcCtx.ElementModel.FindOneByUid(l.ctx, in.Uid)

	if err != nil {
		if errors.Is(err, sqlx.ErrNotFound) {
			return &pb.EleViewRsp{Code: 0}, nil
		}

		return nil, aerror.ErrLog(err)
	}

	eleRsp := pb.Element{}
	copier.Copy(&eleRsp, ele)

	return &pb.EleViewRsp{
		Code:    0,
		Element: &eleRsp,
	}, nil
}
