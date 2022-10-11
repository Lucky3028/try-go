package middlewares

import (
	"log"
	"net/http"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	code int
}

func NewLoggingResponseWriter(writer http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{ResponseWriter: writer, code: http.StatusOK}
}

func (writer *loggingResponseWriter) WriteHeader(code int) {
	writer.code = code
	writer.ResponseWriter.WriteHeader(code)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		traceId := newTraceId()
		logWriter := NewLoggingResponseWriter(writer)

		log.Printf("[%d]%s %s\n", traceId, request.RequestURI, request.Method)

		next.ServeHTTP(logWriter, request)

		log.Printf("[%d]res: %d", traceId, logWriter.code)
	})
}
