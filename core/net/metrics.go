package net

import "github.com/prometheus/client_golang/prometheus"

var handlerSpendStats = prometheus.NewGaugeVec(prometheus.GaugeOpts{
	Namespace: "logic_process",
	Name:      "handler",
	Help:      "a net handler spend how much time,unit nanosecond",
}, []string{"msgid"})

func registerMetrics() {
	prometheus.Register(handlerSpendStats)
}
