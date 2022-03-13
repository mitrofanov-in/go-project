package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var path string
var pirs []string

func tree(path string) []string {
	var dirs []string
	//var files []string

	tmp, _ := ioutil.ReadDir(path)
	for _, t := range tmp {
		if t.IsDir() {
			dirs = append(dirs, t.Name())
		} else {
			origName := filepath.Join(path, t.Name())
			newName := strings.ReplaceAll(origName, "_", " ")
			os.Rename(origName, newName)
		}

		for _, dirName := range dirs {
			tree(path + "/" + dirName)
			//treeF(path + "/" + dirName + "/") //рекурсивный вызов (переход в следующую папку текущей директории)
			pirs = append(pirs, dirName)
		}
	}

	return pirs
}

/*
func treeF(path string) []string {
	//var dirs []string
	var files []string
	tmp, _ := ioutil.ReadDir(path)
	for _, t := range tmp {
		if !t.IsDir() {
			files = append(files, t.Name()+" ("+strconv.Itoa(int(t.Size()))+"b)")

		}

	}
	return files
}
*/

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		path = "good"
	} else {
		path = arguments[1]
	}

	err := tree(path)
	fmt.Println(err)

	/*
		for _, file := range files {
			if !file.IsDir() {
				origName := filepath.Join(e_string, file.Name())
				newName := strings.ReplaceAll(origName, "_", " ")
				os.Rename(origName, newName)
			} else {

					}
				}
	*/

	/*
			tmpName := filepath.Join(e_string, file.Name())
		    if err != nil {
				log.Fatal(err)
			}
			for _, tFile := range tmp {
				origName := filepath.Join(tmpName, tFile.Name())
				newName := strings.ReplaceAll(origName, "_", " ")
				os.Rename(origName, newName)
			}
	*/

}
