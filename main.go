package main

import (
	"fmt"
)

const filepath = "logs.txt"
const rebuildfilepath = "logs3.txt"

func main() {
	store1 := NewKeyValueStorage()
	store1.Set("Customer 1", "$547.45")
	fmt.Println(store1)
	fmt.Println(store1.Get("Tuntu"))
	store1.ReBuildStore(rebuildfilepath)
}
