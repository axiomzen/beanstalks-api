package data

import (
	"github.com/axiomzen/beanstalks-api/model"
)

func (dal *DAL) GetAssessmentsByUserID(userID int, assessments *model.Assessments) error {
	return dal.db.Model(assessments).
		Where("user_id = ?", userID).
		Order("created_at ASC").
		ColumnExpr("assessment.*").
		Relation("Reviewer").
		Select()
}

func (dal *DAL) CreateAssessment(assessment *model.Assessment) error {
	_, err := dal.db.Model(assessment).Insert()
	return err
}

func (dal *DAL) GetAssessmentByPK(assessment *model.Assessment) error {
	return dal.db.Model(assessment).WherePK().Select()
}
