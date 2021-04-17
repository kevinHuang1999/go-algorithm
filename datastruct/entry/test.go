package main

import (
	"algorithm/datastruct/maps"
	"fmt"
)

func main() {
	myMap := maps.NewHashMap()

	myMap.Put("kevin", "1")
	fmt.Printf("key = %v, value = %v\n", "kevin", myMap.Get("kevin"))
	myMap.Put("kevin", "2")
	fmt.Printf("key = %v, value = %v\n", "kevin", myMap.Get("kevin"))

}
