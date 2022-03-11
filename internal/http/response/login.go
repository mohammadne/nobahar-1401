package response

type LoginResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
