package commands

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	success_message = "Build succeeded!"
	command_name    = "go"
	command_arg     = "build"
	name_arg        = "-o"
)

type Builder struct{}

type BinaryBuilder interface {
	prepareToCompilation(config *Configuration) (string, []string, []string)
	execCommand(name string, args []string, env []string) (error, string, string)
}

func buildBinary(config *Configuration, builder BinaryBuilder) (error, string, string) {
	name, args, env := builder.prepareToCompilation(config)
	return builder.execCommand(name, args, env)
}

// cross-compile binary using provided config
func (builder Builder) prepareToCompilation(config *Configuration) (string, []string, []string) {
	name := command_name
	var args []string
	if blank(config.BinName) {
		args = []string{command_arg}
	} else {
		args = []string{command_arg, name_arg, config.BinName}
	}
	env := os.Environ()
	env = append(env, fmt.Sprintf("GOOS=%s", config.Goos))
	env = append(env, fmt.Sprintf("GOARCH=%s", config.Goarch))
	return name, args, env
}

func (builder Builder) execCommand(name string, args []string, env []string) (error, string, string) {
	cmd := exec.Command(name, args...)
	cmd.Env = env
	fmt.Println("Building binary...")
	err := cmd.Run()
	return err, success_message, name
}
