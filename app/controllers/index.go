package controllers

import "github.com/revel/revel"

type Index struct {
	*revel.Controller
}

func (c Index) Index() revel.Result {
	username, ok := c.Session["username"]
	greeting := "Hello"
	if ok {
		greeting += " " + username
	} else {
		greeting += " Boddy"
	}
	return c.Render(greeting)
}