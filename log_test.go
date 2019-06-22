package captainslog

import "testing"

func TestDemo(_ *testing.T) {
	log := NewLogger()
	log.SetLevel(LogLevelSilly)

	log.Silly("Test %d", 1)
	log.Debug("Test %d", 2)
	log.Verbose("Test %d", 3)
	log.Info("Test %d", 4)
	log.Warn("Test %d", 5)
	log.Error("Test %d", 6)
}
