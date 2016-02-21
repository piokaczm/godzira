package commands

import (
	"github.com/piokaczm/godeploy/Godeps/_workspace/src/github.com/codegangsta/cli"
	"io"
	"os"
)

func Config(c *cli.Context) {
	os.Mkdir("./config", 0777)

	f, err := os.Create("config/deploy.yml")
	checkErr(err)
	defer f.Close()
	const (
		comment = "# see example config file at github.com/piokaczm/godeploy"
	)

	io.WriteString(f, comment)
}
