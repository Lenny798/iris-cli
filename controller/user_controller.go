package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris-cli/core"
	"iris-cli/repo"
)

type UserController struct {
	// iris 框架自动为每个请求都绑定上下文对象：可作为接受参数
	Context iris.Context
	// session 对象： 储存session信息
	Session *sessions.Session

	// admin功能实体：引入Service 接口
	UserRepo repo.UserRepo
}

func NewUserController() *UserController {
	return &UserController{
		UserRepo: repo.NewUserRepo(),
	}
}

func (dc *UserController) Add() (result *core.Result) {
	err := dc.UserRepo.Add()
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Msg = "ok"
	}
	return result
}

func (dc *UserController) Login() (result *core.Result) {
	fmt.Println("user login")
	return result
}
