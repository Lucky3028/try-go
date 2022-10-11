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
		log.Println(request.RequestURI, request.Method)

		logWriter := NewLoggingResponseWriter(writer)

		next.ServeHTTP(logWriter, request)

		log.Println("res:", logWriter.code)
	})
}
