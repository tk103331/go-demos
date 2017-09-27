package mylog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
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
	default_log_format = "2006-01-02 15:04:05.000|camel|short"

	default_time_format  = "2006-01-02 15:04:05.000"
	default_level_format = "lower" //lower upper camel
	default_file_format  = "short" //short long

	sep = ' '
)

var (
	levelFmtOps = []string{"upper", "lower", "camel"}
	fileFmtOps  = []string{"long", "short"}
)

var std *myLogger = New(os.Stdout, default_log_format, LevelInfo)

type myLogger struct {
	mu      sync.Mutex
	out     io.Writer
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
	arr := strings.SplitN(format+"||||", "|", 4)
	fi := genFmt(arr[0], arr[1], arr[2])
	return &myLogger{out: writer, lvl: level, fmtInfo: fi}
}

func (l *myLogger) SetOutput(writer io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.out = writer
}

func (l *myLogger) SetFormat(timeFmt string, levelFmt string, fileFmt string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.fmtInfo = genFmt(timeFmt, levelFmt, fileFmt)
}

func (l *myLogger) SetLevel(level int) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.lvl = level
}

func genFmt(timeFmt string, levelFmt string, fileFmt string) fmtInfo {
	level := default_level_format
	if contains(levelFmtOps, levelFmt) {
		level = levelFmt
	}
	file := default_file_format
	if contains(fileFmtOps, fileFmt) {
		file = fileFmt
	}
	return fmtInfo{
		time:  timeFmt,
		level: level,
		file:  file,
	}
}

func contains(arr []string, it string) bool {
	for _, v := range arr {
		if v == it {
			return true
		}
	}
	return false
}

func (l *myLogger) output(lvl int, msg string) {
	if lvl >= l.lvl {
		timeStr := l.timeStr()
		lvlStr := l.levelStr(lvl)
		fileStr, lineStr := l.posStr()

		buf := bytes.NewBuffer([]byte{})
		buf.WriteString(timeStr)
		buf.WriteByte(sep)
		buf.WriteString(lvlStr)
		buf.WriteByte(sep)
		buf.WriteString(fileStr)
		buf.WriteByte(sep)
		buf.WriteString(lineStr)
		buf.WriteByte(sep)
		buf.WriteString(msg)
		buf.WriteByte('\n')
		l.out.Write([]byte(buf.String()))
	}
}

func (l *myLogger) timeStr() string {
	return time.Now().Format(l.fmtInfo.time)
}

func (l *myLogger) levelStr(lvl int) string {
	lvlStr := ""
	switch lvl {
	case LevelTrace:
		lvlStr = "trace"
	case LevelDebug:
		lvlStr = "debug"
	case LevelInfo:
		lvlStr = "info"
	case LevelWarn:
		lvlStr = "warn"
	case LevelError:
		lvlStr = "error"
	case LevelFatal:
		lvlStr = "fatal"
	}
	switch l.fmtInfo.level {
	case "upper":
		lvlStr = strings.ToUpper(lvlStr)
	case "lower":
		lvlStr = strings.ToLower(lvlStr)
	case "camel":
		lvlStr = strings.ToUpper(lvlStr[0:1]) + strings.ToLower(lvlStr[1:])
	}
	return lvlStr
}

func (l *myLogger) posStr() (string, string) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file, line = "???", 0
	}
	fileStr := file
	switch l.fmtInfo.file {
	case "long":
		fileStr = file
	case "short":
		fileStr = path.Base(file)
	}
	lineStr := strconv.Itoa(line)

	return fileStr, lineStr
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
