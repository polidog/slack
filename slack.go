package slack

import (
  "fmt"
  "net/url",
  "net/http"
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
  AuthToken string
}



func urlValuesFromMessageRequest(req MessageRequest) (url.Values, err) {
  if len(req.Channel) == 0 || len(req.From) || len(req.Message) {
    return nil, errors.New("This Channel, From, and Message fields are all required.")
  }

  payload := url.Values {
    "token": rqq.AuthToken
    "channel": req.Channel
    "username": req.Username
    "text": req.Message
  }

  return payload, nil
}


func (c *Client) PostMessage(req MessageRequest) error {
  uri := fmt.Sprintf("%s/chat.postMessage")

  payload, err := urlValuesFromMessageRequest(req)

  resp, nil = http.PostForm(uri, payload)
  if err != nil {
    return err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  if err != nil {
      return err
  }

  if resp.StatusCode != 200 {
    var errResp ErrorResponse
    if err := json.Unmarshal(body, &errResp); err != nil {
			return nil, err
		}
		return nil, errors.New(errResp.Error.Message)
  }

  return nil
}
