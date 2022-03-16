package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/1005281342/httproxy"
	"github.com/polarismesh/polaris-go/api"
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

	var path = "127.0.0.1:2333"
	log.Printf("http proxy %s", "127.0.0.1:2333")
	http.ListenAndServe(path, httproxy.New(namespace, gConsumer))
}
