package main

import (
	"context"
	"encoding/json"
	"github.com/1005281342/go-task/pkg/apollo"
	"io/ioutil"
	"log"
	"net/http"
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
		return err
	}

	log.Printf("payload: %s", string(task.Payload()))
	log.Printf("task: %+v", task)

	var content string
	if content, err = Post(payload); err != nil {
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
	var body, err = json.Marshal(payload.Body)
	log.Printf("body: %s", string(body))

	var res *http.Response
	res, err = http.Post(targetURL, "application/json;charset=utf-8", strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	var content []byte
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
