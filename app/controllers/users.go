package controllers

import (
	"github.com/hackedu/hackedu/app/routes"
	"github.com/robfig/revel"
)

type Users struct {
	Authentication
}

func (c Users) Index() revel.Result {
	user := c.connected()
	if user == nil {
		c.Flash.Error("Please log in first.")
		return c.Redirect(routes.App.Index())
	}
	return c.Render(user)
}
