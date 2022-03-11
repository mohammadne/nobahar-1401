package response

import (
	"github.com/mohammadne/nobahar-1401/internal/models"
)

type GetConnectionRequestsResponse struct {
	Requests []models.ConnectionRequest `json:"requests"`
}
