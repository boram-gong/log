package log

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

const (
	COMMON_FORMAT = ""
	DATE_FORMAT   = "2006-01-02"
	HOUR_FORMAT   = "2006010215"
	MINUTE_FROMAT = "200601021504"
)

const (
	Info int = 0 << iota
	Debug
	Warn
	Error
	Fatal
)

var LogLevel int

func formatCheck(s string) bool {
	if s == DATE_FORMAT || s == HOUR_FORMAT || s == MINUTE_FROMAT || s == COMMON_FORMAT {
		return true
	} else {
		return false
	}
}

func InitLog(log_path string, log_name string, format string) {
	if log_path == "" || log_name == "" || !formatCheck(format) {
		log.Fatalln("logger init fail")
	}
	CreateDir(log_path)
	if format != COMMON_FORMAT {
		var name = fmt.Sprintf("%s/%s_%s.log", log_path, log_name, time.Now().Format(format))
		gongLog.Start(name)
		go logManage(log_path, log_name, format)

	} else {
		var name = fmt.Sprintf("%s/%s.log", log_path, log_name)
		gongLog.Start(name)
	}
}

func LogLevelFilter(level int) {
	if level > Fatal || level < Info{
		return
	}
	LogLevel = level
}

func logManage(log_path string, log_name string, format string) {
	var lastDay = time.Now().Format(format)
	for {
		time.Sleep(1 * time.Second)
		nowDay := time.Now().Format(format)
		if nowDay > lastDay {
			lastDay = nowDay
			gongLog.Close()
			name := fmt.Sprintf("%s/%s_%s.log", log_path, log_name, nowDay)
			gongLog.Start(name)
		}

	}
}

func INFO(content interface{}) {
	if LogLevel > Info {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f, "[INFO]", line, content)
}

func DEBUG(content interface{}) {
	if LogLevel > Debug {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f,"[DEBUG]", line, content)
}

func WARN(content interface{}) {
	if LogLevel > Debug {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f,"[WARN]", line, content)
}

func ERROR(content interface{}) {
	if LogLevel > Error {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f,"[ERROR]", line, content)
}

func FATAL(content interface{}) {
	_, f, line, _ := runtime.Caller(1)
	gongLog.fatalOut(f,"[FATAL]", line, content)
}
