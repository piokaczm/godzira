package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "godeploy"
	app.Usage = "Simply deploy binaries to remote server"
	app.Commands = []cli.Command{
		{
			Name:   "init",
			Usage:  "build config directory and config file",
			Action: initConfig(),
		},
		{
			Name:   "deploy",
			Usage:  "build and deploy binary to remote server",
			Action: deployBinary(),
		},
	}

	app.Run(os.Args)
}

func initConfig() {

}

func deployBinary() {

}
