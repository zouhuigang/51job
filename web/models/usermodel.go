package models

import (
	"51job/model"
	// "fmt"
	"github.com/astaxie/beego/orm"
	"strconv"
)

var o = orm.NewOrm()

func init() {
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
}

func ListUser(sex string, keyword string, limit interface{}, args ...interface{}) []*model.User {
	if keyword == "" {
		var users []*model.User
		user := new(model.User)
		us := o.QueryTable(user)
		if sex == "1" {
			us.Limit(limit, args...).Filter("Sex", "女").OrderBy("-Date51").All(&users)
		} else if sex == "0" {
			us.Limit(limit, args...).Filter("Sex", "男").OrderBy("-Date51").All(&users)
		} else {
			us.Limit(limit, args...).OrderBy("-Date51").All(&users)
		}
		return users

	}

	keywordid, e := strconv.Atoi(keyword)
	if e != nil {
		return nil
	}

	var lists []orm.ParamsList
	var num = int64(0)
	if sex == "1" {
		num, _ = o.Raw("select a.id51 from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=? and a.sex=? order by a.date51 desc,b.created limit ?,?", keywordid, "女", args[0], limit).ValuesList(&lists)
	} else if sex == "0" {
		num, _ = o.Raw("select a.id51 from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=? and a.sex=? order by a.date51 desc,b.created limit ?,?", keywordid, "男", args[0], limit).ValuesList(&lists)
	} else {
		num, _ = o.Raw("select a.id51 from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=? order by a.date51 desc,b.created limit ?,?", keywordid, args[0], limit).ValuesList(&lists)
	}
	numi := int(num)
	if numi > 0 {
		var userss = make([]*model.User, 0, numi)
		for i := 0; i < numi; i++ {
			u := model.User{Id51: lists[i][0].(string)}
			o.Read(&u, "Id51")
			// fmt.Println(lists[i][0].(string))
			userss = append(userss, &u)

		}
		/*		for i, j := range userss {
				fmt.Printf("%d:%v", i, j)
			}*/
		return userss
	}
	return nil
}

func UserBrotherByKeyword(kid string, id string) []string {
	returns := make([]string, 0, 2)
	if kid == "" {
		tempu := model.User{Id51: id}
		o.Read(&tempu, "Id51")

		var listsu []orm.ParamsList
		num, err := o.Raw("SELECT  distinct id51  FROM 51job_user WHERE id<? order by id desc limit 1", tempu.Id).ValuesList(&listsu)
		numi := int(num)
		if numi == 1 && err == nil {
			returns = append(returns, listsu[0][0].(string))
		} else {
			returns = append(returns, "")
		}

		var lists1u []orm.ParamsList
		num1, err1 := o.Raw("SELECT  distinct id51  FROM 51job_user WHERE id>? limit 1", tempu.Id).ValuesList(&lists1u)
		numi1 := int(num1)
		if numi1 == 1 && err1 == nil {
			returns = append(returns, lists1u[0][0].(string))
		} else {
			returns = append(returns, "")
		}
		return returns
	}
	/**/
	/**/
	/**/

	temp := model.UserKeyword{Id51: id}
	o.Read(&temp, "Id51")

	var lists []orm.ParamsList
	num, err := o.Raw("SELECT  distinct id51  FROM 51job_user_keyword WHERE keyword_id = ? and id<? order by id desc limit 1", kid, temp.Id).ValuesList(&lists)
	numi := int(num)
	if numi == 1 && err == nil {
		returns = append(returns, lists[0][0].(string))
	} else {
		returns = append(returns, "")
	}

	var lists1 []orm.ParamsList
	num1, err1 := o.Raw("SELECT  distinct id51  FROM 51job_user_keyword WHERE keyword_id = ? and id>? limit 1", kid, temp.Id).ValuesList(&lists1)
	numi1 := int(num1)
	if numi1 == 1 && err1 == nil {
		returns = append(returns, lists1[0][0].(string))
	} else {
		returns = append(returns, "")
	}
	return returns

}
func ListKeyword() []*model.Keyword {
	var keywords []*model.Keyword
	user := new(model.Keyword)
	us := o.QueryTable(user)
	us.OrderBy("-Time51", "-Created").All(&keywords)
	return keywords
}

func ListOneKeyword(id int) *model.Keyword {
	temp := &model.Keyword{Id: id}
	o.Read(temp)
	return temp
}

func CountUser(keyword string, sex string) int64 {
	var num = int64(0)
	if keyword == "" {
		user := new(model.User)
		us := o.QueryTable(user)
		if sex == "1" {
			num, _ = us.Filter("Sex", "女").Count()
		} else if sex == "0" {
			num, _ = us.Filter("Sex", "男").Count()
		} else {
			num, _ = us.Count()
		}
		return num
	}

	keywordid, e := strconv.Atoi(keyword)
	if e != nil {
		return 0
	}
	var lists []orm.ParamsList
	if sex == "1" {
		num, _ = o.Raw("select count(distinct a.id) from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=? and a.sex=?", keywordid, "女").ValuesList(&lists)

	} else if sex == "0" {
		num, _ = o.Raw("select count(distinct a.id) from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=? and a.sex=?", keywordid, "男").ValuesList(&lists)
	} else {
		num, _ = o.Raw("select count(distinct a.id) from 51job_user as a join 51job_user_keyword  as b where b.id51=a.id51 and b.keyword_id=?", keywordid).ValuesList(&lists)

	}
	tnn, _ := strconv.Atoi(lists[0][0].(string))
	num = int64(tnn)
	return num

}

func ListUserHistroy(user string) []*model.UserKeyword {
	var userkeyword []*model.UserKeyword
	userhistry := new(model.UserKeyword)
	oo := o.QueryTable(userhistry)
	oo.Filter("Id51", user).OrderBy("-Created").RelatedSel().All(&userkeyword)
	return userkeyword
}

func ListUserInfo(user string) model.Userinfo {
	var userinfo model.Userinfo
	info := new(model.Userinfo)
	o.QueryTable(info).Filter("Id51", user).RelatedSel().One(&userinfo)
	return userinfo
}
