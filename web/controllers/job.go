package controllers

import (
	"51job/cons"
	"51job/web/models"
	_ "github.com/beego/i18n"
	"io/ioutil"
)

type JobController struct {
	baseController
}

func (this *JobController) Get() {
	id := this.GetString("id")
	kid := this.GetString("k")
	if id == "" {
		this.Abort("404")
	}
	brother := models.UserBrotherByKeyword(kid, id)
	uh := models.ListUserHistroy(id)
	ui := models.ListUserInfo(id)
	/*


	 */
	for _, uhone := range uh {
		temp := uhone.Keyword.Kind
		temp1, ok := cons.KindMap[temp]
		if ok {
			uhone.Keyword.Kind = temp1
		}

	}

	/**/
	this.Data["UserHistroy"] = uh
	this.Data["User"] = ui
	this.Data["Brotherp"] = brother[0]
	this.Data["Brothern"] = brother[1]
	this.Data["Pk"] = kid
	this.TplName = "userhistroy.tpl"
}

func (this *JobController) Intro() {
	this.TplName = "intro.tpl"
}

func (this *JobController) Img() {
	id := this.GetString("id")
	if id == "" {
		this.StopRun()
	}
	data, e := ioutil.ReadFile("../data/img/" + id + ".jpg")
	if e != nil {
		this.StopRun()
	}
	this.Ctx.ResponseWriter.Write(data)
}

func (this *JobController) ImgPika() {
	id := this.GetString("id")
	if id == "" {
		data, _ := ioutil.ReadFile("static/img/pika1.jpg")
		this.Ctx.ResponseWriter.Write(data)
		this.StopRun()
	}
	data, e := ioutil.ReadFile("../data/img/" + id + ".jpg")
	if e != nil {
		data, _ := ioutil.ReadFile("static/img/pika1.jpg")
		this.Ctx.ResponseWriter.Write(data)
		this.StopRun()
	}
	this.Ctx.ResponseWriter.Write(data)
}
