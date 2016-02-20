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

// msg_error=":crocodile: Something went *wrong*!"
func errorMsg(config *Configuration) {
	msg := ":crocodile: Something went *wrong*!"

	sendMsg(msg, "danger", config)
}

// prepare payload to be send via post request
func prepareMsg(msg string, msgType string) map[string][1]map[string]string {
	// { "attachments": [{"mrkdwn_in": ["text"], "color": "good", "text": "'"$msg_start"'"}
	data := map[string][1]map[string]string{}
	options := [1]map[string]string{}
	options[0]["mrkdwn_in"] = "text"
	options[0]["color"] = msgType
	options[0]["text"] = msg
	data["attachments"] = options

	return data
}

// post request to slack
func sendMsg(msg string, msgType string, config *Configuration) {
	json, err := json.Marshal(prepareMsg(msg, msgType))
	checkErr(err)
	url := config.Slack["webhook"]

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(json))
	r.Header.Add("Content-Type", "application/json")
}
