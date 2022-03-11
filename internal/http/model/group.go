package model

type Member struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Rule  string `json:"rule"`
}

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DetailedGroup struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Members     []Member
}
