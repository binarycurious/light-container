package main

import (
	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/container"
	"github.com/binarycurious/light-container/routines"
	"github.com/binarycurious/light-container/telemetry"
)

func main() {
	s := (&container.GlobalState{}).NewState(&config.Settings{Hostname: "dummy-hostname", Hardfail: false, LogLevel: telemetry.LogLevelDebug})
	c := container.GlobalContainer{}.NewDefaultContainer(s)
	ch := make(chan container.RoutineMsg)
	r := func(ctx container.Context, cIn chan<- container.RoutineMsg, cOut <-chan container.RoutineMsg) error {
		log := "This is a test log - 1, using ctx"
		ctx.GetLogger().Log(&log)

		return nil
	}
	cr, err := routines.StandardRoutine{}.NewRoutine("test1", ch, r)
	if err != nil {
		log := "Failed to register routine"
		c.GetLogger().LogError(&log)
	}
	c.AddRoutine(cr)
	c.Start()
}
