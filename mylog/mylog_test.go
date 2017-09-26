package mylog

import (
	"fmt"
	"os"
	"testing"
)

const (
	default_prefix = "testing: %s"
)

func TestStdLog(t *testing.T) {
	std.Info(fmt.Sprintf(default_prefix, t.Name()))
}

func TestSetOutput(t *testing.T) {
	w, err := os.Create("test.log")
	if err != nil {
		fmt.Printf("open file error:%v\n", err)
	}
	l := New(w, default_log_format, LevelInfo)
	l.Infof(default_prefix, t.Name())
	l.Info("test")
}

func TestSetFormat(t *testing.T) {
	format := "#time 在 #file 文件 #line 行，输出：#msg"
	l := New(os.Stdout, format, LevelInfo)
	l.Infof(default_prefix, t.Name())
	l.Info("test")
}

func TestSetLevel(t *testing.T) {
	l := New(os.Stdout, default_log_format, LevelInfo)
	l.Warnf(default_prefix, t.Name())
	l.Trace("trace log")
	l.Debug("debug log")
	l.Info("info log")
	l.Warn("warn log")
	l.Error("error log")
	l.Fatal("fatal log")
}
