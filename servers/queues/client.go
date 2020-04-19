package queues

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"fmt"
)

func Enqueue(conn redis.Conn, job *Job, prefix string) error {
	buffer, err := json.Marshal(job.Payload)
	if err != nil {
		return err
	}
	err = conn.Send("RPUSH", fmt.Sprintf("%s_queue:%s", prefix, job.Queue), buffer)
	if err != nil {
		return err
	}
	return conn.Flush()
}
