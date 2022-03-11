package jwt

type Payload struct {
	Email  string `json:"email"`
	UserId uint   `json:"userid"`
}

func (payload *Payload) Valid() error {
	return nil
}
