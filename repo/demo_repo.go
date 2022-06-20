package repo

import (
	"iris-cli/db"
	"iris-cli/model"
)

type DemoRepo interface {
	GetList() (*[]model.Book, error)
}

func NewDemoRepo() DemoRepo {
	return &demoRepoImpl{}
}

type demoRepoImpl struct{}

func (d *demoRepoImpl) GetList() (*[]model.Book, error) {
	var books *[]model.Book
	if err := db.Db.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}
