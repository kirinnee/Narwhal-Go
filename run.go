package main

import (
	"fmt"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"time"
)

func run(c *cli.Context) error {
	a := c.Args()
	image, name, context, file, cmd, rest := c.String("image"), c.String("name"), c.String("context"), c.String("file"), a.First(), a.Tail()

	if image == "" {
		seed := time.Now().UTC().UnixNano()
		ng := ng.NewNameGenerator(seed)
		image = ng.Generate()
	}
	if context == "" {
		context = "."
	}
	if file == "" {
		file = "Dockerfile"
	}
	if cmd == "-" {
		cmd = ""
	}

	err := n.Run(context, file, image, name, cmd, rest)
	if len(err) > 0 {
		return e(err)
	}
	return nil

}

func runComplete(c cli.Context) {
	if c.NArg() < 2 {
		seed := time.Now().UTC().UnixNano()
		ng := ng.NewNameGenerator(seed)
		if c.NArg() == 0 {
			for i := 0; i < 5; i++ {
				fmt.Println(ng.Generate())
			}
		}
	}
}
