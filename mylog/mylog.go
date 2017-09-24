package mylog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	LevelAll = iota
	LevelTrace
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)

const (
	default_log_format = "#time #level #file:#line #msg"

	default_time_format  = "2006-01-02 15:04:05.000"
	default_level_format = "lower" //lower upper camel
	default_file_format  = "short" //short long package
	default_line_format  = "0000"
)

var std *myLogger = New(os.Stdout, default_log_format, LevelInfo)

type myLogger struct {
	mu      sync.Mutex
	out     io.Writer
	fmt     string
	lvl     int
	fmtInfo fmtInfo
}

type fmtInfo struct {
	time  string
	level string
	file  string
	line  string
}

func New(writer io.Writer, format string, level int) *myLogger {
	return &myLogger{out: writer, fmt: format, lvl: level}
}

func (l *myLogger) SetOutput(writer io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = writer
}

func (l *myLogger) SetFormat(format string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fmt = format
}

func (l *myLogger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lvl = level
}

func (l *myLogger) output(lvl int, msg string) {
	if lvl >= l.lvl {
		timeStr := l.timeStr()
		lvlStr := levelStr(lvl)
		fileStr, lineStr := l.posStr()
		str := l.fmt
		str = strings.Replace(str, "#time", timeStr, -1)
		str = strings.Replace(str, "#level", lvlStr, -1)
		str = strings.Replace(str, "#file", fileStr, -1)
		str = strings.Replace(str, "#line", lineStr, -1)
		str = strings.Replace(str, "#msg", msg, -1)
		l.out.Write([]byte(str))
		l.out.Write([]byte("\n"))
	}
}

func (l *myLogger) timeStr() string {
	return time.Now().Format(default_time_format)
}

func levelStr(lvl int) string {
	switch lvl {
	case LevelTrace:
		return "trace"
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	default:
		return ""
	}
}

func (l *myLogger) posStr() (string, string) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file, line = "???", 0
	}
	return file, strconv.Itoa(line)
}
func (l *myLogger) Trace(v ...interface{}) {
	l.output(LevelTrace, fmt.Sprint(v...))
}

func (l *myLogger) Tracef(s string, v ...interface{}) {
	l.output(LevelTrace, fmt.Sprintf(s, v...))
}
func (l *myLogger) Debug(v ...interface{}) {
	l.output(LevelDebug, fmt.Sprint(v...))
}

func (l *myLogger) Debugf(s string, v ...interface{}) {
	l.output(LevelDebug, fmt.Sprintf(s, v...))
}
func (l *myLogger) Info(v ...interface{}) {
	l.output(LevelInfo, fmt.Sprint(v...))
}

func (l *myLogger) Infof(s string, v ...interface{}) {
	l.output(LevelInfo, fmt.Sprintf(s, v...))
}
func (l *myLogger) Warn(v ...interface{}) {
	l.output(LevelWarn, fmt.Sprint(v...))
}

func (l *myLogger) Warnf(s string, v ...interface{}) {
	l.output(LevelWarn, fmt.Sprintf(s, v...))
}
func (l *myLogger) Error(v ...interface{}) {
	l.output(LevelError, fmt.Sprint(v...))
}

func (l *myLogger) Errorf(s string, v ...interface{}) {
	l.output(LevelError, fmt.Sprintf(s, v...))
}

func (l *myLogger) Fatal(v ...interface{}) {
	l.output(LevelFatal, fmt.Sprint(v...))
}

func (l *myLogger) Fatalf(s string, v ...interface{}) {
	l.output(LevelFatal, fmt.Sprintf(s, v...))
}
