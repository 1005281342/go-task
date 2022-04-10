package main

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/1005281342/go-task/dispatcher/sdk"
	"github.com/1005281342/go-task/manager/manager"
	"github.com/1005281342/go-task/pkg/http"
	"github.com/1005281342/go-task/sendmessage/sendmessage"

	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
)

var wg = &sync.WaitGroup{}

var (
	duration   = flag.Int("d", 300, "持续时间")
	concurrent = flag.Int("c", 2, "并发数")
)

func main() {
	rand.Seed(time.Now().Unix())

	for i := 0; i < *duration; i++ {
		for j := 0; j < *concurrent; j++ {
			wg.Add(1)
			go task()
		}
		time.Sleep(time.Second)
	}
	wg.Wait()
}

func task() {
	defer wg.Done()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var (
		at = time.Now().Add(time.Second * time.Duration(rand.Intn(5)-2)).Unix()

		senders = [...]string{"goodman", "a", "xxx", "xiaoming"}

		sendReq = &sendmessage.SendReq{
			Msg:      &sendmessage.Message{Message: "hi"},
			Sender:   &sendmessage.Sender{Id: senders[rand.Intn(len(senders))]},
			Receiver: &sendmessage.Receiver{Id: "man"},
		}
		//sendBody []byte
		err error
	)
	//if sendBody, err = json.Marshal(sendReq); err != nil {
	//	panic(err)
	//}

	var (
		payload     = &sdk.Payload{ServiceName: "sendmessage.rpc", MethodName: "send", Body: sendReq}
		payloadBody []byte
	)
	if payloadBody, err = json.Marshal(payload); err != nil {
		panic(err)
	}

	var (
		addReq = &manager.AddReq{
			Task: &manager.TaskInfo{
				Queue:   "default",
				Name:    uuid.NewString(),
				At:      at,
				Payload: string(payloadBody),
			},
		}
		addBody []byte
	)
	if addBody, err = json.Marshal(addReq); err != nil {
		panic(err)
	}
	log.Printf("addBody: %s", addBody)

	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)+1))

	var content string
	if content, err = http.PostJSON("http://127.0.0.1:2333/manager.rpc/add", string(addBody)); err != nil {
		log.Printf("err: %+v", err)
	}
	log.Printf("content: %s", content)
}
