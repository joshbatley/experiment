package utils

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"reflect"
	"strings"
	"time"
)

const (
	LoggerTag = "log"
	Omit      = "omitempty"
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

func LogTags(obj interface{}) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() == reflect.Ptr {
		objValue = objValue.Elem()
	}
	objType := objValue.Type()

	l := log.Info()
	for i := 0; i < objType.NumField(); i++ {
		f := objType.Field(i)
		v := objValue.Field(i)
		if strings.Contains(f.Tag.Get(LoggerTag), Omit) && IsStructEmpty(v.Interface()) {
			continue
		}
		switch f.Type.Kind() {
		case reflect.String:
			l.Str(f.Name, v.String())
		case reflect.Float64:
			l.Float64(f.Name, v.Float())
		default:
			l.Interface(f.Name, v.Interface())
		}

	}
	l.Send()
}
