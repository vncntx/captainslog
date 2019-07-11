package captainslog_test

import (
	"testing"

	"github.com/vincentfiestada/captainslog"
)

var log *captainslog.Logger

func init() {
	log = captainslog.NewLogger()
	log.SetLevel(captainslog.LogLevelTrace)
}

func TestDemo(t *testing.T) {
	log.Trace("%s", "trace")
	log.Debug("%s", "debug")
	log.Info("%s", "info")
	log.Warn("%s", "warn")
	log.Error("%s", "error")
}

func TestPanic(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Fail()
		}
	}()

	log.Panic("%s", "panic")
}
