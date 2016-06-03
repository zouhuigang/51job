//抓取工具类
package util

import (
	"51job/cons"
	"51job/log"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 克隆头部
func CloneHeader(h http.Header) http.Header {
	h2 := make(http.Header, len(h))
	for k, vv := range h {
		vv2 := make([]string, len(vv))
		copy(vv2, vv)
		h2[k] = vv2
	}
	return h2
}

// Post附带信息
func Post(url string, postValues url.Values, header map[string][]string) (body []byte, e error) {
	Wait()
	log.Println("POST链接:" + url)
	OutputArgs("POST 参数:", postValues)
	var request = &http.Request{}
	if postValues != nil {
		pr := ioutil.NopCloser(strings.NewReader(postValues.Encode()))
		request, _ = http.NewRequest("POST", url, pr)
	} else {
		request, _ = http.NewRequest("POST", url, nil)
	}
	request.Header = CloneHeader(header)

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	OutputMaps("-----------request携带头部-----------", request.Header)

	response, err := cons.Client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer response.Body.Close()

	body, e = ioutil.ReadAll(response.Body)

	OutputMaps("---------response携带头部-------------", response.Header)
	log.Printf("状态：%v:%v", response.Status, response.Proto)
	// log.Println(response.Location())
	//设置新Cookie
	MergeCookie(cons.Cookieb, response.Cookies())

	return
}

func Get(url string, resheader map[string][]string) (body []byte, e error) {
	Wait()
	log.Println("GET链接:" + url)

	//新建请求
	request, _ := http.NewRequest("GET", url, nil)

	//带头部，并发不影响，所以克隆
	request.Header = CloneHeader(resheader)

	OutputMaps("---------request携带头部--------", request.Header)

	//开始请求
	response, err := cons.Client.Do(request)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer response.Body.Close()

	OutputMaps("----------response携带头部-----------", response.Header)
	log.Printf("状态：%v:%v", response.Status, response.Proto)
	// log.Println(response.Location())
	//设置新Cookie
	cons.Cookieb = MergeCookie(cons.Cookieb, response.Cookies())

	//返回内容
	body, e = ioutil.ReadAll(response.Body)

	return
}

//合并Cookie，后来的覆盖前来的
func MergeCookie(before []*http.Cookie, after []*http.Cookie) []*http.Cookie {
	cs := make(map[string]*http.Cookie)

	for _, b := range before {
		cs[b.Name] = b
	}

	for _, a := range after {
		if a.Value != "" {
			cs[a.Name] = a
		}
	}

	res := make([]*http.Cookie, 0, len(cs))

	for _, q := range cs {
		res = append(res, q)

	}

	return res

}

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

func sanitizeCookieName(n string) string {
	return cookieNameSanitizer.Replace(n)
}

func sanitizeCookieValue(v string) string {
	v = sanitizeOrWarn("Cookie.Value", validCookieValueByte, v)
	if len(v) == 0 {
		return v
	}
	if v[0] == ' ' || v[0] == ',' || v[len(v)-1] == ' ' || v[len(v)-1] == ',' {
		return `"` + v + `"`
	}
	return v
}

func validCookieValueByte(b byte) bool {
	return 0x20 <= b && b < 0x7f && b != '"' && b != ';' && b != '\\'
}

func sanitizeOrWarn(fieldName string, valid func(byte) bool, v string) string {
	ok := true
	for i := 0; i < len(v); i++ {
		if valid(v[i]) {
			continue
		}
		log.Printf("net/http: invalid byte %q in %s; dropping invalid bytes", v[i], fieldName)
		ok = false
		break
	}
	if ok {
		return v
	}
	buf := make([]byte, 0, len(v))
	for i := 0; i < len(v); i++ {
		if b := v[i]; valid(b) {
			buf = append(buf, b)
		}
	}
	return string(buf)
}

//解析cookie为字符串
func ParseCookie(cookie []*http.Cookie) string {
	strcookie := ""
	for _, c := range cookie {
		if len(c.Value) == 0 {
			continue
		}
		s := fmt.Sprintf("%s=%s", sanitizeCookieName(c.Name), sanitizeCookieValue(c.Value))
		strcookie = strcookie + "; " + s
	}
	if strcookie != "" {
		strcookie = strcookie[1:]
	}
	log.Println("解析的cookie：" + strcookie)
	return strcookie
}

func Wait() {
	log.Printf("暂停%d秒～～", cons.DeadTime)
	time.Sleep(cons.DeadTime * time.Second)
}
