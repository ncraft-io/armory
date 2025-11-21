package model

import (
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

var d *db.DB
var dOnce sync.Once

type DBConfig struct {
	MaxIdleConn int `json:"MaxIdleConn" default:"2"` // 最大空闲链接
}

func GetDB() *db.DB {
	dOnce.Do(func() {
		cfg := &db.Config{}
		err := config.Get("db").Scan(cfg)
		if err != nil {
			logs.Errorw("failed to get the db config", "error", err.Error())
			panic("failed to get the db config")
		}

		if d = db.New(cfg); d == nil {
			panic("create the db failed")
		}

		// use these options when creating tables
		if cfg.Driver == db.MysqlDriverName {
			d.DB = d.DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")
			sqlDb, err := d.DB.DB()
			if err != nil {
				logs.Warnw("get sqlDb error")
			} else {
				if cfg.GetMaxIdleConnections() != 0 {
					sqlDb.SetMaxIdleConns(int(cfg.GetMaxIdleConnections()))
					logs.Debugw("set mysql max idle connections", "val", cfg.GetMaxIdleConnections())
				}
				if cfg.GetMaxOpenConnections() != 0 {
					sqlDb.SetMaxOpenConns(int(cfg.GetMaxOpenConnections()))
					logs.Debugw("set mysql max open connections", "val", cfg.GetMaxOpenConnections())
				}
			}
		}
	})
	return d
}
