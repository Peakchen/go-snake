package ulog

import (
	"fmt"
)

// 默认文本输出格式
// 颜色优先级从高到低:
// 1. 开启颜色(EnableColor)
// 2. 指定颜色(WithColorName, WithColor)
// 3. 日志文本匹配规则(ParseColorRule)
// 4. 日志级别对应颜色(GetColorDefineByLevel)
type TextFormatter struct {
	EnableColor     bool // 命令行着色
	TimestampFormat string
	CallerFileLevel int // 调用栈文件名显示长度, -1: 完整, 0: 文件名, 1: 文件+1级文件夹
	rule            *ColorRuleSet
}

// 从颜色规则文本读取规则
func (self *TextFormatter) ParseColorRule(ruleText string) error {
	self.rule = NewColorRuleSet()
	err := self.rule.Parse(ruleText)
	if err != nil {
		return fmt.Errorf("parse color rule failed: %w", err)
	}

	return nil
}

func (self *TextFormatter) matchText(text string) *ColorDefine {

	if self.rule == nil {
		return nil
	}

	return self.rule.MatchText(text)
}

// 取得颜色前缀
func (self *TextFormatter) GetPrefix(entry *Entry) *ColorDefine {

	var cdef *ColorDefine

	if self.EnableColor {

		cdef = entry.ColorDef

		if cdef == nil {
			cdef = self.matchText(entry.Message)

			if cdef == nil {
				cdef = GetColorDefineByLevel(entry.Level)
			}
		}
	} else {
		cdef = WhiteColorDef
	}

	return cdef
}

// 取得颜色后缀
func (self *TextFormatter) GetSuffix() string {
	if self.EnableColor {
		return consoleColorSuffix
	}
	return ""
}

// 取得时间
func (self *TextFormatter) GetTime(entry *Entry) string {
	var timeFormat string
	if self.TimestampFormat != "" {
		timeFormat = self.TimestampFormat
	} else {
		timeFormat = TextTimeFormat
	}
	return entry.Time.Format(timeFormat)
}

// 取得调用者
func (self *TextFormatter) GetCaller(entry *Entry) string {
	if entry.HasCaller() {
		return fmt.Sprintf("%s:%d", trimPath(entry.Caller.File, self.CallerFileLevel), entry.Caller.Line)
	}

	return ""
}

func (self *TextFormatter) Format(entry *Entry) ([]byte, error) {

	b := entry.Buffer

	b.WriteString(self.GetPrefix(entry).Prefix)
	b.WriteString(entry.Level.String())
	b.WriteString("[")
	b.WriteString(self.GetTime(entry))
	b.WriteString("] ")
	b.WriteString(self.GetCaller(entry))
	b.WriteString(" ")
	b.WriteString(entry.Message)

	for k, v := range entry.Data {
		fmt.Fprintf(b, " %s=%v", k, v)
	}

	b.WriteString(self.GetSuffix())

	b.WriteByte('\n')
	return b.Bytes(), nil
}
