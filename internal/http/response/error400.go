package response

type err struct {
	EnMessage string `json:"enMessage"`
}

type Err400Response struct {
	Err err `json:"error"`
}

func NewErr400Response(message string) *Err400Response {
	return &Err400Response{Err: err{EnMessage: message}}
}
