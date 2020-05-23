![](./logo.png)

# Captain's Log

[![](https://github.com/vincentfiestada/captainslog/workflows/Unit%20Tests/badge.svg)](https://github.com/vincentfiestada/captainslog/actions?query=workflow%3A%22Unit+Tests%22)
[![](https://github.com/vincentfiestada/captainslog/workflows/Style%20Checks/badge.svg)](https://github.com/vincentfiestada/captainslog/actions?query=workflow%3A%22Style+Checks%22)
[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://img.shields.io/badge/godoc-reference-0047ab?labelColor=16161b)](https://pkg.go.dev/github.com/vincentfiestada/captainslog/v2?tab=doc)
[![Conventional Commits](https://img.shields.io/badge/commits-conventional-0047ab.svg?labelColor=16161b)](https://conventionalcommits.org)
[![License: BSD-3](https://img.shields.io/github/license/vincentfiestada/captainslog.svg?labelColor=16161b&color=0047ab)](./license)

A simple logging library for [Go](https://golang.org/)

- Colors
- Multiple levels
- Structured logging
- Function name detection

![Screenshot of captainslog in action](./assets/screenshot.png)

## Installation

```
go get github.com/vincentfiestada/captainslog/v2
```

## Usage

This library is designed to provide a familiar yet powerful interface for logging. Each logging method accepts a format string and arguments. Structured logging is supported right out of the box. The Logger allows you to turn colors on/off, specify a datetime format, set the logging threshold, and even provide your own function to control how logs are written.

```go
package main

import (
	"github.com/vincentfiestada/captainslog/v2"
)

var log = captainslog.NewLogger()

func main() {

	log.Trace("this is %s", "captainslog")
	log.Debug("this is %s", "captainslog")
	log.Info("this is %s", "captainslog")
	log.Warn("this is %s", "captainslog")
	log.Error("this is %s", "captainslog")

	log.Fields(

		log.I("captain", "picard"),
		log.I("first officer", "riker"),
		log.I("science officer", "data"),
		log.I("medical officer", "crusher"),
		log.I("chief engineer", "la forge"),
		log.I("security officer", "worf"),

	).Info("starship enterprise")

}
```

## Performance

The main goals of this library are convenience and familiarity for programmers, but it should have reasonable performance for most projects. To see for yourself, run the benchmarks using `Invoke-Benchmarks`.

## Development

Please read the [Contribution Guide](./CONTRIBUTING.md) before you proceed. This project uses [Powershell Core](https://microsoft.com/PowerShell) to run tasks. To get started,

```ps1
Import-Module .\tasks.psm1
Get-Command -Module tasks
```

## Copyright

Copyright 2019-2020 [Vincent Fiestada](mailto:vincent@vincent.click). This project is released under a [BSD-style license](./license).

Icon made by <a href="https://www.flaticon.com/authors/good-ware" title="Good Ware">Good Ware</a>.