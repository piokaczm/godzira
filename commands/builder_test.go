package commands

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type stubBuilder struct{}

func (builder stubBuilder) execCommand(name string, args []string, env []string) (error, string) {
	args = append(args, name)
	return nil, strings.Join([]string{args}, ":")
}

func (builder stubBuilder) buildBinary(goarch string, goos string) string {
	realBuilder = Builder{}
	return realBuilder.buildBinary(goarch, goos)
}

func TestBinaryBuilding(t *testing.T) {

}
