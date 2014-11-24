package jsonschema

import (
	"encoding/json"
	"reflect"
	"strings"
)

const DEFAULT_SCHEMA = "http://json-schema.org/schema#"

type JSONSchema struct {
	Schema               string                 `json:"$schema,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	Items                *JSONSchemaItems       `json:"items,omitempty"`
	Properties           map[string]*JSONSchema `json:"properties,omitempty"`
	Required             []string               `json:"required,omitempty"`
	AdditionalProperties bool                   `json:"additionalProperties,omitempty"`
}

type JSONSchemaItems struct {
	Type string `json:"type,omitempty"`
}

func (j *JSONSchema) Marshal() ([]byte, error) {
	return json.MarshalIndent(j, "", "    ")
}

func (j *JSONSchema) String() string {
	json, _ := j.Marshal()
	return string(json)
}

func (j *JSONSchema) Load(variable interface{}) {
	j.setDefaultSchema()

	value := reflect.ValueOf(variable)
	j.doLoad(value.Type(), tagOptions(""))
}

func (j *JSONSchema) setDefaultSchema() {
	if j.Schema == "" {
		j.Schema = DEFAULT_SCHEMA
	}
}

func (j *JSONSchema) doLoad(t reflect.Type, opts tagOptions) {
	kind := t.Kind()

	if jsType := getTypeFromMapping(kind); jsType != "" {
		j.Type = jsType
	}

	switch kind {
	case reflect.Slice:
		j.doLoadFromSlice(t)
	case reflect.Map:
		j.doLoadFromMap(t)
	case reflect.Struct:
		j.doLoadFromStruct(t)
	case reflect.Ptr:
		j.doLoad(t.Elem(), opts)
	}
}

func (j *JSONSchema) doLoadFromSlice(t reflect.Type) {
	k := t.Elem().Kind()
	if k == reflect.Uint8 {
		j.Type = "string"
	} else {
		if jsType := getTypeFromMapping(k); jsType != "" {
			j.Items = &JSONSchemaItems{Type: jsType}
		}
	}
}

func (j *JSONSchema) doLoadFromMap(t reflect.Type) {
	k := t.Elem().Kind()

	if jsType := getTypeFromMapping(k); jsType != "" {
		j.Properties = make(map[string]*JSONSchema, 0)
		j.Properties[".*"] = &JSONSchema{Type: jsType}
	} else {
		j.AdditionalProperties = true
	}
}

func (j *JSONSchema) doLoadFromStruct(t reflect.Type) {
	j.Type = "object"
	j.Properties = make(map[string]*JSONSchema, 0)
	j.AdditionalProperties = false

	count := t.NumField()
	for i := 0; i < count; i++ {
		field := t.Field(i)

		tag := field.Tag.Get("json")
		name, opts := parseTag(tag)
		if name == "" {
			name = field.Name
		}

		j.Properties[name] = &JSONSchema{}
		j.Properties[name].doLoad(field.Type, opts)

		if !opts.Contains("omitempty") {
			j.Required = append(j.Required, name)
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
