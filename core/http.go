package core

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/daison12006013/gorvel/config"
	"github.com/gorilla/mux"
)

func HttpApplication() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	srv := &http.Server{
		Addr: config.HOST + ":" + config.PORT,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: config.WRITE_TIMEOUT,
		ReadTimeout:  config.READ_TIMEOUT,
		IdleTimeout:  config.IDLE_TIMEOUT,
		Handler:      routes(), // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}

func routes() *mux.Router {
	handler := mux.NewRouter()

	for r, c := range config.Routes {
		handler.HandleFunc(r, c.(func(http.ResponseWriter, *http.Request)))
	}

	return handler
}
