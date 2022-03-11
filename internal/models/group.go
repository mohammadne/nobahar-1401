package models

type Group struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Users       []User `gorm:"foreignKey:UserRefer"`
}
