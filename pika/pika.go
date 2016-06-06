//模拟点击主函数在此
package pika

import (
	"51job/cons"
	"51job/util"
	// "fmt"
	// "github.com/PuerkitoBio/goquery"
	"51job/log"
	"io/ioutil"
	"net/url"
	// "os"
	// "strings"
)

// 构造真正https登录页数据
func LoginData() url.Values {
	//登陆页面所需数据
	logindata, e := util.Get(cons.Loginurl, cons.Requestheader)
	if e != nil {
		log.Fatal(e.Error())
		Close()
	}
	ioutil.WriteFile(cons.LoginData, logindata, 0644)
	// 找出隐藏字段
	oldAccessKey, sc, ec := FindHideen(logindata)
	if oldAccessKey == "" || sc == "" || ec == "" {
		log.Println("莫名失败！！！")
		log.Println("继续尝试连接，可能网络不好")
		return LoginData()
	} else {
		postdata.Set("checkCode", "")
		postdata.Set("oldAccessKey", oldAccessKey)
		postdata.Set("langtype", "Lang=&Flag=1")
		postdata.Set("isRememberMe", "false")
		postdata.Set("sc", sc)
		postdata.Set("ec", ec)
		postdata.Set("returl", "")
		return postdata
	}
}

//自动测试https POST登陆
func Login() []byte {
	stop = false

	postdata = LoginData()
	// 头部加cookie
	/*	cons.S1header["Cookie"] = []string{
		util.ParseCookie(cons.Cookieb),
	}*/
	// time.Sleep(time.Second * 10)
	data, e := util.Post(cons.Realloginurl, postdata, cons.S1header)
	if e != nil {
		log.Fatal(e.Error())
		Close()
	}
	ioutil.WriteFile(cons.RealLoginData, data, 0644)
	forceurl, forcedata := ForceDown(data)
	if forceurl == "" {
		log.Println("成功登陆了额！！！！！！！")
	} else {
		d, _ := util.Post(forceurl, forcedata, cons.Requestheader)
		ioutil.WriteFile(cons.ForceData, d, 0644)
		return d
	}
	return data
}

//搜索首页
func SearchIndex() []byte {
	// cons.Requestheader["Referer"] = []string{
	// 	"http://ehire.51job.com/Candidate/SearchResumeIndex.aspx",
	// }
	data, e := util.Get(cons.Searchurl, cons.Requestheader)
	if e != nil {
		log.Fatal(e.Error())
		Close()
	}
	charst := []byte("<meta charset='utf-8' />")
	data = append(charst, data...)
	ioutil.WriteFile(cons.SearchIndexData, data, 0644)
	return data
}

//搜索 真正页
func Search(postdata url.Values) []byte {
	data, _ := util.Post(cons.Searchurl, postdata, cons.Requestheader)
	return data
}

func Close() {
	stop = true
}

func IsClose() bool {
	return stop
}
