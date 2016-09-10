package commands

import (
	"github.com/codegangsta/cli"
	"io"
	"os"
)

func Config(c *cli.Context) {
	os.Mkdir("./config", 0777)

	f, err := os.Create("config/deploy.yml")
	checkErr(err)
	defer f.Close()
	const (
		comment = `# find more information at github.com/piokaczm/godeploy
# goos: linux # cross-compilation variables
# goarch: amd64
# vendor: true # if you're using vendor experiment and want to filter out vendor dir (damn you go!)
# test: true # only if you want to run all tests before deploy, defaults to false
# godep: true # only if you're using godep and want to run godep restore before building a binary, defaults to false
# strategy: scp # optional
# binary_name: name # optional - if blank the tool will run plain go build
#
# environments:
#   staging: # it's [environment] for deploy command, name it as you wish
#     host: example.net # no need for adding a digit when deploying to one server only
#     user: example_user
#     path: binaries/
#     restart_command: etc/init.d/daemon restart
#   production:
#     host_1: anotherexample.net # Godzira matches host and user using the provided digit, so make sure to fill it properly
#     user_1: user_1
#     host_2: anotherexample2.net
#     user_2: user_2
#     path: current/binaries/
#     restart_command: etc/init.d/daemon restart
#
# slack: # optional
#   token: slack token
#   appname: AppName
#   start_emoji: rocket
#   finish_emoji: heart
#   error_emoji: troll
#   channel: bot_checks`
  )

	io.WriteString(f, comment)
}
