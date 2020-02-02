package msg_test

import (
	"os"
	"testing"

	"github.com/vincentfiestada/captainslog/v2/format"
	"github.com/vincentfiestada/captainslog/v2/levels"
	"github.com/vincentfiestada/captainslog/v2/msg"
	"github.com/vincentfiestada/captainslog/v2/preflight"
)

func TestProps(test *testing.T) {
	t := preflight.Unit(test)

	names := []string{
		"trace",
		"debug",
		"info",
		"warn",
		"error",
		"fatal",
	}

	for l := levels.Trace; l < levels.Quiet; l++ {
		message := createMessage(l)
		stream, level, _ := message.Props()

		t.Expect(level).Equals(names[l])
		if l < levels.Warn {
			t.Expect(stream).Equals(os.Stdout)
		} else {
			t.Expect(stream).Equals(os.Stderr)
		}
	}
}

func TestLogs(test *testing.T) {
	t := preflight.Unit(test)

	message := createMessage(levels.Info)

	message.Format = func(input *msg.Message) {
		t.Expect(input).Equals(message)
		t.Expect(input.Data).HasLength(0)
	}

	message.Trace("captainslog")
	message.Debug("captainslog")
	message.Info("captainslog")
	message.Warn("captainslog")
	message.Error("captainslog")
}

func TestExit(test *testing.T) {
	t := preflight.Unit(test)

	message := createMessage(levels.Info)

	t.ExpectExitCode(func() {
		message.Exit(2, "captainslog")
	}).Equals(2)
}

func TestFatal(test *testing.T) {
	t := preflight.Unit(test)

	message := createMessage(levels.Info)

	t.ExpectExitCode(func() {
		message.Fatal("captainslog")
	}).Equals(1)
}

func TestPanic(test *testing.T) {
	t := preflight.Unit(test)

	defer func() {
		t.Expect(recover().(error).Error()).Equals("x")
	}()

	message := createMessage(levels.Info)
	message.Panic("x")
}

func TestField(test *testing.T) {
	t := preflight.Unit(test)

	message := createMessage(levels.Info)
	message.Field("science officer", "data")

	t.Expect(message.Data).HasLength(2)
}

func TestFields(test *testing.T) {
	t := preflight.Unit(test)

	message := createMessage(levels.Info)
	message.Fields(
		msg.Field{"science officer", "data"},
		msg.Field{"chief engineer", "geordi la forge"},
	)

	t.Expect(message.Data).HasLength(4)
}

/**
 * Test Helpers
 */
func createMessage(level int) *msg.Message {
	return &msg.Message{
		Time:      "07-23-1996 07:23:00 PST",
		Name:      "captainslog",
		Level:     level,
		Threshold: levels.Trace,
		Stdout:    os.Stdout,
		Stderr:    os.Stderr,
		Format:    format.Flat,
		Data:      []interface{}{},
	}
}
