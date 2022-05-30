package mylog

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

//日志级别
type LogLevel uint16

type Logger struct {
	Level LogLevel
}

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s {
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		return UNKNOWN, errors.New("无效的日志级别错误")
	}
}

func getinfo(skip int) (funcName string, fileName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	funcName = strings.Split(funcName, ".")[1]
	fileName = filepath.Base(file)
	lineNo = line
	return funcName, fileName, lineNo
}

func NewLog(levelStr string) Logger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		fmt.Printf("%v", err)

	}
	return Logger{
		Level: level,
	}
}

func newFileLog() *os.File {
	fileObj, err := os.OpenFile("wlog/text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Printf("Openfile failed,err:%v", err)
	}
	return fileObj

}

func (l Logger) Debug(msg string) {
	if l.Level <= DEBUG {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}

func (l Logger) Trace(msg string) {
	if l.Level <= TRACE {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][Trace] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}

func (l Logger) Info(msg string) {
	if l.Level <= INFO {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][Info] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}

func (l Logger) Warning(msg string) {
	if l.Level <= WARNING {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][Warning] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}

func (l Logger) Error(msg string) {
	if l.Level <= ERROR {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][Error] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}

func (l Logger) Fatal(msg string) {
	if l.Level <= FATAL {
		//now := time.Now()
		funcName, fileName, lineNo := getinfo(2)
		f := newFileLog()
		var cstSH, _ = time.LoadLocation("Asia/Shanghai")
		timeObj := time.Now().In(cstSH).Format("2006-01-02 15:04:05")
		//fmt.Fprintf(f, "[%s][DEBUG] [%s: %s: %d]%s\n", now.Format("2006-01-02 15:04:05"), funcName, fileName, lineNo, msg)
		fmt.Fprintf(f, "[%s][Fatal] [%s: %s: %d]%s\n", timeObj, funcName, fileName, lineNo, msg)
		f.Close()
	}
}
