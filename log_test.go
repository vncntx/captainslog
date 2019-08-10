package captainslog_test

import (
	"os"
	"testing"

	"github.com/vincentfiestada/captainslog"
	"github.com/vincentfiestada/captainslog/preflight"
)

// ANSI Color Codes
const (
	red    = "31m"
	green  = "32m"
	yellow = "33m"
	blue   = "34m"
	purple = "35m"
)

// Patterns
const (
	timePattern = "([0-9]{2}-){2}[0-9]{4} ([0-9]{2}:){2}[0-9]{2} .{2,3}"
	funcPattern = "func[0-9]+"
)

func getLogger() *captainslog.Logger {
	log := captainslog.NewLogger()
	log.Level = captainslog.LogLevelTrace
	return log
}

func TestNewLogger(test *testing.T) {
	t := preflight.Unit(test)

	log := captainslog.NewLogger()

	// should return a new Logger instance
	t.Expect(log).Is().Not().Nil()

	// should use default values
	t.Expect(log.HasColor).Is().EqualTo(true)
	t.Expect(log.LogFormat()).Is().Not().Nil()
	t.Expect(log.Level).Equals(captainslog.LogLevelDebug)
	t.Expect(log.TimeFormat).Equals(captainslog.DefaultTimeFormat)
	t.Expect(log.MaxNameLength).Equals(captainslog.DefaultMaxNameLength)
	t.Expect(log.Stdout).Equals(os.Stdout)
	t.Expect(log.Stderr).Equals(os.Stderr)
}

func TestLogStderr(test *testing.T) {
	t := preflight.Unit(test)

	logs := t.ExpectLogs(func(out *os.File) {
		log := getLogger()
		log.Stderr = out

		log.Trace("message %d", 1)
		log.Debug("message %d", 2)
		log.Info("message %d", 3)
		log.Warn("message %d", 4)
		log.Error("message %d", 5)
	})

	t.Expect(logs).HasLength(2)

	logs[0].Fields.Is().Empty()
	logs[0].Name.Matches(funcPattern)
	logs[0].Time.Matches(timePattern)
	logs[0].Level.Equals("warn")
	logs[0].Message.Equals("message 4")

	logs[0].Fields.Is().Empty()
	logs[1].Name.Matches(funcPattern)
	logs[1].Time.Matches(timePattern)
	logs[1].Level.Equals("error")
	logs[1].Message.Equals("message 5")
}

func TestLogStdout(test *testing.T) {
	t := preflight.Unit(test)

	logs := t.ExpectLogs(func(out *os.File) {
		log := getLogger()
		log.Stdout = out

		log.Trace("message %d", 1)
		log.Debug("message %d", 2)
		log.Info("message %d", 3)
		log.Warn("message %d", 4)
		log.Error("message %d", 5)
	})

	t.Expect(logs).HasLength(3)

	logs[0].Fields.Is().Empty()
	logs[0].Name.Matches(funcPattern)
	logs[0].Time.Matches(timePattern)
	logs[0].Level.Equals("trace")
	logs[0].Message.Equals("message 1")

	logs[0].Fields.Is().Empty()
	logs[1].Name.Matches(funcPattern)
	logs[1].Time.Matches(timePattern)
	logs[1].Level.Equals("debug")
	logs[1].Message.Equals("message 2")

	logs[0].Fields.Is().Empty()
	logs[2].Name.Match(funcPattern)
	logs[2].Time.Matches(timePattern)
	logs[2].Level.Equals("info")
	logs[2].Message.Equals("message 3")
}

//func TestLog(test *testing.T) {

//t := x.NewUnit(test)
//
//t.Expect(1 + 1).Equals(2)

//// intercept stdout and stderr
//out, outChan, readOut := t.intercept(os.Stdout)
//err, errChan, readErr := t.intercept(os.Stderr)
//log.Stdout = out
//log.Stderr = err
//
//// read from streams to channels
//go readOut()
//go readErr()
//
//name := "TestLog"
//message := "x 1"
//// perform log assertions
//for _, logF := range []func(string, ...interface{}){
//	log.Trace,
//	log.Debug,
//	log.Info,
//	log.Warn,
//	log.Error,
//} {
//	logF("x %d", 1)
//}
//out.Close()
//err.Close()
//
//logs := strings.Split(<-outChan, "\n")
//expect.Equal(t, 3, len(logs)-1)
//t.assertLog(logs[0], purple, "trace", name, message)
//t.assertLog(logs[1], green, "debug", name, message)
//t.assertLog(logs[2], blue, "info", name, message)
//
//logs = strings.Split(<-errChan, "\n")
//assert.Equal(t, 2, len(logs)-1)
//t.assertLog(logs[0], yellow, "warn", name, message)
//t.assertLog(logs[1], red, "error", name, message)
//}

// func TestFatal(test *testing.T) {

// 	t := x.NewUnit(test)

// 	log := captainslog.NewLogger()
// 	log.SetLevel(captainslog.LogLevelTrace)

// 	t.Expect(1 + 1).Equals(2)
// 	//t := x.Unit{test}
// 	//
// 	//if os.Getenv("EXIT") == "1" {
// 	//	log.Fatal("x")
// 	//	return
// 	//}
// 	//cmd := exec.Command(os.Args[0], "-x.run=TestFatal")
// 	//cmd.Env = append(os.Environ(), "EXIT=1")
// 	//t.assertExitCode(cmd., 1)
// }

// func TestExit(test *testing.T) {

// 	t := x.NewUnit(test)

// 	log := captainslog.NewLogger()
// 	log.SetLevel(captainslog.LogLevelTrace)

// 	t.Expect(1 + 1).Equals(2)
// 	//t := unit{test}
// 	//
// 	//if os.Getenv("EXIT") == "1" {
// 	//	log.Exit(2, "x")
// 	//	return
// 	//}
// 	//cmd := exec.Command(os.Args[0], "-x.run=TestExit")
// 	//cmd.Env = append(os.Environ(), "EXIT=1")
// 	//t.assertExitCode(cmd, 2)
// }

// func TestPanic(test *testing.T) {

// 	t := x.NewUnit(test)

// 	log := captainslog.NewLogger()
// 	log.SetLevel(captainslog.LogLevelTrace)

// 	t.Expect(1 + 1).Equals(2)
// 	//t := unit{test}
// 	//
// 	//log := captainslog.NewLogger()
// 	//
// 	//// intercept stderr
// 	//err, errChan, readErr := t.intercept(os.Stderr)
// 	//log.Stderr = err
// 	//
// 	//defer func() {
// 	//	// must log fatal error
// 	//	err.Close()
// 	//	logs := <-errChan
// 	//	t.assertLog(logs, red, "fatal", "TestPanic", "x 1")
// 	//
// 	//	// must panic with error
// 	//	panicLog := recover().(error)
// 	//	assert.Equal(t, "x 1", panicLog.Error())
// 	//}()
// 	//
// 	//go readErr()
// 	//log.Panic("x %d", 1)
// }

// func TestSetName(test *testing.T) {

// 	t := x.NewUnit(test)

// 	log := captainslog.NewLogger()
// 	log.SetLevel(captainslog.LogLevelTrace)

// 	t.Expect(1 + 1).Equals(2)
// 	//t := unit{test}
// 	//
// 	//log := captainslog.NewLogger()
// 	//log.SetName("captain")
// 	//
// 	//// should set the name
// 	//assert.Equal(t, "captain", log.GetName())
// }

// func TestSetTimeFormat(test *testing.T) {

// 	t := x.NewUnit(test)

// 	log := captainslog.NewLogger()
// 	log.SetLevel(captainslog.LogLevelTrace)

// 	t.Expect(1 + 1).Equals(2)

// 	//t := unit{test}
// 	//
// 	//log := captainslog.NewLogger()
// 	//log.SetTimeFormat("2006/01/02 03:04 PM")
// 	//
// 	//// intercept stdout
// 	//out, ch, read := t.intercept(os.Stdout)
// 	//log.Stdout = out
// 	//go read()
// 	//
// 	//log.Info("x")
// 	//out.Close()
// 	//
// 	//// datetime must be formatted correctly
// 	//logs := <-ch
// 	//datetime := strings.Split(logs, " :: ")[1]
// 	//t.matches("[0-9]{4}(/[0-9]{2}){2} [0-9]{2}:[0-9]{2} (AM|PM)", datetime)
// }

// // -------- x utils --------

// // intercept creates a channel for intercepting a stream
// //func (test unit) intercept(target *os.FileExpectation) (w *os.FileExpectation, ch chan string, read func()) {
// //	r, w, err := os.Pipe()
// //	if err != nil {
// //		test.Error(err)
// //	}
// //	ch = make(chan string)
// //	read = func() {
// //		var buf bytes.Buffer
// //		_, err := io.Copy(&buf, r)
// //		if err != nil {
// //			test.Error(err)
// //		}
// //		ch <- buf.String()
// //		r.Close()
// //	}
// //	return w, ch, read
// //}

// // colorless removes ANSI color codes from text and returns the color
// //func (test unit) colorless(text string) (colorless string, color string) {
// //	rxp, err := regexp.Compile("\\x1b\\[[0-9;]*m")
// //	if err != nil {
// //		test.Error(err)
// //	}
// //	colorCode := rxp.FindString(text)
// //	if len(colorCode) >= 3 {
// //		// get the color code if it is present
// //		color = colorCode[len(colorCode)-3:]
// //	}
// //	return rxp.ReplaceAllString(text, ""), color
// //}

// // matches asserts that a string matches a regular expression
// //func (test unit) matches(pattern string, text string) {
// //	matches, err := regexp.MatchString(pattern, text)
// //	if err != nil {
// //		test.Error(err)
// //	}
// //	assert.True(test, matches)
// //}

// // assertLog asserts the contents of a log message
// //func (test unit) assertLog(log string, expectedColor string, expectedLevel string, expectedName string, expectedMsg string) {
// //	text, color := test.colorless(log)
// //	parts := strings.Split(text, " :: ")
// //	level, time, name, message := parts[0], parts[1], parts[2], parts[3]
// //	if len(color) > 0 {
// //		// assert color only if log has colors
// //		assert.Equal(test, expectedColor, color)
// //	}
// //	assert.Equal(test, expectedLevel, strings.TrimLeft(level, " "))
// //	assert.Equal(test, expectedName, name)
// //	assert.Equal(test, expectedMsg, strings.TrimRight(message, "\n"))
// //	test.matches("([0-9]{2}-){2}[0-9]{4} ([0-9]{2}:){2}[0-9]{2} .{3}", time)
// //}
// //
// //func (test unit) assertExitCode(cmd *exec.Cmd, expectedCode int) {
// //	exit := cmd.Run()
// //	if e, ok := exit.(*exec.ExitError); ok {
// //		assert.Equal(test, expectedCode, e.ExitCode())
// //		return
// //	}
// //	// failed to retrieve exit code
// //	test.Error("could not get exit code")
// //}
