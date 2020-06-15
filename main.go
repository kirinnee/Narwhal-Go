package main

import (
	"fmt"
	"github.com/fatih/color"
	ng "github.com/goombaio/namegenerator"
	"github.com/urfave/cli/v2"
	"gitlab.com/kiringo/narwhal_lib"
	"log"
	"math/rand"
	"os"
	"strconv"
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
			Name:    "deploy",
			Aliases: []string{"up"},
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
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "stack name to remove",
				},
			},
			ArgsUsage: "[stack-file(omit if stack.yml or docker-compose.yml present)]",
			Usage:     "Deploys an extended docker-compose file with images",
			Action:    deploy,
		},
		{
			Name: "run",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "context",
					Aliases: []string{"c"},
					Usage:   "context folder of the docker build. Default: .",
				},
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "Dockerfile to use relative to the context. Default: Dockerfile",
				},
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "Name of the container created. Default: docker-daemon random generation",
				},
				&cli.StringFlag{
					Name:    "image",
					Aliases: []string{"i"},
					Usage:   "image name created. Default: random generation",
				},
				&cli.StringSliceFlag{
					Name:    "build-args",
					Aliases: []string{"ba"},
					Usage:   "build arguments",
				},
				&cli.StringSliceFlag{
					Name:    "environment",
					Aliases: []string{"e"},
					Usage:   "run environment",
				},
				&cli.StringSliceFlag{
					Name:    "volume",
					Aliases: []string{"v"},
					Usage:   "run volume",
				},
				&cli.StringFlag{
					Name:    "network",
					Aliases: []string{"net"},
					Usage:   "run network to connect to",
				},
			},
			Aliases:   []string{"r"},
			ArgsUsage: "[command(default: Dockerfile definition)] [addition docker flags....]",
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
			Name:    "copy",
			Aliases: []string{"cp"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "file",
					Aliases: []string{"f"},
					Usage:   "dockerfile name to use, relative to context. Default: Dockerfile",
				},
				&cli.StringFlag{
					Name:    "context",
					Aliases: []string{"c"},
					Usage:   "docker context to use. Default: .",
				},
			},
			ArgsUsage: "[file from container] [file to host] [command to execute(default:sh)]",
			Usage:     "copy a file out of a static image after executing a command",
			Action: func(ctx *cli.Context) error {
				seed := time.Now().UTC().UnixNano()
				ng := ng.NewNameGenerator(seed)
				s1 := rand.NewSource(seed)
				r := rand.New(s1)

				from, to, cmd, context, file := ctx.Args().Get(0), ctx.Args().Get(1), ctx.Args().Get(2), ctx.String("context"), ctx.String("file")
				if from == "" {
					return e1("please enter file to copy from")
				}
				if to == "" {
					return e1("please enter file to copy to")
				}
				if cmd == "" {
					cmd = "sh"
				}
				if context == "" {
					context = "."
				}
				if file == "" {
					file = "."
				}
				image := ng.Generate() + ":" + strconv.Itoa(r.Int()) + "-" + strconv.Itoa(r.Int()) + "-" + strconv.Itoa(r.Int())

				err := n.MoveOut(context, file, image, from, to, cmd)
				if len(err) > 0 {
					return e(err)
				}
				return nil

			},
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
			Name:    "undeploy",
			Aliases: []string{"down"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "name",
					Aliases: []string{"n"},
					Usage:   "stack name to remove",
				},
				&cli.BoolFlag{
					Name:    "auto",
					Aliases: []string{"a"},
					Usage:   "automatically leaves the swarm by force",
				},
			},
			ArgsUsage: "[stack-file(omit if stack.yml or docker-compose.yml present)]",
			Usage:     "remove the list of images",
			Action:    stopStack,
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
	app.UseShortOptionHandling = true
	app.EnableBashCompletion = true
	app.Name = "Narwhal"
	app.Description = "A docker utility CLI that allows you to save time"
	app.Version = "0.3.1r3"
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

func e(s []string) error {
	for _, v := range s {
		_, _ = fmt.Fprintf(color.Output, color.RedString(v)+"\n")
	}

	return cli.Exit("", 1)
}

func e1(s string) error {
	return e([]string{s})
}

func ee(err error) error {
	return e1(err.Error())
}
