package synchro

import (
	"sync"

	"github.com/mojo-lang/mojo/go/pkg/mojo/db"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/config"
	"github.com/ncraft-io/ncraft/go/pkg/ncraft/logs"
)

var defaultDB string
var ds map[string]*db.DB
var dOnce sync.Once

type DataDBCfg struct {
	Dbs     map[string]*db.Config `json:"dbs"`
	Default string                `json:"default"`
}

func GetDataDB(name string) *db.DB {
	dOnce.Do(func() {
		cfgs := &DataDBCfg{}
		err := config.Get("dataDB").Scan(cfgs)
		if err != nil {
			logs.Errorw("failed to get the data db config", "error", err.Error())
			panic("failed to get the data db config")
		}
		defaultDB = cfgs.Default

		ds = make(map[string]*db.DB)

		for k, cfg := range cfgs.Dbs {
			d := db.New(cfg)
			if d == nil {
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
			ds[k] = d
		}
	})

	if d, ok := ds[name]; ok {
		return d
	} else if d, ok = ds[defaultDB]; ok {
		return d
	} else {
		if len(ds) == 1 {
			for _, d = range ds {
				return d
			}
		}
	}

	return nil
}
