package idwork

import "github.com/caixos/tokit/configs"

func ID() string {
	return getID(int64(configs.EnvInt("server_id", 512)))
}
