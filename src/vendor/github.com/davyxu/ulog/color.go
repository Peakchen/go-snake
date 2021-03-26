package ulog

import (
	"strings"
)

type Color int

// 命令行着色
const (
	White Color = iota
	Black
	Red
	Green
	Yellow
	Blue
	Purple
	DarkGreen
	Gray
)

func (self Color) String() string {
	return GetColorDefine(self).Name
}

type ColorDefine struct {
	Name   string
	Color  Color
	Prefix string
}

// 每种颜色定义的名称, 命令行加色
var colorDefines = []*ColorDefine{
	{Name: "white", Color: White, Prefix: ""},
	{Name: "black", Color: Black, Prefix: "\x1b[030m"},
	{Name: "red", Color: Red, Prefix: "\x1b[031m"},
	{Name: "green", Color: Green, Prefix: "\x1b[032m"},
	{Name: "yellow", Color: Yellow, Prefix: "\x1b[033m"},
	{Name: "blue", Color: Blue, Prefix: "\x1b[034m"},
	{Name: "purple", Color: Purple, Prefix: "\x1b[035m"},
	{Name: "darkgreen", Color: DarkGreen, Prefix: "\x1b[036m"},
	{Name: "gray", Color: Gray, Prefix: "\x1b[037m"},
}

var (
	consoleColorSuffix = "\x1b[0m"
	WhiteColorDef      *ColorDefine
	YellowColorDef     *ColorDefine
	RedColorDef        *ColorDefine
	GrayColorDef       *ColorDefine
)

// 从名称取得定义
func GetColorDefineByName(name string) *ColorDefine {

	lower := strings.ToLower(name)

	for _, d := range colorDefines {

		if d.Name == lower {
			return d
		}
	}
	return nil
}

// 从枚举取得定义
func GetColorDefine(c Color) *ColorDefine {

	for _, d := range colorDefines {

		if d.Color == c {
			return d
		}
	}
	return nil
}

// 根据日志级别, 获得命令行颜色
func GetColorDefineByLevel(l Level) *ColorDefine {
	switch l {
	case DebugLevel:
		return GrayColorDef
	case WarnLevel:
		return YellowColorDef
	case ErrorLevel:
		return RedColorDef
	}

	return WhiteColorDef
}

func init() {
	WhiteColorDef = GetColorDefine(White)
	YellowColorDef = GetColorDefine(Yellow)
	RedColorDef = GetColorDefine(Red)
	GrayColorDef = GetColorDefine(Gray)
}
