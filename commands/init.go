package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/codegangsta/cli"
)

// Init creates config dir along with empty deploy.yml file with link to example configuration.
func Init(c *cli.Context) {
	err := os.Mkdir("./config", 0777)
	if err != nil {
		printErrorsAndTerminate(fmt.Errorf("an error occured while creating config dir : %s", err))
	}

	f, err := os.Create(configPath)
	if err != nil {
		printErrorsAndTerminate(fmt.Errorf("an error occured while trying to create '%s' : %s", configPath, err))
	}
	defer close(f)

	_, err = io.WriteString(f, "# see example config file at github.com/piokaczm/godzira")
	if err != nil {
		printErrorsAndTerminate(fmt.Errorf("an error occured while trying to write a comment to config file : %s", err))
	}
}

func close(c io.Closer) {
	err := c.Close()
	if err != nil {
		printErrorsAndTerminate(err)
	}
}
