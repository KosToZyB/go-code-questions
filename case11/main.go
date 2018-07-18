package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	x := 0
	go func(p *int) {
		for i := 1; i <= 20000000000; i++ {
			*p = i
		}
	}(&x)

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("x = %d.\n", x)
}
