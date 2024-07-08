package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func searchDirectory(con *cli.Context) error {
	directory := con.Args().First()

	dirs, err := ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
	}

	recurseDir(directory, dirs)
	return nil
}

func recurseDir(path string, dirs []fs.FileInfo) error {
	for _, f := range dirs {
		if f.IsDir() {
			subPath := filepath.Join(path, f.Name())
			fmt.Println(subPath)

			subdirs, err := ioutil.ReadDir(subPath)
			if err != nil {
				log.Fatal(err)
			}

			err = recurseDir(subPath, subdirs)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return nil
}
