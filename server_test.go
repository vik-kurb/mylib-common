package common

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoggingMiddleware(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)
	defer log.SetOutput(os.Stderr)

	called := false
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusTeapot)
		w.Write([]byte("I'm a teapot"))
	})

	handlerToTest := LoggingMiddleware(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.RemoteAddr = "127.0.0.1:1234"
	rr := httptest.NewRecorder()

	handlerToTest.ServeHTTP(rr, req)

	assert.Equal(t, rr.Code, http.StatusTeapot)

	logOutput := logBuffer.String()
	assert.Contains(t, logOutput, http.MethodGet)
	assert.Contains(t, logOutput, "/test")
	assert.Contains(t, logOutput, "418")
	assert.True(t, called)
}

func TestCORSMiddleware(t *testing.T) {
	corsAllowedOrigin := "http://localhost:5173"
	os.Setenv("CORS_ALLOWED_ORIGIN", corsAllowedOrigin)

	type testCase struct {
		name               string
		method             string
		expectedCalled     bool
		expectedStatusCode int
	}
	testCases := []testCase{
		{
			name:               "not_options_method",
			method:             http.MethodGet,
			expectedCalled:     true,
			expectedStatusCode: http.StatusTeapot,
		},
		{
			name:               "options_method",
			method:             http.MethodOptions,
			expectedCalled:     false,
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, tc := range testCases {
		called := false
		nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			called = true
			w.WriteHeader(http.StatusTeapot)
			w.Write([]byte("I'm a teapot"))
		})
		handlerToTest := CORSMiddleware(nextHandler)

		req := httptest.NewRequest(tc.method, "/test", nil)
		w := httptest.NewRecorder()

		handlerToTest.ServeHTTP(w, req)
		resp := w.Result()

		assert.Equal(t, w.Code, tc.expectedStatusCode)

		assert.Equal(t, resp.Header.Get("Access-Control-Allow-Origin"), corsAllowedOrigin)
		assert.Equal(t, resp.Header.Get("Access-Control-Allow-Methods"), "POST, OPTIONS, PUT, GET, DELETE")
		assert.Equal(t, resp.Header.Get("Access-Control-Allow-Headers"), "Content-Type, Authorization")
		assert.Equal(t, resp.Header.Get("Access-Control-Allow-Credentials"), "true")
		assert.Equal(t, called, tc.expectedCalled)
	}
}
