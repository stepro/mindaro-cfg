package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
	"github.com/stepro/mindaro-cfg/jsons"
)

// TypeSpec structure
type TypeSpec struct {
	Name string
	jsons.ObjectSchema
	Exports map[string]struct {
		Usage      *jsons.Schema
		Properties map[string]jsons.Schema
	}
}

// ComponentSpec structure
type ComponentSpec struct {
}

func parseTypeSpec(raw interface{}) TypeSpec {
	var typeSpec TypeSpec
	properties := raw.(map[string]interface{})
	typeSpec.Name = properties["name"].(string)
	typeSpec.
	return typeSpec
}

func main() {
	var kind = os.Args[1]
	var filename = os.Args[2]
	bytes, _ := ioutil.ReadFile(filename)
	var templateString string
	if len(os.Args) < 4 {
		templateString = "{{.}}"
	} else {
		templateString = os.Args[3]
	}

	parsedTemplate, err := template.New("template").Parse(templateString)
	if err != nil {
		log.Fatal(err)
	}

	var spec interface{}
	if kind == "type" {
		var typeSpec TypeSpec
		var rawSpec interface{}
		err = yaml.Unmarshal(bytes, &rawSpec)
		typeSpec = parseTypeSpec(rawSpec)
		spec = typeSpec
	}
	if kind == "component" {
		var componentSpec ComponentSpec
		var rawSpec interface{}
		err = yaml.Unmarshal(bytes, &rawSpec)
		spec = componentSpec
	}
	if err != nil {
		log.Fatal(err)
	}

	err = parsedTemplate.Execute(os.Stdout, spec)
	if err != nil {
		log.Fatal(err)
	}
}
