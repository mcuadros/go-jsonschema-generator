go-jsonschema-generator [![Build Status](https://travis-ci.org/mcuadros/go-jsonschema-generator.png?branch=master)](https://travis-ci.org/mcuadros/go-jsonschema-generator) [![GoDoc](http://godoc.org/github.com/mcuadros/go-jsonschema-generator?status.png)](http://godoc.org/github.com/mcuadros/go-jsonschema-generator)
==============================

[json-schema](http://json-schema.org/) generator based on Go types


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

```

```json
{
  "type": "object",
  "properties": {
    "Bar": {
      "type": "string"
    },
    "Qux": {
      "type": "integer"
    },
    "foo": {
      "type": "bool"
    }
  },
  "required": [
    "foo",
    "Qux"
  ]
}
```

License
-------

MIT, see [LICENSE](LICENSE)
