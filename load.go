package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func load(c *cli.Context) error {
	if c.NArg() < 2 {
		return errors.New("need at least 2 arguments")
	} else {
		n.Load(c.Args().Get(1), c.Args().Get(0))
		return nil
	}
}

func loadComplete(c *cli.Context) {
	if c.NArg() != 1 {
		return
	}
	ls := n.Cmd.Create("docker", "volume", "ls", "-q")
	volumes := make([]string, 0, 10)
	ls.CustomRun(func(s string) {
		volumes = append(volumes, s)
	}, NOOP)
	for _, t := range volumes {
		fmt.Println(t)
	}
}
