package commands

import (
	"fmt"
	"os/exec"
)

const (
	success_message = "Build succeeded!"
)

type Builder struct{}

type BinaryBuilder interface {
	buildBinary(goarch string, goos string) error
	execCommand(name string, args []string, start_msg string, finish_msg string) error
}

// cross-compile binary using provided config
func (builder Builder) buildBinary(goarch string, goos string) string {
	name := "go"
	args := []string{"build"}
	env := os.Environ()
	env = append(env, fmt.Sprintf("GOOS=%s", goos))
	env = append(env, fmt.Sprintf("GOARCH=%s", goarch))
	message, err := builder.execCommand(name, args, env)
	checkErr(err)
	return message
}

func (builder Builder) execCommand(name string, args []string, env []string) (error, string) {
	cmd := exec.Command(name, args...)
	cmd.Env = env
	fmt.Println("Building binary...")
	err := cmd.Run()
	return err, success_message
}
