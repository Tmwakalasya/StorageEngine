package main

import (
	"flag"
	"fmt"
	"os"
)

const filepath = "logs.txt"

func main() {

	//initialize()
	//http.Handle("/metrics", promhttp.Handler())
	//go func() {
	//	fmt.Println("Starting HTTP server on : 9090")
	//	if err := http.ListenAndServe(":9090", nil); err != nil {
	//		fmt.Println("Error starting HTTP server %v\n", err)
	//	}
	//}()
	//store1 := NewKeyValueStorage()
	//
	//store1.Set("Customer1", "$547.45")
	//store1.Set("Customer2", "$123.45")
	//store1.Delete("Customer1")
	//store1.Get("Customer2")

	if len(os.Args) < 2 {
		fmt.Println("Error: No command given, available commands (GET, SET, DELETE, EXISTS)")
		return
	}

	// Print the command-line arguments
	fmt.Println("Arguments:", os.Args)

	// Define the SET command
	SETcmd := flag.NewFlagSet("SET", flag.ExitOnError)
	SetKV := SETcmd.String("set", "", "Set a key and value to the store")

	switch os.Args[1] {
	case "SET":
		SETcmd.Parse(os.Args[2:])
		if *SetKV == "" {
			fmt.Println("Error: No key-value pair provided for SET command")
			return
		}
		fmt.Println("SET command called with:", *SetKV)
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}

}
