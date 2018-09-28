package main

import (
	"flag"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	file := flag.String("filename", "", "filename")

	flag.Parse()

	yamlFile, err := os.OpenFile(*file, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer yamlFile.Close()

	var out interface{}
	if err = yaml.NewDecoder(yamlFile).Decode(&out); err != nil {
		log.Fatal(err)
	}

	if err = yamlFile.Truncate(0); err != nil {
		log.Fatal(err)
	}
	if _, err = yamlFile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	err = yaml.NewEncoder(yamlFile).Encode(out)
	if err != nil {
		log.Fatal(err)
	}
}
