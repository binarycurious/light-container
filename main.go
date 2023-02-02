package main

import (
	"fmt"
	"sync"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/container"
	"github.com/binarycurious/light-container/routines"
	"github.com/binarycurious/light-container/telemetry"
)

func main() {
	s := (&container.GlobalState{}).NewState(&config.Settings{Hostname: "dummy-hostname", Hardfail: false, LogLevel: telemetry.LogLevelDebug})
	c := container.NewDefaultContainer(s)
	rn1 := "test-routine-1"
	rn2 := "test-routine-2"

	r := func(ctx container.Context) error {
		ctx.GetLogger().Log("This is a test log - 1, using ctx")
		msg := routines.NewMessage("testid", "testname", "test msg")

		k := ctx.GetRoutineKey(&rn2)

		ctx.Send(&k, container.RoutineMsg(&msg))

		return nil
	}

	cr, err := routines.NewStandardRoutine(rn1, r)
	if err != nil {
		c.GetLogger().LogError("Failed to register routine")
	}
	c.AddRoutine(cr)

	r2 := func(ctx container.Context) error {

		logger := ctx.GetLogger()
		logger.Log("This is a test log - 2, using ctx")

		ch, err := ctx.GetReceiver()

		if err != nil {
			s := fmt.Sprint(err)
			logger.LogError(s)
		}

		for ctx.ContainerIsRunning() {
			select {
			case msg := <-ch:
				logger.Log("Received msg on routine: " + ctx.GetRoutineName())
				logger.Log(fmt.Sprint((msg).GetMsg()))
			}
		}

		return nil
	}
	cr2, err := routines.NewStandardRoutine(rn2, r2)

	if err != nil {
		fmt.Printf(err.Error())
	}

	c.AddRoutine(cr2)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go c.Start(&wg)
	wg.Wait()

}
