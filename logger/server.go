package main

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
)

type server struct {
	port int
	mux  *http.ServeMux
	srv  *http.Server
}

func newServer() *server {
	mux := http.NewServeMux()

	return &server{
		port: 9000,
		mux:  mux,
		srv: &http.Server{
			Addr:    ":9000",
			Handler: mux,
		},
	}
}

func (s *server) addHandler(url string, handler http.HandlerFunc) {
	s.mux.Handle(url, handler)
}

func (s *server) serve() {
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error().Msgf("ListenAndServe: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	<-stop
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Error().Msgf("server Shutdown Failed: %+v", err)
	}

	log.Info().Msg("server stopped gracefully")
}
