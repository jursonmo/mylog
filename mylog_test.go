package mylog

import (
	"fmt"
	"testing"
)

func TestRegistLog(t *testing.T) {
	//"debug, info, notice, warn, error"
	fmt.Println("suport level:", ShowSupportLevels())

	log1, err := RegisterLog("module_aa", "info")
	if err != nil {
		t.Fatal(err)
		return
	}

	log2, err := RegisterLog("module_bb", "info")
	if err != nil {
		t.Fatal(err)
		return
	}

	//test register a log for exist module
	_, err = RegisterLog("module_bb", "info")
	if err == nil {
		t.Fatal(err)
		return
	}

	_, _ = log1, log2
}
func TestRegistLogFunc(t *testing.T) {

	log1, err := RegisterLog("module1", "info")
	if err != nil {
		t.Fatal(err)
		return
	}

	log2, err := RegisterLog("module2", "info")
	if err != nil {
		t.Fatal(err)
		return
	}

	InitLog("info", "./logfile.log") //set log level/logfile for all modules
	log1.Info("test %s", "log1")
	log2.Info("test %s", "log2")

	err = SetModuleLogLevel("module1", "notice") //or log1.SetModuleLogLevel("notice")
	if err != nil {
		t.Fatal(err)
		return
	}
	log1.Info("test level, cannot write to log ")
	log1.Notice("test level, can write to log  ")

	err = SetLogLevel("warn") //set log level for all modules log
	if err != nil {
		t.Fatal(err)
		return
	}
	log1.Info("test level, cannot write to log ")
	log2.Info("test level, cannot write to log  ")

	log1.Warning("test level, must write to log ")
	log2.Warning("test level, must write to log  ")

}
