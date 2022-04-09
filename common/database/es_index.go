package database

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"strings"
	"time"
)

const ES_ACGER_PAIR = "acger_pair"

var AcgerPairBody = `
{
  "settings": {
    "number_of_replicas": 0,
    "number_of_shards": 5
  },
  "mappings": {
    "properties": {
      "uid": {
        "type": "keyword"
      },
      "boost": {
        "type": "integer"
      },
      "start": {
        "type": "boolean"
      },
      "skill": {
        "type": "text",
        "analyzer": "smartcn"
      },
      "skill_need": {
        "type": "text",
        "analyzer": "smartcn"
      },
      "create_time": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis||yyyy-MM-dd HH:mm:ss"
      },
      "update_time": {
        "type": "date",
        "format": "strict_date_optional_time||epoch_millis||yyyy-MM-dd HH:mm:ss"
      }
    }
  }
}
`

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}
type Total struct {
	Value    int64  `json:"value"`
	Relation string `json:"relation"`
}
type Source struct {
	Skill      string    `json:"skill"`
	SkillNeed  string    `json:"skill_need"`
	UID        int64     `json:"uid"`
	Boost      int64     `json:"boost"`
	Star       int64     `json:"star"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
type Highlight struct {
	Skill     []string `json:"skill"`
	SkillNeed []string `json:"skill_need"`
}
type Hits struct {
	Index     string    `json:"_index"`
	Type      string    `json:"_type"`
	ID        string    `json:"_id"`
	Score     float64   `json:"_score"`
	Source    Source    `json:"_source"`
	Highlight Highlight `json:"highlight"`
}
type Hit struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hits  `json:"hits"`
}

type EsSearchPairResult struct {
	Took     int64  `json:"took"`
	TimedOut bool   `json:"timed_out"`
	Shards   Shards `json:"_shards"`
	Hits     Hit    `json:"hits"`
}

func InitIndex(client *es.Client, index string, body string) {
	rsp, err := client.Indices.Get([]string{index})

	if err != nil {
		logx.Error("init es index fail")
		logx.Error(err.Error())
		logx.Error(rsp)
		return
	}

	if rsp.StatusCode == http.StatusNotFound {
		_, e := client.Indices.Create(index, client.Indices.Create.WithBody(strings.NewReader(body)))
		if e != nil {
			logx.Error("init es index fail.")
			logx.Error(e.Error())
			return
		}
	}
}
