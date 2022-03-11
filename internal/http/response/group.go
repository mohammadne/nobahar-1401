package response

import "github.com/mohammadne/nobahar-1401/internal/http/model"

type CreateGroupResponse struct {
	Group struct {
		ID int `json:"id"`
	} `json:"group"`
	Message string `json:"message"`
}

type GetGroupsResponse struct {
	Groups []model.Group `json:"groups"`
}

type GetMyGroupsResponse struct {
	Groups []model.DetailedGroup `json:"groups"`
}
