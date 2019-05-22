package data

import "github.com/axiomzen/beanstalks-api/model"

func (dal *DAL) CreateUser(user *model.User) error {
	_, err := dal.db.Model(user).Insert()
	return err
}

func (dal *DAL) GetUserByID(user *model.User) error {
	return dal.db.Model(user).WherePK().Select()
}

func (dal *DAL) GetAllUsers(users *model.Users) error {
	return dal.db.Model(users).Select()
}

func (dal *DAL) GetUserByEmail(user *model.User) error {
	return dal.db.Model(user).Where("email = ?email").Select()
}
