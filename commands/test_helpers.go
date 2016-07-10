package commands

var data = `
goos: linux
goarch: amd64
test: true
godep: true
strategy: scp
binary_name: 'test_name'

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
  appname: AppName
`

var dataNoStrategy = `
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
  appname: AppName
`
var dataNoSlack = `
goos: linux
goarch: amd64

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
`
