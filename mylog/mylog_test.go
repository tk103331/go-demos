package mylog

import (
	"testing"
)

const (
	default_prefix = "  %s  "
)

func TestStdLog(t *testing.T) {
	std.Info(t.Name())
}
