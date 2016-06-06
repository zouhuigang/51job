//解析函数
package pika

import (
	// "51job/cons"
	"51job/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	// "io/ioutil"
	"51job/log"
	"net/url"
	// "os"
	"51job/model"
	"strings"
	// "time"
)

//寻找登陆首页隐藏字段
func FindHideen(r []byte) (oldAccessKey string, sc string, ec string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r)))
	if err != nil {
		log.Println(err)
	}
	oldAccessKey, _ = doc.Find("#hidAccessKey").Attr("value")
	sc, _ = doc.Find("#fksc").Attr("value")
	ec, _ = doc.Find("#hidEhireGuid").Attr("value")
	log.Printf("登陆首页隐藏字段:oldAccessKey:%s\tsc:%s\tec:%s\n", oldAccessKey, sc, ec)
	return oldAccessKey, sc, ec
}

// 我要死了，强制下线
/*
http://ehire.51job.com/Member/UserOffline.aspx?tokenId=07504412-98df-4087-807c-e1c54934&errorFlag=0&dbid=4&val=c793cab10b9743a5&isRememberMe=False&sc=199c12176cdabbab&Lang=&Flag=1
http://ehire.51job.com/Member/UserOffline.aspx?tokenId=1b62b7c8-9748-4d9f-aa32-f88052ee&errorFlag=0&dbid=4&val=7085df05b9d634b3&isRememberMe=False&sc=199c12176cdabbab&Lang=&Flag=1
*/
func ForceDown(r []byte) (downhref string, downpost map[string][]string) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r)))
	if err != nil {
		log.Println(err)
		return
	}
	title := doc.Find("title").Text()
	ok := strings.Contains(title, "在线用户管理")

	log.Println("标题：", title)
	if !ok {
		log.Println("貌似成功登陆！！！！")
		return
	} else {
		log.Println("必须强制下线，用户多地点登陆或者账户达到上限！！！！")
	}
	token, _ := doc.Find("#__VIEWSTATE").Attr("value")
	action, _ := doc.Find("#form1").Attr("action")

	downpost = map[string][]string{
		"__VIEWSTATE":     {token},
		"__EVENTTARGET":   {"gvOnLineUser"},
		"__EVENTARGUMENT": {"KickOut$0"},
	}
	downhref = "http://ehire.51job.com/Member/" + action
	log.Printf("下线地址：%s", downhref)
	util.OutputMaps("强制下线参数", downpost)
	return
}

// 查找隐藏字段，构造post
func FindSearchHidden(data []byte) url.Values {
	searchdata := url.Values{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	doc.Find("input[type=hidden]").Each(func(i int, node *goquery.Selection) {
		ss, _ := node.Attr("name")
		vv, _ := node.Attr("value")
		if ss != "" && vv != "" {
			searchdata.Set(ss, vv)
		}
	})
	return searchdata
}

//查看搜索页页数
func FindPageNum(data []byte) string {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	nodet := doc.Find("#pagerTop_previousButton").Next().Text()
	s := strings.Split(nodet, "/")
	if len(s) == 2 {
		return s[1]
	}
	return ""

}

//简历浓缩哈
func CutInfo(r []byte, u string) []byte {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(r)))
	if err != nil {
		log.Println(err)
	}
	title := doc.Find("title").Text()
	h, _ := doc.Find("#divResume >table").Html()
	if h == "" {
		return []byte("")
	}
	h = strings.Replace(h, "#→start→#", "##", -1)
	h = strings.Replace(h, "#←end←#", "##", -1)
	title = `<!Doctype><html><head><meta charset="utf-8"/><title>舜飞爬取|` + title + `</title>`
	h = title + `<link href="http://js.51jobcdn.com/ehire2007/css/20160519/default/rschart.css" type="text/css" rel="stylesheet" />
<link href="http://ehire.51job.com/App_Themes/Default/style.css" type="text/css" rel="stylesheet"/>
<link href="http://js.51jobcdn.com/ehire2007/css/20160519/default/candidate.css" type="text/css" rel="stylesheet" />
<link href="http://js.51jobcdn.com/ehire2007/css/20160519/default/jquery-ui-1.8.4.thin.css" type="text/css" rel="stylesheet" />
<link href="http://js.51jobcdn.com/ehire2007/css/20160519/default/message.css" type="text/css" rel="stylesheet" />
</head><body style="width:700;align:center;margin:0 auto"><table> <thead><tr><th>简历更新时间:` + u + `</th></tr></thead>` + h + `</table></body></html>`
	return []byte(h)

}

/*type User struct {
	Id           int
	Url          string         `orm:"-"`      //链接
	Id51         string         `orm:"unique"` //ID姓名
	Jobyear      string         `orm:"null"`
	Age          string         `orm:"null"`
	Sex          string         `orm:"null"`
	Address      string         `orm:"null"`
	Major        string         `orm:"null"`
	Study        string         `orm:"null"`
	Date51       time.Time      `orm:"not null;type(date)"`         //简历更新时间
	Created      time.Time      `orm:"auto_now_add;type(datetime)"` //第一次抓到的时间
	Updated      time.Time      `orm:"auto_now;type(datetime)"`     //
	UserKeywords []*UserKeyword `orm:"reverse(many)"`               // 一个简历对应多个检索条件
	Userinfo     *Userinfo      `orm:"null;rel(one);on_delete(set_null)"`
}*/
func PrintUser(m *model.User) string {
	s := fmt.Sprintf("地址:%s编号:%s-工作年限:%s-年龄:%s-性别:%s-地址:%s-专业:%s-学历:%s-更新时间:%s", m.Url, m.Id51, m.Jobyear, m.Age, m.Sex, m.Address, m.Major, m.Study, m.Date51)
	return s
}

//访问太过频繁，死！！！
func DeadPika(data []byte) {
	if float64(len(data))/1000 < 56.2 {
		log.Printf("文件大小:%d字节", len(data))
		log.Printf("%s", string(data))
		//还有保密的简历要检测大小
		log.Fatal("速度太快了，死掉了！！")
		Close()
	}
}

//被禁掉了，哈哈哈哈哈哈哈哈哈
//查看搜索页页数
func DeadSearch(data []byte) {
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	nodet := doc.Find("title").Text()
	if strings.Contains(nodet, "错误") {
		log.Fatal("不能搜索，被禁掉了！！哇哈哈哈")
		Close()
	}

}
