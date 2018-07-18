package main

import "fmt"

func main() {
	var num int

	for i := 0; i < 10000; i++ {
		go func() {
			num = i
		}()
	}
	fmt.Printf("NUM is %d", num)
}
