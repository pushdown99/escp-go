package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
)

type App struct {
	controller.BaseController
	Routes []string
}

func (a *App) Setting() {
	a.Ctx.Template = "Application/Index"
	a.Ctx.Data["title"] = "Sign in"
	a.HTML(http.StatusOK)
}

func NewApp() controller.Controller {
	return &App{
		Routes: []string{
			"get;/;Setting",
		},
	}
}
