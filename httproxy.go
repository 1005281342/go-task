package main

import (
	"flag"
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

	http.ListenAndServe("127.0.0.1:2333", httproxy.New(namespace, gConsumer))
}
