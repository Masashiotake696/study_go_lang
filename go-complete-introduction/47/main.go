package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := &Person{
		Name: "otake",
		Age:  24,
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String())

	var p2 Person
	dec := json.NewDecoder(&buf)
	if err := dec.Decode(&p2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", p2)
}
