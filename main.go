package main

import (
	"fmt"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/container"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

type troutine struct{}

func (tr *troutine) Execute(routineKey *container.RoutineKey, c container.Context) <-chan interface{} {
	panic("not implemented") // TODO: Implement
}

func main() {

	fmt.Println(Hello("api-sneaky"))

	/*Init global state*/
	// gstate := container.GlobalState{}
	// container := container.GlobalContainer{}
	// // gstate.Init()
	// // container.SetState(gstate)

	s := (&container.GlobalState{}).NewState(&config.Settings{Hostname: "dummy-hostname", Hardfail: false})
	c := container.CreateDefaultGlobalContainer(s)
	fmt.Println(c.GetState())

}
