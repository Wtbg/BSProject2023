// @title        Golang Service Template
// @version      0.1
// @description  Golang back-end service template, get started with back-end projects quickly
// @BasePath     /api

package main

import (
	"github.com/sirupsen/logrus"
	"go-svc-tpl/app"
	_ "go-svc-tpl/docs"
	"go-svc-tpl/model"
	"go-svc-tpl/mqtt"
)

func main() {
	logrus.SetReportCaller(true)
	logrus.SetLevel(logrus.ErrorLevel)
	model.Init()
	app.InitWebFramework()
	mqtt.Init()
	app.StartServer()
}
