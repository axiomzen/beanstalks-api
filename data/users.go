package data

import "github.com/axiomzen/beanstalks-api/model"

func (dal *DAL) CreateUser(user *model.User) error {
	_, err := dal.db.Model(user).Insert()
	return err
}
