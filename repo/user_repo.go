package repo

import (
	"fmt"
	"iris-cli/model"
)

type UserRepo interface {
	Add() error
}

func NewUserRepo() UserRepo {
	return &UserRepoImpl{}
}

type UserRepoImpl struct{}

func (u *UserRepoImpl) Add() error {
	var user *model.User
	user.UserName = ""
	//err := db.Db.Save(&user).Error
	fmt.Println("user add one account!!!")
	return nil
}
