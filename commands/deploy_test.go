package commands

import (
	"reflect"
	"testing"
)

func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

var data = `
goos: linux
goarch: amd64

environments:
  environment_1:
    name: staging
    server: pizda.net
    user: pizdek
    path: binaries/
    restart_command: etc/dupa/daemon restart
  environment_2:
    name: production
    server: real-pizda.net
    user: pizdekmaster
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
	Expect(t, result.Slack["webhook"], "https://hooks.slack.com/services/sth/more")
	Expect(t, result.Slack["emoji"], ":rocket:")
	Expect(t, result.Slack["botname"], "bot")
	Expect(t, result.Environments["environment_1"]["server"], "pizda.net")
	Expect(t, result.Environments["environment_1"]["name"], "staging")
	Expect(t, result.Environments["environment_1"]["user"], "pizdek")
	Expect(t, result.Environments["environment_1"]["path"], "binaries/")
	Expect(t, result.Environments["environment_1"]["restart_command"], "etc/dupa/daemon restart")
	Expect(t, result.Environments["environment_2"]["server"], "real-pizda.net")
	Expect(t, result.Environments["environment_2"]["name"], "production")
	Expect(t, result.Environments["environment_2"]["user"], "pizdekmaster")
	Expect(t, result.Environments["environment_2"]["path"], "current/binaries/")
	Expect(t, result.Environments["environment_2"]["restart_command"], "etc/prod/dupa/daemon restart")
}
