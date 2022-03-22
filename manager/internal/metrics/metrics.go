package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	rkprom "github.com/rookie-ninja/rk-prom"
)

var metricsSet *rkprom.MetricsSet

type Cmd string

const (
	AddFailed Cmd = "add_failed"
	Add       Cmd = "add"
)

var cmdList = [...]Cmd{Add, AddFailed}

func Init(namespace string, registerer prometheus.Registerer) {
	metricsSet = rkprom.NewMetricsSet(namespace, "manager", registerer)

	for _, cmd := range cmdList {
		var err = metricsSet.RegisterCounter(string(cmd))
		if err != nil {
			panic(err)
		}
	}
}

func GetMetricsSet() *rkprom.MetricsSet {
	return metricsSet
}

func Report(cmd Cmd) {
	metricsSet.GetCounterWithValues(string(cmd)).Inc()
}
