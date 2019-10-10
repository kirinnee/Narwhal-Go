package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:      "load",
			Aliases:   []string{"l"},
			Usage:     "loads tarball into a docker volume",
			ArgsUsage: "[path-to-tar] [volume-name]",
			Action: func(c *cli.Context) error {
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:      "save",
			Aliases:   []string{"s"},
			Usage:     "saves a docker volume as a tarball",
			ArgsUsage: "[volume-name] [tar-name] [path-to-save]",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	app.Name = "Narwhal"
	app.Description = "A command line interface that allows you to load and save docker volumes as tarballs"
	app.Version = "0.0.1"
	app.Usage = "A docker volume helper"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		{
			Name:  "Kirinnee",
			Email: "kirinnee97@gmail.com",
		},
	}

	cli.AppHelpTemplate = `{{.Name}} - {{.Usage}}
VERSION: {{.Version}}
USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}

   {{end}}
`

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
