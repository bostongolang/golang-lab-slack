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

// cambridge
func GetWeather(place string) (*openweathermap.CurrentWeatherData, error) {
	w, err := openweathermap.NewCurrent("F", "en")

	// w == openweathermap.Weather{
	// 	Temperature int
	// 	Clouds string
	// 	Language string
	// 	Weather
	// 	CurrentByName() &weather
	// }

	if err != nil {
		return nil, fmt.Errorf("Could not get weather: %s", err)
	}

	err = w.CurrentByName(place)
	if err != nil {
		return nil, fmt.Errorf("Weather fetch fail:", err)
	}
	return w, nil
}

func ListenForWeather() {
	bot := slackbot.New(os.Getenv("SLACK_API_TOKEN"))
	bot.Hear("(?i)clouds (.*)").MessageHandler(WeatherHandler)
	bot.Run()
}

func WeatherHandler(ctx context.Context, bot *slackbot.Bot, msg *slack.MessageEvent) {
	parts := strings.Split(msg.Msg.Text, " ") // 'weather cambridge' -> ['weather','cambridge']
	if len(parts) != 2 {                      // 2 == 2
		return
	}

	weather, err := GetWeather(parts[1]) // 'cambridge'
	if err != nil {
		fmt.Println("Could not get weather:", err)
		return
	}

	description := ""
	if len(weather.Weather) > 0 {
		description = weather.Weather[0].Description
	}
	bot.Reply(msg, fmt.Sprintf("The current temperature for %s is %.0f degrees farenheight (%s)",
		weather.Name,
		weather.Main.Temp,
		description),
		slackbot.WithTyping)
}
