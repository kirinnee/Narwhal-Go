package main

import (
	"fmt"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"time"
)

func run(c *cli.Context) error {
	a := c.Args()
	image, name, context, file := a.Get(0), a.Get(1), a.Get(2), a.Get(3)

	if image == "" {
		return e1("please enter an image name")
	}
	if context == "" {
		context = "."
	}
	if file == "" {
		file = "Dockerfile"
	}

	err := n.Run(context, file, image, name)
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
