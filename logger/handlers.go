package main

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func IntakeHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Debug().Msg(string(body))
}

func generateAPIKey() (string, error) {
	// Define the length of the API key you want to generate
	apiKeyLength := 32

	// Create a byte slice with the specified length
	key := make([]byte, apiKeyLength)

	// Read random bytes into the byte slice
	if _, err := rand.Read(key); err != nil {
		return "", err
	}

	// Encode the byte slice to a hexadecimal string
	apiKey := hex.EncodeToString(key)

	return apiKey, nil
}
