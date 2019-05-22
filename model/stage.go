package model

type Stage struct {
	TableName struct{} `sql:"stages" json:"-"`

	ID          int    `sql:",pk" json:"id"`
	TrackID     int    `sql:"track_id" json:"trackId"`
	Description string `sql:"description" json:"description"`
	Level       int    `json:"level"`
}
