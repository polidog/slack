package main

import (
  "github.com/polidog/slack"
  )

func main() {
	c := hipchat.Client{AuthToken: "xxxoxp-2466313728-2466313730-2514368816-382e90"}
	req := hipchat.MessageRequest{
		Channel:        "general",
		Username:          "polidog bot",
		Message:       "カップラーメンおいしいお"
	}
	if err := c.PostMessage(req); err != nil {
		log.Printf("Expected no error, but got %q", err)
	}
}
