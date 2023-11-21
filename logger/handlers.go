package main

import (
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func intakeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Debug().Msg(string(body))
}
