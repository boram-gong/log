package log

import (
	"fmt"
	"runtime"
	"time"
)

const (
	COMMON_FORMAT = ""
	MONTH_FROMAT  = "200601"
	DATE_FORMAT   = "20060102"
	HOUR_FORMAT   = "2006010215"
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
	if s == COMMON_FORMAT || s == DATE_FORMAT || s == HOUR_FORMAT || s == MONTH_FROMAT {
		return true
	} else {
		return false
	}
}

func InitLog(log_path string, log_name string, format string) {
	if log_path == "" {
		log_path = "./default_log"
	}
	if log_name == "" {
		log_name = "default"
	}
	if !formatCheck(format) {
		format = COMMON_FORMAT
	}
	CreateDir(log_path)
	gongLog.lockLogDir(log_path, log_name, format)
	if format != COMMON_FORMAT {
		var name = fmt.Sprintf("%s/%s_%s.log", log_path, log_name, time.Now().Format(format))
		gongLog.start(name)
		go logManage(log_path, log_name, format)

	} else {
		var name = fmt.Sprintf("%s/%s.log", log_path, log_name)
		gongLog.start(name)
	}
}

func LogLevelFilter(level int) {
	if level > Fatal || level < Info {
		return
	}
	LogLevel = level
}

func LogFileSweeper(l int) {
	gongLog.sweeper(l)
}

func GetLogSize() int64 {
	return gongLog.getSize()
}

func logManage(log_path string, log_name string, format string) {
	var lastDay = time.Now().Format(format)
	for {
		time.Sleep(1 * time.Second)
		nowDay := time.Now().Format(format)
		if nowDay > lastDay {
			lastDay = nowDay
			gongLog.close()
			name := fmt.Sprintf("%s/%s_%s.log", log_path, log_name, nowDay)
			gongLog.start(name)
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
	gongLog.commonOut(f, "[DEBUG]", line, content)
}

func WARN(content interface{}) {
	if LogLevel > Debug {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f, "[WARN]", line, content)
}

func ERROR(content interface{}) {
	if LogLevel > Error {
		return
	}
	_, f, line, _ := runtime.Caller(1)
	gongLog.commonOut(f, "[ERROR]", line, content)
}

func FATAL(content interface{}) {
	_, f, line, _ := runtime.Caller(1)
	gongLog.fatalOut(f, "[FATAL]", line, content)
}
