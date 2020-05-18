package main

import (
	"errors"
	"fmt"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"strings"
	"time"
)

func deploy(c *cli.Context) error {

	app := c.Args().Get(0)
	stack := c.Args().Get(1)

	if app == "" {
		return errors.New("please enter stack name")
	}
	if stack == "" {
		files, err := ioutil.ReadDir("./")
		if err != nil {
			return err
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
		return errors.New("cannot find docker-compose file")
	}

	auto, unsafe := c.Bool("auto"), c.Bool("unsafe")

	if !auto && unsafe {
		return errors.New("you have to use --auto if you want to use --unsafe")
	}
	if auto {
		err := n.DeployAuto(app, stack, unsafe)

		if len(err) > 0 {
			return errors.New(strings.Join(err, "\n"))
		}
	} else {
		err := n.Deploy(app, stack)
		if len(err) > 0 {
			return errors.New(strings.Join(err, "\n"))
		}
	}
	return nil
}

func deployComplete(c *cli.Context) {
	seed := time.Now().UTC().UnixNano()
	ng := ng.NewNameGenerator(seed)
	if c.NArg() == 0 {
		for i := 0; i < 5; i++ {
			fmt.Println(ng.Generate())
		}
	}
}
