package models

type User struct {
	ID      uint   `gorm:"primarykey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	RoleID  uint
	Role    Role
	GroupId uint
	Group   Group
}
