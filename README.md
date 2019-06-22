# Captain's Log

[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://godoc.org/github.com/vincentfiestada/captainslog?status.svg)](https://godoc.org/github.com/vincentfiestada/captainslog)

A simple logging library for [Go](https://golang.org/).

- Support for multiple logging levels
- Colored output, even on Windows
- Output name of calling function

## Usage

```go
package main

import log "github.com/vincentfiestada/captainslog"

func main() {
	log.SetLevel(LogLevelSilly)

	log.Silly("This is %s", "Silly")
	log.Debug("This is %s", "Debug")
	log.Verbose("This is %s", "Verbose")
	log.Info("This is %s", "Info")
	log.Warn("This is %s", "Warning")
	log.Error("This is %s", "Error")
}
```
