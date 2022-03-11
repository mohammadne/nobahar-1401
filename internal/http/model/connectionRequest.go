package model

import "time"

type ConnectionRequest struct {
	ID      int       `json:"connectionRequestId"`
	GroupID int       `json:"groupId"`
	Sent    time.Time `json:"sent"`
}
