package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris-cli/core"
	"iris-cli/repo"
)

type PicController struct {
	// iris 框架自动为每个请求都绑定上下文对象：可作为接受参数
	Context iris.Context
	// session 对象： 储存session信息
	Session *sessions.Session

	// admin功能实体：引入Service 接口
	UserRepo repo.UserRepo
}

func NewPicController() *PicController {
	return &PicController{}
}

func (dc *PicController) Add() (result *core.Result) {

	return result
}
