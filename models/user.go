package models

type User struct {
	ID                   string      `json:"id"`
	Name                 string      `json:"name"`
	Email                string      `json:"email"`
	Role                 string      `json:"role"`
	Tags                 []string    `json:"tags"`
	MostRecentAssessment *Assessment `json:"mostRecentAssessment"`
}
