package main

import (
	"fmt"
)

const filepath = "logs.txt"

func main() {
	store1 := NewKeyValueStorage()
	store1.Set("Customer 1", "$547.45")
	fmt.Println(store1)
	fmt.Println(store1.Get("Tuntu"))
	store1.ReBuildStore(LogFilePath)
}
