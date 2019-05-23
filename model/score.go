package model

type Score struct {
	TableName struct{} `sql:"scores" json:"-"`

	AssessmentID int `sql:",pk" json:"assessmentId"`
	TrackID      int `sql:"track_id,pk" json:"trackId"`
	StageID      int `sql:"stage_id,pk" json:"stageId"`
	Score        int `sql:"score,notnull" json:"score"`

	Track *Track `json:"track"`
	Stage *Stage `json:"stage"`
}

type Scores []*Score
