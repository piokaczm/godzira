package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsing(t *testing.T) {
	result := parseConfig([]byte(data))
	assert.Equal(t, result.Goos, "linux")
	assert.Equal(t, result.Strategy, "scp")
	assert.Equal(t, result.Goarch, "amd64")
	assert.Equal(t, result.Test, true)
	assert.Equal(t, result.Godep, true)
	assert.Equal(t, result.Slack["webhook"], "https://hooks.slack.com/services/sth/more")
	assert.Equal(t, result.Slack["appname"], "AppName")
	assert.Equal(t, result.Environments["staging"]["host"], "pizda.net")
	assert.Equal(t, result.Environments["staging"]["user"], "pizdek")
	assert.Equal(t, result.Environments["staging"]["path"], "binaries/")
	assert.Equal(t, result.Environments["staging"]["restart_command"], "etc/dupa/daemon restart")
	assert.Equal(t, result.Environments["production"]["host_1"], "real-pizda.net")
	assert.Equal(t, result.Environments["production"]["user_1"], "pizdekmaster")
	assert.Equal(t, result.Environments["production"]["host_2"], "real-pizda2.net")
	assert.Equal(t, result.Environments["production"]["user_2"], "pizdekmaster2")
	assert.Equal(t, result.Environments["production"]["path"], "current/binaries/")
	assert.Equal(t, result.Environments["production"]["restart_command"], "etc/prod/dupa/daemon restart")
}

func TestParseServer(t *testing.T) {
	assert.Equal(t, parseServer("pizdek", "pizda.net"), "pizdek@pizda.net")
}

func TestSetServerWithTwoHosts(t *testing.T) {
	config := parseConfig([]byte(data))
	result, _ := getServers(config.Environments, "production")
	assert.Equal(t, result, []string{"pizdekmaster@real-pizda.net", "pizdekmaster2@real-pizda2.net"})
	assert.Equal(t, len(result), 2)
}

func TestSetServerWithOneHost(t *testing.T) {
	config := parseConfig([]byte(data))
	result, _ := getServers(config.Environments, "staging")
	assert.Equal(t, len(result), 1)
	assert.Equal(t, result[0], "pizdek@pizda.net")
}

func TestGetStrategy(t *testing.T) {
	config := parseConfig([]byte(data))
	result := config.getStrategy()
	assert.Equal(t, result, "scp")

	otherConfig := parseConfig([]byte(dataNoStrategy))
	otherResult := otherConfig.getStrategy()
	assert.Equal(t, otherResult, "rsync")
}
