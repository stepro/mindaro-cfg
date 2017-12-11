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
	Kind string
}

// HelmReleaseBuild represents common build configuration
type HelmReleaseBuild struct {
	Args                map[string]string
	CacheFrom           string `yaml:"cacheFrom"`
	Context             string
	DisableContentTrust *bool `yaml:"disableContentTrust"`
	Dockerfile          string
	Labels              map[string]string
	Target              string
}

// HelmReleaseBuildConfig represents specific build configuration
type HelmReleaseBuildConfig struct {
	HelmReleaseBuild `yaml:",inline"`
	Ignore           []string
}

// HelmReleaseContainer represents common container configuration
type HelmReleaseContainer struct {
	ExposePorts []string `yaml:"exposePorts"`
	HTTPPorts   []string `yaml:"httpPorts"`
}

// HelmReleaseContainerConfig represents specific container configuration
type HelmReleaseContainerConfig struct {
	HelmReleaseContainer `yaml:",inline"`
	Command              []string
	Entrypoint           []string
	Env                  map[string]string
	Imports              map[string]string
	Sync                 []string
	SyncTarget           string `yaml:"syncTarget"`
	Unchecked            *bool
	Workdir              string
}

// HelmReleaseInstall represents common install configuration
type HelmReleaseInstall struct {
	Chart  string
	Values []string
	Set    map[string]interface{}
}

// HelmReleaseInstallConfig represents specific install configuration
type HelmReleaseInstallConfig struct {
	HelmReleaseInstall `yaml:",inline"`
}

// HelmReleaseDocument represents a configuration file for a helm-release workload
type HelmReleaseDocument struct {
	Document       `yaml:",inline"`
	Build          HelmReleaseBuild
	Container      HelmReleaseContainer
	Install        HelmReleaseInstall
	Configurations map[string]struct {
		Build     HelmReleaseBuildConfig
		Container HelmReleaseContainerConfig
		Install   HelmReleaseInstallConfig
	}
}

func main() {
	tmpl, err := template.New("template").Parse(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	var filename string
	if len(os.Args) > 2 {
		filename = os.Args[2]
	} else {
		filename = "mindaro.yaml"
	}
	bytes, _ := ioutil.ReadFile(filename)

	var doc Document
	err = yaml.Unmarshal(bytes, &doc)
	if err != nil {
		log.Fatal(err)
	}
	if doc.Kind == "" {
		log.Fatal("Missing kind")
	}

	if doc.Kind == "helm-release" {
		var doc HelmReleaseDocument
		err = yaml.UnmarshalStrict(bytes, &doc)
		if err != nil {
			log.Fatal(err)
		}
		err = tmpl.Execute(os.Stdout, doc)
	} else {
		log.Fatal("Unknown kind '" + doc.Kind + "'")
	}
	if err != nil {
		log.Fatal(err)
	}
}
