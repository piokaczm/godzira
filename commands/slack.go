package commands

import (
  "fmt"
  "github.com/nlopes/slack"
  "os"
)

func startMsg(slackConfig map[string]string, env string) {
  user := os.Getenv("USER")
  name := slackConfig["appname"]
	msg := fmt.Sprintf(":%s: %s has *started* deploying %s to *%s*!", slackConfig["start_emoji"], user, name, env)

	sendMsg(msg, slackConfig)
}

func finishMsg(slackConfig map[string]string, env string) {
  user := os.Getenv("USER")
  name := slackConfig["appname"]
	msg := fmt.Sprintf(":%s: %s has *finished* deploying %s to *%s*!", slackConfig["finish_emoji"], user, name, env)

	sendMsg(msg, slackConfig)
}

func errorMsg(slackConfig map[string]string) {
	msg := fmt.Sprintf(":%s: Something went *wrong*!", slackConfig["error_emoji"])

	sendMsg(msg, slackConfig)
}

func sendMsg(msg string, slackConfig map[string]string) {
  api := slack.New(slackConfig["token"])

  _, _, err := api.PostMessage(slackConfig["channel"], msg, slack.PostMessageParameters{})
	if err != nil {
		panic(err)
	}
}
