package mq

import (
	"pair/service/chat/model"
	"github.com/jinzhu/copier"
	json "github.com/json-iterator/go"
	"github.com/zeromicro/go-queue/kq"
	"gorm.io/gorm"
	"strconv"
)

func CreateChatMessageConsumer(db *gorm.DB) kq.ConsumeHandle {
	return func(k, v string) error {
		m := model.ClientChatMsg{}
		json.Unmarshal([]byte(v), &m)

		c := model.Chats{}
		copier.Copy(&c, &m)
		c.Uid, _ = strconv.ParseInt(m.Uid, 10, 64)
		c.ToUid, _ = strconv.ParseInt(m.ToUid, 10, 64)
		db.Create(&c)

		return nil
	}
}

func GetKqueueList(db *gorm.DB) []kq.ConsumeHandle {
	return []kq.ConsumeHandle{
		CreateChatMessageConsumer(db),
	}
}
