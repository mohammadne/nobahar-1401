package response

type SignupResponse struct {
	Token   string `json:"token"`
	Message string `json:"message"`
}
