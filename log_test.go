package captainslog_test

import (
	"testing"

	"github.com/vincentfiestada/captainslog"
)

func TestDemo(t *testing.T) {
	log := captainslog.NewLogger()
	log.SetLevel(captainslog.LogLevelSilly)

	log.Silly("%s", "silly")
	log.Debug("%s", "debug")
	log.Verbose("%s", "verbose")
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

	log := captainslog.NewLogger()
	log.Panic("%s", "panic")
}
