package main

import (
	"51job/cons"
	"51job/log"
	"51job/model"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "io"
	"os"
	"time"
)

func init() {
	/*	参数1        数据库的别名，用来在ORM中切换数据库使用
		参数2        driverName
		参数3        对应的链接字符串
		set default database
		orm.RegisterDataBase("default", "mysql", "root:6833066@/51job?charset=utf8", 30)
		orm.SetMaxIdleConns("default", 30)
		orm.SetMaxOpenConns("default", 30)
		参数4(可选)  设置最大空闲连接
		参数5(可选)  设置最大数据库连接 (go >= 1.2)*/
	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "mysql", cons.Db, maxIdle, maxConn)

	// 设置为 UTC 时间
	orm.DefaultTimeLoc = time.UTC
	//打出查询语句
	orm.Debug = true
	w, _ := os.OpenFile("./db.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	orm.DebugLog = orm.NewLog(w)

}

func main() {
	ts := "20150808"
	date51, _ := time.Parse("20060102", ts)
	fmt.Println(date51)
	user := &model.User{Id51: "324324", Age: "28", Jobyear: "5-7年", Sex: "男", Address: "广州", Major: "物流管理", Date51: date51} //19844446|28|5-7年|男|广州-越秀区|物流管理|大专|2016-04-10
	data := []byte("ssssss")
	keepfile := "/data/sds"
	keyword := "keyword"
	address := "广州"
	kind := "3"
	e := SaveUser(user, data, keepfile, keyword, address, kind)
	log.Println(e)
}
func SaveUser(u *model.User, data []byte, keepfile string, keyword string, address string, kind string) error {
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数
	err := o.Begin()
	uk := new(model.UserKeyword)
	uk.Id51 = u.Id51
	uk.Keyword = keyword
	uk.FileAddress = keepfile
	uk.Address = address
	uk.Kind = kind
	uk.Date51 = u.Date51

	ui := new(model.Userinfo)
	ui.Id51 = u.Id51
	ui.Date51 = u.Date51
	ui.Content = string(data)

	r := o.Read(u, "Id51")
	if r == orm.ErrNoRows {
		log.Println("数据库不存在该简历记录")

		_, err1 := o.Insert(u)

		if err1 != nil {
			err = o.Rollback()
			return err
		}

	} else if r == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		log.Println("数据库存在该简历记录")
		u.Date51 = uk.Date51
		if _, err := o.Update(u); err != nil {
			err = o.Rollback()
			return err
		}
	}

	uk.User = u
	_, err2 := o.Insert(uk)
	if err2 != nil {
		err = o.Rollback()
		return err
	}

	rui := o.Read(ui, "Id51")

	if rui == orm.ErrNoRows {
		log.Println("数据库不存在该简历详情页")
		ui.User = u
		_, err3 := o.Insert(ui)
		if err3 != nil {
			err = o.Rollback()
			return err
		}

	} else if r == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		log.Println("数据库存在该简历详情页")
		ui.Content = string(data)
		ui.Date51 = u.Date51
		ui.User = u
		if _, err := o.Update(ui); err != nil {
			err = o.Rollback()
			return err
		}
	}

	err = o.Commit()
	return err
}
