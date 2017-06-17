package slack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("with no webhook provided", func(t *testing.T) {
		slack := New("", "channel", "emoji", "name")
		assert.False(t, slack.available, "marks slack client as unavailable")
	})

	t.Run("with webhook provided", func(t *testing.T) {
		slack := New("testwebhook", "channel", "emoji", "name")
		assert.True(t, slack.available, "marks slack client as available")
	})
}

func TestPostDanger(t *testing.T) {
	t.Run("when client disabled", func(t *testing.T) {
		slack := New("", "channel", "emoji", "name")
		assert.NotPanics(t, func() { slack.PostDanger("test") }, "doesnt try to send the message")
	})

	t.Run("when client available", func(t *testing.T) {
		slack := New("testwebhook", "channel", "emoji", "name")
		assert.Panics(t, func() { slack.PostDanger("test") }, "tries to send the message")
	})
}

func TestPostSuccess(t *testing.T) {
	t.Run("when client disabled", func(t *testing.T) {
		slack := New("", "channel", "emoji", "name")
		assert.NotPanics(t, func() { slack.PostSuccess("test") }, "doesnt try to send the message")
	})

	t.Run("when client available", func(t *testing.T) {
		slack := New("testwebhook", "channel", "emoji", "name")
		assert.Panics(t, func() { slack.PostSuccess("test") }, "tries to send the message")
	})
}
