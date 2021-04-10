package ulog

import "fmt"

// 命令行格式, 只有输出文本和可选颜色
type ConsoleFormatter struct {
	EnableColor bool // 命令行着色
}

// 取得颜色前缀
func (self *ConsoleFormatter) GetPrefix(entry *Entry) *ColorDefine {

	var cdef *ColorDefine

	if self.EnableColor {

		cdef = entry.ColorDef

		if cdef == nil {
			cdef = GetColorDefineByLevel(entry.Level)
		}
	} else {
		cdef = WhiteColorDef
	}

	return cdef
}

// 取得颜色后缀
func (self *ConsoleFormatter) GetSuffix() string {
	if self.EnableColor {
		return consoleColorSuffix
	}
	return ""
}
func (self *ConsoleFormatter) Format(entry *Entry) ([]byte, error) {

	b := entry.Buffer

	b.WriteString(self.GetPrefix(entry).Prefix)
	b.WriteString(entry.Message)

	for k, v := range entry.Data {
		fmt.Fprintf(b, " %s=%v", k, v)
	}

	b.WriteString(self.GetSuffix())

	b.WriteByte('\n')
	return b.Bytes(), nil
}
