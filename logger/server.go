package main

import (
	"context"
	"errors"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
)

type Server struct {
	port int
	mux  *http.ServeMux
	srv  *http.Server
}

const (
	APIKeyHeader = "x-api-key"
	APIKeyTemp   = "JBSWY3DPEHPK3PXP"
)

func NewServer() *Server {
	mux := http.NewServeMux()
	mux.Handle("/intake", AuthMiddleware(IntakeHandler, APIKeyTemp))

	return &Server{
		port: 9000,
		mux:  mux,
		srv: &http.Server{
			Addr:    ":9000",
			Handler: mux,
		},
	}
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

	// Create a context with a timeout for the graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Call Shutdown to gracefully stop the server
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server Shutdown Failed: %+v", err)
	}

	log.Info().Msg("Server stopped gracefully")
}

func AuthMiddleware(next http.HandlerFunc, apiKey string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pKey := r.Header.Get(APIKeyHeader)
		if pKey == apiKey {
			next(w, r)
			return
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})
}
