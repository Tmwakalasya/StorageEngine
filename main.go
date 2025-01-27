package main

import "fmt"

const filepath = "logs.txt"

func main() {
	// Create the original key-value store
	store1 := NewKeyValueStorage()

	// Perform some operations on the original store
	store1.Set("Customer1", "$547.45")
	store1.Set("Customer2", "$123.45")
	store1.Delete("Customer1")

	// Replicate the store
	store1.Replication()

	// Create a new replica store
	replica := NewKeyValueStorage()
	replica.ReBuildStore(LogFilePath)

	// Compare the original store with the replica
	isEqual := store1.CompareReplica(replica)
	if isEqual {
		fmt.Println("Replication successful: The replica matches the original store.")
	} else {
		fmt.Println("Replication failed: The replica does not match the original store.")
	}
}
