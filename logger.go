package log

// 自主来发的轻量级日志框架, 直接融入到agent中, 不以第三方的形式
// 功能: 1. 多级别输出 2.可自主创建log目录级文件 3.支持按日期\小时\分钟的日志分割

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

type logger struct {
	LogFile       *os.File
	LogPath       string
	LogName       string
	LogFormat     string
	MonthFlag     string
	LogMonthCache []string
}

func (this *logger) start(log_name string) {
	var logF = flag.String(log_name, log_name, "log file")
	flag.Parse()
	this.LogFile, _ = os.OpenFile(*logF, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
}

func (this *logger) lockLogDir(path, name, format string) {
	this.LogPath = path
	this.LogName = name
	this.LogFormat = format
}

func (this *logger) close() {
	if this.LogFile != nil {
		this.LogFile.Close()
	}
}

func (this *logger) sweeper(l int) {
	if this.LogFormat == COMMON_FORMAT {
		return
	}
	last := time.Now().Format("200601")
	this.MonthFlag = last
	this.LogMonthCache = append(this.LogMonthCache, last)
	for {
		time.Sleep(60 * time.Second)
		now := time.Now().Format("200601")
		if now > this.MonthFlag {
			this.LogMonthCache = append(this.LogMonthCache, now)
			if len(this.LogMonthCache) > l {
				RmFile(this.LogPath, fmt.Sprintf("%s_%s", this.LogName, this.LogMonthCache[0]))
				this.LogMonthCache = this.LogMonthCache[1:]
			}
		}
	}
}

func (this *logger) getSize() int64 {
	s, _ := DirSize(this.LogPath)
	return s
}

func (this *logger) commonOut(content ...interface{}) {
	if this.LogFile == nil {
		return
	}
	log.SetOutput(this.LogFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Println(content...)
}

func (this *logger) fatalOut(content ...interface{}) {
	if this.LogFile == nil {
		return
	}
	log.SetOutput(this.LogFile)
	log.SetFlags(log.Ldate | log.Ltime)
	log.Fatalln(content...)
}

var gongLog = new(logger)
