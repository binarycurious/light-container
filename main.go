package main

import (
	"fmt"
	"sync"
	"time"

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
		logger := ctx.GetLogger()
		logger.Log("This is a test log - 1, using ctx")
		msg := routines.NewMessage("testid", "testname", "test msg")

		r2key := ctx.GetRoutineKey(&rn2)

		ctx.Send(&r2key, container.RoutineMsg(&msg))

		for ctx.ContainerIsRunning() {
			r2ch, err := ctx.Subscribe(&r2key)
			if err != nil {
				logger.LogFatal("Failed to subscribe to r2 out channel")
			}
			select {
			case r2msg := <-r2ch:
				switch msgVal := r2msg.GetMsg().(type) {
				case string:
					logger.LogDebug("Message received: " + msgVal)
				}
			}
		}

		return nil
	}

	cr, err := routines.NewStandardRoutine(rn1, r)
	if err != nil {
		c.GetLogger().LogError("Failed to register routine")
	}
	_ = c.AddRoutine(cr)

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
				time.Sleep(1000)

				msgOut := routines.NewMessage("r2msgId:"+time.Now().String(), "r2Msg", "This is a message published from R2")
				ctx.Publish(&msgOut)
			}
		}

		return nil
	}
	cr2, err := routines.NewStandardRoutine(rn2, r2)

	if err != nil {
		fmt.Printf(err.Error())
	}

	_ = c.AddRoutine(cr2)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go c.Start(&wg)
	wg.Wait()

}
