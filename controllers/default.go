package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "peng.edy@gmail.com"
  c.Data["EmailName"] = "Alvin Qi"
	c.TplName = "index.tpl"
}

func (c *MainController) HelloSitepoint() {
  c.Data["Website"] = "beego.me"
  c.Data["Email"] = "peng.edy@gmail.com"
  c.Data["EmailName"] = "Alvin Qi"
  c.TplName = "default/hello-sitepoint.tpl"
}
