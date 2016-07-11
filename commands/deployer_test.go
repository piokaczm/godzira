package commands

import (
	// "github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type stubDeployer struct{}

func (deployer stubDeployer) preparePath(config *Configuration, env string, server string) string {
	realDeployer := Deployer{}
	return realDeployer.preparePath(config, env, server)
}

func (deployer stubDeployer) prepareCommand(binary string, path string, strategy string) (error, string, []string) {
	realDeployer := Deployer{}
	return realDeployer.prepareCommand(binary, path, strategy)
}

func (deployer stubDeployer) execCopy(command string, args []string) (string, error) {
	values := append([]string{command}, args...)
	return strings.Join(values, ":"), nil
}

func (deployer stubDeployer) execRestart(server string, command string) error {
	return nil
}

func (deployer stubDeployer) execCommand(name string, args []string, start_msg string, finish_msg string) (string, error) {
	return "", nil
}

func TestSingleBinaryDeployment(t *testing.T) {

}
