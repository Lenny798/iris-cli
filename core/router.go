package core

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-cli/controller"
)

func InitRouter(app *iris.Application) {
	baseUrl := "/api/v1"
	mvc.New(app.Party(baseUrl + "/demo")).Handle(controller.NewDemoController())
}
