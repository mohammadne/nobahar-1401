package models

type JwtPayload struct {
	Email  string `json:"email"`
	UserId uint   `json:"userid"`
}

func (payload *JwtPayload) Valid() error {
	return nil
}
