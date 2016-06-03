//主函数
package pika

import (
	"51job/cons"
	"51job/log"
	"51job/model"
	"51job/util"
	"io/ioutil"
	// "net/url"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	stop bool
)

func init() {
	os.MkdirAll("../data/img/", 0777)
	stop = true
}

// 搜索连续翻页
func SearchPika(keyword string, address string, addressid string, kind string) {
	//访问搜索页内页
	data := Search(nil)

	//死掉吧，搜索
	DeadSearch(data)

	//构造post
	searchargs := FindSearchHidden(data)

	//打印
	// util.Output("", searchargs)

	//第一次点击查询
	log.Notice("关键字:%s-地点:%s-第%d页开始\n", keyword, address, 1)
	searchargs = SetPostData(searchargs, keyword, address, addressid, kind)
	data = Search(searchargs)

	//计算总页数
	pagenum := FindPageNum(data)
	if strings.EqualFold(pagenum, "") || strings.EqualFold(pagenum, "0") {
		return
	}

	//解析隐藏参数
	searchargs = FindSearchHidden(data)

	//保存
	charst := []byte("<meta charset='utf-8' />")
	data = append(charst, data...)
	today := util.TodayDate()
	list := cons.UserListKeepPath + "/" + today + "/" + keyword
	os.MkdirAll(list+"/", 0777)
	fe := ioutil.WriteFile(list+"/1.html", data, 0644)
	if fe != nil {
		log.Println(fe.Error())
	} else {
		log.Notice("保存简历列表：%s\n", list+"/1.html")
	}

	//--------------------
	CatchPika("1", ListPika(data), keyword, address, kind)
	//--------------------
	//开始翻页
	for i := 1; ; i++ {
		if stop {
			break
		}

		//构造post
		//  第几页
		log.Notice("关键字:%s-地点:%s-第%d页开始\n", keyword, address, i+1)
		pages := strconv.FormatInt((int64)(i+1), 10)
		//来源页
		initpage := strconv.FormatInt((int64)(i), 10)
		searchargs = SetPostData(searchargs, keyword, address, addressid, kind)
		searchargs.Set("__EVENTTARGET", "")
		searchargs.Set("pagerBottom$nextButton", "下一页")
		searchargs.Set("pagerBottom$txtGO", initpage)

		//打印
		// util.Output("", searchargs)

		//第i+1次点击
		data = Search(searchargs)
		pagenum = FindPageNum(data)
		if strings.EqualFold(pagenum, "") || strings.EqualFold(pagenum, "0") {
			break
		}

		//保存
		charst := []byte("<meta charset='utf-8' />")
		data = append(charst, data...)
		fe = ioutil.WriteFile(list+"/"+initpage+"-"+pages+".html", data, 0644)
		if fe != nil {
			log.Println(fe.Error())
		} else {
			log.Notice("保存简历列表：%s\n", list+"/"+initpage+"-"+pages+".html")
		}
		//解析隐藏参数
		searchargs = FindSearchHidden(data)

		//--------------------
		CatchPika(fmt.Sprint(i+1), ListPika(data), keyword, address, kind)
		//--------------------

	}
}

//解析
func ListPika(data []byte) map[int]*model.User {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	if err != nil {
		log.Println(err)
	}

	//存放提取数据
	userlist := make(map[int]*model.User)
	//单个人
	user := new(model.User)
	num := 0

	//查找简历每行
	doc.Find("#divGridData tr").Each(func(i int, node *goquery.Selection) {
		nodeid, _ := node.Attr("id")
		//只检索基础信息列
		if strings.Contains(nodeid, "trBaseInfo") {
			user = &model.User{}
			num++ //记录第几个人
			//每列
			jinhan := -1 //判断User字段
			node.Find("td").Each(func(j int, node1 *goquery.Selection) {
				//检索跳转链接
				node1id, _ := node1.Attr("class")
				if strings.Contains(node1id, "inbox_td22") {
					node1href := node1.Find("a")
					k := node1href.Text()              //ID51
					v, exist := node1href.Attr("href") //链接
					if exist {
						user.Url = v
						user.Id51 = k
					}
				}
				//检索其他信息
				node1style, _ := node1.Attr("style")
				if strings.Contains(node1style, "width:auto;") {
					ddd := strings.TrimSpace(node1.Text())
					jinhan++
					if jinhan == 0 {
						user.Age = ddd
					} else if jinhan == 1 {
						user.Jobyear = ddd
					} else if jinhan == 2 {
						user.Sex = ddd
					} else if jinhan == 3 {
						user.Address = ddd
					} else if jinhan == 4 {
						user.Major = ddd
					} else if jinhan == 5 {
						user.Study = ddd
					} else if jinhan == 6 {
						dddd, _ := time.ParseInLocation("2006-01-02", ddd, time.Local)
						user.Date51 = dddd
					} else {

					}
				}
			})
			//加入豪华套餐
			userlist[num] = user
		}

	})
	return userlist

}

func CatchPika(goname string, userlist map[int]*model.User, keyword string, address string, kind string) {
	//创建文件夹，格式20160202，今天抓取的文件
	today := util.TodayDate()
	dir := cons.UserKeepPath + today

	//时间戳子目录，放400个文件
	firstdir := dir + "/" + keyword + "/" + goname
	err1 := os.MkdirAll(firstdir+"/", 0777)
	if err1 != nil {
		log.Notice("%s\n", err1.Error())
	} else {
		log.Notice("进程go:%s-创建或进入目录:%s\n", goname, firstdir)
	}

	for i, u := range userlist {
		if stop {
			break
		}

		who := PrintUser(u)
		// 打印简短信息
		log.Printf("进程go:%s--第%d人:%s", goname, i, who)

		//查看文件是否存在
		//Db查看
		//不存在

		//保存位置
		ee := u.Url
		keepfile := ""
		if ee != "" {
			keepfile = firstdir + "/" + u.Id51 + ".html"
		} else {
			continue
		}
		_, err := os.Stat(keepfile)
		if err == nil {
			log.Warning("进程go:%s本地文件已经存在,还是要重新抓取", keepfile)
		} else {
			log.Warning("进程go:待保存文件:%s", keepfile)
		}
		//是否存在这个文件
		data, _ := util.Get("http://ehire.51job.com/"+u.Url, cons.Requestheader)

		//死掉吧!
		DeadPika(data)

		//写入文件，写入成功保存进数据库
		//文件数量是否超出
		if util.SizeofDir(firstdir) >= cons.MaxFileNum {
			log.Printf("进程go:%s:这里是不会出现的", goname)
		}
		data = CutInfo(data, u.Date51.Format("20060102"))
		if string(data) == "" {
			log.Warning("简历保密")
			continue
		}
		e := ioutil.WriteFile(keepfile, data, 0644)

		//写入成功保存数据库
		if e == nil {
			log.Notice("进程go:%s-保存文件：%s", goname, keepfile)
			img, _ := util.Get("http://ehire.51job.com/Candidate/ReadAttach.aspx?UserID="+u.Id51, cons.Requestheader)
			imgpath := "../data/img/" + u.Id51 + ".jpg"
			os.MkdirAll("../data/img/", 0777)
			ioutil.WriteFile(imgpath, img, 0644)
			//DB  用户 用户数据 文件地址
			dbe := model.SaveUser(today, u, data, keepfile, keyword, address, kind)
			if dbe == nil {
				log.Alert("%s\n", "保存数据库成功")
			} else {
				log.Warning("%s:%s\n", "保存数据库失败", dbe.Error())
			}
		}

	}
}
