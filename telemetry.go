package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	kvstoreReadsTotal = prometheus.NewCounter(

		prometheus.CounterOpts{
			Name: "kvstore_reads_total",
			Help: "Total number of reads from the key-value store.",
		},
	)
)

var (
	kvstoreWritesTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kvstore_total_writes",
			Help: "Total number of writes from the key-value store",
		},
	)
)

var (
	kvstoreErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kvstore_total_errors",
			Help: "Total number of failed operations",
		},
	)
)

var (
	kvstoreReadsLatencySeconds = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "kvstore_read_latency_seconds",
			Help: "This histogram represents the read latency of the key-value store",
		},
	)
)
var (
	kvstoreWritesLatencySeconds = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "kvstore_write_latency_seconds",
			Help: "This histogram represents the write latency of the key-value store",
		},
	)
)

func initialize() {
	prometheus.MustRegister(kvstoreReadsTotal)
	prometheus.MustRegister(kvstoreWritesTotal)
	prometheus.MustRegister(kvstoreErrorsTotal)
	prometheus.MustRegister(kvstoreReadsLatencySeconds)
	prometheus.MustRegister(kvstoreWritesLatencySeconds)
	fmt.Println("Prometheus metrics registered successfully")

}
