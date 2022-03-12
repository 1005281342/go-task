package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/1005281342/httproxy"
	"github.com/polarismesh/polaris-go/api"
)

const (
	errMsgKey = httproxy.ErrMsgKey
)

var (
	namespace string
	gConsumer api.ConsumerAPI
)

func initArgs() {
	flag.StringVar(&namespace, "namespace", "default", "namespace")
}

func main() {
	initArgs()
	flag.Parse()

	var consumer, err = api.NewConsumerAPI()
	if err != nil {
		panic(err)
	}
	gConsumer = consumer

	go func() {
		http.HandleFunc("/err", errHandle)
		http.ListenAndServe("127.0.0.1:2334", nil)
	}()
	http.ListenAndServe("127.0.0.1:2333", httproxy.New(namespace, gConsumer))
}

func errHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, r.Header.Get(errMsgKey)) //这个写入到w的是输出到客户端的
	r.Header.Del(errMsgKey)
}
