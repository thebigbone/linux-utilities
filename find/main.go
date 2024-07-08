package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "finder",
		Usage: "find utility in go",
		Commands: []*cli.Command{
			{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "directory name",
				Action: func(con *cli.Context) error {
					return searchDirectory(con)
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
