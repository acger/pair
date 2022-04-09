package logic

import (
	"pair/common/aerror"
	"pair/common/database"
	"pair/service/pair/model"
	"pair/service/pair/rpc/internal/svc"
	"pair/service/pair/rpc/pair/pb"
	"bytes"
	"context"
	"database/sql"
	"github.com/jinzhu/copier"
	json "github.com/json-iterator/go"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type ElementSaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewElementSaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ElementSaveLogic {
	return &ElementSaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ElementSaveLogic) ElementSave(in *pb.EleSaveReq) (*pb.Response, error) {
	ele, err := l.svcCtx.ElementModel.FindOneByUid(l.ctx, in.Uid)

	if err == nil && ele.Skill == in.Element.Skill && ele.SkillNeed == in.Element.SkillNeed {
		return &pb.Response{Code: 0}, nil
	}

	in.Element.Uid = in.Uid
	data := model.Elements{}
	copier.Copy(&data, in.Element)

	var opErr error

	if err == nil {
		data.Id = ele.Id
		opErr = l.svcCtx.ElementModel.Update(l.ctx, &data)
	} else {
		var rsp sql.Result
		rsp, opErr = l.svcCtx.ElementModel.Insert(l.ctx, &data)
		data.Id, _ = rsp.LastInsertId()
	}

	if opErr != nil {
		return nil, aerror.ErrLog(opErr, in)
	}

	//if canal not work
	go func() error {
		ctx := context.Background()
		currentEle, fErr := l.svcCtx.ElementModel.FindOne(ctx, data.Id)
		if fErr != nil {
			return aerror.ErrLog(fErr, in)
		}

		eleJson := bytes.Buffer{}
		json.NewEncoder(&eleJson).Encode(currentEle)
		es := l.svcCtx.ES
		docId := strconv.FormatInt(currentEle.Uid, 10)
		_, esErr := es.Index(database.ES_ACGER_PAIR, &eleJson, es.Index.WithDocumentID(docId))

		if esErr != nil {
			return aerror.ErrLog(esErr)
		}

		return nil
	}()

	return &pb.Response{Code: 0}, nil
}
