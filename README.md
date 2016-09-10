#Godzira

##Smash your apps to servers just like Godzira would smash a city!
[![Build Status](https://travis-ci.org/piokaczm/godzira.svg?branch=master)](https://travis-ci.org/piokaczm/godzira)

Simple deploy tool:
- build a binary using cross-compilation
- copy over to your remote server (or multiple servers at once)
- restart binary over ssh
- send notifications to Slack

###Installation

```
go get github.com/piokaczm/godzira
```

###Usage

In your app directory run

```
godzira init
```

It creates config directory with empty `deploy.yml` config file.

After setting it up just run

```
godzira deploy [environment]
```

Depending on your config it restores dependencies, runs tests, builds binary, copies it over to your server(s) and sends notification to your Slack.

###Configuration

You need to properly set your configuration file. Structure should look something like this:

```
goos: linux # cross-compilation variables
goarch: amd64
vendor: true # if you're using vendor experiment and want to filter out vendor dir (damn you go!)
test: true # only if you want to run all tests before deploy, defaults to false
godep: true # only if you're using godep and want to run godep restore before building a binary, defaults to false
strategy: scp # optional
binary_name: name # optional - if blank the tool will run plain go build

environments:
  staging: # it's [environment] for deploy command, name it as you wish
    host: example.net # no need for adding a digit when deploying to one server only
    user: example_user
    path: binaries/
    restart_command: etc/init.d/daemon restart
  production:
    host_1: anotherexample.net # Godzira matches host and user using the provided digit, so make sure to fill it properly
    user_1: user_1
    host_2: anotherexample2.net
    user_2: user_2
    path: current/binaries/
    restart_command: etc/init.d/daemon restart

slack: # optional
  token: slack token
  channel: bot_checks
  appname: AppName
  start_emoji: rocket
  finish_emoji: heart
  error_emoji: troll
```

###Strategy

Godzira provides two deploy strategies: `scp` and `rsync`.
You can choose which one to use in your config file. If no strategy specified, the tool will use `rsync`.
