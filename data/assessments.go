package data

import (
	"github.com/axiomzen/beanstalks-api/model"
)

func (dal *DAL) GetAssessmentsByUserID(userID int, assessments *model.Assessments) error {
	return dal.db.Model(assessments).
		Where("user_id = ?", userID).
		Order("created_at DESC").
		ColumnExpr("assessment.*").
		Relation("Reviewer").
		Select()
}
