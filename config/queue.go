package config

import "time"

type QueueConfig struct {
	Prefix      string
	Listen      []string
	Interval    time.Duration
	Concurrency int
}

func LoadQueueConfig() *QueueConfig {
	interval := EnvInt("queue.interval", 1)
	config := &QueueConfig{
		Prefix:      EnvString("queue.prefix", "wego"),
		Listen:      EnvStringSlice("queue.listen"),
		Interval:    time.Duration(interval) * time.Second,
		Concurrency: EnvInt("queue.concurrency", 1),
	}
	return config
}
