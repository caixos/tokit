package idwork

import "caixin.app/tokit/config"

func ID() string {
	return getID(int64(config.EnvInt("server_id", 512)))
}
