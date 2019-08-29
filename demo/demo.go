package main

import (
	"github.com/vincentfiestada/captainslog"
	"github.com/vincentfiestada/captainslog/levels"
)

var log *captainslog.Logger

func init() {
	log = captainslog.NewLogger()
	log.Name = "captainslog"
	log.Level = levels.Trace
}

func main() {
	log.Trace("%d", 1)
	log.Debug("%d", 2)
	log.Info("%d", 3)
	log.Warn("%d", 4)
	log.Error("%d", 5)

	log.Fields(
		log.I("captain", "picard"),
	).Trace("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
	).Debug("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
	).Info("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
	).Warn("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
	).Error("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
		log.I("first officer", "riker"),
	).Info("starship enterprise")

}
