package main

import (
	"os"

	slackbot "github.com/BeepBoopHQ/go-slackbot"
)

func main() {
	bot := slackbot.New(os.Getenv("SLACK_TOKEN"))
	// do something with bot!
}
