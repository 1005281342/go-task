package main

import (
	"github.com/1005281342/go-task/dispatcher/sdk"
	"log"
	"time"

	"github.com/1005281342/go-task/manager/manager"
	"github.com/1005281342/go-task/pkg/http"
	"github.com/1005281342/go-task/sendmessage/sendmessage"
	"github.com/google/uuid"

	"github.com/json-iterator/go"
)

func main() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	var (
		sendReq = &sendmessage.SendReq{
			Msg:      &sendmessage.Message{Message: "hi"},
			Sender:   &sendmessage.Sender{Id: "goodman"},
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
				At:      time.Now().Add(time.Second * 5).Unix(),
				Payload: string(payloadBody),
			},
		}
		addBody []byte
	)
	if addBody, err = json.Marshal(addReq); err != nil {
		panic(err)
	}
	log.Printf("addBody: %s", addBody)

	var content string
	if content, err = http.PostJSON("http://127.0.0.1:2333/manager.rpc/add", string(addBody)); err != nil {
		panic(err)
	}
	log.Printf("content: %s", content)
}
