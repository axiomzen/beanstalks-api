package data

import "github.com/axiomzen/beanstalks-api/model"

func (dal *DAL) GetScoresByAssessmentID(assessmentID int, scores *model.Scores) error {
	return dal.db.Model(scores).
		Where("assessment_ID = ?", assessmentID).
		Relation("Track").
		Relation("Stage").
		Select()
}

func (dal *DAL) CreateScore(score *model.Score) error {
	_, err := dal.db.Model(score).Insert()
	return err
}
