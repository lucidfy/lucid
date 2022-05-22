package kernel

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/lucidfy/lucid/pkg/facade/logger"
	"github.com/lucidfy/lucid/pkg/facade/urls"
	"github.com/lucidfy/lucid/pkg/helpers"
	"github.com/lucidfy/lucid/registrar"
)

type App struct {
	Server   *http.Server
	Deadline time.Duration
}

func New() *App {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	var handler http.Handler

	engine := helpers.Getenv("LUCID_ROUTER_ENGINE", "mux")
	engine_handler, ok := registrar.Engines[engine]
	if !ok {
		panic(fmt.Errorf(`%s engine does not exists`, engine))
	}
	handler = engine_handler()

	write_timeout, _ := strconv.Atoi(helpers.Getenv("LUCID_WRITE_TIMEOUT", "10"))
	read_timeout, _ := strconv.Atoi(helpers.Getenv("LUCID_READ_TIMEOUT", "10"))
	idle_timeout, _ := strconv.Atoi(helpers.Getenv("LUCID_IDLE_TIMEOUT", "60"))

	srv := &http.Server{
		Addr:         urls.GetAddr(),
		WriteTimeout: time.Second * time.Duration(write_timeout),
		ReadTimeout:  time.Second * time.Duration(read_timeout),
		IdleTimeout:  time.Second * time.Duration(idle_timeout),
		Handler:      handler,
	}

	return &App{
		Server:   srv,
		Deadline: wait,
	}
}

func (h *App) Run() *App {
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			logger.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	return h
}

func (h *App) WithGracefulShutdown() *App {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), h.Deadline)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	h.Server.Shutdown(ctx)

	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)

	return h
}
