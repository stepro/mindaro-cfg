package main

import (
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Document represents a configuration file
type Document struct {
	Workload string
}

// HelmReleaseContainer represents container configuration for the helm-release workload
type HelmReleaseContainer struct {
	Build struct {
		Args       map[string]string
		Context    string
		Dockerfile string
		Labels     map[string]string
		Target     string
	}
	Command    []string
	Entrypoint []string
	Env        map[string]string
	Imports map[string]string
	Ports   map[string]string
	Sync struct {
		Enabled *bool
		Target string
	}
	Tracing *bool
	Routing struct {
		ToMainline *bool `yaml:"toMainline"`
		ToSideline *bool `yaml:"toSideline"`
	}
	HttpPorts struct {
		Include []string
		Exclude []string
	} `yaml:"httpPorts"`
}

// HelmReleaseProperties represents configuration for the helm-release workload
type HelmReleaseProperties struct {
	HelmReleaseContainer `yaml:",inline"`
	Containers map[string]HelmReleaseContainer
	Install struct {
		Chart string
		Values []string
		Set map[string]interface{}
		Wait *bool
	}
	SkipHealthChecks *bool `yaml:"skipHealthChecks"`
	KillOnTerminate  *bool `yaml:"killOnTerminate"`
}

// HelmReleaseDocument represents a configuration file for a helm-release workload
type HelmReleaseDocument struct {
	Workload string
	Properties     HelmReleaseProperties
	Configurations map[string]HelmReleaseProperties
}

func main() {
	tmpl, err := template.New("template").Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	bytes, _ := ioutil.ReadFile(os.Args[2])

	var doc Document
	err = yaml.Unmarshal(bytes, &doc)
	if err != nil {
		log.Fatal(err)
	}
	if doc.Workload == "" {
		log.Fatal("Missing workload")
	}

	if doc.Workload == "helm-release" {
		var doc HelmReleaseDocument
		err = yaml.UnmarshalStrict(bytes, &doc)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(os.Stdout, doc)
	} else {
		log.Fatal("Unknown workload '" + doc.Workload + "'")
	}
	if err != nil {
		log.Fatal(err)
	}
}
