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

func (tr *troutine) Execute(c container.Context) <-chan interface{} {
	panic("not implemented") // TODO: Implement
}

func main() {
	s := (&container.GlobalState{}).NewState(&config.Settings{Hostname: "dummy-hostname", Hardfail: false})
	c := container.GlobalContainer{}.NewDefaultContainer(s)
	fmt.Println(c.GetState())

}
