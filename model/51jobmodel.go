package model

import (
	// "github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	_ "io"
	"time"
)

//19844446|28|5-7年|男|广州-越秀区|物流管理|大专|2016-04-10
//用户简历表
type User struct {
	Id           int
	Url          string         `orm:"-"`      //链接
	Id51         string         `orm:"unique"` //ID姓名
	Jobyear      string         `orm:"null"`
	Age          string         `orm:"null"`
	Sex          string         `orm:"null"`
	Address      string         `orm:"null"`
	Major        string         `orm:"null"`
	Study        string         `orm:"null"`
	Date51       time.Time      `orm:"not null;type(datetime)"`     //简历更新时间
	Created      time.Time      `orm:"auto_now_add;type(datetime)"` //第一次抓到的时间
	Updated      time.Time      `orm:"auto_now;type(datetime)"`     //
	UserKeywords []*UserKeyword `orm:"reverse(many)"`               // 一个简历对应多个检索条件
	Userinfo     *Userinfo      `orm:"reverse(one)"`
}

type Userinfo struct {
	Id      int
	Id51    string    `orm:"not null;unique"` //ID姓名
	Date51  time.Time `orm:"type(datetime)"`  //简历更新时间
	Content string    `orm:" type(text)"`
	Created time.Time `orm:"auto_now_add;type(datetime)"` //第一次抓到的时间
	Updated time.Time `orm:"auto_now;type(datetime)"`     //
	User    *User     `orm:"null;rel(one);on_delete(set_null)"`
}

//关键字表-简历表
type UserKeyword struct {
	Id          int
	Id51        string
	FileAddress string    `orm:"not null`                     //文件地址
	Date51      time.Time `orm:"type(datetime)"`              //简历更新时间
	Created     time.Time `orm:"auto_now_add;type(datetime)"` //保证今天只保留一个本地文件
	User        *User     `orm:"rel(fk)"`                     //设置一对多关系
	Keyword     *Keyword  `orm:"rel(fk)"`
}

type Keyword struct {
	Id           int
	Keyword      string         `orm:"not null`
	Address      string         `orm:"not null`
	Kind         string         `orm:"not null`
	Created      time.Time      `orm:"null;type(datetime)"` //
	Updated      time.Time      `orm:"null;type(datetime)"` //
	Time51       int            //几次？
	UserKeywords []*UserKeyword `orm:"reverse(many)"` // 一个简历对应多个关键字
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}

// 设置引擎为 INNODB
func (u *UserKeyword) TableEngine() string {
	return "INNODB"
}
