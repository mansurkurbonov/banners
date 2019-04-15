package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crucial/banner/app/config"
	"crucial/banner/app/http/router"
)

// Listen -
func Listen() {
	var (
		cfg    = config.Peek().Server
		quitCh = make(chan os.Signal)
		srv    *http.Server
	)

	srv = &http.Server{
		Addr:         cfg.Port,
		Handler:      router.Peek(),
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
	}

	signal.Notify(quitCh, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	go graceShutdown(quitCh, srv)

	log.Printf("server: starting on 127.0.0.1%s\n", cfg.Port)
	log.Fatalln(srv.ListenAndServe())
}

func graceShutdown(quitCh chan os.Signal, srv *http.Server) {
	var (
		ctx        context.Context
		cancelFunc context.CancelFunc
		dur        time.Time
		err        error
	)

	s := <-quitCh
	log.Printf("server: received signal %+v\n", s)

	dur = time.Now().Add(30 * time.Second) // dummy deadline
	ctx, cancelFunc = context.WithDeadline(context.Background(), dur)
	defer cancelFunc()

	err = srv.Shutdown(ctx)
	if err != nil {
		log.Panicln("server: couldn't shutdown because of " + err.Error())
	}
}
