package commands

import (
	"fmt"
	"os/exec"
	"strings"
)

type Deployer struct{}

type BinaryDeployer interface {
	runDeploy(config *Configuration, servers []string, env string)
	copyBinary(binary string, path string, strategy string) error
	runCopy(command string, args []string) error
	runRestart(command string, args []string) error
	execCommand(name string, args []string, start_msg string, finish_msg string) error
}

// actual deployment
func (deployer Deployer) runDeploy(config *Configuration, servers []string, env string) {
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
func (deployer Deployer) runRestart(server string, command string) error {
	args := append([]string{server}, strings.Split(command, " ")...)
	err := runCommand(
		"ssh",
		args,
		"Restarting binary...",
		"Binary restarted!")
	return err
}

// rsync binary to server(s) listed in the config file
func (deployer Deployer) copyBinary(binary string, path string, strategy string) error {
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

func (deployer Deployer) runCopy(command string, args []string) error {
	err := runCommand(
		command,
		args,
		"Deploying...",
		"Deploy succeeded!")
	return err
}

func (deployer Deployer) execCommand(name string, args []string, start_msg string, finish_msg string) error {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()
	if err != nil {
		return err
	} else {
		fmt.Println(finish_msg)
		return nil
	}
}
