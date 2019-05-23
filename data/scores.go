package data

import (
	"github.com/axiomzen/beanstalks-api/model"
)

func (dal *DAL) GetScoresByAssessmentID(assessmentID int, scores *model.Scores) error {
	return dal.db.Model(scores).
		Where("assessment_id = ?", assessmentID).
		Relation("Stage").
		Relation("Track").
		Select()
}

func (dal *DAL) CreateScore(score *model.Score) error {
	_, err := dal.db.Model(score).Insert()
	return err
}

func (dal *DAL) UpdateScoreByPK(score *model.Score) error {
	_, err := dal.db.Model(score).WherePK().Update()
	return err
}
