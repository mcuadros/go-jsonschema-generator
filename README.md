go-jsonschema-generator [![Build Status](https://travis-ci.org/mcuadros/go-jsonschema-generator.png?branch=master)](https://travis-ci.org/mcuadros/go-jsonschema-generator) [![GoDoc](http://godoc.org/github.com/mcuadros/go-jsonschema-generator?status.png)](http://godoc.org/github.com/mcuadros/go-jsonschema-generator)
==============================

json-schema generator based on Go types


Installation
------------

The recommended way to install go-jsonschema-generator

```
go get github.com/mcuadros/go-jsonschema-generator
```

Examples
--------

A basic example:

```go
import (
    "fmt"
    "github.com/mcuadros/go-jsonschema-generator"
)

type ExampleBasic struct {
    Foo bool   `json:"foo"`
    Bar string `json:",omitempty"`
    Qux int8
}

func NewExampleBasic() *ExampleBasic {
    example := new(ExampleBasic)
    SetDefaults(example) //<-- This set the defaults values

    return example
}

...

test := NewExampleBasic()
fmt.Println(test.Foo) //Prints: true
fmt.Println(test.Bar) //Prints: 33
fmt.Println(test.Qux) //Prints:

```

License
-------

MIT, see [LICENSE](LICENSE)
