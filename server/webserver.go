// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"51job/cons"
	"51job/log"
	"51job/pika"
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

var addr = flag.String("addr", ":8099", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade:" + err.Error())
		return
	}
	log.Conn = c
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:" + err.Error())
			pika.Close()
			return
		}
		if pika.IsClose() {
			go func() {
				c.WriteMessage(1, []byte("开始爬虫"))
				a := strings.Split(string(message), "|")
				keyword := a[0]
				address := a[1]
				addressid, e := cons.Address[address]
				if !e {
					c.WriteMessage(1, []byte("ERROR:城市不存在"))
					return
				}
				kind := a[2]
				//登陆
				pika.Login()
				pika.SearchPika(keyword, address, addressid, kind)
				c.WriteMessage(1, []byte("爬虫结束"))
			}()
		} else {
			c.WriteMessage(1, []byte("请先关闭运行中的爬虫再运行"))
		}
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/", home)
	go func() {
		fmt.Println("开始启动：127.0.0.1:8099")
		err := http.ListenAndServe(*addr, nil)
		if err != nil {
			log.Printf("http server error: %s\nExit...", err.Error())
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("结束")
}

func home(w http.ResponseWriter, r *http.Request) {
	var homeTemplate = template.Must(template.New("").Parse(`<!Doctype html>
<html>
<head>
    <meta charset="utf-8" />
    <title>爬虫启动台</title>
    <style type="text/css">    
    #jobinput {
        margin: 0 auto;
        width: 300px;
        float: left;
        padding:0 10px 0 10px;
    }
    #jobinput label {
        font-size: 1.2em;
        display: block;
    }
    #jobinput input {}
    #button {
        font-size: 1.2em;
        display: block
    }

    #main{   float: right;width:800px;height:402px;overflow-y:auto;border:1px solid #ddd;padding:0 10px 0 10px;}

    </style>
    <script>
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var keyword = document.getElementById("keyword");
    var address = document.getElementById("address");
    var kind = document.getElementById("kind");
    // var clear = document.getElementById("clear");
    var ws;
  var hid=document.getElementById('msg_end');//隐藏在消息框下面的元素

    var print = function(message) {
        var d = document.createElement("div");
        d.innerHTML = message;
        output.appendChild(d);
        hid.scrollIntoView(false);//方式1通过调用隐藏元素的scrollIntoView()方法使其可见
    };

    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            alert("爬虫已经关闭");
            ws = null;
        }
        ws.onmessage = function(evt) {
            if(evt.data=="请先关闭运行中的爬虫再运行"){
                alert("请先关闭运行中的爬虫再运行")
            }
            print("" + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };

    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        var s = keyword.value + "|" + address.value + "|" + kind.value
        print("SEND: 关键字:" + s);
        ws.send(s);
        return false;
    };

    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };

    document.getElementById("clear").onclick = function(evt) {
        output.innerHTML = ""
        return false;
    };

});
</script>
</head>
<body>
    <div id="jobinput">
        <button id="open">开启爬虫</button>
        <button id="close">停止爬虫</button>
        <button id="clear">清屏</button>
        <form>
            <label for="keyword">关键字</label>
            <input name="keyword" type="text" id="keyword" />
            <select name="kind" id="kind">
                <option value="0" selected="selected">全文</option>
                <option value="2">职务</option>
                <option value="3">公司</option>
                <option value="4">学校</option>
                <option value="5">证书</option>
                <option value="1">工作</option>
            </select>
            <label for="address">地点</label>
            <input name="address" type="text" id="address" />
        </form>
        <input type="submit" id="send" value="Run" />
    </div>
    <div id="main">
    <div id="output"></div>
     <span id="msg_end" style="overflow:hidden"></span>
     </div>
</body>
</html>
`))
	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
}
