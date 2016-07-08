package commands

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"os/exec"
	"strings"
)

const (
	scp       = "scp"
	rsync     = "rsync"
	rsyncArgs = "-chavzP"
)

func Deploy(c *cli.Context) {
	deploy_env := c.Args()[0] // try to extract it somehow

	config := getConfig()
	servers, err := getServers(config.Environments, deploy_env)
	checkErr(err)

	if config.Godep {
		e := restoreDependencies()
		checkErr(e)
	}

	if config.Test {
		e := runTests(config.Vendor)
		checkErr(e)
	}

	buildBinary(config.Goarch, config.Goos)
	runDeploy(&config, servers, deploy_env)
}

// cross-compile binary using provided config
func buildBinary(goarch string, goos string) {
	// try to rewrite runCommand so there is not so much duplication
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
		checkErr(err)
	} else {
		fmt.Println("Build succeeded!")
	}
}

// actual deployment
func runDeploy(config *Configuration, servers []string, env string) {
	binary := getDir() // that's stupid, compile named file

	fmt.Println("Starting deployment!")
	if slackEnabled(config.Slack) {
		startMsg(config.Slack, env)
	}
	strategy := config.getStrategy()

	for _, value := range servers {
		path := strings.Join([]string{value, config.Environments[env]["path"]}, ":")
		err := copyBinary(binary, path, strategy)
		checkErrWithMsg(err, config.Slack)
		e := runRestart(value, config.Environments[env]["restart_command"])
		checkErr(e)
	}

	fmt.Println(deployed)
	if slackEnabled(config.Slack) {
		finishMsg(config.Slack, env)
	}
}

// restart binary via ssh
func runRestart(server string, command string) error {
	args := append([]string{server}, strings.Split(command, " ")...)
	err := runCommand(
		"ssh",
		args,
		"Restarting binary...",
		"Binary restarted!")
	return err
}

// rsync binary to server(s) listed in the config file
func copyBinary(binary string, path string, strategy string) error {
	var command string
	args := make([]string, 0, 3)
	if strategy == "scp" {
		command = scp
		args = append(args, []string{binary, path}...)
	} else if strategy == "rsync" {
		command = rsync
		args = append(args, []string{rsyncArgs, binary, path}...)
	} else {
		return errors.New("Unknown strategy, please select scp or rsync")
	}
	err := runCopy(command, args)
	return err
}

func runCopy(command string, args []string) error {
	err := runCommand(
		command,
		args,
		"Deploying...",
		"Deploy succeeded!")
	return err
}

// run all tests before deploy
// if one of them fails stop deploying
func runTests(vendor bool) error {
	args := []string{"test", "-v"}
	if vendor {
		dirs, e := filterVendor()
		checkErr(e)
		args = append(args, dirs...)
	} else {
		args = append(args, "./...")
	}

	err := runCommand(
		"go",
		args,
		"Running tests...",
		"Tests passed!")
	return err
}

// filter out /vendor dir for tests
// if app uses vendor experiment
func filterVendor() ([]string, error) {
	list := exec.Command("go", "list", "./...")
	grep := exec.Command("grep", "-v", "/vendor/")
	listOut, _ := list.StdoutPipe()
	list.Start()
	grep.Stdin = listOut

	out, err := grep.Output()
	if err != nil {
		return nil, err
	} else {
		dirs := strings.Split(string(out), "\n")
		dirs = dirs[:len(dirs)-1]
		return dirs, nil
	}
}

// restore all dependencies before deploy
func restoreDependencies() error {
	err := runCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!")
	return err
}

func runCommand(name string, args []string, start_msg string, finish_msg string) error {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()
	if err != nil {
		return err
	} else {
		fmt.Println(finish_msg)
		return nil
	}
}
