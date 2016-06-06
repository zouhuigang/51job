//自定义日志处理类
package log

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

var log *logs.BeeLogger
var Conn *websocket.Conn

func TodayDate() string {
	return time.Now().Format("20060102")
}

func init() {
	log = logs.NewLogger(1000)
	timetoday := "../log/" + TodayDate()
	os.MkdirAll(timetoday, 0777)
	//异步
	// log.Async()
	//允许行
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
	// log.SetLogger("console", ``)

	log.SetLogger("multifile", `{"filename":"`+timetoday+`/51job.log","separate":[ "error", "warning", "info", "debug"]}`)

}

func Printf(format string, v ...interface{}) {
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
	log.Info(format, v...)
}

func Println(v ...interface{}) {
	l := len(v)
	format := ""
	for i := 1; i <= l; i++ {
		format = format + "%v"
	}
	log.Info(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}

func Fatal(format string, v ...interface{}) {
	log.Error(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
	// os.Exit(1)

}
func Info(format string, v ...interface{}) {
	log.Info(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}
func Error(format string, v ...interface{}) {
	log.Error(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
	// os.Exit(1))
}

func Trace(format string, v ...interface{}) {
	log.Trace(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}
func Notice(format string, v ...interface{}) {
	log.Notice(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}
func Warning(format string, v ...interface{}) {
	log.Warning(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}

func Emergency(format string, v ...interface{}) {
	log.Emergency(format, v...)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}

func Debug(format string, v ...interface{}) {
	log.Debug(format, v)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}

func Alert(format string, v ...interface{}) {
	log.Alert(format, v)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}

func Critical(format string, v ...interface{}) {
	log.Critical(format, v)
	Conn.WriteMessage(1, []byte(fmt.Sprintf(format, v...)))
}
