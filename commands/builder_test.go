package commands

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type stubBuilder struct{}

func (builder stubBuilder) execCommand(name string, args []string, env []string) (error, string) {
	values := append([]string{name}, args...)
	return nil, strings.Join(values, ":")
}

func (builder stubBuilder) prepareToCompilation(config *Configuration) (string, []string, []string) {
	realBuilder := Builder{}
	return realBuilder.prepareToCompilation(config)
}

func TestBinaryBuilding(t *testing.T) {
	config := parseConfig([]byte(dataNoStrategy))
	builder := stubBuilder{}
	_, result := buildBinary(&config, builder)
	assert.Equal(t, "go:build", result)
}

func TestNamedBinaryBuilding(t *testing.T) {
	config := parseConfig([]byte(data))
	builder := stubBuilder{}
	_, result := buildBinary(&config, builder)
	assert.Equal(t, "go:build:-o:test_name", result)
}
