package main

import "github.com/vincentfiestada/captainslog"

var log *captainslog.Logger

func init() {
	log = captainslog.NewLogger()
	log.Name = "captainslog"
	log.Level = captainslog.LevelTrace
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

}
