package request

type CreateGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
