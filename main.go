package main

import (
	"fmt"
	"time"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/container"
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
		msg := container.NewMessage("testid", "testname", "test msg")

		r2key := ctx.GetRoutineKey(&rn2)

		ctx.Send(&r2key, container.RoutineMsg(&msg))

		r2ch, err := ctx.Subscribe(&r2key)
		if err != nil {
			logger.LogFatal("Failed to subscribe to r2 out channel")
		}

		for ctx.ContainerIsRunning() {
			select {
			case r2msg := <-r2ch:
				switch msgVal := r2msg.GetMsg().(type) {
				case string:
					logger.LogDebug("Message received: " + msgVal)
				}
			}
		}
		fmt.Printf("Go routine (%s) ending...", ctx.GetRoutineName())
		ctx.EndRoutine()
		return nil
	}

	cr, err := container.NewStandardRoutine(rn1, r)
	if err != nil {
		c.GetLogger().LogError("Failed to register routine")
	}
	_ = c.AddRoutine(cr)
	_ = c.AddStandardRoutine(rn2, func(ctx container.Context) error {

		logger := ctx.GetLogger()
		logger.Log("This is a test log - 2, using ctx")

		ch, err := ctx.GetReceiver()

		if err != nil {
			s := fmt.Sprint(err)
			logger.LogError(s)
		}

		select {
		case msg := <-ch:
			logger.Log("Received msg on routine: " + ctx.GetRoutineName())
			logger.Log(fmt.Sprint((msg).GetMsg()))
			msgOut := container.NewMessage("r2msgId:"+time.Now().String(), "r2Msg", "This is a message published from R2")
			ctx.Publish(&msgOut)
		}
		for ctx.ContainerIsRunning() {
		}

		fmt.Printf("Go routine (%s) ending...", ctx.GetRoutineName())
		ctx.EndRoutine()
		return nil
	})
	go c.Start()

	go func() {
		time.Sleep(time.Second * 5)
		c.Stop()
	}()

	c.Wait()

}
