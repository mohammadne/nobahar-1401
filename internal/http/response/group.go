package response

import "github.com/mohammadne/nobahar-1401/internal/models"

type CreateGroupResponse struct {
	Group struct {
		ID int `json:"id"`
	} `json:"group"`
	Message string `json:"message"`
}

type GetGroupsResponse struct {
	Groups []models.Group `json:"groups"`
}

type GetMyGroupsResponse struct {
	Groups []models.Group `json:"groups"`
}
