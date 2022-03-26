package main

import (
	"context"
	"encoding/json"
	"github.com/1005281342/go-task/pkg/apollo"
	"github.com/1005281342/go-task/pkg/http"
	"log"
	"reflect"
	"runtime"
	"strings"

	"github.com/1005281342/go-task/dispatcher/sdk"

	"github.com/hibiken/asynq"
)

func main() {

	apollo.Init(apollo.EnvDev, "dispatcher")

	var srv = asynq.NewServer(
		sdk.DefaultRedisClientOpt,
		asynq.Config{
			Concurrency: runtime.NumCPU(),
			Queues: map[string]int{
				"high":    6,
				"default": 3,
				"low":     1,
			},
		},
	)

	if err := srv.Run(asynq.HandlerFunc(handler)); err != nil {
		log.Fatal(err)
	}
}

func handler(ctx context.Context, task *asynq.Task) error {
	var (
		payload sdk.Payload
		err     error
	)
	if err = json.Unmarshal(task.Payload(), &payload); err != nil {
		log.Printf("payload: %s Unmarshal failed: %+v", string(task.Payload()), err)
		return err
	}

	log.Printf("payload: %s", string(task.Payload()))
	log.Printf("task: %+v", task)

	var content string
	if content, err = Post(payload); err != nil {
		log.Printf("payload: %s Post failed: %+v", string(task.Payload()), err)
		return err
	}
	log.Printf("content: %s", content)
	return nil
}

func Post(payload sdk.Payload) (string, error) {
	var targetURL = strings.Join([]string{
		apollo.Config().GetStringValue("url", "http://127.0.0.1:2333"),
		payload.ServiceName,
		payload.MethodName,
	}, "/")
	log.Printf("targetURL: %s", targetURL)

	var (
		body []byte
		err  error
	)
	switch reflect.TypeOf(payload.Body).Kind() {
	case reflect.String:
		body = []byte(payload.Body.(string))
	default:
		body, err = json.Marshal(payload.Body)
		if err != nil {
			return "", err
		}
	}

	log.Printf("body: %s", string(body))

	return http.PostJSON(targetURL, string(body))
}
