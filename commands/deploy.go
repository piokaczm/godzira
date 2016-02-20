package commands

import (
	"github.com/codegangsta/cli"
)

func Deploy(c *cli.Context) {
	deploy_env := c.Args()[1]

	config := getConfig()
	servers, err := getServers(&config, deploy_env)
	checkErr(err)

	if config.Godep {
		restoreDependencies()
	}

	if config.Test {
		runTests()
	}

	buildBinary(config)
	runDeploy(&config, servers, deploy_env)
}

func runDeploy(config *Configuration, servers map[string]string, env string) {
	if len(config.Slack) != 0 {
		startMsg(config, env)
	}

}

// run all tests before deploy
// if any of them fails stop deploying
func runTests() {
	runCommand(
		"go",
		[]string{"test", "-v", "./..."},
		"Running tests...",
		"Tests Passed")
}

// restore all dependencies before deploy
func restoreDependencies() {
	runCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!")
}
