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
	dir, err := os.Mkdir(config, d)
	checkErr(err)

	f, err := os.Create("config/config.yml")
	checkErr(err)
	defer f.Close()

	w, err := f.Write("goos: \ngoarch: \nenvironment: \nserver: \npath: \nslack_webhook: \nstart_msg: \nfinish_msg: \nslack_name: \nslack_emoji: ")
	checkErr(err)
}

func deployBinary() {
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
