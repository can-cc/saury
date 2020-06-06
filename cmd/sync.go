package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

const DefaultGalleries = "./galleries"

func main() {
	args:= os.Args
	var galleriesDir string
	if len(args) < 2 {
		galleriesDir = DefaultGalleries
	} else {
		galleriesDir = args[1]
	}
	entries, err := ioutil.ReadDir(galleriesDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range entries {
		if !f.IsDir() {
			continue
		}
		fmt.Println(f.Name())
		photos, err := ioutil.ReadDir(path.Join(galleriesDir, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		for _, p := range photos {
			if p.IsDir() {
				continue
			}
			fmt.Println(p.Name())
		}

	}
}