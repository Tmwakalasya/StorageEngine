package main

import (
	"fmt"
)

const filepath = "logs.txt"

func main() {
	store1 := NewKeyValueStorage()
	store1.Set("Customer 1", "$547.45")
	fmt.Println(store1)

	fmt.Println(store1.Get("tuntu"))
	store1.Delete("tuntu")
	fmt.Println(store1.Exists("tuntu"))
	store1.Set("Customer 2", "552.65")
	store1.Get("Customer 2")
	store1.Exists("Customer 2")
	store1.ReBuildStore(filepath)
}
