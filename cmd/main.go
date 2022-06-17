package main

import (
	"github.com/kataras/iris/v12"
	"github.com/spf13/viper"
	"iris-cli/core"
)

func main() {
	app := iris.New()
	app.Configure(iris.WithOptimizations)
	app.AllowMethods(iris.MethodOptions)

	core.InitRouter(app)
	core.InitConfig()

	err := app.Run(iris.Addr(":"+viper.GetString("Port")), iris.WithoutServerError(iris.ErrServerClosed))
	if err != nil {
		panic(err)
	}
}
