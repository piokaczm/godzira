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
)

type Builder struct{}

type BinaryBuilder interface {
	prepareToCompilation(goarch string, goos string) (string, []string, []string)
	execCommand(name string, args []string, env []string) (error, string)
}

func buildBinary(goarch string, goos string, builder BinaryBuilder) (error, string) {
	name, args, env := builder.prepareToCompilation(goarch, goos)
	return builder.execCommand(name, args, env)
}

// cross-compile binary using provided config
func (builder Builder) prepareToCompilation(goarch string, goos string) (string, []string, []string) {
	name := command_name
	args := []string{command_arg}
	env := os.Environ()
	env = append(env, fmt.Sprintf("GOOS=%s", goos))
	env = append(env, fmt.Sprintf("GOARCH=%s", goarch))
	return name, args, env
}

func (builder Builder) execCommand(name string, args []string, env []string) (error, string) {
	cmd := exec.Command(name, args...)
	cmd.Env = env
	fmt.Println("Building binary...")
	err := cmd.Run()
	return err, success_message
}
