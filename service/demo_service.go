package service

import (
	"iris-cli/model"
	"iris-cli/repo"
)

// DemoService 接口
type DemoService interface {
	GetList() (*[]model.Book, error)
}

func NewDemoService(demoRepo repo.DemoRepo) DemoService {
	return &demoServiceImpl{DemoRepo: demoRepo}
}

type demoServiceImpl struct {
	DemoRepo repo.DemoRepo
}

func (d *demoServiceImpl) GetList() (*[]model.Book, error) {
	// Do logic
	list, err := d.DemoRepo.GetList()
	return list, err
}
