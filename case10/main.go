package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	result := make([]*int, len(a))
	for i, v := range a {
		result[i] = &v
	}
	for _, u := range result {
		fmt.Printf("%d ", *u)
	}
}
