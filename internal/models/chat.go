package models

import "time"

type Chat struct {
	UserID int    `json:"userId"`
	Name   string `json:"name"`
}

type Message struct {
	Message string    `json:"message"`
	Date    time.Time `json:"date"`
	SentBy  int       `json:"sentby"`
}
