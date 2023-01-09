package main

import (
	"fmt"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/container"
)

func main() {
	s := (&container.GlobalState{}).NewState(&config.Settings{Hostname: "dummy-hostname", Hardfail: false})
	c := container.GlobalContainer{}.NewDefaultContainer(s)
	c.AddRoutine()

	fmt.Println(c.GetState())

}
