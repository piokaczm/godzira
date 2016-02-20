###GoDeploy *work in progress*

*it's actually a spec for me while it's in development :)*
Simple ssh deploy tool:
- build a binary using cross-compilation
- copy over to a remote server (or multiple servers at once)
- restart binary over ssh
- integrate with slack

###Installation

`go get github.com/piokaczm/godeploy`

then

`go install`

###Usage

In your app directory run

`godeploy init`

It creates config directory with empty `deploy.yml` config file.
After setting it up just run

`godeploy deploy [environment]`

Depending on your config it restores dependencies, runs tests, builds binary, copies it over to your server(s) and sends notification to your Slack.

###Configuration

First of all you need to properly set your configuration file. Structure should look something like this:

```
appname: NewApp # used for slack integration - it'll be used in msgs
goos: linux # variables needed to properly crosscompile for your machine
goarch: amd64
test: true # only if you want to run all test before deploy, defaults to false
godep: true # only if you're using godep and want to run godep restore before building a binary, defaults to false

environments:
  staging: # this is cli argument you'll be using to deploy to choosen env
    host: example.net # no need for specyfing number when deploying to one host only
    user: example_user
    path: binaries/
    restart_command: etc/init.d/daemon restart
  production:
    host_1: anotherexample.net # the tool matches host and user using the provided digit, so make sure to fill it properly
    user_1: user_1
    host_2: anotherexample2.net
    user_2: user_2
    path: current/binaries/
    restart_command: etc/init.d/daemon restart

slack: # optional
  webhook: https://hooks.slack.com/services/xxx/xxx # no more custom settings for now, please select emoji, name etc. via Slack
```
