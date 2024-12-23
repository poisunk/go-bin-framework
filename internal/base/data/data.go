package data

import (
	"errors"
	"go-gin-framework/internal/base/conf"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

type Data struct {
	DB *xorm.Engine
}

func NewData(engine *xorm.Engine) *Data {
	return &Data{
		DB: engine,
	}
}

func NewDB(debug bool, dataConf *conf.Database) (*xorm.Engine, error) {
	if dataConf == nil {
		return nil, errors.New("dataConf is nil")
	}

	// 默认使用mysql
	if dataConf.Driver == "" {
		dataConf.Driver = "mysql"
	}

	// 创建引擎
	engine, err := xorm.NewEngine(dataConf.Driver, dataConf.Connection)
	if err != nil {
		return nil, err
	}

	if debug {
		engine.ShowSQL(true)
	} else {
		engine.SetLogLevel(log.LOG_ERR)
	}

	// 设置连接池
	engine.SetMaxOpenConns(dataConf.MaxOpenConns)
	engine.SetMaxIdleConns(dataConf.MaxIdleConns)
	engine.SetConnMaxLifetime(time.Duration(dataConf.MaxLifeSecs) * time.Second)

	return engine, nil
}
