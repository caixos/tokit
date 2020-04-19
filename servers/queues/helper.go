package queues

import (
	"github.com/caixos/tokit/configs"
	"github.com/caixos/tokit/clients"
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
