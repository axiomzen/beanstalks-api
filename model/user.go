package model

import "time"

type User struct {
	TableName struct{} `sql:"users" json:"-"`

	ID             int       `sql:",pk" json:"id"`
	Name           string    `sql:"name" json:"name"`
	Email          string    `sql:"email" json:"email"`
	HashedPassword string    `sql:"hashed_password" json:"-"`
	Role           string    `sql:"role" json:"role"`
	Tags           []string  `sql:",array" json:"tags"`
	CreatedAt      time.Time `sql:"created_at" json:"createdAt"`
}

type Users []*User
