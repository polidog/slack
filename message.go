package slack

import (
  "time"
  )

const (
  ISO8601 = "2006-01-02T15:04:05-0700"
  )

type Message struct {
  ISODate string `json:"date"`

  From struct {
    Name: string
    UserId: interface{} `json:"date"`
  }

  Message string

  File struct {
    Name string
    Size int
    URL string
  }
}

func (m *Message) Time() (time.Time, error){
  return time.Parse(ISO8601, m.ISODate)
}
