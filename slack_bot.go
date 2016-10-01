package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	key := os.Getenv("SLACK_BOT_API_KEY")
	api := slack.New(key)

	users, err := api.GetUsers()

	if err != nil {
		fmt.Printf("Erorr: %s", err)
	}

	for _, user := range users {
		fmt.Printf("Name: %s\tID: %s\n", user.RealName, user.ID)
	}
}
