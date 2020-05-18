package main

import (
	"errors"
	"fmt"
	"github.com/urfave/cli/v2"
)

func teardown(c *cli.Context) error {
	if c.NArg() != 1 && c.NArg() != 2 {
		return errors.New("only have 1 or 2 arguments, either 'bash', 'zsh', 'powershell $PROFILE")
	}
	t := c.Args().Get(0)

	if t == "bash" {
		if c.NArg() != 1 {
			return errors.New("'teardown bash' has no other arguments")
		}
		return tearDownBash()
	} else if t == "zsh" {
		if c.NArg() != 1 {
			return errors.New("'teardown zsh' has no other arguments")
		}
		return tearDownZsh()
	} else if t == "powershell" {
		if c.NArg() != 2 {
			return errors.New("'teardown zsh $profile' has no other arguments")
		}
		profile := c.Args().Get(1)
		return tearDownPowerShell(profile)
	} else {
		return errors.New("only have 1 argument, either 'bash' or 'zsh'")
	}

}

func setup(c *cli.Context) error {
	if c.NArg() != 1 && c.NArg() != 2 {
		return errors.New("only have 1 or 2 arguments, either 'bash', 'zsh', 'powershell $PROFILE")
	}
	t := c.Args().Get(0)

	if t == "bash" {
		if c.NArg() != 1 {
			return errors.New("'setup bash' has no other arguments")
		}
		return setupBash()
	} else if t == "zsh" {
		if c.NArg() != 1 {
			return errors.New("'setup zsh' has no other arguments")
		}
		return setupZSH()
	} else if t == "powershell" {
		if c.NArg() != 2 {
			return errors.New("'setup zsh $profile' has no other arguments")
		}
		profile := c.Args().Get(1)
		return setupPowerShell(profile)
	} else {
		return errors.New("only have 1 argument, either 'bash' or 'zsh'")
	}
}

func teardownComplete(c *cli.Context) {
	if c.NArg() == 0 {
		fmt.Println("bash")
		fmt.Println("zsh")
		fmt.Println("powershell")
	} else if c.NArg() == 1 {
		if c.Args().Get(0) == "powershell" {
			fmt.Println("$profile")
		} else {
			fmt.Println("")
		}
	} else {
		fmt.Println("")
	}

}
