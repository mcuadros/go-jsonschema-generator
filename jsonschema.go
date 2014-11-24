package jsonschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

const DEFAULT_SCHEMA = "http://json-schema.org/schema#"

type Schema struct {
	Schema string `json:"$schema,omitempty"`
	property
}

type item struct {
	Type string `json:"type,omitempty"`
}

func (s *Schema) Marshal() ([]byte, error) {
	return json.MarshalIndent(s, "", "    ")
}

func (s *Schema) String() string {
	json, _ := s.Marshal()
	return string(json)
}

func (s *Schema) Load(variable interface{}) {
	s.setDefaultSchema()

	value := reflect.ValueOf(variable)
	s.doLoad(value.Type(), tagOptions(""))
}

func (s *Schema) setDefaultSchema() {
	if s.Schema == "" {
		s.Schema = DEFAULT_SCHEMA
	}
}

type property struct {
	Type                 string               `json:"type,omitempty"`
	Items                *item                `json:"items,omitempty"`
	Properties           map[string]*property `json:"properties,omitempty"`
	Required             []string             `json:"required,omitempty"`
	AdditionalProperties bool                 `json:"additionalProperties,omitempty"`
}

func (p *property) doLoad(t reflect.Type, opts tagOptions) {
	kind := t.Kind()

	if jsType := getTypeFromMapping(kind); jsType != "" {
		p.Type = jsType
	}

	switch kind {
	case reflect.Slice:
		p.doLoadFromSlice(t)
	case reflect.Map:
		p.doLoadFromMap(t)
	case reflect.Struct:
		p.doLoadFromStruct(t)
	case reflect.Ptr:
		p.doLoad(t.Elem(), opts)
	}
}

func (p *property) doLoadFromSlice(t reflect.Type) {
	k := t.Elem().Kind()
	if k == reflect.Uint8 {
		p.Type = "string"
	} else {
		if jsType := getTypeFromMapping(k); jsType != "" {
			p.Items = &item{Type: jsType}
		}
	}
}

func (p *property) doLoadFromMap(t reflect.Type) {
	k := t.Elem().Kind()

	if jsType := getTypeFromMapping(k); jsType != "" {
		p.Properties = make(map[string]*property, 0)
		p.Properties[".*"] = &property{Type: jsType}
	} else {
		p.AdditionalProperties = true
	}
}

func (p *property) doLoadFromStruct(t reflect.Type) {
	p.Type = "object"
	p.Properties = make(map[string]*property, 0)
	p.AdditionalProperties = false

	count := t.NumField()
	for i := 0; i < count; i++ {
		field := t.Field(i)

		tag := field.Tag.Get("json")
		name, opts := parseTag(tag)
		if name == "" {
			name = field.Name
		}

		p.Properties[name] = &property{}
		p.Properties[name].doLoad(field.Type, opts)

		if !opts.Contains("omitempty") {
			p.Required = append(p.Required, name)
		}
	}
}

var mapping = map[reflect.Kind]string{
	reflect.Bool:    "bool",
	reflect.Int:     "integer",
	reflect.Int8:    "integer",
	reflect.Int16:   "integer",
	reflect.Int32:   "integer",
	reflect.Int64:   "integer",
	reflect.Uint:    "integer",
	reflect.Uint8:   "integer",
	reflect.Uint16:  "integer",
	reflect.Uint32:  "integer",
	reflect.Uint64:  "integer",
	reflect.Float32: "number",
	reflect.Float64: "number",
	reflect.String:  "string",
	reflect.Slice:   "array",
	reflect.Struct:  "object",
	reflect.Map:     "object",
}

func getTypeFromMapping(k reflect.Kind) string {
	if t, ok := mapping[k]; ok {
		return t
	}

	return ""
}

type tagOptions string

func parseTag(tag string) (string, tagOptions) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tagOptions(tag[idx+1:])
	}
	return tag, tagOptions("")
}

func (o tagOptions) Contains(optionName string) bool {
	if len(o) == 0 {
		return false
	}

	s := string(o)
	for s != "" {
		var next string
		i := strings.Index(s, ",")
		if i >= 0 {
			s, next = s[:i], s[i+1:]
		}
		if s == optionName {
			return true
		}
		s = next
	}
	return false
}
