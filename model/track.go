package model

type Track struct {
	TableName struct{} `sql:"tracks" json:"-"`

	ID          int      `sql:",pk" json:"id"`
	Name        string   `sql:"name" json:"name"`
	Description string   `sql:"description" json:"description"`
	Tags        []string `sql:",array" json:"tags"`
}
