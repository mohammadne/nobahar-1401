package models

type Group struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AdminID     uint
	Users       []User `gorm:"foreignKey:UserRefer"`
}

type DetailedGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Users       []User
}
