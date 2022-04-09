package ele

import (
	"pair/service/pair/rpc/pair/pb"
	"context"
	"encoding/json"
	"github.com/jinzhu/copier"
	"strconv"

	"pair/service/pair/api/internal/svc"
	"pair/service/pair/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) ElementLogic {
	return ElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ElementLogic) Element(in *types.EleViewReq) (resp *types.EleRsp, err error) {
	var uid int64

	if in.Uid != "" {
		uid, _ = strconv.ParseInt(in.Uid, 10, 64)
	} else {
		uid, _ = l.ctx.Value("userId").(json.Number).Int64()
	}

	r, err := l.svcCtx.PairRPC.ElementView(l.ctx, &pb.EleViewReq{Uid: uid})

	if err != nil {
		return nil, err
	}

	ele := &types.Element{}
	copier.Copy(ele, r.Element)

	return &types.EleRsp{Code: 0, Element: ele}, nil
}
