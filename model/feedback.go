package model

type Feedback struct {
	TableName struct{} `sql:"feedback" json:"-"`

	AssessmentID    int `sql:",pk"`
	TrackID         int `sql:"track_id,pk"`
	Feedback        string
	Examples        string
	Recommendations string

	Assessment *Assessment
	Track      *Track
}
