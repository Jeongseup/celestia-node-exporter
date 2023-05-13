package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// fail counter
	failCount = promauto.NewCounter(
		prometheus.CounterOpts{
			Namespace: "celestia_node_exporter",
			Name:      "failed_counter",
			Help:      "Example metric with a string value.",
		},
	)
	// expose metrics
	chainIDMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "celestia_node_exporter",
			Name:      "current_chain_id",
			Help:      "exposing currrent your node's chain id for checking correct network.",
		},
		[]string{"chain_id"},
	)
	heightMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "celestia_node_exporter",
			Name:      "current_height",
			Help:      "exposing currrent your node's height.",
		})
	syncTimeMetric = promauto.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "celestia_node_exporter",
			Name:      "current_synctime_interval",
			Help:      "exposing currrent your node's time interval to check synced node with now and blocktimestamp.",
		})
)
