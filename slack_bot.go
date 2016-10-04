package main

import (
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {

	var key = flag.String("k", os.Getenv("SLACK_BOT_API_KEY"), "The Slack API key for the bot")
	flag.Parse()

	if *key == "" {
		fmt.Println("You must supply an API key via the -k flag or by setting the SLACK_BOT_API_KEY environment variable")
		os.Exit(1)
	}

	api := slack.New(*key)

	users, err := api.GetUsers()

	if err != nil {
		fmt.Printf("Erorr: %s", err)
	}

	for _, user := range users {
		fmt.Printf("Name: %s\tID: %s\n", user.RealName, user.ID)
	}
}
