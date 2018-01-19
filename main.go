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
	CacheFrom           []string `yaml:"cacheFrom"`
	Context             string
	DisableContentTrust *bool `yaml:"disableContentTrust"`
	Dockerfile          string
	Labels              map[string]string
	Target              string
}

// HelmReleaseBuildConfig represents specific build configuration
type HelmReleaseBuildConfig struct {
	HelmReleaseBuild `yaml:",inline"`
	UseGitIgnore     *bool `yaml:"useGitIgnore"`
	Ignore           []string
}

// HelmReleaseContainer represents common container configuration
type HelmReleaseContainer struct {
	ExposePorts []string `yaml:"exposePorts"`
	HTTPPorts   []string `yaml:"httpPorts"`
	WSPorts     []string `yaml:"wsPorts"`
}

// HelmReleaseContainerConfig represents specific container configuration
type HelmReleaseContainerConfig struct {
	HelmReleaseContainer `yaml:",inline"`
	Args                 []string
	Command              []string
	Env                  map[string]string
	Imports              []string
	Sync                 []string
	SyncTarget           string `yaml:"syncTarget"`
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
	var filename = os.Args[1]
	bytes, _ := ioutil.ReadFile(filename)

	tmpl, err := template.New("template").Parse(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

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
