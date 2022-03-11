package request

type JoinRequest struct {
	GroupID int `json:"groupId"`
}

type AcceptJoinRequest struct {
	JoinRequestID int `json:"joinRequestId"`
}
