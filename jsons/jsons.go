package jsons

import (
	"regexp"
)

// GenericKeywords structure
type GenericKeywords struct {
	Title       string
	Description string
	Default     interface{}
	Examples    []interface{}
}

// ArrayKeywords structure
type ArrayKeywords struct {
	Items           *Schema
	ItemsList       []Schema
	AdditionalItems *Schema
	Contains        *Schema
	MinItems        int
	MaxItems        int
	UniqueItems     bool
}

// NumericKeywords structure
type NumericKeywords struct {
	MultipleOf       int
	Minimum          int
	ExclusiveMinimum int
	Maximum          int
	ExclusiveMaximum int
}

// ObjectKeywords structure
type ObjectKeywords struct {
	Properties           map[string]Schema
	Required             []string
	AdditionalProperties *Schema
	PropertyNames        *StringSchema
	MinProperties        int
	MaxProperties        int
	Dependencies         map[string]struct {
		Properties []string
		Schema     *ObjectSchema
	}
	PatternProperties map[string]Schema
}

// StringFormats
const (
	StringFormatDateTime     = "date-time"
	StringFormatEmail        = "email"
	StringFormatHostName     = "hostname"
	StringFormatIPV4         = "ipv4"
	StringFormatIPV6         = "ipv6"
	StringFormatJSONPointer  = "json-pointer"
	StringFormatURI          = "uri"
	StringFormatURIReference = "uri-reference"
	StringFormatURITemplate  = "uri-template"
)

// StringKeywords structure
type StringKeywords struct {
	MinLength int
	MaxLength int
	Pattern   *regexp.Regexp
	Format    string
}

// GenericValueKeywords structure
type GenericValueKeywords struct {
	Const interface{}
	Enum  []interface{}
}

// ArraySchema structure
type ArraySchema struct {
	GenericKeywords
	ArrayKeywords
	GenericValueKeywords
	AllOf []ArraySchema
	AnyOf []ArraySchema
	OneOf []ArraySchema
	Not   *Schema
}

func ParseArraySchema(o interface{}) ArraySchema {
	return ArraySchema{}
}

// NumericSchema structure
type NumericSchema struct {
	GenericKeywords
	NumericKeywords
	GenericValueKeywords
	AllOf []NumericSchema
	AnyOf []NumericSchema
	OneOf []NumericSchema
	Not   *Schema
}

func ParseNumericSchema(o interface{}) NumericSchema {
	return NumericSchema{}
}

// ObjectSchema structure
type ObjectSchema struct {
	GenericKeywords
	ObjectKeywords
	GenericValueKeywords
	AllOf []ObjectSchema
	AnyOf []ObjectSchema
	OneOf []ObjectSchema
	Not   *Schema
}

func ParseObjectSchema(o interface{}) ObjectSchema {
	return ObjectSchema{}
}

// StringSchema structure
type StringSchema struct {
	GenericKeywords
	ObjectKeywords
	GenericValueKeywords
	AllOf []StringSchema
	AnyOf []StringSchema
	OneOf []StringSchema
	Not   *Schema
}

func ParseStringSchema(o interface{}) StringSchema {
	return StringSchema{}
}

// Types
const (
	TypeArray   = "array"
	TypeBoolean = "boolean"
	TypeInteger = "integer"
	TypeNull    = "null"
	TypeNumber  = "number"
	TypeObject  = "object"
	TypeString  = "string"
)

// Schema structure
type Schema struct {
	GenericKeywords
	Types []string
	ArrayKeywords
	NumericKeywords
	ObjectKeywords
	StringKeywords
	GenericValueKeywords
	AllOf []Schema
	AnyOf []Schema
	OneOf []Schema
	Not   *Schema
}

// ParseSchema parses a JSON schema
func ParseSchema(o interface{}) Schema {
	var schema Schema
	properties := o.(map[string]interface{})

	schemaType := properties["type"].(string)
	if schemaType == TypeArray {
		parseArraySchema()
	}
	if schemaType == TypeBoolean {

	}
	if schemaType == TypeInteger {

	}
	if schemaType == TypeNull {

	}
	if schemaType == TypeNumber {

	}
	if schemaType == TypeObject {

	}
	if schemaType == TypeString {

	}
	if schemaType == nil {

	}
	return schema
}
