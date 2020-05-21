package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

func stopStack(c *cli.Context) error {
	app := c.String("name")
	auto := c.Bool("auto")
	stack := c.Args().Get(0)

	if stack == "" {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			return ee(err)
		}
		for _, f := range files {
			fn := f.Name()
			if fn == "stack.yml" {
				stack = fn
				fmt.Println("Using " + fn)
				break
			}

			if fn == "docker-compose.yml" {
				stack = fn
				fmt.Println("Using " + fn)
				break
			}
		}
	}
	err := n.StopStack(app, stack)
	if len(err) > 0 {
		return e(err)
	}

	if auto {
		return e(n.Cmd.Create("docker", "swarm", "leave", "--force").Run())
	}
	return nil

}
