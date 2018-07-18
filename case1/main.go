package main

import "fmt"

func main() {
	const (
		a = iota
		b = iota
		c = -1
		d = iota
	)

	fmt.Println(a, b, c, d)
}
