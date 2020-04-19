package client

import (
	"caixin.app/caixos/tokit/client/mysql"
	"caixin.app/caixos/tokit/client/redis"
	"github.com/go-xorm/xorm"
	redigo "github.com/gomodule/redigo/redis"
)

func DB() *xorm.Engine {
	return mysql.GetDB()
}

func Redis() redigo.Conn {
	return redis.GetRedisPool().Get()
}

func RedisPool() *redigo.Pool {
	return redis.GetRedisPool()
}
