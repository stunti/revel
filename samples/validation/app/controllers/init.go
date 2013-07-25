package controllers

import "github.com/robfig/revel"

func init() {
	revel.InterceptMethod((*Sample1).Setup, revel.BEFORE)
}
