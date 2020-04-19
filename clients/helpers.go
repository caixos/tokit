package clients

import (
	"github.com/caixos/tokit/clients/mysql"
	"github.com/caixos/tokit/clients/redis"
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
