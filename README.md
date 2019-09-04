![](./logo.png)

# Captain's Log

[![Build Status](https://vincentofearth.visualstudio.com/captainslog/_apis/build/status/vincentfiestada.captainslog?branchName=dev)](https://vincentofearth.visualstudio.com/captainslog/_build/latest?definitionId=5&branchName=dev)
[![GoReportCard](https://goreportcard.com/badge/github.com/vincentfiestada/captainslog)](https://goreportcard.com/report/github.com/vincentfiestada/captainslog)
[![GoDoc](https://godoc.org/github.com/vincentfiestada/captainslog?status.svg)](https://godoc.org/github.com/vincentfiestada/captainslog)
[![Conventional Commits](https://img.shields.io/badge/commits-conventional-00b6ff.svg?labelColor=1F6CB4)](https://conventionalcommits.org)
[![License: BSD-3](https://img.shields.io/github/license/vincentfiestada/captainslog.svg?labelColor=1F6CB4&color=00b6ff)](https://github.com/vincentfiestada/captainslog/blob/master/LICENSE)

A simple logging library for [Go](https://golang.org/)

- Colors
- Multiple levels
- Structured logging
- Function name detection

![Screenshot of captainslog in action](./assets/screenshot.png)

## Usage

This library is designed to provide a familiar yet powerful interface for logging. Each logging method accepts a format string and arguments. Structured logging is supported right out of the box. The Logger allows you to turn colors on/off, specify a datetime format, set the logging threshold, and even provide your own function to control how logs are written.

```go
package main

import (
	"github.com/vincentfiestada/captainslog"
)

var log = captainslog.NewLogger()

func main() {

	log.Trace("this is %s", "captainslog")
	log.Debug("this is %s", "captainslog")
	log.Info("this is %s", "captainslog")
	log.Warn("this is %s", "captainslog")
	log.Error("this is %s", "captainslog")

	log.Fields(

		log.I("captain",          "picard"),
		log.I("first officer", 	  "riker"),
		log.I("science officer",  "data"),
		log.I("medical officer",  "crusher"),
		log.I("chief engineer",   "la forge"),
		log.I("security officer", "worf"),

	).Info("starship enterprise")

}
```

## Development

Please read the [Contribution Guide](./CONTRIBUTING.md) before you proceed.

This project uses [Task runner](https://taskfile.dev/). List all available tasks by running `task -l`. To get started, run:

```
task install
```

## Copyright

Copyright 2019 Vincent Fiestada. This project is released under a BSD-style [license](./LICENSE).
