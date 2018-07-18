package main

import "fmt"

func main() {
	const (
		a, b, c = iota, iota, iota
	)

	fmt.Println(a, b, c)
}
