package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlackEnabled(t *testing.T) {
	config := Configuration{}
	assert.Equal(t, slackEnabled(config.Slack), false)
	config.Slack = map[string]string{"webhook": "test"}
	assert.Equal(t, slackEnabled(config.Slack), true)
}
