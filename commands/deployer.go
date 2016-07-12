package commands

import (
	"errors"
	"fmt"
	"strings"
)

const (
	deploy_succeeded_msg = "Deployment finished!"
	deploy_started_msg   = "Starting deployment!"
)

type Deployer struct{}

type BinaryDeployer interface {
	preparePath(config *Configuration, env string, server string) string
	prepareCommand(binary string, path string, strategy string) (error, string, []string)
	execCopy(command string, args []string) (string, error)
	execRestart(server string, command string) (string, error)
	execCommand(name string, args []string, start_msg string, finish_msg string) (string, error)
}

// actual deployment
func runDeploy(config *Configuration, server string, env string, binary string, deployer BinaryDeployer) string {
	deployPrint(server, deploy_started_msg)

	path := deployer.preparePath(config, env, server)
	strategy := config.getStrategy()
	commandErr, command, args := deployer.prepareCommand(binary, path, strategy)
	checkErr(commandErr)
	finishMsg, copyErr := deployer.execCopy(command, args)
	checkErr(copyErr)

	if notBlank(config.Environments[env]["restart_command"]) {
		restartMsg, restartErr := deployer.execRestart(server, config.Environments[env]["restart_command"])
		checkErr(restartErr)
		fmt.Println(restartMsg)
	}

	return finishMsg
}

func (deployer Deployer) preparePath(config *Configuration, env string, server string) string {
	return strings.Join([]string{server, config.Environments[env]["path"]}, ":")
}

func (deployer Deployer) execRestart(server string, command string) (string, error) {
	args := append([]string{server}, strings.Split(command, " ")...)
	return execCommand(
		"ssh",
		args,
		"Restarting binary...",
		"Binary restarted!")
}

func (deployer Deployer) prepareCommand(binary string, path string, strategy string) (error, string, []string) {
	var command string
	args := make([]string, 0, 3)
	if strategy == "scp" {
		command = scp
		path := strings.Join([]string{path, binary}, "")
		args = append(args, []string{binary, path}...)
	} else if strategy == "rsync" {
		command = rsync
		args = append(args, []string{rsyncArgs, binary, path}...)
	} else {
		return errors.New("Unknown strategy, please select scp or rsync"), "", []string{}
	}
	return nil, command, args
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
