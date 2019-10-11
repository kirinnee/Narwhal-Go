package main

import (
	"errors"
	"github.com/urfave/cli"
	"gitlab.com/kiringo/narwhal_lib"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	n := narwhal_lib.Narwhal{}
	app.Commands = []cli.Command{
		{
			Name:      "load",
			Aliases:   []string{"l"},
			Usage:     "loads tarball into a docker volume",
			ArgsUsage: "[path-to-tar] [volume-name]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					return errors.New("need at least 2 arguments")
				} else {
					n.Load(c.Args()[1], c.Args()[0])
					return nil
				}
			},
		},
		{
			Name:      "save",
			Aliases:   []string{"s"},
			Usage:     "saves a docker volume as a tarball",
			ArgsUsage: "[volume-name] [tar-name] [path-to-save]",
			Action: func(c *cli.Context) error {
				l := len(c.Args())
				var volume string
				tarName := "data"
				path := "./"
				if l == 0 {
					return errors.New("need at least 1 arguments")
				}
				if l > 0 {
					volume = c.Args()[0]
				}
				if l > 1 {
					tarName = c.Args()[1]
				}
				if l > 2 {
					path = c.Args()[2]
				}
				n.Save(volume, tarName, path)
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
