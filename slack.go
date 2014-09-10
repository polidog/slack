package slack

import (
  "fmt"
  "net/url"
  "net/http"
  "errors"
  "io/ioutil"
  "encoding/json"
)

const (
	baseURL = "https://slack.com/api"

	ColorYellow = "yellow"
	ColorRed    = "red"
	ColorGreen  = "green"
	ColorPurple = "purple"
	ColorGray   = "gray"
	ColorRandom = "random"

	FormatText = "text"
	FormatHTML = "html"

	ResponseStatusSent = "sent"
)


type MessageRequest struct {
  // hipcaht to RoomID
  Channel string

  Username string

  AuthToken string

  Message string

  MessageFormat string

  Notify bool

  Color string

  Token string

  IconUrl string
}


type ErrorResponse struct {
  Error struct {
    Code int
    Type string
    Message string
  }
}


type Client struct {

}



func urlValuesFromMessageRequest(req MessageRequest) (url.Values, error) {

  if len(req.Channel) == 0 || len(req.Username) == 0 || len(req.Message) == 0 {
    return nil, errors.New("This Channel, From, and Message fields are all required.")
  }

  payload := url.Values{
    "token": {req.AuthToken},
    "channel": {req.Channel},
    "username": {req.Username},
    "text": {req.Message},
  }

  return payload, nil
}


func (c *Client) PostMessage(req MessageRequest) error {
  uri := fmt.Sprintf("%s/chat.postMessage")

  payload, err := urlValuesFromMessageRequest(req)
  if err != nil {
    return err
  }

  resp, err := http.PostForm(uri, payload)
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
      return err
  }

  msgResp := &struct{ Status string }{}
	if err := json.Unmarshal(body, msgResp); err != nil {
		return err
	}
	if msgResp.Status != ResponseStatusSent {
		return errors.New("PostMessage: response 'status' field was not 'sent'.")
	}

  return nil
}
