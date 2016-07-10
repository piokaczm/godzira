package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	deploy_succeeded_msg = "Deployment finished!"
	deploy_started_msg   = "Starting deployment!"
)

type Deployer struct{}

type BinaryDeployer interface {
	preparePath(config *Configuration, env string, server string) (string, error)
	prepareCommand(binary string, path string, strategy string) (string, []string)
	execCopy(command string, args []string) (string, error)
	execRestart(server string, command string) error
	execCommand(name string, args []string, start_msg string, finish_msg string) (string, error)
}

// actual deployment
func runDeploy(config *Configuration, server string, env string, binary string, deployer BinaryDeployer) {
	deployPrint(server, deploy_started_msg)

	path, prepareErr := deployer.preparePath(config, env, server)
	checkErr(prepareErr)

	command, args := deployer.prepareCommand(binary, path, strategy)
	finishMsg, copyErr := deployer.execCopy(command, args)
	checkErr(copyErr)

	restartErr := deployer.execRestart(server, config.Environments[env]["restart_command"])
	checkErr(restartErr)

	return finishMsg
}

func (deployer Deployer) preparePath(config *Configuration, env string, server string) (string, error) {
	return strings.Join([]string{server, config.Environments[env]["path"]}, ":")
}

func (deployer Deployer) execRestart(server string, command string) error {
	args := append([]string{server}, strings.Split(command, " ")...)
	return execCommand(
		"ssh",
		args,
		"Restarting binary...",
		"Binary restarted!")
}

func (deployer Deployer) prepareCommand(binary string, path string, strategy string) (string, []string) {
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
	return command, args
}

func (deployer Deployer) execCopy(command string, args []string) (string, error) {
	return execCommand(
		command,
		args,
		"Deploying...",
		"Deploy succeeded!")
}

func (deployer Deployer) execCommand(name string, args []string, start_msg string, finish_msg string) (string, error) {
	return execCommand(name, args, start_msg, finish_msg)
}
