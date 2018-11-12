package main

import (
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"log"
	_ "statik"
)

func main() {
	statikFS, err := fs.New()
	fs.Walk()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := fs.ReadFile(statikFS, "/Dockerfile")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("microTemplate", bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
