package ele

import (
	"pair/service/pair/rpc/pair/pb"
	"context"
	"github.com/jinzhu/copier"
	"strconv"

	"pair/service/pair/api/internal/svc"
	"pair/service/pair/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListElementLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListElementLogic(ctx context.Context, svcCtx *svc.ServiceContext) ListElementLogic {
	return ListElementLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListElementLogic) ListElement(in *types.EleListReq) (resp *types.EleListRsp, err error) {
	r, err := l.svcCtx.PairRPC.ElementList(l.ctx, &pb.EleListReq{Page: in.Page, PageSize: in.PageSize})

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
