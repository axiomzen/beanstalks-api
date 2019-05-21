package models

type Assessment struct {
	ID        string    `json:"id"`
	UserId    string    `json:"userId"`
	CreatedAt string    `json:"createdAt"`
	Reviewer  *Reviewer `json:"reviewer"`
	Comments  string    `json:"comments"`
	State     string    `json:"state"`
	Tracks    []*Track  `json:"tracks"`
}

type Reviewer struct {
}
