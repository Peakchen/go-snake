package akLog

import "testing"

func TestLog(t *testing.T) {
	Error("1111")
	Info("%v", "1111")
}
