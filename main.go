package main

import (
	"flag"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type YAMLContainer struct {
	Groups []struct {
		Name  string `yaml:"name"`
		Rules []struct {
			Alert       string `yaml:"alert,omitempty"`
			Annotations struct {
				Message string `yaml:"message"`
			} `yaml:"annotations,omitempty"`
			Expr   string `yaml:"expr"`
			For    string `yaml:"for,omitempty"`
			Labels struct {
				Severity string `yaml:"severity"`
			} `yaml:"labels,omitempty"`
			Record string `yaml:"record,omitempty"`
		} `yaml:"rules"`
	} `yaml:"groups"`
}

func main() {
	dir := flag.String("dir", "./yamlfiles", "path to yaml files")
	flag.Parse()
	filesInDir, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".yml") || strings.HasSuffix(file.Name(), ".yaml") {
			yamlFile, err := ioutil.ReadFile(filepath.Join(*dir, file.Name()))
			if err != nil {
				log.Fatal(err)
			}

			yamlsave := YAMLContainer{}

			err = yaml.Unmarshal(yamlFile, &yamlsave)
			if err != nil {
				log.Fatal(err)
			}
			out, err := yaml.Marshal(yamlsave)
			if err != nil {
				log.Fatal(err)
			}

			err = ioutil.WriteFile(filepath.Join(*dir, file.Name()), out, 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
