package models

type Role struct {
	ID   uint     `gorm:"primarykey"`
	Type RoleType `json:"name"`
}

type RoleType uint8

const (
	Admin RoleType = iota + 1
	Member
	None
)
