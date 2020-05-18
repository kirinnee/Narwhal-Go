package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func save(c *cli.Context) error {
	l := c.NArg()
	var volume string
	tarName := "data"
	path := "./"
	if l == 0 {
		return errors.New("need at least 1 arguments")
	}
	if l > 0 {
		volume = c.Args().Get(0)
	}
	if l > 1 {
		tarName = c.Args().Get(1)
	}
	if l > 2 {
		path = c.Args().Get(2)
	}
	n.Save(volume, tarName, path)
	return nil
}

func saveComplete(c *cli.Context) {
	// This will complete if no args are passed

	if c.NArg() == 1 {
		fmt.Println("data")
	}

	if c.NArg() == 2 {
		fmt.Println("./")
	}

	if c.NArg() == 0 {
		ls := n.Cmd.Create("docker", "volume", "ls", "-q")
		volumes := make([]string, 0, 10)
		ls.CustomRun(func(s string) {
			volumes = append(volumes, s)
		}, func(s string) {

		})
		for _, t := range volumes {
			fmt.Println(t)
		}
	}

}
