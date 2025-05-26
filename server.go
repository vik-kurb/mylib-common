package common

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)
	responseData, err := json.Marshal(ErrorResponse{Error: msg})
	if err != nil {
		log.Print("Failed to build error response: ", err)
	}
	_, writeErr := w.Write(responseData)
	if writeErr != nil {
		log.Print("Failed to response: ", writeErr)
	}
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}, cookie *http.Cookie) {
	if cookie != nil {
		http.SetCookie(w, cookie)
	}
	w.WriteHeader(code)
	responseData, err := json.Marshal(payload)
	if err != nil {
		log.Print("Failed to build response: ", err)
	}
	_, writeErr := w.Write(responseData)
	if writeErr != nil {
		log.Print("Failed to response: ", writeErr)
	}
}

func CloseResponseBody(response *http.Response) {
	err := response.Body.Close()
	if err != nil {
		log.Printf("Failed to close response body: %v", err)
	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := &loggingResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		start := time.Now()
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		log.Printf("[%s] %s %s %d %s",
			r.Method,
			r.RemoteAddr,
			r.URL.Path,
			lrw.statusCode,
			duration,
		)
	})
}
