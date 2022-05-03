package apollo

import (
	"log"

	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
)

var (
	apolloClient agollo.Client
	namespace    = "application"

	EnvDev = "default"
	EnvPro = "pro"
)

func Config() *storage.Config {
	return apolloClient.GetConfig(namespace)
}

func Init(env string, appID string) {
	c := &config.AppConfig{
		AppID:          appID,
		Cluster:        env,
		IP:             "http://127.0.0.1:9180",
		NamespaceName:  namespace,
		IsBackupConfig: true,
		Secret:         "",
	}

	var err error
	apolloClient, err = agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return c, nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
