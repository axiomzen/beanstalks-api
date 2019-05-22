package model

import "time"

type Assessment struct {
	TableName struct{} `sql:"assessments" json:"-"`

	ID         int       `sql:",pk" json:"id"`
	UserID     int       `sql:"user_id" json:"userId"`
	ReviewerID int       `sql:"reviewer_id" json:"reviewerId"`
	State      string    `sql:"state" json:"state"`
	CreatedAt  time.Time `sql:"created_at" json:"createdAt"`

	Reviewer *User `json:"reviewer"`
}

type Assessments []*Assessment
