import "github.com/nlopes/slack"

func (s *SecurityAlertWorkflow) NotifyUser() {
	params := slack.NewPostMessageParameters()
	params.Username = "deputy"
	s.Status = SENT
	params.Markdown = true
	params.IconEmoji = ":scream:"

	attachments := make([]slack.Attachment, 2)
	attachments[0] = slack.Attachment{
		Pretext: "A suspicious command was ran.",
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Message",
				Value: s.Alert.Message,
			},
		},
	}
	attachments[1] = slack.Attachment{
		Color:   "#36a64f",
		Pretext: "Was this you? [yes/no]?",
	}

	params.Attachments = attachments
	s.bot.Client.PostMessage(s.ChannelID, "", params)
}
