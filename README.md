# Golang Boston Lab #3 - Building a Slack bot 

Welcome to Boston Golang Lab #3 - Building a Slack bot

Format of this meetup:

1. (10-15 minute) introduction to building bots in Slack

1. (60 minutes) Choose your own adventure:
   - Go through the example `weatherbot` exercise
     
     *or* 

   - Build your own slack bot (form teams of 2 or more!)

1. (15 minutes) Battle of the bots! Show off what you've done

## Prerequisites 

1. PLEASE BRING:  A laptop with a Go language environment with Go 1.6 already set up. Please see 'Vagrant setup' below to get started with one easily.
  * If you are having trouble setting up your Go programming environment, please join the #lab-help channel in the Boston Golang Slack group. You can signup for the Slack group [here](http://bostongolang-slack-invite.herokuapp.com/).

2. PLEASE BRING: A text editor or IDE suitable for writing Go code.
  * For beginners, [Sublime](http://www.sublimetext.com) is a good option. Make sure you install the Go plugin [here](https://github.com/DisposaBoy/GoSublime).
  * For VIM users, there is a pretty nice VIM setup [here](https://github.com/fatih/vim-go).  
  * If you'd like, you can also use the Vagant setup which has a Go environment set up (See [VAGRANT_SETUP.md](VAGRANT_SETUP.md)).

3. PLEASE HAVE:  Some basic Go language exposure.  You should be familiar with the Go basics: e.g., Go's types, structs, control flow structures, goroutines, and channels.   
Other programming language experience and concepts (such as networking, etc) will be helpful. A good introduction to the basics of Go for people familiar with 
other programming language is available at: [https://tour.golang.org](https://tour.golang.org). If you can get through this tour, you will be well-prepared for this meetup!

### General setup

1. Fork this repository in Github.

1. Clone the repository into a directory
  
  ```bash
  # open a terminal window and type:
  $ git clone https://github.com/bostongolang/golang-lab-slack.git
  ```

1. Sign up for a Slack account on our Slack group: [http://bostongolang-slack-invite.herokuapp.com/](http://bostongolang-slack-invite.herokuapp.com/)

   * Log in at [bostongolang.slack.com](https://bostongolang.slack.com).

1. Join the #bots channel on Slack with your main account (for communicating).

1. Create a bot integration in Slack:

  1. Go to [https://bostongolang.slack.com/apps/build/custom-integration](https://bostongolang.slack.com/apps/build/custom-integration)
  1. Choose 'bots'
  1. Select a unique username to create your bot under, e.g. 'jenbot'
  1. Click 'Add bot integration to finish adding your integration' and view the settings to get the API key for your bot (required)

## Labs 

### For beginners: weatherbot 

Weatherbot responds to requests for the weather.

Build a Slack bot that responds to `weather <zip code|name of place>` and return the weather for the zip code.

[Try it out now!](exercises/weatherbot/README.md)

### For advanced: build your own!

Ideas:

  * untableflip: resets a 'tableflip' emoji
  * TBD
