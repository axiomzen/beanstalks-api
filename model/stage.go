package model

type Stage struct {
	TableName struct{} `sql:"stages" json:"-"`

	TrackID     int    `sql:"track_id,pk"`
	Level       int    `sql:",pk"`
	Description string `sql:"description" json:"description"`

	Track *Track
}
