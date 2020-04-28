package log

// 自主来发的轻量级日志框架, 直接融入到agent中, 不以第三方的形式
// 功能: 1. 多级别输出 2.可自主创建log目录级文件 3.支持按日期\小时\分钟的日志分割

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type logger struct {
	LogFile *os.File
}

func (this *logger) Start(log_name string) {
	var logF = flag.String(log_name, log_name, "log file")
	flag.Parse()
	this.LogFile, _ = os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
}

func (this *logger) Close() {
	if this.LogFile != nil {
		this.LogFile.Close()
	}

}

func (this *logger) commonOut(f, level string, line int, content interface{}) {
	if this.LogFile == nil {
		return
	}
	log.SetOutput(this.LogFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(fmt.Sprintf("%v_%v %v %v", f, line, level, content))
}

func (this *logger) fatalOut(f, level string, line int, content interface{}) {
	if this.LogFile == nil {
		return
	}
	log.SetOutput(this.LogFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Fatalln(fmt.Sprintf("%v_%v %v %v", f, line, level, content))
}

var gongLog = new(logger)
