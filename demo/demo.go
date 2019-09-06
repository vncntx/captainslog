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

	log.Field("captain", "picard").Trace("starship enterprise")
	log.Field("captain", "picard").Debug("starship enterprise")
	log.Field("captain", "picard").Info("starship enterprise")
	log.Field("captain", "picard").Warn("starship enterprise")
	log.Field("captain", "picard").Error("starship enterprise")

	log.Fields(
		log.I("captain", "picard"),
		log.I("first officer", "riker"),
		log.I("science officer", "data"),
		log.I("medical officer", "crusher"),
		log.I("chief engineer", "la forge"),
		log.I("security officer", "worf"),
	).Info("starship enterprise")

}
