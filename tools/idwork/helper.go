package idwork

import "caixin.app/caixos/tokit/config"

func ID() string {
	return getID(int64(config.EnvInt("server_id", 512)))
}
