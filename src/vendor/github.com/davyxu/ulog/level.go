package ulog

import "strings"

type Level uint32

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (self Level) String() string {
	return levelString[self]
}

func ParseLevelString(str string) (Level, bool) {

	switch strings.ToLower(str) {
	case "debug", "debu":
		return DebugLevel, true
	case "info":
		return InfoLevel, true
	case "warn", "warnning":
		return WarnLevel, true
	case "error", "erro":
		return ErrorLevel, true
	}

	return DebugLevel, false
}

var levelString = [...]string{
	"DEBU",
	"INFO",
	"WARN",
	"ERRO",
}
