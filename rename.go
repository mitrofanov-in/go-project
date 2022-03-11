package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	e_string := ""
	arguments := os.Args

	if len(arguments) == 1 {
		e_string = "good"
	} else {
		e_string = arguments[1]
	}

	files, err := ioutil.ReadDir(e_string)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if !file.IsDir() {

			origName := filepath.Join(e_string, file.Name())
			newName := strings.ReplaceAll(origName, "_", " ")
			os.Rename(origName, newName)
		}
	}
}
