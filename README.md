# Captain's Log

[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://godoc.org/github.com/vincentfiestada/captainslog?status.svg)](https://godoc.org/github.com/vincentfiestada/captainslog)
[![Conventional Commits](https://img.shields.io/badge/commits-conventional-00b6ff.svg?labelColor=1F6CB4)](https://conventionalcommits.org)
[![License: BSD-3](https://img.shields.io/github/license/vincentfiestada/captainslog.svg?labelColor=1F6CB4&color=00b6ff)](https://github.com/vincentfiestada/captainslog/blob/master/LICENSE)

A simple logging library for [Go](https://golang.org/).

- Support for multiple logging levels
- Colored output, even on Windows
- Print the calling function name

![Screenshot of captainslog in action](./assets/screenshot.png)

## Usage

```go
package main

import (
	"math"

	"github.com/vincentfiestada/captainslog"
)

var log *captainslog.Logger

func main() {
	log.Info("π = %v", math.Pi)
}

func init() {
	log = captainslog.NewLogger()
}

```

## Development

Run tests recursively with coverage:

```
go test ./... --cover
```

<small>Gopher artwork by [Ashley McNamara](https://twitter.com/ashleymcnamara) on [Gopherize.me](https://gopherize.me/gopher/5dcbe4dc48ab6fbf903aae352f8742cb59e7099b)</small>
