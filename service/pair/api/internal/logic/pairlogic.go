package logic

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

type PairLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPairLogic(ctx context.Context, svcCtx *svc.ServiceContext) PairLogic {
	return PairLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PairLogic) Pair(in *types.EleListReq) (resp *types.EleListRsp, err error) {
	uid, _ := l.ctx.Value("userId").(json.Number).Int64()

	r, err := l.svcCtx.PairRPC.ElementPair(l.ctx, &pb.ElePairReq{Uid: uid, Page: in.Page, PageSize: in.PageSize})

	if err != nil {
		return nil, err
	}

	userEleRsp := make([]*types.UserElement, len(r.UserElement))

	for i, ue := range r.UserElement {
		userEleRsp[i] = &types.UserElement{}
		copier.Copy(userEleRsp[i], ue)
		userEleRsp[i].Id = strconv.FormatInt(ue.Id, 10)
	}

	return &types.EleListRsp{Code: 0, UserElement: userEleRsp}, nil
}
