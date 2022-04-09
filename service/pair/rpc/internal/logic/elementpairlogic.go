package logic

import (
	"pair/common/aerror"
	"pair/common/database"
	"pair/service/user/rpc/user"
	"bytes"
	"context"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"strconv"
	"time"

	"pair/service/pair/rpc/internal/svc"
	"pair/service/pair/rpc/pair/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ElementPairLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementPairLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementPairLogic {
	return &ElementPairLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func Pair(ctx context.Context, svcCtx *svc.ServiceContext, body io.Reader) ([]*pb.UserElement, error) {
	es := svcCtx.ES
	rsp, err := es.Search(es.Search.WithIndex(database.ES_ACGER_PAIR), es.Search.WithBody(body))
	if err != nil {
		return nil, aerror.ErrLog(err)
	}

	//读取搜索结果
	buff := &bytes.Buffer{}
	buff.ReadFrom(rsp.Body)
	result := database.EsSearchPairResult{}
	jsoniter.ConfigCompatibleWithStandardLibrary.Unmarshal(buff.Bytes(), &result)
	len := result.Hits.Total.Value

	if len == 0 {
		return nil, nil
	}

	//RPC获取用户信息
	var uids []int64
	uidHighlight := make(map[int64]database.Highlight)
	uidSource := make(map[int64]database.Source)

	for _, h := range result.Hits.Hits {
		uid := h.Source.UID
		uids = append(uids, int64(uid))
		uidHighlight[uid] = h.Highlight
		uidSource[uid] = h.Source
	}

	usersReq := user.UserListReq{Id: uids}
	usersRsp, userListErr := svcCtx.UserRPC.UserList(ctx, &usersReq)

	if userListErr != nil {
		return nil, userListErr
	}

	users := make(map[int64]pb.UserElement)

	for _, u := range usersRsp.User {
		users[u.Id] = pb.UserElement{
			Id:      u.Id,
			Name:    u.Name,
			Account: u.Account,
			Avatar:  u.Avatar,
			Element: new(pb.Element),
		}
	}

	//给搜索结果填补上用户信息
	var ue []*pb.UserElement
	for uid, source := range uidSource {
		id := uid
		e := users[id]
		e.Element.Uid = id
		e.Element.Skill = source.Skill
		e.Element.SkillNeed = source.SkillNeed

		if uidHighlight[uid].Skill != nil {
			e.Element.HighLightSkill = uidHighlight[uid].Skill[0]
		}

		if uidHighlight[uid].SkillNeed != nil {
			e.Element.HighLightSkillNeed = uidHighlight[uid].SkillNeed[0]
		}

		ue = append(ue, &e)
	}

	return ue, nil
}

func (l *ElementPairLogic) ElementPair(in *pb.ElePairReq) (*pb.ELePairRsp, error) {
	ele, err := l.svcCtx.ElementModel.FindOneByUid(l.ctx, in.Uid)

	if err != nil {
		return nil, aerror.ErrLog(err)
	}

	from := (in.Page - 1) * in.PageSize
	size := in.PageSize
	fromStr := strconv.FormatInt(from, 10)
	sizeStr := strconv.FormatInt(size, 10)
	uidStr := strconv.FormatInt(in.Uid, 10)
	boost := strconv.FormatInt(ele.Boost, 10)
	timeStr := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))

	body := &bytes.Buffer{}
	body.WriteString(`
		{
		  "from": ` + fromStr + `,
		  "size": ` + sizeStr + `,
		  "_source": ["uid", "skill", "skill_need"], 
		  "query": {
			"function_score": {
			  "query": {
				"bool": {
				  "must": [
					{
					  "multi_match": {
						"fields": [
						  "skill",
						  "skill_need"
						],
						"query": "` + ele.SkillNeed + `"
					  }
					}
				  ],
				  "must_not": [
					{
					  "term": {
						"uid": {
						  "value": ` + uidStr + `
						}
					  }
					}
				  ]
				}
			  },
			  "functions": [
				{
				  "gauss": {
					"update_time": {
					  "origin": "` + timeStr + `",
					  "scale": "10d",
					  "offset": "90d",
					  "decay": 0.5
					}
				  }
				},
				{
				  "filter": {
					"range": {
					  "boost": {
						"gt": 0
					  }
					}
				  },
				  "weight": ` + boost + `
				},
				{
				  "filter": {
					"term": {
					  "star": 1
					}
				  },
				  "weight": 5
				}
			  ],
			  "score_mode": "sum"
			}
		  },
		  "highlight": {
			"pre_tags": "<b>",
			"post_tags": "</b>", 
			"fields": {
			  "skill": {},
			  "skill_need": {}
			}
		  }
		}
    `)

	ue, err := Pair(l.ctx, l.svcCtx, body)

	if err != nil {
		return nil, aerror.ErrLog(err, in)
	}

	return &pb.ELePairRsp{Code: 0, UserElement: ue}, nil
}
