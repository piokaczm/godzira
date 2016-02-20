package commands

import (
	"github.com/codegangsta/cli"
)

func Deploy(c *cli.Context) {
	// deploy_env := c.Args()[1]

	config := getConfig()
	// servers, err := getServers(config, deploy_env)
	// checkErr(err)

	if config.Godep {
		restoreDependencies()
	}

	if config.Test {
		runTests()
	}

	// startDeploy()
}

// func startDeploy() {

// }

func runTests() {
	runCommand(
		"go",
		[]string{"test", "-v", "./..."},
		"Running tests...",
		"Tests Passed")
}

func restoreDependencies() {
	runCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!")
}
