package response

import (
	"github.com/mohammadne/nobahar-1401/internal/http/model"
)

type GetConnectionRequestsResponse struct {
	Requests []model.ConnectionRequest `json:"requests"`
}
