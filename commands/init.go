package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

func Config(c *cli.Context) {
	os.Mkdir("./config", 0777)

	f, err := os.Create(configPath)
	if err != nil {
		fmt.Printf("there was en error while trying to create '%s'", configPath)
		os.Exit(1)
	}
	defer f.Close()

	io.WriteString(f, "# see example config file at github.com/piokaczm/godzira")
}
