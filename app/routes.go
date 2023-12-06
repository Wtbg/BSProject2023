package app

import (
	echoSwagger "github.com/swaggo/echo-swagger"
	"go-svc-tpl/app/controller"
	"go-svc-tpl/app/controller/user"
)

func addRoutes() {
	api := e.Group("api")

	api.GET("/doc/*", echoSwagger.WrapHandler)

	api.GET("/foo", controller.Foo)

	api.POST("/user/register", user.Register)

	api.POST("/user/login", user.Login)
}
