package commands

import (
	"testing"
)

func TestSlackEnabled(t *testing.T) {
	config := Configuration{}
	Expect(t, slackEnabled(config.Slack), false)
	config.Slack = map[string]string{"webhook": "test"}
	Expect(t, slackEnabled(config.Slack), true)
}
