package svc

import (
	"pair/common/database"
	"pair/service/user/model"
	"pair/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB:        database.NewMysql(c.DB.DataSource),
		UserModel: model.NewUsersModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
