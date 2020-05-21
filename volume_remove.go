package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func volumeRemove(c *cli.Context) error {
	force, names := c.Bool("force"), c.Args().Slice()
	args := []string{"volume", "remove"}
	if force {
		args = append(args, "-f")
	}
	for _, v := range names {
		args = append(args, v)
	}
	err := n.Cmd.Create("docker", args...).Run()
	if len(err) > 0 {
		return e(err)
	}
	return nil
}

func volumeRemoveComplete(c *cli.Context) {
	volumes := n.Cmd.Create("docker", "volume", "ls", "-q").Run()
	for _, v := range volumes {
		fmt.Println(v)
	}

}
