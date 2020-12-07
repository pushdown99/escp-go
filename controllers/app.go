package controllers

import (
	"net/http"

	"github.com/gernest/utron/controller"
)

type App struct {
	controller.BaseController
	Routes []string
}

func (a *App) Index() {
	a.Ctx.Template = "index"
	a.HTML(http.StatusOK)
}

func (a *App) SignIn() {
	a.Ctx.Template = "signin"
	a.HTML(http.StatusOK)
}

func NewApp() controller.Controller {
	return &App{
		Routes: []string{
			"get;/;Index",
		},
	}
}
