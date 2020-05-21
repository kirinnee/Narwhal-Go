package main

import (
	"github.com/urfave/cli/v2"
)

func remove(c *cli.Context) error {

	if c.NArg() > 0 {
		return e1("too many arguments")
	}
	err := n.RemoveAll()
	if len(err) > 0 {
		return e(err)
	}
	return nil
}
