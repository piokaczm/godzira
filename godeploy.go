package main

import (
	"github.com/codegangsta/cli"
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

				io.WriteString(f, "goos: \ngoarch: \n\nenvironment: \nserver: \npath: \n\nslack_webhook: \nstart_msg: \nfinish_msg: \nslack_name: \nslack_emoji: ")
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
