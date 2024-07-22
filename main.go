package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kheob/slack-status-app/slack"
)

func main() {
	token := os.Getenv("SLACK_STATUS_APP_TOKEN")
	if token == "" {
		fmt.Println("SLACK_STATUS_APP_TOKEN environment variable not set")
		return
	}

	emoji := flag.String("emoji", "", "status emoji to set (e.g. :lunch: or 🥪)")
	text := flag.String("status", "", "status text to set")
	expires := flag.Int("expires", 60, "status expiration in minutes")
	flag.Parse()

	s := slack.New(token)
	err := s.SetStatus(*emoji, *text, *expires)
	if err != nil {
		fmt.Println("error setting status:", err)
	}

	fmt.Println("Status updated successfully!")
}
