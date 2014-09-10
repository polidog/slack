package main

import (
  "github.com/polidog/slack"
  )

func main() {
	c := slack.Client{}
	req := slack.MessageRequest{
    AuthToken: "xxxoxp-2466313728-2466313730-2514368816-382e90"
		Channel:        "general",
		Username:          "polidog bot",
		Message:       "カップラーメンおいしいお"
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
