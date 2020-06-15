package main

import (
	"fmt"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"log"
	"time"
)

func run(c *cli.Context) error {
	a := c.Args()
	image, name, context, file, args := c.String("image"), c.String("name"), c.String("context"), c.String("file"), a.Slice()

	ba := c.StringSlice("build-args")
	env := c.StringSlice("environment")
	vol := c.StringSlice("volume")
	net := c.String("network")

	rest := make([]string, 0, len(args))
	cmd := make([]string, 0, len(args))

	for _, v := range ba {
		rest = append(rest, "b:--build-arg")
		rest = append(rest, "b:"+v)

	}
	for _, v := range env {
		rest = append(rest, "r:-e")
		rest = append(rest, "r:"+v)

	}

	for _, v := range vol {
		rest = append(rest, "r:-v")
		rest = append(rest, "r:"+v)

	}

	log.Println("net", net)
	if net != "" {
		rest = append(rest, "r:--net")
		rest = append(rest, "r:"+net)
	}

	after := false

	for _, v := range args {
		if v == "---" {
			after = true
			continue
		}
		if after {
			rest = append(rest, v)
		} else {
			cmd = append(cmd, v)
		}
	}

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
