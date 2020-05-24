package middleware

import (
	"net/http"

	"mix/pkg/log"
	"mix/pkg/rand"
)

const (
	HeaderRequestId = "X-Request-Id"
)

// RequestId is a middleware that injects a request id into the context of each request.
func RequestId(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(HeaderRequestId)
		if requestId == "" {
			requestId = rand.String(16)
		}
		ctx := log.PutRequestId(r.Context(), requestId)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
