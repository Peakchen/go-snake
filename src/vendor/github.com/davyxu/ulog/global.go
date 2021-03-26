package ulog

import (
	"io"
)

var (
	glog = New()
)

// 获取全局日志
func Global() *Logger {
	return glog
}

// 设置日志格式化器
func SetFormatter(formatter Formatter) {
	glog.SetFormatter(formatter)
}

// 设置日志输出
func SetOutput(out io.Writer) {
	glog.SetOutput(out)
}

// 设置当前日志输出的级别
func SetLevel(level Level) {
	glog.SetLevel(level)
}

// 获取当前日志输出的级别
func GetLevel() Level {
	return glog.GetLevel()
}

// 判断当前级别的日志是否开启
func IsLevelEnabled(level Level) bool {
	return glog.IsLevelEnabled(level)
}

// 指定当前行的日志的颜色, 覆盖按级别以及从日志文本应有的颜色
func WithColorName(colorName string) *Entry {
	return glog.WithColorName(colorName)
}

// 指定当前行的日志的颜色, 覆盖按级别以及从日志文本应有的颜色
func WithColor(color Color) *Entry {
	return glog.WithColor(color)
}

// 设置当前行日志的附加字段
func WithField(key string, value interface{}) *Entry {
	return glog.WithField(key, value)
}

// 设置当前行日志的附加字段
func WithFields(fields Fields) *Entry {
	return glog.WithFields(fields)
}

func Debugf(format string, args ...interface{}) {
	glog.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	glog.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	glog.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	glog.Errorf(format, args...)
}

func Debugln(args ...interface{}) {
	glog.Debugln(args...)
}

func Infoln(args ...interface{}) {
	glog.Infoln(args...)
}

func Warnln(args ...interface{}) {
	glog.Warnln(args...)
}

func Errorln(args ...interface{}) {
	glog.Errorln(args...)
}
