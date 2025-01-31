package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	kvstore_reads_total = prometheus.NewCounter(

		prometheus.CounterOpts{
			Name: "kvstore_reads_total",
			Help: "Total number of reads from the key-value store.",
		},
	)
)

var (
	kvstore_writes_total = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kvstore_total_writes",
			Help: "Total number of writes from the key-value store",
		},
	)
)

var (
	kvstore_errors_total = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "kvstore_total_errors",
			Help: "Total number of failed operations",
		},
	)
)

func initialize() {
	prometheus.MustRegister(kvstore_reads_total)
	prometheus.MustRegister(kvstore_writes_total)
	prometheus.MustRegister(kvstore_errors_total)
	fmt.Println("Prometheus metrics registered successfully")

}
