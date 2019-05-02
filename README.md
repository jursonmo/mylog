### sample log
1. 可以给整个日志打印设置打印等级, 支持 "debug, info, notice, warn, error"
2. 日志文件名称自动加入当天的时间，每天都产生一个日志文件(比如logfile_2016-4-24.log)，默认保留最近三天的日志
3. 每个模块都可以注册，这样每个模块的log 的等级都可以分别设置，所有模块的日志都写到一个日志文件
4. 日志的打印受 当前模块的log 的等级和整个日志打印等级 影响

```go
package main

import (
	"fmt"
	"mylog"
)

func main() {
	//"debug, info, notice, warn, error"
	fmt.Println("suport level:", mylog.ShowSupportLevels())

	log1, err := mylog.RegisterLog("module1", "info")
	if err != nil {
		fmt.Println(err)
		return
	}

	log2, err := mylog.RegisterLog("module2", "info")
	if err != nil {
		fmt.Println(err)
		return
	}

	mylog.InitLog("info", "./logfile.log")//set log level/logfile for all modules
	log1.Info("test %s", "log1")
	log2.Info("test %s", "log2")

	err = mylog.SetModuleLogLevel("module1", "notice") // log1.SetModuleLogLevel("notice")
	if err != nil {
		fmt.Println(err)
		return
	}
	log1.Info("test level, cannot write to logfile ")
	log1.Notice("test level, can write to logfile ")

	err = mylog.SetLogLevel("warn") //set log level for all modules log
	if err != nil {
		fmt.Println(err)
		return
	}
	log1.Info("test level, cannot write to logfile")
	log2.Info("test level, cannot write to logfile")

	log1.Warning("test level,  write to logfile ")
	log2.Warning("test level,  write to logfile ")
}
