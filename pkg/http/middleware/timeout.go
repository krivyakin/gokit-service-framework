package middleware

import (
	"net/http"
	"time"
)

func NewTimeoutMiddleware(duration time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.TimeoutHandler(next, duration, "Timeout")
	}
}