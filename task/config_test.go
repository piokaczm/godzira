package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestYmlParsing(t *testing.T) {
	config, _, _, _ := loadConfig("../fixtures/config_basic.yml", "production")

	assert.Equal(t, "production", config.CurrentEnv)
	assert.Equal(t, "basicApp", config.Name)
	assert.Equal(t, "rsync", config.Strategy)
	assert.Equal(t, "linux", config.Goos)
	assert.Equal(t, "amd64", config.Goarch)
	assert.Equal(t, true, config.Test)
	assert.Equal(t, true, config.Vendor)
	assert.Equal(t, "basic", config.Mode)

	// check if all envs are parsed properly
	assert.Equal(t, 2, len(config.Environments["production"]["hosts"]))
	assert.Equal(t, 1, len(config.Environments["staging"]["hosts"]))
}
