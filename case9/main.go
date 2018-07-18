package main

import (
	"fmt"
	"time"
)

func main() {
	var dataMap map[string]int

	dataMap = make(map[string]int)
	for i := 0; i < 10000; i++ {
		go func(d map[string]int, num int) {
			d[fmt.Sprintf("%d", num)] = num
		}(dataMap, i)
	}
	time.Sleep(5 * time.Second)
	fmt.Println(len(dataMap))
}
