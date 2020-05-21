package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

const BASH_ALIAS = "\nalias nw='narwhal'"
const ZSH_ALIAS = "\nalias nw='narwhal'"
const POWERSHELL_ALIAS = "\nSet-Alias -Name nw -Value narwhal"

func removeAlias(c *cli.Context) error {
	shell := c.Args().Get(0)
	if shell == "bash" {
		f, err := homeFile(".bashrc")
		if err != nil {
			return ee(err)
		}
		return removeFromFile(f, BASH_ALIAS)

	} else if shell == "zsh" {
		f, err := homeFile(".zshrc")
		if err != nil {
			return ee(err)
		}
		return removeFromFile(f, ZSH_ALIAS)
	} else if shell == "powershell" {
		profile := c.Args().Get(1)
		return removeFromFile(profile, POWERSHELL_ALIAS)
	} else {
		return e1("unknown command")
	}

}

func addAlias(c *cli.Context) error {
	shell := c.Args().Get(0)
	err := removeAlias(c)
	if err != nil {
		return err
	}
	if shell == "bash" {
		err := appendTo(".bashrc", BASH_ALIAS, true)
		if err != nil {
			return err
		}

		fmt.Println("Please either restart your shell or run:")
		fmt.Println("\tsource ~/.bashrc")
		return nil
	} else if shell == "zsh" {
		err := appendTo(".zshrc", ZSH_ALIAS, true)
		if err != nil {
			return err
		}

		fmt.Println("Please either restart your shell or run:")
		fmt.Println("\tsource ~/.zshrc")
		return nil
	} else if shell == "powershell" {
		profile := c.Args().Get(1)
		err := appendTo(profile, POWERSHELL_ALIAS, false)
		if err != nil {
			return err
		}

		fmt.Println("Please either restart your shell or run:")
		fmt.Println("\t& $profile")
		return nil
	} else {
		return e1("unknown command")
	}

}

func addAliasComplete(c *cli.Context) {
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

func removeAliasComplete(c *cli.Context) {
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
