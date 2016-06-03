//杂工具包
package util

import (
	"51job/log"
	"net/url"
)

//打印url参数
func OutputArgs(info string, args url.Values) {
	/*	log.Println(info)
		for i, v := range args {
			log.Printf("%s:%v", i, v)
		}*/
}

//打印映射
func OutputMaps(info string, args map[string][]string) {
	/*	log.Println(info)
		for i, v := range args {
			log.Printf("%s:%v", i, v)
		}*/
}

//打印映射
func Output(info string, args map[string][]string) {
	log.Println(info)
	for i, v := range args {
		log.Printf("%s:%v", i, v)
	}
}
