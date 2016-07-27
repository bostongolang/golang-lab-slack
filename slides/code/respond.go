package main

import (
	"os"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))

	toMe := bot.Messages(slackbot.DirectMessage, slackbot.DirectMention).Subrouter()
	toMe.Hear("(?i)(hi|hello).*").MessageHandler(HelloHandler)
	bot.Run()
}

func HelloHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
	bot.Reply(evt, "Oh hello!", slackbot.WithTyping)
}
