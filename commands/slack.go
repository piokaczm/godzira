package commands

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io/ioutil"
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
	data := []byte(fmt.Sprintf(`{"attachments": [{"mrkdwn_in": ["text"], "color": "%s", "text": "%s"}]}`, msgType, msg))
	url := config.Slack["webhook"]

	r, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
