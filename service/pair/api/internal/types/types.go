// Code generated by goctl. DO NOT EDIT.
package types

type EleListRsp struct {
	Code        int64          `json:"code"`
	Message     string         `json:"message"`
	UserElement []*UserElement `json:"userElement"`
	Page        int64          `json:"page"`
	PageSize    int64          `json:"pageSize"`
	Total       int64          `json:"total"`
}

type EleListReq struct {
	Page     int64  `form:"page,default=1"`
	PageSize int64  `form:"pageSize,default=30"`
	Keyword  string `form:"keyword,optional"`
}

type EleSaveReq struct {
	Element *Element `json:"element"`
}

type EleViewReq struct {
	Uid string `form:"uid,optional"`
}

type EleRsp struct {
	Code    int64    `json:"code"`
	Message string   `json:"message"`
	Element *Element `json:"element"`
}

type Rsp struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
}

type Element struct {
	Skill              string `json:"skill,optional"`
	SkillNeed          string `json:"skill_need,optional"`
	Star               bool   `json:"star,optional"`
	Boost              int64  `json:"boost,optional"`
	HighLightSkill     string `json:"highlight_skill,optional"`
	HighLightSkillNeed string `json:"highlight_skill_need,optional"`
}

type UserElement struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Account string   `json:"account"`
	Avatar  string   `json:"avatar"`
	Element *Element `json:"element"`
}
