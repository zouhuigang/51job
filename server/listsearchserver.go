//主函数
package main

import (
	"51job/cons"
	"51job/log"
	"51job/pika"
)

func main() {
	//登陆
	pika.Login()
	keyword := "前端"
	address := "广州"
	addressid, e := cons.Address[address]
	if !e {
		log.Fatal("城市不存在")
	}
	kind := cons.KeywordAll
	pika.SearchPika(keyword, address, addressid, kind)

}
