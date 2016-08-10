//常量包
//保存PikaPika项目下的常量
package cons

import (
	"51job/log"
	"net/http"
	"net/http/cookiejar"
	_ "net/url"
)

//数据库配置
const (
	//数据库账号:数据库密码@tcp(数据库ip:端口号)/数据库名?编码等
	Db = "root:6ff@tcp(localhost:3306)/51job?charset=utf8&loc=Local"
	//数据库日志
	LogPath = "../log/db.log"
	//开启日志？
	OpenDbLog = false

	//用户头像保留地
	ImagePath = "../data/img/"
)

const (
	//暂停时间
	DeadTime = 3
	//登陆首页
	Loginurl = "http://ehire.51job.com/MainLogin.aspx"
	//真正登陆页面
	Realloginurl = "https://ehirelogin.51job.com/Member/UserLogin.aspx"
	Testindex    = "http://ehire.51job.com/Navigate.aspx?ShowTips=11&PwdComplexity=N"
	//搜索真正页
	Searchurl = "http://ehire.51job.com/Candidate/SearchResume.aspx"
	//搜索首页
	SearchIndex = "http://ehire.51job.com/Candidate/SearchResumeIndex.aspx"
	//验证码
	Code = "http://ehire.51job.com/CommonPage/RandomNumber.aspx?type=login"
)

const (
	//登陆首页数据
	LoginData = "../data/MainLogin.html"
	//真正登陆页数据，可能被定向到强制下线
	RealLoginData = "../data/UserLogin.html"
	P2            = "../data/output2.html"
	T1            = "../data/Navigate.html"
	//登陆成功首页
	IndexData = "../data/Index.html"
	//强制下线后数据
	ForceData = "../data/Force.html"
	//搜索内页
	SearchData = "../data/Search.html"
	//搜索首页
	SearchIndexData = "../data/SearchIndex.html"

	//简历列表存放处
	UserListKeepPath = "../data/userlist"
	//简历存放处
	UserKeepPath = "../data/user/"
	//每个文件夹最多存放的文件数量
	MaxFileNum = 400
)

func NewJar() *cookiejar.Jar {
	cookieJar, _ := cookiejar.New(nil)
	return cookieJar
}

var (
	Client = &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			log.Printf("------------------自动跳转跳转%v-------------------\n", req.URL)
			return nil
		},
		Jar: NewJar(),
	}
	//每次访问携带的cookie
	Cookieb = []*http.Cookie{} //map[string][]string
)
var (
	Ua = "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:46.0) Gecko/20100101 Firefox/46.0"
	//http://ehire.51job.com/MainLogin.aspx通用请求头部
	Requestheader = map[string][]string{
		"User-Agent": {
			Ua,
		},
		"Host": {
			"ehire.51job.com",
		},
		"Accept": {
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		},
		/*		"Accept-Language": {
					"en-US,en;q=0.5",
				},
				"Accept-Encoding": {
					"gzip, deflate",
				},*/
		"Connection": {
			"keep-alive",
		},
	}

	//https://ehirelogin.51job.com/Member/UserLogin.aspx请求头部
	S1header = map[string][]string{
		"User-Agent": {
			Ua,
		},
		"Host": {
			"ehirelogin.51job.com",
		},
		"Accept": {
			"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		},
		// "Accept-Language": {
		// 	"en-US,en;q=0.5",
		// },
		// "Accept-Encoding": {
		// 	"gzip, deflate, br",
		// },
		"Referer": {
			"http://ehire.51job.com/MainLogin.aspx",
		},
		"Connection": {
			"keep-alive",
		},
	}
)
