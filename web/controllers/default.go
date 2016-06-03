package controllers

import (
	"51job/cons"
	"51job/web/models"
	"github.com/astaxie/beego/utils/pagination"
	_ "github.com/beego/i18n"
	"strconv"
)

type MainController struct {
	baseController
}

func (c *MainController) Get() {
	keywordid := c.GetString("k")
	sex := c.GetString("s")
	k := models.ListKeyword()
	for _, kw := range k {
		temp, ok := cons.KindMap[kw.Kind]
		if ok {
			kw.Kind = temp
		}
	}
	usercount := models.CountUser(keywordid, sex)
	postsPerPage := 20
	paginator := pagination.SetPaginator(c.Ctx, postsPerPage, usercount)
	offset := paginator.Offset()

	users := models.ListUser(sex, keywordid, postsPerPage, offset)

	c.Data["User"] = users
	c.Data["Count"] = usercount
	c.Data["Keyword"] = k
	c.Data["Pk"] = keywordid

	if keywordid != "" {
		ss, e := strconv.Atoi(keywordid)
		if e == nil {
			kone := models.ListOneKeyword(ss)
			temp, ok := cons.KindMap[kone.Kind]
			if ok {
				kone.Kind = temp
			}
			c.Data["K"] = kone
		}
	}
	c.TplName = "index.tpl"
}
