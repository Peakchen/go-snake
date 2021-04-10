package ulog

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var (
	bufferPool *sync.Pool
)

// 一行日志的内容, 提供给Hook和Formatter的参数
type Entry struct {
	Logger *Logger

	Data Fields // KV日志字段

	Time time.Time // 日志时间

	Level Level // 日志行的级别

	ColorDef *ColorDefine // 指定的日志行颜色

	Message string // 日志内容

	Buffer *bytes.Buffer // 最终生成的日志完整内容

	Caller *runtime.Frame // 日志调用者

	needFree bool
}

func (self *Entry) Reset() {
	if len(self.Data) > 0 {
		self.Data = map[string]interface{}{}
	}
	self.Time = time.Time{}
	self.Level = DebugLevel
	self.ColorDef = nil
	self.Message = ""
	self.Caller = nil
}

func (self *Entry) GetColorDefine() *ColorDefine {
	if self.ColorDef == nil {
		return WhiteColorDef
	}

	return self.ColorDef
}

func (self *Entry) WithColorName(colorName string) *Entry {
	self.ColorDef = GetColorDefineByName(colorName)
	return self
}

func (self *Entry) WithColor(color Color) *Entry {
	self.ColorDef = GetColorDefine(color)
	return self
}

func (self *Entry) WithField(key string, value interface{}) *Entry {
	self.Data[key] = value
	return self
}

func (self *Entry) WithFields(fields Fields) *Entry {

	for k, v := range fields {
		self.Data[k] = v
	}

	return self
}

func (self *Entry) Log(level Level, msg string) {
	self.log(level, msg)
	self.Logger.freeEntry(self)
}

func (self *Entry) Logln(level Level, args ...interface{}) {
	msg := fmt.Sprintln(args...)
	self.log(level, msg[:len(msg)-1])
	self.Logger.freeEntry(self)
}

func (self Entry) HasCaller() (has bool) {
	return self.Logger != nil &&
		self.Logger.ReportCaller &&
		self.Caller != nil
}

// 这里每次初始化
func (self Entry) log(level Level, msg string) {
	var buffer *bytes.Buffer

	// 如果没有With过Time, 这里的Time永远初始化
	if self.Time.IsZero() {
		self.Time = time.Now()
	}

	self.Level = level
	self.Message = msg

	if self.Logger.ReportCaller {
		self.Caller = getCaller()
	}

	buffer = bufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer bufferPool.Put(buffer)
	self.Buffer = buffer

	self.write()

	self.Buffer = nil
}
func (self *Entry) write() {
	self.Logger.guard.Lock()
	defer self.Logger.guard.Unlock()
	serialized, err := self.Logger.Formatter.Format(self)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
	} else {
		_, err = self.Logger.Output.Write(serialized)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
		}
	}
}

func (self *Entry) Debugf(format string, args ...interface{}) {
	self.Log(DebugLevel, fmt.Sprintf(format, args...))
}

func (self *Entry) Infof(format string, args ...interface{}) {
	self.Log(InfoLevel, fmt.Sprintf(format, args...))
}

func (self *Entry) Warnf(format string, args ...interface{}) {
	self.Log(WarnLevel, fmt.Sprintf(format, args...))
}

func (self *Entry) Errorf(format string, args ...interface{}) {
	self.Log(ErrorLevel, fmt.Sprintf(format, args...))
}

func (self *Entry) Debugln(args ...interface{}) {
	self.Logln(DebugLevel, args...)
}

func (self *Entry) Infoln(args ...interface{}) {
	self.Logln(InfoLevel, args...)
}

func (self *Entry) Warnln(args ...interface{}) {
	self.Logln(WarnLevel, args...)
}

func (self *Entry) Errorln(args ...interface{}) {
	self.Logln(ErrorLevel, args...)
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
		Data:   make(Fields, 6),
	}
}

func init() {
	bufferPool = &sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
}
