package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
)

func deploy(c *cli.Context) error {

	app := c.String("name")
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
	if stack == "" {
		return e1("cannot find docker-compose file")
	}

	auto, unsafe := c.Bool("auto"), c.Bool("unsafe")

	if !auto && unsafe {
		return e1("you have to use --auto if you want to use --unsafe")
	}
	if auto {
		err := n.DeployAuto(app, stack, unsafe)

		if len(err) > 0 {
			return e(err)
		}
	} else {
		err := n.Deploy(app, stack)
		if len(err) > 0 {
			return e(err)
		}
	}
	return nil
}
