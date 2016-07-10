package commands

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type stubBuilder struct{}

func (builder stubBuilder) execCommand(name string, args []string, env []string) (error, string) {
	args = append(args, name)
	return nil, strings.Join(args, ":")
}

func (builder stubBuilder) prepareToCompilation(goarch string, goos string) (string, []string, []string) {
	realBuilder := Builder{}
	return realBuilder.prepareToCompilation(goarch, goos)
}

func TestBinaryBuilding(t *testing.T) {
	goarch, goos := "amd64", "linux"
	builder := stubBuilder{}
	_, result := buildBinary(goarch, goos, builder)
	assert.Equal(t, "build:go", result)
}
