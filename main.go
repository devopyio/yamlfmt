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

	var (
		yamlFile      *os.File
		yamlFileOut   *os.File
		err           error
		overwriteFile bool
	)
	if *file != "" {
		overwriteFile = true
		yamlFile, err = os.OpenFile(*file, os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("could not open file: %v", err)
		}

		yamlFileOut = yamlFile
		defer yamlFile.Close()
	} else {
		yamlFile = os.Stdin
		yamlFileOut = os.Stdout
	}

	var out interface{}
	if err = yaml.NewDecoder(yamlFile).Decode(&out); err != nil {
		log.Fatalf("could not decode file: %v", err)
	}

	if overwriteFile {
		if err = yamlFileOut.Truncate(0); err != nil {
			log.Fatalf("could not truncate file: %v", err)
		}
		if _, err = yamlFileOut.Seek(0, 0); err != nil {
			log.Fatalf("could not seek file: %v", err)
		}
	}

	err = yaml.NewEncoder(yamlFileOut).Encode(out)
	if err != nil {
		log.Fatalf("could not encode file: %v", err)
	}
}
