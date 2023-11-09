package utils

import (
	"bytes"
	"net/http"
)

type HttpWriter struct {
	apiKey string
	url    string
}

func NewHttpWriter(apiKey string, url string) *HttpWriter {
	return &HttpWriter{
		apiKey,
		url,
	}
}

func (w HttpWriter) Write(p []byte) (n int, err error) {
	_, err = http.Post(w.url, "application/json", bytes.NewBuffer(p))
	return len(p), err
}
