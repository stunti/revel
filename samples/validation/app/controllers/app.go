package controllers

import "github.com/robfig/revel"

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
	c.Layout = "layout.html"
	return c.Render()
}
