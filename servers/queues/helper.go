package queues

import (
	"caixin.app/caixos/tokit/configs"
	"caixin.app/caixos/tokit/clients"
)

func Fire(name string, router string, params map[string]interface{}) error {
	conn := clients.Redis()
	defer conn.Close()
	job := &Job{
		Queue: name,
		Payload: Payload{
			Route:  router,
			Params: params,
		},
	}
	prefix := configs.EnvString("queue.prefix", "wego")
	return Enqueue(conn, job, prefix)

}
