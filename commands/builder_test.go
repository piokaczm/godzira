package commands

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type stubBuilder struct{}

func (builder stubBuilder) execCommand(name string, args []string, env []string) (error, string, string) {
	return nil, name, strings.Join(args, ":")
}

func (builder stubBuilder) prepareToCompilation(config *Configuration) (string, []string, []string) {
	realBuilder := Builder{}
	return realBuilder.prepareToCompilation(config)
}

func TestBinaryBuilding(t *testing.T) {
	config := parseConfig([]byte(dataNoStrategy))
	builder := stubBuilder{}
	_, command, args := buildBinary(&config, builder)
	assert.Equal(t, "go", command)
	assert.Equal(t, "build", args)
}

func TestNamedBinaryBuilding(t *testing.T) {
	config := parseConfig([]byte(data))
	builder := stubBuilder{}
	_, command, args := buildBinary(&config, builder)
	assert.Equal(t, "go", command)
	assert.Equal(t, "build:-o:test_name", args)
}
