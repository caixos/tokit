package mysql

import (
	"caixin.app/caixos/tokit/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"sync"
)

var db *xorm.Engine
var onceMysql sync.Once

func GetDB() *xorm.Engine {
	onceMysql.Do(func() {
		db = newMySql()
	})
	return db
}
func newMySql() *xorm.Engine {
	conf := config.LoadMySqlConfig()
	engine, _ := xorm.NewEngine(
		conf.Driver,
		conf.DataSource,
	)
	engine.SetMaxIdleConns(conf.MaxIdleConns)
	engine.SetMaxOpenConns(conf.MaxOpenConns)
	engine.ShowSQL(conf.ShowSQL)
	return engine
}
