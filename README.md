go-jsonschema-generator [![Build Status](https://img.shields.io/github/workflow/status/mcuadros/go-jsonschema-generator/Test.svg)](https://github.com/mcuadros/go-jsonschema-generator/actions) [![GoDoc](http://godoc.org/github.com/mcuadros/go-jsonschema-generator?status.png)](https://pkg.go.dev/github.com/mcuadros/go-jsonschema-generator)
==============================

Basic [json-schema](http://json-schema.org/) generator based on Go types, for easy interchange of Go structures across languages.


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

type EmbeddedType struct {
  Zoo string
}

type Item struct {
  Value string
}

type ExampleBasic struct {
  Foo bool   `json:"foo"`
  Bar string `json:",omitempty"`
  Qux int8
  Baz []string
  EmbeddedType
  List []Item
}

func main() {
  s := &jsonschema.Document{}
  s.Read(&ExampleBasic{})
  fmt.Println(s)
}
```

```json
{
    "$schema": "http://json-schema.org/schema#",
    "type": "object",
    "properties": {
        "Bar": {
            "type": "string"
        },
        "Baz": {
            "type": "array",
            "items": {
                "type": "string"
            }
        },
        "List": {
            "type": "array",
            "items": {
                "type": "object",
                "properties": {
                    "Value": {
                        "type": "string"
                    }
                },
                "required": [
                    "Value"
                ]
            }
        },
        "Qux": {
            "type": "integer"
        },
        "Zoo": {
            "type": "string"
        },
        "foo": {
            "type": "boolean"
        }
    },
    "required": [
        "foo",
        "Qux",
        "Baz",
        "Zoo",
        "List"
    ]
}
```

License
-------

MIT, see [LICENSE](LICENSE)
