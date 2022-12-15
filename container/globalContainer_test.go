package container

import (
	"fmt"
	"sync"
	"testing"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/telemetry"
)

func TestGlobalContainer_registerOutChannel(t *testing.T) {
	type fields struct {
		logger        telemetry.Logger
		state         *GlobalState
		routines      map[string]Routine
		inChans       map[string]chan interface{}
		outChans      map[string][]chan interface{}
		containerLock *sync.Mutex
	}
	type args struct {
		in0 *RoutineKey
		in1 chan<- interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := GlobalState{settings: config.Settings{Hardfail: false}}
			c := CreateDefaultGlobalContainer(&s)
			fmt.Print(c.GetState().intitialized)

		})
	}
}
