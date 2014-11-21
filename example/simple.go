package main

import (
	"fmt"
	"github.com/mcuadros/go-jsonschema-generator"
)

type ExampleBasic struct {
	Foo bool   `json:"foo"`
	Bar string `json:",omitempty"`
	Qux int8
}

func main() {
	j := &jsonschema.JSONSchema{}
	j.Load(&ExampleBasic{})

	fmt.Println(j)
}
