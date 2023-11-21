package main

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"shared/clients"
)

type Server struct {
	port int
	mux  *http.ServeMux
	srv  *http.Server
}

func NewServer() *Server {
	mux := http.NewServeMux()

	return &Server{
		port: 9000,
		mux:  mux,
		srv: &http.Server{
			Addr:    ":9000",
			Handler: mux,
		},
	}
}

func (s *Server) AddHandler(url string, handler http.HandlerFunc) {
	s.mux.Handle(url, handler)
}

func (s *Server) Serve() {
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
		log.Error().Msgf("Server Shutdown Failed: %+v", err)
	}

	log.Info().Msg("Server stopped gracefully")
}

func AuthMiddleware(next http.HandlerFunc, apiKey string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pKey := r.Header.Get(clients.APIKeyHeader)
		if pKey == apiKey {
			next(w, r)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
}
