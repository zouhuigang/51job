package routers

import (
	"51job/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//简历查看
	beego.Router("/", &controllers.MainController{})
	beego.Router("/job", &controllers.JobController{})
	beego.Router("/intro", &controllers.JobController{}, "*:Intro")
	beego.Router("/img", &controllers.JobController{}, "*:Img")
	beego.Router("/imgp", &controllers.JobController{}, "*:ImgPika")
}
