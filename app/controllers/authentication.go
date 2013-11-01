package controllers

import "github.com/robfig/revel"

type Authentication struct {
	*revel.Controller
}

func (c Authentication) Login() revel.Result {
	return c.Render()
}

func (c Authentication) SignUp() revel.Result {
	return c.Render()
}
