package controllers

import (
	"golang_tutorial/blogsystem/models"

	"github.com/astaxie/beego"
)

// IndexController struct
type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	beego.BConfig.WebConfig.ViewsPath = "views"
	this.Data["blogs"] = models.GetAll()
	this.Layout = "layout.tpl"
	this.TplName = "views/index.tpl"
}
