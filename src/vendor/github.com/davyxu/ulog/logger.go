package ulog

import (
	"fmt"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Logger struct {
	Output io.Writer // 日志输出

	Formatter Formatter // 格式化器

	ReportCaller bool

	level Level // 日志输出级别

	entryPool sync.Pool // 分配Entry的池

	guard sync.RWMutex
}

// 设置日志格式化器
func (self *Logger) SetFormatter(formatter Formatter) {
	self.guard.Lock()
	defer self.guard.Unlock()
	self.Formatter = formatter
}

// 设置日志输出
func (self *Logger) SetOutput(output io.Writer) {
	self.guard.Lock()
	defer self.guard.Unlock()
	self.Output = output
}

// 设置是否显示调用者
func (self *Logger) SetReportCaller(reportCaller bool) {
	self.guard.Lock()
	defer self.guard.Unlock()
	self.ReportCaller = reportCaller
}

// 设置当前日志输出的级别
func (logger *Logger) SetLevel(level Level) {
	atomic.StoreUint32((*uint32)(&logger.level), uint32(level))
}

// 获取当前日志输出的级别
func (self *Logger) GetLevel() Level {
	return Level(atomic.LoadUint32((*uint32)(&self.level)))
}

func (self *Logger) IsLevelEnabled(level Level) bool {
	return self.GetLevel() <= level
}

// 指定当前行的日志的颜色, 覆盖按级别以及从日志文本应有的颜色
func (self *Logger) WithColorName(colorName string) *Entry {

	entry := self.allocEntry()
	return entry.WithColorName(colorName)
}

// 指定当前行的日志的颜色, 覆盖按级别以及从日志文本应有的颜色
func (self *Logger) WithColor(color Color) *Entry {

	entry := self.allocEntry()
	return entry.WithColor(color)
}

// 设置当前行日志的附加字段
func (self *Logger) WithFields(fields Fields) *Entry {

	entry := self.allocEntry()
	return entry.WithFields(fields)
}

// 设置当前行日志的附加字段
func (self *Logger) WithField(key string, value interface{}) *Entry {
	entry := self.allocEntry()
	return entry.WithField(key, value)
}

func (self *Logger) Logf(level Level, format string, args ...interface{}) {
	if self.IsLevelEnabled(level) {
		entry := self.allocEntry()
		entry.Log(level, fmt.Sprintf(format, args...))
	}
}

func (self *Logger) Logln(level Level, args ...interface{}) {
	if self.IsLevelEnabled(level) {
		entry := self.allocEntry()
		entry.Logln(level, args...)
	}
}

func (self *Logger) Debugf(format string, args ...interface{}) {
	self.Logf(DebugLevel, format, args...)
}

func (self *Logger) Infof(format string, args ...interface{}) {
	self.Logf(InfoLevel, format, args...)
}

func (self *Logger) Warnf(format string, args ...interface{}) {
	self.Logf(WarnLevel, format, args...)
}

func (self *Logger) Errorf(format string, args ...interface{}) {
	self.Logf(ErrorLevel, format, args...)
}

func (self *Logger) Debugln(args ...interface{}) {
	self.Logln(DebugLevel, args...)
}

func (self *Logger) Infoln(args ...interface{}) {
	self.Logln(InfoLevel, args...)
}

func (self *Logger) Warnln(args ...interface{}) {
	self.Logln(WarnLevel, args...)
}

func (self *Logger) Errorln(args ...interface{}) {
	self.Logln(ErrorLevel, args...)
}

func (self *Logger) allocEntry() *Entry {
	entry, ok := self.entryPool.Get().(*Entry)
	if !ok {
		entry = NewEntry(self)
	}

	entry.Reset()

	entry.needFree = true
	return entry
}

func (self *Logger) freeEntry(entry *Entry) {
	if entry.needFree {

		self.entryPool.Put(entry)
	}
}

func New() *Logger {
	return &Logger{
		Output:    os.Stdout,
		Formatter: &TextFormatter{},
		level:     DebugLevel,
	}
}
