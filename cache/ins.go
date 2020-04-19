package cache

import (
	"caixin.app/tokit/config"
	"caixin.app/tokit/constant"
	"github.com/coocood/freecache"
	"encoding/json"
	"runtime/debug"
	"sync"
	"errors"
)

var ins *freecache.Cache
var once sync.Once

func GetIns() *freecache.Cache {
	once.Do(func() {
		ins = initCache()
	})
	return ins
}

func initCache() *freecache.Cache {
	config := config.LoadCacheConfig()
	if config.Size != 0 {
		c := freecache.NewCache(config.Size)
		//根据cache的大小进行设置
		debug.SetGCPercent(20)
		return c
	}
	return nil
}

func Set(key string, value interface{}, exp int) error {
	if GetIns() == nil {
		return errors.New(constant.ErrCacheInit)
	}

	k := []byte(key)
	v, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = GetIns().Set(k, v, exp)
	if err != nil {
		return err
	}
	return nil
}
func Get(key string) ([]byte, error) {
	if GetIns() == nil {
		return nil, errors.New(constant.ErrCacheInit)
	}
	k := []byte(key)
	return GetIns().Get(k)
}
