<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->


- [Part 1: Setup the project (5-10 minutes)](#part-1-setup-the-project-5-10-minutes)
  - [Goals](#goals)
  - [Steps](#steps)
- [Part 2: Get the weather (5-10 minutes)](#part-2-get-the-weather-5-10-minutes)
  - [Goals](#goals-1)
  - [Steps](#steps-1)
- [Part 3: Respond to the `weather <place>` command (5-10 minutes)](#part-3-respond-to-the-weather-place-command-5-10-minutes)
  - [Goals](#goals-2)
  - [Steps](#steps-2)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->




# Part 1: Setup the project (5-10 minutes)

## Goals

The goal of the portion is to setup a project `weatherbot` which will
hold the code for your weather bot project. 

We're going to use two libraries to build our bot:

1. nlopes Slack library - https://github.com/nlopes/slack
1. Open Weather Map http://briandowns.github.io/openweathermap/

## Steps

1. :star2: Create a directory `weatherbot` for your code in your $GOPATH, and
a file`main.go` that will contain the code.

  E.g. under `$GOPATH/bostongolang`
  
    ```bash
    mkdir -p weatherbot 
    cd weatherbot 
    touch weatherbot/weatherbot.go
    ```

1. Create an account on openweathermap and setup your API key.

   1. :star2: Sign up to openweathermap.org and get an API key.
   1. :star2: Set your `OWM_API_KEY` environment variable to the API key you just created, e.g. 

   ```bash
   export OWM_API_KEY=5634bf4c007be5d2f95aceeae356a2ba
   ```

1. Setup your weatherbot Slack API token.

   1. :star2: Set your Slack bot's `SLACK_API_TOKEN` environment variable:

   ```bash
   export SLACK_API_TOKEN=my-token-foo-bar
   ```

1. Get the necessary Go libraries:

  
    1. :star2: Get the slack libraries

    ```bash
    go get github.com/BeepBoopHQ/go-slackbot 
    go get github.com/nlopes/slack
    ```

    1. :star2: Get the openweathermap library

    ```bash
    go get github.com/briandowns/openweathermap 
    ```

1. Populate `weatherbot.go` with a shell to get started:

   ```go
   package main
   
   import (
   	"fmt"
   	"os"
   	"strings"
   
   	"golang.org/x/net/context"
   
   	slackbot "github.com/BeepBoopHQ/go-slackbot"
     "github.com/briandowns/openweathermap"
   	"github.com/nlopes/slack"
   )
   
   func main() {
   	  ListenForWeather()
   }
   
   func GetWeather(place string) (*openweathermap.CurrentWeatherData, error) {
     // TODO: get the weather for the place
   	r
   	turn nil, nil
   }
   
   func ListenForWeather() {
   	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
   	// TODO: add a bot.Hear listener here for a `weather <place>` request
   	bot.Run()
   }
   
   func WeatherHandler(ctx context.Context, bot *slackbot.Bot, evt *slack.MessageEvent) {
     // TODO: Respond to a `weather <place>` request 
   }
   ```
  
   Type `go run weatherbot.go` and verify the code works.

# Part 2: Get the weather (5-10 minutes)

## Goals

In this section, we'll populate the `GetWeather` function that
will connect to openweather using the openweathermap library.

## Steps

1. :star2: Find the `GetWeather` function and add some code
that connects to openweathermap. 

   * Hint: To connect, you'll need to use `NewCurrent` method from the openweathermap library.
   Check out the [README](https://github.com/briandowns/openweathermap) for examples.

1. :star2: Get the weather object for the `place` provided

   * Hint: You'll need to use the `CurrentByName` function. Check out the [README](https://github.com/briandowns/openweathermap) for examples.

1. :star2: To test, you can add a temporary `GetWeather("02145")` call to the `main()` function and print the results.  Run `go run weatherbot.go` to test.

# Part 3: Respond to the `weather <place>` command (5-10 minutes)

## Goals

We need to add a bot.Hear handler that will look for a `weather <place>` command.

## Steps

1. Populate the `ListenForWeather` handler to respond to messages

  * :star2: In the `ListenForWeather` function, identify the "TODO:" statement and add a `bot.Hear` handler that looks for any `weather (.*)` message from Slack.  It should call the `WeatherHandler` response.

  * You can look at [this code](https://github.com/BeepBoopHQ/go-slackbot/blob/master/examples/simple/simple.go#L17) for an example on how to listen for a message.

1. Populate the `WeatherHandler` function

  * :star2: Parse the `place` parameter for weather lookup. Hint: you'll need to look at `evt.Msg.Text` and grab the second part.

  * :star2: Once you have the place, call `GetWeather` for the place to get a response.

  * :star2: Format the response and add a `bot.Reply` that will return something like this:

  ```
  The current temperature for Somerville is 70 degrees farenheight (light rain)
  ```

     * Hint: Look at the `Name`, `Main`, and `Weather[0].Description` fields in the GetWeather() response 

1. :star2: To test, run `go run weatherbot.go` and open a direct message to your bot name.  Type `weather 02145` and see what happens!

    ![img](https://sc-cdn.scaleengine.net/i/3cad57674be60bd6f971a6b1925025b8.png)


 
