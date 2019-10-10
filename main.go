package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:        "load",
			Aliases:     []string{"l"},
			Description: "loads tarball into a docker volume",
			Usage:       "load <path to tar> <volume name>",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:        "save",
			Aliases:     []string{"s"},
			Description: "saves a docker volume as a tarball",
			Usage:       "save <volume name> <tar name> <path to save to>",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	app.Name = "Narwhal"
	app.Description = "A command line interface that allows you to load and save docker volumes as tarballs"
	app.Version = "0.0.1"
	app.Usage = "A docker volume helper"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Kirinnee",
			Email: "kirinnee97@gmail.com",
		},
	}

	app.EnableBashCompletion = true
	app.BashComplete = func(c *cli.Context) {
		if c.NArg() > 0 {
			return
		}
		fmt.Println("load")
		fmt.Println("save")
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
