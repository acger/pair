package database

import (
	es "github.com/elastic/go-elasticsearch/v8"
	"github.com/zeromicro/go-zero/core/logx"
)

type ElasticsearchConf struct {
	Addresses []string
	Username  string
	Password  string
}

func NewElasticsearch(conf *ElasticsearchConf) *es.Client {
	client, err := es.NewClient(es.Config{
		Addresses: conf.Addresses,
		Username:  conf.Username,
		Password:  conf.Password,
	})

	if err != nil {
		logx.Error("elasticsearch connect fail.")
		return nil
	}

	go InitIndex(client, ES_ACGER_PAIR, AcgerPairBody)

	return client
}
