package main

import (
	"fmt"
	"github.com/mcuadros/go-jsonschema-generator"
)

type ExampleBasic struct {
	Foo bool   `json:"foo"`
	Bar string `json:",omitempty"`
	Qux int8
	Baz []string
}

func main() {
	s := &jsonschema.Schema{}
	s.Load(&ExampleBasic{})
	fmt.Println(s)
}
