package models

import (
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	Message string `json:"message"`
	SentBy  int    `json:"sentby"`
}
