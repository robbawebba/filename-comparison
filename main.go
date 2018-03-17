package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

const rootDir string = "testDir"

func main() {
	fmt.Println("Starting search with ioutil.ReadDir")
	fileInfos, err := ioutil.ReadDir(rootDir)
	if err != nil {
		panic(fmt.Sprint("Cannot search directory: ", err.Error()))
	}

	for _, f := range fileInfos {
		fmt.Println(f.Name())
	}

	fmt.Println("Starting search with filepath.Walk")
	err = filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			fmt.Println("Searching for files in ", info.Name())
			return nil
		}
		fmt.Println(info.Name())
		return nil
	})
}
