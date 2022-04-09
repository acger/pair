package logic

import (
	"pair/common/aerror"
	"bytes"
	"context"
	"strconv"

	"pair/service/pair/rpc/internal/svc"
	"pair/service/pair/rpc/pair/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementListLogic {
	return &ElementListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementListLogic) ElementList(in *pb.EleListReq) (*pb.ELePairRsp, error) {
	from := (in.Page - 1) * in.PageSize
	size := in.PageSize
	fromStr := strconv.FormatInt(int64(from), 10)
	sizeStr := strconv.FormatInt(int64(size), 10)

	body := &bytes.Buffer{}
	body.WriteString(`
		{
		  "from": ` + fromStr + `,
		  "size": ` + sizeStr + `,
		  "_source": ["uid", "skill", "skill_need"],
		  "query": {
			"match_all": {}
		  },
		  "sort": [
			{
			  "update_time": {
				"order": "desc"
			  }
			}
		  ]
		}
    `)

	ue, err := Pair(l.ctx, l.svcCtx, body)

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	return &pb.ELePairRsp{Code: 0, UserElement: ue}, nil
}
