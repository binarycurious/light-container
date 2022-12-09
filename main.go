package main

import (
	"fmt"

	"github.com/binarycurious/light-container/container"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func main() {

	fmt.Println(Hello("api-sneaky"))

	/*Init global state*/
	gstate := container.GlobalState{}
	container := container.GlobalContainer{}
	gstate.Init()
	container.SetState(gstate)

}
