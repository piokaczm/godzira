package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// these are just wrappers for sendMsg
func startMsg(config *Configuration, env string) {
	user := os.Getenv("USER")
	name := config.AppName
	msg := fmt.Sprintf(":rocket: %s has *started* deploying %s to *%s*!", user, name, env)

	sendMsg(msg, "good", config)
}

func finishMsg(config *Configuration, env string) {
	user := os.Getenv("USER")
	name := config.AppName
	msg := fmt.Sprintf(":star2: %s has *finished* deploying %s to *%s*!", user, name, env)

	sendMsg(msg, "good", config)
}

func errorMsg(config *Configuration) {
	msg := ":crocodile: Something went *wrong*!"

	sendMsg(msg, "danger", config)
}

// post request to slack
func sendMsg(msg string, msgType string, config *Configuration) {
	json, err := json.Marshal(fmt.Sprintf(`{ "attachments": [{"mrkdwn_in": ["text"], "color": "%s", "text": "%s"}`, msgType, msg))
	checkErr(err)
	url := config.Slack["webhook"]

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(json))
	r.Header.Add("Content-Type", "application/json")
}
