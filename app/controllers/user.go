package controllers

import (
	"github.com/revel/revel"
	"proj/girls/app/routes"
	//"proj/girls/app/models"
)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	_, ok := c.Session["username"]
	if !ok {
		c.RenderArgs["message"] = "Welcome to Login!!!"
		return c.RenderTemplate("user/login.html")
	}
	return c.Redirect(routes.Index.Index())
}

func (c User) Register() revel.Result {
	c.RenderArgs["message"] = "Welcome to Register!!!"
	return c.RenderTemplate("user/register.html")
}

func (c User) Login() revel.Result {
	username := c.Params.Get("username")
	password := c.Params.Get("password")
	if !(username == "anaven" || password == "123456") {
		c.Flash.Success("Login Failed")
		return c.Redirect(routes.Index.Index())
	}
	c.Session["username"] = username
	c.Flash.Success("Welcome, " + username)
	return c.Redirect(routes.Index.Index())
}

func (c User) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Index.Index())
}


