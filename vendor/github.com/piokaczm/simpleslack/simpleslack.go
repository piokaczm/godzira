package simpleslack

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type Slack struct {
	Webhook  string
	Channel  string
	Emoji    string
	Username string
}

const (
	danger  = "danger"
	success = "good"
	noColor = ""
)

// Create slack wrapper for sending messages. You can leave channel/emoji empty
// Then it will just use settings from your integration
func Init(webhook string, channel string, emoji string, name string) *Slack {
	return &Slack{
		Webhook:  webhook,
		Channel:  channel,
		Emoji:    emoji,
		Username: name,
	}
}

// Send default message to slack
func (slack *Slack) Post(msg string) {
	payload := slack.prepareMsg(msg, noColor)
	slack.send(payload)
}

// Send red colored message to slack
func (slack *Slack) PostDanger(msg string) {
	payload := slack.prepareMsg(msg, danger)
	slack.send(payload)
}

// Send green colored message to slack
func (slack *Slack) PostSuccess(msg string) {
	payload := slack.prepareMsg(msg, success)
	slack.send(payload)
}

// Add all needed options to payload and return it as byte slice
func (slack *Slack) prepareMsg(msg string, color string) []byte {
	data := fmt.Sprintf(`"mrkdwn_in": ["text"], "text": "%s"`, sanitize(msg))
	data = appendOption(data, "color", color)
	data = appendOption(data, "channel", slack.Channel)
	data = appendOption(data, "icon_emoji", slack.Emoji)
	data = appendOption(data, "username", slack.Username)
	return []byte(fmt.Sprintf(`{"attachments": [{%s}]}`, data))
}

// Post request to slack via webhook
func (slack *Slack) send(payload []byte) {
	r, err := http.NewRequest("POST", slack.Webhook, bytes.NewBuffer(payload))
	check(err)
	r.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	defer resp.Body.Close()
	check(err)
}

// Add proper keys to payload
func appendOption(data string, k string, v string) string {
	if present(v) {
		option := fmt.Sprintf(`"%s": "%s"`, k, v)
		return strings.Join([]string{data, option}, ", ")
	}
	return data
}

func present(i string) bool {
	return len(i) > 0
}

func sanitize(s string) string {
	return strings.Replace(s, `"`, "", -1)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
