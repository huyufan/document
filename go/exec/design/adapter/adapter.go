package main

import "fmt"

//目标接口  定义客户端所期望的接口
type Logger interface {
	WriteLog(message string)
}

//源接口  需要被适配的接口
type LogWriter interface {
	Write(message string)
}

//适配器
type LogAdapter struct {
	logWriter LogWriter
}

func (adapter *LogAdapter) WriteLog(message string) {
	adapter.logWriter.Write(message)
}

type ExistingLogWriter struct{}

func (writer *ExistingLogWriter) Write(message string) {
	// 已有日志库的写入逻辑
	fmt.Println("Writing log:", message)
}

func main() {
	log := &LogAdapter{logWriter: &ExistingLogWriter{}}

	log.WriteLog("huyufan")
}
