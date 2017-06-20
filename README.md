#Godzira

##Smash your apps to servers just like Godzira would smash a city!
[![Build Status](https://travis-ci.org/piokaczm/godzira.svg?branch=master)](https://travis-ci.org/piokaczm/godzira)

Simple deploy tool:
- copy your app to remote servers
- configure tasks to run before or after deployment
- send notifications to Slack

###Installation

```
go install github.com/piokaczm/godzira
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

###Configuration

You need to properly set your configuration file. Structure should look something like this:

```
test: true
strategy: scp
name: 'sample_bin'
binary_path: 'bin/sample'

environments:
  staging:
    - host: stg.net
      user: stgapp
      path: /home/stgapp/app/
  production:
    - host: real.net
      user: app1
      path: /home/app/binaries/
    - host: real2.net
      user: app2
      path: /home/app/binaries/

pretasks:
  - name: copy env
    path: /home/sampleapp/.env
    destination: /home/app/.env
    type: copy

  - name: echo
    command: echo test
    type: local

posttasks:
  - name: restart
    command: /etc/sample/daemon restart
    type: remote

  - name: echo
    command: echo test
    type: local

```

As for now godzira supports 3 `type` labels for tasks:

- `local` for tasks which are supposed to run on a local machine
- `remote` tasks which are supposed to run on each host for given environment
- `copy` for copying files from local machine to each host (where `path` is local path to the file and `destination` is where it should be copied to)

###Strategy

Godzira provides two deploy strategies: `scp` and `rsync`.
You can choose which one to use in your config file. If no strategy specified, the tool will use `rsync`.
