package response

import "time"

type JoinResponse struct {
	ID      int       `json:"id"`
	GroupID int       `json:"groupId"`
	UserID  int       `json:"userId"`
	Date    time.Time `json:"date"`
}

type JoinListResponse struct {
	JoinRequests []JoinResponse `json:"joinRequests"`
}
