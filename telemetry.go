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
			Name: "kvstore_writes_total",
			Help: "Total number of writes from the key-value store",
		},
	)
)

var (
	kvstoreErrorsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kvstore_errors_total",
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
	prometheus.MustRegister(
		kvstoreReadsTotal,
		kvstoreErrorsTotal,
		kvstoreWritesLatencySeconds,
		kvstoreReadsLatencySeconds,
		kvstoreWritesTotal,
	)
	fmt.Println("Prometheus metrics registered successfully")

}
