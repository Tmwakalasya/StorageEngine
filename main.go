package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

const filepath = "logs.txt"

func main() {

	initialize()
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		fmt.Println("Starting HTTP server on : 9090")
		if err := http.ListenAndServe(":9090", nil); err != nil {
			fmt.Println("Error starting HTTP server %v\n", err)
		}
	}()
	store1 := NewKeyValueStorage()

	store1.Set("Customer1", "$547.45")
	store1.Set("Customer2", "$123.45")
	store1.Delete("Customer1")
	store1.Get("Customer2")

	select {}
}
