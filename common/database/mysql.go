package database

import (
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MysqlConf struct {
	DataSource string
}

func NewMysql(datasource string) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               datasource,
		DefaultStringSize: 255,
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		logx.Error("mysql connect fail")
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
