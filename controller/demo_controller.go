package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
	"iris-cli/core"
	"iris-cli/repo"
	"iris-cli/service"
)

type DemoController struct {
	// iris 框架自动为每个请求都绑定上下文对象：可作为接受参数
	Context iris.Context
	// session 对象： 储存session信息
	Session *sessions.Session

	// admin功能实体：引入Service 接口
	DemoService service.DemoService
}

var demoRepo repo.DemoRepo

func NewDemoController() *DemoController {
	return &DemoController{
		DemoService: service.NewDemoService(demoRepo),
	}
}

func (dc *DemoController) GetList() (result *core.Result) {
	list, err := dc.DemoService.GetList()
	if err != nil {
		result.Code = 400
		result.Msg = err.Error()
	} else {
		result.Code = 200
		result.Data = list
		result.Msg = "ok"
	}
	return result
}
