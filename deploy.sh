#!/bin/bash
# deploy positionly-reports binary

webhook_url="https://hooks.slack.com/services/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"
user=$USER
env="$1"
msg_start=":rocket: $user has *started* deploying Reports-API to $env"
msg_finish=":star2: $user *finished* deploying Reports-API to $env"

start_deploy() {
  echo "Deploy to $env started"
  curl -X POST --data-urlencode 'payload={ "attachments": [{"mrkdwn_in": ["text"], "color": "good", "text": "'"$msg_start"'"}] }' $webhook_url; echo
}

finish_deploy() {
  echo "Deploy succeeded"
  curl -X POST --data-urlencode 'payload={ "attachments": [{"mrkdwn_in": ["text"], "color": "good", "text": "'"$msg_finish"'"}] }' $webhook_url; echo
}

err_report() {
  msg_error=":crocodile: Something went *wrong*!"
  echo "Error, exiting with status 1"
  curl -X POST --data-urlencode 'payload={ "attachments": [{"mrkdwn_in": ["text"], "color": "danger", "text": "'"$msg_error"'"}] }' $webhook_url; echo
  exit 1
}

git checkout $env
go get github.com/tools/godep
godep restore
go install
go test -v ./...
if [ "$?" != 0 ]; then
  exit 1
fi
echo "Tests passed"

GOOS=linux GOARCH=amd64 go build
echo "Binary build"

if [ "$1" == "staging" ]; then
  start_deploy
  rsync -chavzP --stats positionly-reports ap15@pizdus.pstnly.net:/home/ap15/apireports
  if [ "$?" == 0 ]; then
    ssh ap15@pizdus.pstnly.net '/etc/init.d/daemon_apireports restart'
    finish_deploy
  else
    err_report
  fi
elif [ "$1" == "production" ]; then
  start_deploy
  # rsync -chavzP --stats positionly-reports prod-server
  if [ "$?" == 0 ]; then
    # ssh prd@prd '/etc/init.d/daemon_apireports restart' restart on prd
    finish_deploy
  else
    err_report
  fi
else
  echo "Please specify proper env (staging/production)"
fi
