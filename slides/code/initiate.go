func (s *SecurityAlertWorkflow) OpenChat() {
	if users, err := s.bot.Client.GetUsers()
	if err != nil {
		panic(err)
	}

	userId := ""
	for _, user := range users {
		if user.Name == s.User.Username {
			userId = user.ID
			break
		}
	}
	if userId == "" {
		panic("User not found")
	}

	_, _, channel, err := s.bot.Client.OpenIMChannel(userId)
	if err != nil {
		panic(err)
	}
	s.ChannelID = channel
}

