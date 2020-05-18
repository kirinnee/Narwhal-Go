package main

import (
	"errors"
	"github.com/urfave/cli/v2"
	"gitlab.com/kiringo/narwhal_lib"
	"log"
	"os"
	"strings"
	"time"
)

var n = narwhal_lib.New(false)

var NOOP = func(s string) {}

func main() {
	app := cli.NewApp()

	app.Commands = []*cli.Command{
		{
			Name:      "setup",
			Usage:     "enable auto complete for bash, zsh or PowerShell",
			ArgsUsage: "[shell type (bash|zsh|powershell)]",
			Action:    setup,
		},
		{
			Name:         "teardown",
			Aliases:      []string{"td"},
			Usage:        "disable auto complete for bash, zsh or PowerShell",
			ArgsUsage:    "[shell type (bash|zsh|powershell)]",
			Action:       teardown,
			BashComplete: teardownComplete,
		},
		{
			Name:      "alias",
			Aliases:   []string{"al"},
			Usage:     "manage the 'nw' alias",
			ArgsUsage: "[add|remove] [shell type (bash|zsh|powershell)]",
			Subcommands: []*cli.Command{
				{
					Name:         "add",
					Aliases:      []string{"a"},
					Usage:        "add the 'nw' alias",
					ArgsUsage:    "[shell type (bash|zsh|powershell)] [$profile(for PowerShell Users)]",
					Action:       addAlias,
					BashComplete: addAliasComplete,
				},
				{
					Name:         "remove",
					Aliases:      []string{"r"},
					Usage:        "remove the 'nw' alias",
					ArgsUsage:    "[shell type (bash|zsh|powershell)] [$profile(for PowerShell Users)]",
					Action:       removeAlias,
					BashComplete: removeAliasComplete,
				},
			},
		},
		{
			Name:         "load",
			Aliases:      []string{"l"},
			Usage:        "loads tarball into a docker volume",
			ArgsUsage:    "[path-to-tar] [volume-name]",
			Action:       load,
			BashComplete: loadComplete,
		},
		{
			Name:         "save",
			Aliases:      []string{"s"},
			Usage:        "saves a docker volume as a tarball",
			ArgsUsage:    "[volume-name] [tar-name] [path-to-save]",
			Action:       save,
			BashComplete: saveComplete,
		},
		{
			Name:    "kill",
			Aliases: []string{"k"},
			Usage:   "kills all running containers",
			Action:  kill,
		},
		{
			Name:   "stop",
			Usage:  "stop all running containers",
			Action: stop,
		},
		{
			Name:    "remove",
			Aliases: []string{"rma"},
			Usage:   "remove all containers",
			Action:  remove,
		},
		{
			Name: "deploy",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "auto",
					Aliases: []string{"a"},
					Usage:   "Automatically initialize swarm if not in swarm",
				},
				&cli.BoolFlag{
					Name:    "unsafe",
					Aliases: []string{"u"},
					Usage:   "Restart the swarm forcefully if deploy fails to try again",
				},
			},
			ArgsUsage:    "[stack-name] [stack-file(omit if stack.yml or docker-compose.yml present)]",
			Usage:        "Deploys an extended docker-compose file with images",
			Action:       deploy,
			BashComplete: deployComplete,
		},
		{
			Name:      "run",
			Aliases:   []string{"r"},
			ArgsUsage: "[image-name] [container-name(default:random)] [context(default:.)] [dockerfile(default:Dockerfile)]",
			Usage:     "builds and runs the image immediately",
			Action:    run,
		},
		{
			Name:      "image",
			Aliases:   []string{"img"},
			ArgsUsage: "[<key>=<value>] [<key>=<value>]....",
			Usage:     "get a list of images",
			Action:    image,
		},
		{
			Name:    "remove-image",
			Aliases: []string{"rmi"},
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "force",
					Aliases: []string{},
					Usage:   "do not prompt confirmation",
				},
			},
			ArgsUsage: "[<key>=<value>] [<key>=<value>]....",
			Usage:     "remove the list of images",
			Action:    rmi,
		},
		{
			Name:      "d",
			ArgsUsage: "[any docker command]",
			Usage:     "docker alias",
			Action: func(c *cli.Context) error {
				err := n.Cmd.Create("docker", c.Args().Slice()...).Run()
				if len(err) > 0 {
					return errors.New(strings.Join(err, "\n"))
				}
				return nil
			},
		},
		{
			Name:    "volume",
			Aliases: []string{"v"},
			Usage:   "manage volumes",
			Subcommands: []*cli.Command{
				{
					Name:         "Create",
					Aliases:      []string{"c"},
					Usage:        "create volumes",
					Action:       volumeCreate,
					BashComplete: volumeCreateComplete,
				},
				{
					Name:    "Remove",
					Aliases: []string{"r"},
					Usage:   "remove volumes",
					Flags: []cli.Flag{
						&cli.BoolFlag{
							Name:    "force",
							Aliases: []string{"f"},
							Usage:   "Force the removal of one or more volumes",
						},
					},
					Action:       volumeRemove,
					BashComplete: volumeRemoveComplete,
				},
			},
		},
	}
	app.EnableBashCompletion = true
	app.Name = "Narwhal"
	app.Description = "A docker utility CLI that allows you to save time"
	app.Version = "0.2.0"
	app.Usage = "Docker utilities"
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
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
