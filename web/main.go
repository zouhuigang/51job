package main

import (
	_ "51job/web/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func main() {
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.Run()

}
