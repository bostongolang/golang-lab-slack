package main

import (
	"fmt"
	"os"

	"golang.org/x/net/context"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
	"github.com/nlopes/slack"
)

func listenForHumanWeakness() {
	bot := slackbot.New(os.Getenv("SPARKY_SLACK_API_TOKEN"))
	bot.Hear("(fail|tried|err|break|broke|won(|')t|can(|')t)").MessageHandler(humanFailureHandler)
	bot.Run()
}

func humanFailureHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	bot.Reply(msg, fmt.Sprintf("Suckas! For robots, perfection is effortless."), slackbot.WithTyping)
}

func main() {
	fmt.Println("I am sparky and I am sparking.")
	fmt.Println("My slack token is", os.Getenv("SPARKY_SLACK_API_TOKEN"))
	listenForHumanWeakness()
}
