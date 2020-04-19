package queues

import (
	"caixin.app/tokit/config"
	"caixin.app/tokit/client"
)

func Fire(name string, router string, params map[string]interface{}) error {
	conn := client.Redis()
	defer conn.Close()
	job := &Job{
		Queue: name,
		Payload: Payload{
			Route:  router,
			Params: params,
		},
	}
	prefix := config.EnvString("queue.prefix", "wego")
	return Enqueue(conn, job, prefix)

}
