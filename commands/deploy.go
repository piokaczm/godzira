package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/piokaczm/godzira/commands/parser"
	"github.com/piokaczm/godzira/commands/task"
)

const (
	configPath = "config/deploy.yml"
	deployed   = `
                   ,:',:',:'
              __||_||_||_||__
         ____["""""""""""""""]____
         \ " '''''''''''''''''''' \
  ~^~^~^~^~^^~^~^~^~^~^~^~^~~^~^~^^~~^~^
     _            _                      _ 
  __| | ___ _ __ | | ___  _   _  ___  __| |
 / _  |/ _ \ '_ \| |/ _ \| | | |/ _ \/ _  |
| (_| |  __/ |_) | | (_) | |_| |  __/ (_| |
 \__,_|\___| .__/|_|\___/ \__, |\___|\__,_|
           |_|            |___/            
  `
)

// Deploy is a wrapper for deploy process.
func Deploy(c *cli.Context) {
	env := c.Args()[0]
	queue := task.NewQueue()

	errors := parser.Read(queue, configPath, env) // TODO: allow setting custom config path!
	if len(errors) > 0 {
		printErrorsAndTerminate(errors)
	}

	err := queue.Exec()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(deployed)
}

func printErrorsAndTerminate(errors []error) {
	for _, err := range errors {
		fmt.Println(err) // do it in red!
	}
	os.Exit(1)
}
