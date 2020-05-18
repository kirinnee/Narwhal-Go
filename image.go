package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strings"
)

func image(c *cli.Context) error {
	filters := c.Args().Slice()

	images, remain, errs := n.Images(filters...)
	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "\n"))
	}

	args := []string{
		"images",
	}
	for _, v := range images {
		filter := []string{
			"-f",
			"reference=" + v.Name,
		}
		args = append(args, filter...)
	}
	args = append(args, remain...)

	err := n.Cmd.Create("docker", args...).Run()
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil

}
