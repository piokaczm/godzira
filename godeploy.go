package main

import (
	"github.com/codegangsta/cli"
	// "gopkg.in/yaml.v2"
	"io"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "godeploy"
	app.EnableBashCompletion = true
	app.Usage = "Simply deploy binaries to remote server"
	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "build config directory and config file",
			Action: func(c *cli.Context) {
				os.Mkdir("./config", 0777)

				f, err := os.Create("config/config.yml")
				checkErr(err)
				defer f.Close()
				const (
					comment = "# see example config file at github.com/piokaczm/godeploy"
				)

				io.WriteString(f, comment)
			},
		},
		{
			Name:  "deploy",
			Usage: "build and deploy binary to remote server",
			// Action: deployBinary,
		},
	}

	app.Run(os.Args)
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
