package controllers

import (
	"Editor_NaN/EditorApi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"net/http"
)

type HomeController struct {
	beego.Controller
}

func (c *HomeController) Home() {

	if c.Ctx.Request.Method == http.MethodPost {
		user := models.User{
			Email:    c.GetString("email"),
			Password: c.GetString("password"),
		}

		valid := validation.Validation{}
		isOk, _:= valid.Valid(&user)

		if isOk {
			u, _ := models.SelectUser(user.Email)
			if u != nil {
				c.StartSession()
				c.SetSession("user", u)
				c.Ctx.Redirect(http.StatusSeeOther, "/play")
				return
			}
		}
	}

	c.Layout = "layout/index.html"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["HtmlHead"] = "fragment/loginCss.html"
	c.TplName = "login.html"
	c.Render()
}

func (c *HomeController) Play() {

	c.StartSession()
	if c.GetSession("user") == nil {
		c.Ctx.Redirect(http.StatusSeeOther, "/")
		return
	}

	c.Layout = "layout/index.html"
	c.LayoutSections = map[string]string{}
	c.LayoutSections["HtmlHead"] = "fragment/play.html"
	c.TplName = "play.html"
	c.Render()
}