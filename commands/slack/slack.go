package slack

import "github.com/piokaczm/simpleslack"

type slackClient struct {
	*simpleslack.Slack
	available bool
}

// New is a constructor for slackClient struct which is a wrapper
// around simpleslack.Slack. It's here to handle configs without slack info.
func New(webhook, channel, emoji, name string) *slackClient {
	slack := &slackClient{}

	if webhook != "" {
		slack = &slackClient{
			simpleslack.Init(webhook, channel, emoji, name),
			true,
		}
	}
	return slack
}

// PostDanger is a wrapper around orginal implementation of simpleslack.Slack{} PostDanger method.
// It will send a message if and only if available flag on slackClient struct is set to true.
func (s *slackClient) PostDanger(msg string) {
	if s.available {
		s.PostDanger(msg)
	}
}

// PostSuccess is a wrapper around orginal implementation of simpleslack.Slack{} PostSuccess method.
// It will send a message if and only if available flag on slackClient struct is set to true.
func (s *slackClient) PostSuccess(msg string) {
	if s.available {
		s.PostSuccess(msg)
	}
}
