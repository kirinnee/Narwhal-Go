package main

import (
	"errors"
	"fmt"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"strings"
	"time"
)

func volumeCreate(c *cli.Context) error {
	name := c.Args().Get(0)
	if name == "" {
		return errors.New("please enter volume name")
	}
	remain := c.Args().Tail()
	args := []string{"volume", "create"}
	args = append(args, name)
	args = append(args, remain...)
	err := n.Cmd.Create("docker", args...).Run()
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil
}

func volumeCreateComplete(c *cli.Context) {
	if c.NArg() < 1 {
		seed := time.Now().UTC().UnixNano()
		ng := ng.NewNameGenerator(seed)
		if c.NArg() == 0 {
			for i := 0; i < 5; i++ {
				fmt.Println(ng.Generate())
			}
		}
	}
}
