package utils

import (
	"bytes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"os"
	"time"
)

type Logger struct {
	Writers []io.Writer
	Tags    []string
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) WithTags(tags []string) *Logger {
	l.Tags = tags
	return l
}

func (l *Logger) WithWriter(writer io.Writer) *Logger {
	l.Writers = append(l.Writers, writer)
	return l
}

func (l *Logger) WithHttpLogger(url string, apiKey string) *Logger {
	l.Writers = append(l.Writers, NewHttpWriter(apiKey, url))
	return l
}

func (l *Logger) WithConsoleLogger() *Logger {
	l.Writers = append(l.Writers, zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
	return l
}

func (l *Logger) Build() {
	outputs := zerolog.MultiLevelWriter(l.Writers...)
	log.Logger = zerolog.New(outputs).With().Strs("tags", l.Tags).Timestamp().Logger()
}

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
