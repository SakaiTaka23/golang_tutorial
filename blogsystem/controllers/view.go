package controllers

import (
	"golang_tutorial/blogsystem/models"
	"strconv"

	"github.com/astaxie/beego"
)

// ViewController struct
type ViewController struct {
	beego.Controller
}

func (this *ViewController) Get() {
	id, _ := strconv.Atoi(this.Ctx.Input.Params[":id"])
	this.Data["Post"] = models.GetBlog(id)
	this.Layout = "layout.tpl"
	this.TplName = "view.tpl"
}
