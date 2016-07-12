package main

import (
	"github.com/codegangsta/cli"
	"github.com/piokaczm/godeploy/commands"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "godeploy"
	app.Version = "1.0.2"
	// add more precise description, add some better help text for deploy
	app.Usage = "Simply deploy binaries to remote server"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "build config directory and config file",
			Action: commands.Config,
		},
		{
			Name:   "deploy",
			Usage:  "build and deploy binary to remote server",
			Action: commands.Deploy,
		},
	}

	app.Run(os.Args)
}
