package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
	Authentication
}

func (c App) Index() revel.Result {
	if c.connected() != nil {
		return c.RenderTemplate("Users/Index.html")
	}
	return c.RenderTemplate("App/Index.html")
}
