package idwork

import "caixin.app/caixos/tokit/configs"

func ID() string {
	return getID(int64(configs.EnvInt("server_id", 512)))
}
