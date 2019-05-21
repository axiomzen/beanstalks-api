package model

type Score struct {
	TableName struct{} `sql:"scores" json:"-"`

	AssessmentID int `sql:",pk"`
	TrackID      int `sql:"track_id,pk"`
	StageID      int `sql:"stage_id,pk"`
	Score        int

	Assessment *Assessment
	Track      *Track
	Stage      *Stage
}
