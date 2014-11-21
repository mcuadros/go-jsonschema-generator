package jsonschema

import (
	"testing"

	. "launchpad.net/gocheck"
)

func Test(t *testing.T) { TestingT(t) }

type JSONSchemaSuite struct{}

var _ = Suite(&JSONSchemaSuite{})

type ExampleJSONBasic struct {
	Bool       bool    `json:",omitempty"`
	Integer    int     `json:",omitempty"`
	Integer8   int8    `json:",omitempty"`
	Integer16  int16   `json:",omitempty"`
	Integer32  int32   `json:",omitempty"`
	Integer64  int64   `json:",omitempty"`
	UInteger   uint    `json:",omitempty"`
	UInteger8  uint8   `json:",omitempty"`
	UInteger16 uint16  `json:",omitempty"`
	UInteger32 uint32  `json:",omitempty"`
	UInteger64 uint64  `json:",omitempty"`
	String     string  `json:",omitempty"`
	Bytes      []byte  `json:",omitempty"`
	Float32    float32 `json:",omitempty"`
	Float64    float64
}

func (self *JSONSchemaSuite) TestLoad(c *C) {
	j := &JSONSchema{}
	j.Load(&ExampleJSONBasic{})

	c.Assert(*j, DeepEquals, JSONSchema{
		Type:     "object",
		Required: []string{"Float64"},
		Properties: map[string]*JSONSchema{
			"Bool":       &JSONSchema{Type: "bool"},
			"Integer":    &JSONSchema{Type: "integer"},
			"Integer8":   &JSONSchema{Type: "integer"},
			"Integer16":  &JSONSchema{Type: "integer"},
			"Integer32":  &JSONSchema{Type: "integer"},
			"Integer64":  &JSONSchema{Type: "integer"},
			"UInteger":   &JSONSchema{Type: "integer"},
			"UInteger8":  &JSONSchema{Type: "integer"},
			"UInteger16": &JSONSchema{Type: "integer"},
			"UInteger32": &JSONSchema{Type: "integer"},
			"UInteger64": &JSONSchema{Type: "integer"},
			"String":     &JSONSchema{Type: "string"},
			"Bytes":      &JSONSchema{Type: "string"},
			"Float32":    &JSONSchema{Type: "number"},
			"Float64":    &JSONSchema{Type: "number"},
		},
	})
}

type ExampleJSONBasicWithTag struct {
	Bool bool `json:"test"`
}

func (self *JSONSchemaSuite) TestLoadWithTag(c *C) {
	j := &JSONSchema{}
	j.Load(&ExampleJSONBasicWithTag{})

	c.Assert(*j, DeepEquals, JSONSchema{
		Type:     "object",
		Required: []string{"test"},
		Properties: map[string]*JSONSchema{
			"test": &JSONSchema{Type: "bool"},
		},
	})

}

type ExampleJSONBasicSlices struct {
	Strings []string `json:",omitempty"`
}

func (self *JSONSchemaSuite) TestLoadSlice(c *C) {
	j := &JSONSchema{}
	j.Load(&ExampleJSONBasicSlices{})

	c.Assert(*j, DeepEquals, JSONSchema{
		Type: "object",
		Properties: map[string]*JSONSchema{
			"Strings": &JSONSchema{
				Type:  "array",
				Items: &JSONSchemaItems{Type: "string"},
			},
		},
	})
}

type ExampleJSONBasicMaps struct {
	Maps map[string]string `json:",omitempty"`
}

func (self *JSONSchemaSuite) TestLoadMap(c *C) {
	j := &JSONSchema{}
	j.Load(&ExampleJSONBasicMaps{})

	c.Assert(*j, DeepEquals, JSONSchema{
		Type: "object",
		Properties: map[string]*JSONSchema{
			"Maps": &JSONSchema{
				Type: "object",
				Properties: map[string]*JSONSchema{
					".*": &JSONSchema{Type: "string"},
				},
			},
		},
	})
}
