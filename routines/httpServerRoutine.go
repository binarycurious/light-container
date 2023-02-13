package routines

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/binarycurious/light-container/container"
)

const (
	shutdownSignal = "shutdown-signal"
)

type HttpServerRoutine struct {
	name     string
	server   *http.Server
	handlers map[string]http.HandlerFunc
}

// NewHttpServerRoutine - Create a new HttpServerRoutine for distributing http requests to routine channels
func NewHttpServerRoutine(hostname string, port int, readTimeout int, writeTimeout int) HttpServerRoutine {
	if readTimeout < 1 {
		readTimeout = 10
	}
	if writeTimeout < 1 {
		writeTimeout = 10
	}

	addr := fmt.Sprintf("%s:%d", hostname, port)
	srv := &http.Server{
		Addr:           addr,
		Handler:        nil,
		ReadTimeout:    time.Duration(readTimeout) * time.Second,
		WriteTimeout:   time.Duration(writeTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return HttpServerRoutine{server: srv, name: fmt.Sprintf("http-server_%s", addr)}
}

// RegisterHandlerFunc - register a function to handle incomming http requests
func (r *HttpServerRoutine) RegisterHandlerFunc(path string, handlerFunc http.HandlerFunc) {

	//TODO: Abstract the handler to allow channel pipeline wiring in http handler methods
	if r.handlers == nil {
		r.handlers = make(map[string]http.HandlerFunc, 5)
	}

	r.handlers[path] = handlerFunc
}

// GetName -
func (r *HttpServerRoutine) GetName() *string {
	return &r.name
}

// Execute - container execute impl
func (r *HttpServerRoutine) Execute(ctx container.Context) error {
	regWg := sync.WaitGroup{}
	regWg.Add(1)
	go func() {
		for k, v := range r.handlers {
			http.DefaultServeMux.Handle(k, v)
		}
		regWg.Done()

		// await shutdown signal
		for ctx.ContainerIsRunning() {
			ch, err := ctx.GetReceiver()
			if err != nil {
				ctx.GetLogger().LogError(err.Error())
				break
			}
			select {
			case msg, _ := <-ch:
				if *msg.GetName() == shutdownSignal {
					r.server.Shutdown(context.Background())
					return
				}
			}
		}
	}()

	regWg.Wait()
	ctx.GetLogger().Log(fmt.Sprintf("Starting HTTP server : %s : %s", ctx.GetRoutineName(), r.server.Addr))
	err := r.server.ListenAndServe()
	if err != http.ErrServerClosed {
		ctx.GetLogger().LogFatal(err.Error())
		return err
	}
	return nil
}
