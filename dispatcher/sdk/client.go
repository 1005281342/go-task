package sdk

import "github.com/hibiken/asynq"

func NewClient(opt asynq.RedisClientOpt) *asynq.Client {

	return asynq.NewClient(opt)
}

var (
	DefaultRedisClientOpt = asynq.RedisClientOpt{Addr: "127.0.0.1:6379"}
)
