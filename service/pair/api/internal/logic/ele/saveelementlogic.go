package ele

import (
	"pair/service/pair/rpc/pair/pb"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"

	"pair/service/pair/api/internal/svc"
	"pair/service/pair/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) SaveElementLogic {
	return SaveElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveElementLogic) SaveElement(req *types.EleSaveReq) (resp *types.Rsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()
	ele := pb.Element{}
	copier.Copy(&ele, &req.Element)

	_, err = l.svcCtx.PairRPC.ElementSave(l.ctx, &pb.EleSaveReq{
		Uid:     uid,
		Element: &ele,
	})

	if err != nil {
		return nil, err
	}

	return &types.Rsp{Code: 0}, nil
}
