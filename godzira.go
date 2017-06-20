package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/piokaczm/godzira/commands"
)

func main() {
	app := cli.NewApp()
	app.Name = "godzira"
	app.Version = "2.0.0"
	// add more precise description, add some better help text for deploy
	app.Usage = "Smash your apps to servers just like Godzira would smash a city!"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "build config directory and config file",
			Action: commands.Init,
		},
		{
			Name:   "deploy",
			Usage:  "build and deploy binary to remote server",
			Action: commands.Deploy,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err) // one last panic [*]
	}
}
