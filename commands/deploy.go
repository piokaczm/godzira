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

	builder := Builder{}
	deployer := Deployer{}
	deployApp(builder, deployer, config, servers, deploy_env)
}

func deployApp(builder BinaryBuilder, deployer BinaryDeployer, config Configuration, servers []string, env string) {
	// build binary, check if it succeeded, if so print success message
	err, msg := buildBinary(config.Goarch, config.Goos, builder)
	checkErr(err)
	fmt.Println(msg)
	// possibly just move servers fetching to deployer interface? why inject it here as we need it inside it
	// and config is passed anyway?
	deployer.runDeploy(&config, servers, env)
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

	err := execCommand(
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
	err := execCommand(
		"godep",
		[]string{"restore"},
		"Restoring dependencies...",
		"Dependencies restored!")
	return err
}

func execCommand(name string, args []string, start_msg string, finish_msg string) error {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()
	if err != nil {
		return err
	} else {
		fmt.Println(finish_msg)
		return nil
	}
}
