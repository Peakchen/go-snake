package ulog

import (
	"encoding/json"
	"fmt"
)

type JSONFormatter struct {
	TimestampFormat string
	PrettyPrint     bool
}

func (self *JSONFormatter) GetTime(entry *Entry) string {
	var timeFormat string
	if self.TimestampFormat != "" {
		timeFormat = self.TimestampFormat
	} else {
		timeFormat = JsonTimeFormat
	}
	return entry.Time.Format(timeFormat)
}
func (self *JSONFormatter) Format(entry *Entry) ([]byte, error) {

	b := entry.Buffer

	if entry.HasCaller() {
		entry.Data["@file"] = fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	}

	entry.Data["@time"] = self.GetTime(entry)
	entry.Data["@level"] = entry.Level.String()
	entry.Data["@msg"] = entry.Message

	encoder := json.NewEncoder(b)
	if self.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(entry.Data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}

	return b.Bytes(), nil
}
