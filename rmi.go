package main

import (
	"errors"
	"fmt"
	"github.com/olekukonko/tablewriter"
	_ "github.com/olekukonko/tablewriter"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func rmi(c *cli.Context) error {
	force, filters := c.Bool("force"), c.Args().Slice()

	if !force {
		images, remain, errs := n.Images(filters...)
		if len(remain) > 0 {
			return errors.New("unknown filters: " + strings.Join(remain, ", "))
		}
		if len(errs) > 0 {
			return errors.New(strings.Join(errs, "\n"))
		}
		fmt.Println("These will be the images that are removed: ")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"NAME", "ID"})
		for _, v := range images {
			table.Append([]string{v.Name, v.Id})
		}
		table.Render()
		fmt.Print("Are you sure you want to continue? [y/N]: ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return err
		}
		if input == "y" || input == "Y" {

		} else {
			return nil
		}
	}

	err := n.RemoveImage(filters...)
	if len(err) > 0 {
		return errors.New(strings.Join(err, "\n"))
	}
	return nil
}
