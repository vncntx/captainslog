package captainslog_test

import (
	"os"
	"testing"
	"time"

	"github.com/vincentfiestada/captainslog"
	"github.com/vincentfiestada/captainslog/preflight"
)

// Patterns
const (
	iso8601 = "([0-9]{2}-){2}[0-9]{4} ([0-9]{2}:){2}[0-9]{2} .+"
)

func getLogger() *captainslog.Logger {
	log := captainslog.NewLogger()
	log.Level = captainslog.LevelTrace
	return log
}

func TestNewLogger(test *testing.T) {
	t := preflight.Unit(test)

	log := captainslog.NewLogger()

	// should return a new Logger instance
	t.Expect(log).Is().Not().Nil()

	// should use default values
	t.Expect(log.HasColor).Is().EqualTo(true)
	t.Expect(log.Level).Equals(captainslog.LevelDebug)
	t.Expect(log.TimeFormat).Equals(captainslog.ISO8601)
	t.Expect(log.MaxNameLength).Equals(15)
	t.Expect(log.Stdout).Equals(os.Stdout)
	t.Expect(log.Stderr).Equals(os.Stderr)
}

func TestLogs(test *testing.T) {
	t := preflight.Unit(test)

	stdout, stderr := t.CaptureLogs(func(stdout *os.File, stderr *os.File) {
		log := getLogger()
		log.Stdout = stdout
		log.Stderr = stderr

		log.Trace("message %d", 1)
		log.Debug("message %d", 2)
		log.Info("message %d", 3)
		log.Warn("message %d", 4)
		log.Error("message %d", 5)
	})

	check := func(log *preflight.LogExpectation, level string, message string) {
		log.Fields.Is().Empty()
		log.Level.Equals(level)
		log.Message.Equals(message)
		log.Time.Matches(iso8601)
		log.Name.Matches("func[0-9]+")
	}

	// trace, debug, info should go to stdout
	// warn, error should go to stderr
	t.Expect(stdout).HasLength(3)
	t.Expect(stderr).HasLength(2)

	// log messages should use the correct name, level, & message
	check(stdout[0], "trace", "message 1")
	check(stdout[1], "debug", "message 2")
	check(stdout[2], "info", "message 3")

	check(stderr[0], "warn", "message 4")
	check(stderr[1], "error", "message 5")
}

func TestName(test *testing.T) {
	t := preflight.Unit(test)

	expectedName := "captainslog"

	logs, _ := t.CaptureLogs(func(stdout *os.File, _ *os.File) {
		log := getLogger()
		log.Stdout = stdout
		log.Name = expectedName

		log.Info("x")
	})

	logs[0].Name.Equals(expectedName)
}

func TestTimeFormat(test *testing.T) {
	t := preflight.Unit(test)

	rfc822 := "[0-9]{2} [A-Z][a-z]{2} [0-9]{2} [0-9]{2}:[0-9]{2} .+"

	logs, _ := t.CaptureLogs(func(stdout *os.File, _ *os.File) {
		log := getLogger()
		log.Stdout = stdout
		log.TimeFormat = time.RFC822

		log.Info("x")
	})

	logs[0].Time.Matches(rfc822)
}

func TestExit(test *testing.T) {
	t := preflight.Unit(test)

	_, logs := t.CaptureLogs(func(_ *os.File, stderr *os.File) {
		t.ExpectExitCode(func() {
			log := captainslog.NewLogger()
			log.Stderr = stderr

			log.Exit(2, "message")

		}).Equals(2)
	})

	logs[0].Level.Equals("fatal")
	logs[0].Message.Equals("message")
}

func TestFatal(test *testing.T) {
	t := preflight.Unit(test)

	t.ExpectExitCode(func() {
		log := captainslog.NewLogger()

		log.Fatal("message")

	}).Equals(1)
}

func TestPanic(test *testing.T) {
	t := preflight.Unit(test)

	defer func() {
		t.Expect(recover().(error).Error()).Equals("x")
	}()

	_, logs := t.CaptureLogs(func(_ *os.File, stderr *os.File) {
		log := getLogger()
		log.Stderr = stderr

		log.Panic("x")
	})

	logs[0].Message.Equals("x")
	logs[0].Level.Equals("fatal")
}

func TestFields(test *testing.T) {
	t := preflight.Unit(test)

	logs, _ := t.CaptureLogs(func(stdout *os.File, stderr *os.File) {
		log := getLogger()
		log.Stdout = stdout
		log.Stderr = stderr

		log.Field("captain", "picard").Info("energize")
	})

	logs[0].Message.Equals("energize")
	logs[0].Fields.Equals("captain=\"picard\"")
}

func TestLevels(test *testing.T) {
	t := preflight.Unit(test)

	stdout, stderr := t.CaptureLogs(func(stdout *os.File, stderr *os.File) {
		log := getLogger()
		log.Stdout = stdout
		log.Stderr = stderr
		log.Level = captainslog.LevelWarn

		log.Info("x")
		log.Warn("x")
	})

	t.Expect(stdout).HasLength(0)
	t.Expect(stderr).HasLength(1)
}
