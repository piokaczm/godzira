# Simple Slack

Do you need a simple Slack alerting for your golang app, but you don't want to use huge libs with gazilion features?
Then *Simple Slack* is for you!

There are no fancy options, just possibility to send basic messages via webhook, but thanks to it you end up with really lightweight lib.

## Usage

*Initialization*

If you're a lazy-ass you can initialize it like this:

```
slack := slack.Init("webhook", "channel", "emoji" "name")
```

Which will return you `*slack.Slack`, if you don't need `channel`, `emoji`, or `name` you can just pass empty string and these options will be ignored while sending messages.

Otherwise you can initialize it manually:

```
slack := &Slack{
    Webhook:  webhook,
    Channel:  channel,
    Emoji:    emoji,
    Username: name,
  }
```

*Posting messages*

Simple Slack provides three methods for posting:

```
slack := slack.Init("webhook", "channel", "emoji" "name")

// default message
slack.Post("normal message")

// red colored message
slack.PostDanger("danger message")

// green colored message
slack.PostSuccess("success message")
```
