package commands

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os/exec"
	"strings"
)

const (
	scp       = "scp"
	rsync     = "rsync"
	rsyncArgs = "-chavzP"
)

// Deploy is a wrapper for whole deply process.
func Deploy(c *cli.Context) {
	deploy_env := c.Args()[0] // try to extract it somehow
	config := getConfig()

	if config.Godep {
		msg, err := restoreDependencies()
		checkErr(err)
		fmt.Println(msg)
	}

	if config.Test {
		msg, e := runTests(config.Vendor)
		checkErr(e)
		fmt.Println(msg)
	}

	builder := Builder{}
	deployer := Deployer{}
	deployApp(builder, deployer, config, deploy_env)
}

// deployApp is a function which builds a binary using provided builder and then iterates over
// servers and deploys it to them.
func deployApp(builder BinaryBuilder, deployer BinaryDeployer, config Configuration, env string) {
	buildErr, buildMsg := buildBinary(&config, builder)
	checkErr(buildErr)
	fmt.Println(buildMsg)
	var binary string

	servers, err := getServers(config.Environments, env)
	checkErr(err)
	if notBlank(config.BinName) {
		binary = config.BinName
	} else {
		binary = getDir()
	}

	if slackEnabled(config.Slack) {
		startMsg(config.Slack, env)
	}

	for _, server := range servers {
		deployMsg := runDeploy(&config, server, env, binary, deployer)
		fmt.Println(deployMsg)
	}

	// asci art
	fmt.Println(deployed)
	if slackEnabled(config.Slack) {
		finishMsg(config.Slack, env)
	}
}

// runTests runs all availabele tests.
// If one of them fails, deploy stops.
func runTests(vendor bool) (string, error) {
	args := []string{"test", "-v"}
	if vendor {
		dirs, e := filterVendor()
		checkErr(e)
		args = append(args, dirs...)
	} else {
		args = append(args, "./...")
	}

	return execCommand(
		"go",
		args,
		"Running tests...",
		"Tests passed!")
}

// filterVendor filters vendor directory for tests.
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

// restoreDependencies restores all dependencies using godep.
// Probably it should be removed.
func restoreDependencies() (string, error) {
	return execCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!")
}

// execCommand is a wrapper for running shell commands.
// It returns last provided message for testing purposes (no better ide atm).
func execCommand(name string, args []string, start_msg string, finish_msg string) (string, error) {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()
	return finish_msg, err
}
