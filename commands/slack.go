package commands

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

// these are just wrappers for sendMsg
func startMsg(slackConfig map[string]string, env string) {
	user := os.Getenv("USER")
	name := slackConfig["appname"]
	msg := fmt.Sprintf(":rocket: %s has *started* deploying %s to *%s*!", user, name, env)

	sendMsg(msg, "good", slackConfig)
}

func finishMsg(slackConfig map[string]string, env string) {
	user := os.Getenv("USER")
	name := slackConfig["appname"]
	msg := fmt.Sprintf(":star2: %s has *finished* deploying %s to *%s*!", user, name, env)

	sendMsg(msg, "good", slackConfig)
}

func errorMsg(slackConfig map[string]string) {
	msg := ":crocodile: Something went *wrong*!"

	sendMsg(msg, "danger", slackConfig)
}

// post request to slack
func sendMsg(msg string, msgType string, slackConfig map[string]string) {
	data := []byte(fmt.Sprintf(`{"attachments": [{"mrkdwn_in": ["text"], "color": "%s", "text": "%s"}]}`, msgType, msg))
	url := slackConfig["webhook"]

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
