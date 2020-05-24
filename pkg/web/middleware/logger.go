package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"

	"mix/pkg/log"
)

const (
	MsgHandleRequest = "handle http request"
)

// Logger is a middleware to log the http handlers.
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor) // save a response status

		next.ServeHTTP(ww, r)

		log.ExtractLogger(r.Context()).With(
			"uri", fmt.Sprintf("%s %s%s", r.Method, r.Host, r.RequestURI),
			"status", ww.Status(),
			"ip", r.RemoteAddr,
			"latency", time.Since(start),
		).Info(MsgHandleRequest)
	}
	return http.HandlerFunc(fn)
}
