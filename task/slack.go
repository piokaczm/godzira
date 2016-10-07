package main

import (
	"bytes"
	"fmt"
	"net/http"
)

const (
	green = "good"
	red   = "danger"
)

type Slack struct {
	Webhook string `yaml:"webhook"`
	AppName string `yaml:"name"`
	BotName string `yaml:"botname"`
	Emoji   string `yaml:"emoji"`
	User    string
	Env     string
}

func (s *Slack) start() {
	msg := fmt.Sprintf(":rocket: %s has *started* deploying %s to *%s*!", s.User, s.AppName, s.Env)
	s.send(msg, green)
}

func (s *Slack) finish() {
	msg := fmt.Sprintf(":star2: %s has *finished* deploying %s to *%s*!", s.User, s.AppName, s.Env)
	s.send(msg, green)
}

func (s *Slack) err() {
	msg := ":crocodile: Oooops! Something went *wrong*!"
	s.send(msg, red)
}

func (s *Slack) send(msg string, color string) {
	data := fmt.Sprintf(`{"attachments": [{"mrkdwn_in": ["text"], "color": "%s", "text": "%s"}]}`, msg, color)
	r, _ := http.NewRequest("POST", s.Webhook, bytes.NewBuffer([]byte(data)))
	r.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
