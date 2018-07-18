package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type MyData struct {
		One int    `json:"one"`
		two string `json:"two"`
	}

	in := MyData{1, "two"}
	fmt.Printf("%#v\n", in)

	encoded, err := json.Marshal(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(encoded))

	var out MyData
	json.Unmarshal(encoded, &out)
	fmt.Printf("%#v\n", out)
}
