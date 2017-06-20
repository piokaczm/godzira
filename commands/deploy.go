package commands

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/piokaczm/godzira/commands/parser"
	"github.com/piokaczm/godzira/commands/printer"
	"github.com/piokaczm/godzira/commands/slack"
	"github.com/piokaczm/godzira/commands/task"
)

const (
	errorLabel = "[ error ]"
	infoLabel  = "[ info  ]"
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
	printer.PrintInfo(infoLabel, "Godzira is starting deployment...")
	validateCommand(c)

	env := c.Args()[0]
	queue := task.NewQueue()
	config, err := parser.New(configPath, env)
	if err != nil {
		printErrorsAndTerminate(err)
	}

	slackClient := slack.New(
		config.Slack.Webhook,
		config.Slack.Channel,
		config.Slack.Emoji,
		config.Slack.Name,
	)

	errors := parser.Read(config, queue)
	if len(errors) > 0 {
		printErrorsAndTerminate(errors...)
	}

	slackClient.PostSuccess(fmt.Sprintf("*%s* started deploying *%s* to *%s*!", os.Getenv("USER"), config.Name, env))
	err = queue.Exec()
	if err != nil {
		slackClient.PostDanger(fmt.Sprintf("Deployment of *%s* to *%s* failed!", config.Name, env))
		os.Exit(1)
	}

	slackClient.PostSuccess(fmt.Sprintf("*%s* was succesfully deployed to *%s*!", config.Name, env))
	fmt.Println(deployed)
}

func printErrorsAndTerminate(errors ...error) {
	for _, err := range errors {
		printer.PrintWarning(errorLabel, err.Error())
	}
	os.Exit(1)
}

func validateCommand(c *cli.Context) {
	// check if at least env was passed to the command
	if len(c.Args()) < 1 {
		printer.PrintWarning(errorLabel, "please provide deployment env")
		os.Exit(1)
	}

	// check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		printer.PrintWarning(errorLabel, "couldn't find config file 'config/deploy.yml'")
		os.Exit(1)
	}
}
