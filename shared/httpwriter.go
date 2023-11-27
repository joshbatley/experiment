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
	req, err := http.NewRequest("POST", w.url, bytes.NewBuffer(p))
	if err != nil {
		return 0, err
	}
	req.Header.Set("x-api-key", w.apiKey)
	c := &http.Client{}
	_, err = c.Do(req)
	if err != nil {
		return 0, err
	}
	return len(p), err
}
