package commands

import (
	"testing"
)

var data = `
goos: linux
goarch: amd64
test: true
godep: true

environments:
  staging:
    host: pizda.net
    user: pizdek
    path: binaries/
    restart_command: etc/dupa/daemon restart
  production:
    host_1: real-pizda.net
    user_1: pizdekmaster
    host_2: real-pizda2.net
    user_2: pizdekmaster2
    path: current/binaries/
    restart_command: etc/prod/dupa/daemon restart

slack:
  webhook: https://hooks.slack.com/services/sth/more
  emoji: ":rocket:"
  botname: bot
`

func TestParsing(t *testing.T) {
	result := parseConfig([]byte(data))
	Expect(t, result.Goos, "linux")
	Expect(t, result.Goarch, "amd64")
	Expect(t, result.Test, true)
	Expect(t, result.Godep, true)
	Expect(t, result.Slack["webhook"], "https://hooks.slack.com/services/sth/more")
	Expect(t, result.Slack["emoji"], ":rocket:")
	Expect(t, result.Slack["botname"], "bot")
	Expect(t, result.Environments["staging"]["host"], "pizda.net")
	Expect(t, result.Environments["staging"]["user"], "pizdek")
	Expect(t, result.Environments["staging"]["path"], "binaries/")
	Expect(t, result.Environments["staging"]["restart_command"], "etc/dupa/daemon restart")
	Expect(t, result.Environments["production"]["host_1"], "real-pizda.net")
	Expect(t, result.Environments["production"]["user_1"], "pizdekmaster")
	Expect(t, result.Environments["production"]["host_2"], "real-pizda2.net")
	Expect(t, result.Environments["production"]["user_2"], "pizdekmaster2")
	Expect(t, result.Environments["production"]["path"], "current/binaries/")
	Expect(t, result.Environments["production"]["restart_command"], "etc/prod/dupa/daemon restart")
}

func TestParseServer(t *testing.T) {
	Expect(t, parseServer("pizdek", "pizda.net"), "pizdek@pizda.net")
}

func TestSetServerWithTwoHosts(t *testing.T) {
	config := parseConfig([]byte(data))
	result, _ := setServers(&config, "production")
	Expect(t, result["server_1"], "pizdekmaster@real-pizda.net")
	Expect(t, result["server_2"], "pizdekmaster2@real-pizda2.net")
	Expect(t, len(result), 2)
}

func TestSetServerWithOneHost(t *testing.T) {
	config := parseConfig([]byte(data))
	result, _ := setServers(&config, "staging")
	Expect(t, len(result), 1)
	Expect(t, result["server_1"], "pizdek@pizda.net")
}
