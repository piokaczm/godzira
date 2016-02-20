package commands

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)
// msg_start=":rocket: $user has *started* deploying Reports-API to $env"
// msg_finish=":star2: $user *finished* deploying Reports-API to $env"
// curl -X POST --data-urlencode 'payload={ "attachments": [{"mrkdwn_in": ["text"], "color": "danger", "text": "'"$msg_error"'"}] }' $webhook_url; echo
// curl -X POST --data-urlencode 'payload={ "attachments": [{"mrkdwn_in": ["text"], "color": "good", "text": "'"$msg_finish"'"}] }' $webhook_url; echo

// these are just wrappers for sendMsg
func Start(config *Configuration, env string) {
	user := os.Getenv("USER")
	name := config.AppName
	msg := fmt.Sprintf(":rocket: %s has *started* deploying %s to *%s*", user, name, env)

	sendMsg(url, msg, "good")
}

func Finish(config *Configuration) {

}

// prepare payload to be send via post request
func prepareMsg(msg string, msgType string) url.Values {
	data := [1]map[string]string{}
	data[0]["mrkdwn_in"] = "text"
	data[0]["color"] = msgType
	data[0]["text"] = msg

	payload := url.Values{}
	payload.Set("attachements", data)
	return payload
}

// post request to slack
func sendMsg(msg string, msgType string) error {
	payload := prepareMsg(msg, msgType)

	url := config.Slack["webhook"]
	client := &http.Client{}
	r, _ := http.NewRequest("POST", url, bytes.NewBufferString(payload.Encode()))
	r.Header.Add("Content-Type", "application/x-www-data-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(payload.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)
}
