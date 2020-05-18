package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"strings"
)

func remove(c *cli.Context) error {

	if c.NArg() > 0 {
		return errors.New("too many arguments")
	}
	err := n.RemoveAll()
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil
}
