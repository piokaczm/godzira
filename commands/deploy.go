package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strings"
)

func Deploy(c *cli.Context) {
	deploy_env := c.Args()[0]

	config := getConfig()
	servers, err := getServers(&config, deploy_env)
	checkErr(err)

	if config.Godep {
		restoreDependencies(&config)
	}

	if config.Test {
		runTests(&config)
	}

	buildBinary(&config)
	runDeploy(&config, servers, deploy_env)
}

// cross-compile binary using provided config
func buildBinary(config *Configuration) {
	goos := config.Goos
	goarch := config.Goarch

	name := "go"
	args := []string{"build"}
	env := os.Environ()
	env = append(env, fmt.Sprintf("GOOS=%s", goos))
	env = append(env, fmt.Sprintf("GOARCH=%s", goarch))

	cmd := exec.Command(name, args...)
	cmd.Env = env
	fmt.Println("Building binary...")
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("Build succeeded!")
	}
}

// actual deployment
func runDeploy(config *Configuration, servers map[string]string, env string) {
	binary, e := exec.LookPath(getDir())
	checkErr(e)

	fmt.Println("Starting deployment!")
	if slackEnabled(config.Slack) {
		startMsg(config, env)
	}

	for _, value := range servers {
		path := strings.Join([]string{value, config.Environments[env]["path"]}, ":")
		args := []string{"-chavzP", binary, path}
		copyBinary(binary, args, config)
		runRestart(value, config.Environments[env]["restart"], config)
	}

	fmt.Println("Deployment succeeded! ;))))")
	if slackEnabled(config.Slack) {
		finishMsg(config, env)
	}
}

// restart binary via ssh
func runRestart(server string, command string, config *Configuration) {
	args := append([]string{server}, strings.Split(command, " ")...)
	runCommand(
		"ssh",
		args,
		"Restarting binary...",
		"Binary restarted!",
		config)
}

// rsync binary to server(s) listed in the config file
func copyBinary(binary string, args []string, config *Configuration) {
	runCommand(
		"rsync",
		args,
		"Deploying...",
		"Deploy succeeded!",
		config)
}

// run all tests before deploy
// if one of them fails stop deploying
func runTests(config *Configuration) {
	runCommand(
		"go",
		[]string{"test", "-v", "./..."},
		"Running tests...",
		"Tests Passed",
		config)
}

// restore all dependencies before deploy
func restoreDependencies(config *Configuration) {
	runCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!",
		config)
}

func runCommand(name string, args []string, start_msg string, finish_msg string, config *Configuration) {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()

	if err != nil {
		if name == "rsync" {
			errorMsg(config)
			printErr(err)
		} else {
			printErr(err)
		}
	} else {
		fmt.Println(finish_msg)
	}
}
