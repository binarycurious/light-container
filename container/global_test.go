package container

import (
	"bytes"
	"log"
	"reflect"
	"testing"

	"github.com/binarycurious/light-container/config"
	"github.com/binarycurious/light-container/telemetry"
)

func TestGlobalState_CantReInit(t *testing.T) {
	t.Run("Can't initialize the global state twice", func(t *testing.T) {
		var buf bytes.Buffer
		log.SetOutput(&buf)

		gs := &GlobalState{}

		if gs.intitialized {
			t.Error("State should not be initialized!")
		}

		gs.NewState(&config.Settings{}, false)

		if !gs.intitialized {
			t.Error("State should be initialized at this point")
		}

		gs.NewState(&config.Settings{}, false)
		if buf.String() != "FATAL: Attempt to set global state after already intialized" {
			t.Error("Should have error on attempt to reinitialize")
		}
	})
}

func TestGlobalState_Init(t *testing.T) {
	type fields struct {
		intitialized bool
		settings     config.Settings
	}
	type args struct {
		s config.Settings
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Enforce intitialized",
		},
		{
			name: "Enforce hostname",
			args: args{s: config.Settings{
				Hostname: "testing1",
			},
			},
		},
		{
			name: "Enforce hostname 2",
			fields: fields{
				settings: config.Settings{
					Hostname: "nottherighthostname",
				},
			},
			args: args{s: config.Settings{
				Hostname: "testing2",
			},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gs := &GlobalState{
				intitialized: tt.fields.intitialized,
				settings:     tt.fields.settings,
			}
			gs.Init(tt.args.s)
			if !gs.intitialized {
				t.Error("Global State Not Initialized!")
			}
			if gs.settings.Hostname != tt.args.s.Hostname {
				t.Error("Hostname not matching")
			}
		})
	}
}

func TestGlobalLogger_Log(t *testing.T) {
	type fields struct {
		loglevel telemetry.LogLevel
	}
	type args struct {
		log *string
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
			l := &GlobalLogger{
				loglevel: tt.fields.loglevel,
			}
			l.Log(tt.args.log)
		})
	}
}

func TestGlobalLogger_LogDebug(t *testing.T) {
	type fields struct {
		loglevel telemetry.LogLevel
	}
	type args struct {
		log *string
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
			l := &GlobalLogger{
				loglevel: tt.fields.loglevel,
			}
			l.LogDebug(tt.args.log)
		})
	}
}

func TestGlobalLogger_LogWarn(t *testing.T) {
	type fields struct {
		loglevel telemetry.LogLevel
	}
	type args struct {
		log *string
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
			l := &GlobalLogger{
				loglevel: tt.fields.loglevel,
			}
			l.LogWarn(tt.args.log)
		})
	}
}

func TestGlobalLogger_LogError(t *testing.T) {
	type fields struct {
		loglevel telemetry.LogLevel
	}
	type args struct {
		log *string
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
			l := &GlobalLogger{
				loglevel: tt.fields.loglevel,
			}
			l.LogError(tt.args.log)
		})
	}
}

func TestGlobalLogger_SetActiveLogLevel(t *testing.T) {
	type fields struct {
		loglevel telemetry.LogLevel
	}
	type args struct {
		lvl telemetry.LogLevel
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
			l := &GlobalLogger{
				loglevel: tt.fields.loglevel,
			}
			l.SetActiveLogLevel(tt.args.lvl)
		})
	}
}

func TestGlobalContainer_SetState(t *testing.T) {
	type fields struct {
		state    GlobalState
		routines map[string]ContainerRoutine
		logger   telemetry.Logger
	}
	type args struct {
		s GlobalState
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
			c := &GlobalContainer{
				state:    tt.fields.state,
				routines: tt.fields.routines,
				logger:   tt.fields.logger,
			}
			c.SetState(tt.args.s)
		})
	}
}

func TestGlobalContainer_Execute(t *testing.T) {
	type fields struct {
		state    GlobalState
		routines map[string]ContainerRoutine
		logger   telemetry.Logger
	}
	type args struct {
		key string
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
			c := &GlobalContainer{
				state:    tt.fields.state,
				routines: tt.fields.routines,
				logger:   tt.fields.logger,
			}
			c.Execute(tt.args.key)
		})
	}
}

func TestGlobalContainer_AddRoutine(t *testing.T) {
	type fields struct {
		state    GlobalState
		routines map[string]ContainerRoutine
		logger   telemetry.Logger
	}
	type args struct {
		key     string
		routine ContainerRoutine
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
			c := &GlobalContainer{
				state:    tt.fields.state,
				routines: tt.fields.routines,
				logger:   tt.fields.logger,
			}
			c.AddRoutine(tt.args.key, tt.args.routine)
		})
	}
}

func TestGlobalContainer_GetState(t *testing.T) {
	type fields struct {
		state    GlobalState
		routines map[string]ContainerRoutine
		logger   telemetry.Logger
	}
	tests := []struct {
		name   string
		fields fields
		want   GlobalState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &GlobalContainer{
				state:    tt.fields.state,
				routines: tt.fields.routines,
				logger:   tt.fields.logger,
			}
			if got := c.GetState(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GlobalContainer.GetState() = %v, want %v", got, tt.want)
			}
		})
	}
}
