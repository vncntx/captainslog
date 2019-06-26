# Captain's Log

[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://godoc.org/github.com/vincentfiestada/captainslog?status.svg)](https://godoc.org/github.com/vincentfiestada/captainslog)
[![Conventional Commits](https://img.shields.io/badge/commits-conventional-00b6ff.svg?labelColor=1F6CB4)](https://conventionalcommits.org)
[![License: MIT](https://img.shields.io/github/license/vincentfiestada/acrylic.svg?labelColor=1F6CB4&color=00b6ff)](https://github.com/vincentfiestada/acrylic/blob/master/LICENSE)

A simple logging library for [Go](https://golang.org/).

- Support for multiple logging levels
- Colored output, even on Windows
- Output name of calling function

![Screenshot of captainslog in action](./assets/screenshot.png)

## Usage

```go
package main

import log "github.com/vincentfiestada/captainslog"

func main() {
	log.SetLevel(log.LogLevelSilly)

	log.Info("Info %d", 1)
}
```

<small>Gopher artwork by [Ashley McNamara](https://twitter.com/ashleymcnamara) on [Gopherize.me](https://gopherize.me/gopher/5dcbe4dc48ab6fbf903aae352f8742cb59e7099b)</small>
