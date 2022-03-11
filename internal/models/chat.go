package models

type Chat struct {
	ID     uint   `gorm:"primarykey"`
	UserID int    `json:"userId"`
	Name   string `json:"name"`
}
