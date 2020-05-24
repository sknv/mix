package middleware

import (
	"net/http"

	"github.com/pkg/errors"

	"mix/pkg/log"
	"mix/pkg/response"
	webresp "mix/pkg/web/response"
)

// Recoverer is a middleware to recover panic in http handlers.
func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil && rvr != http.ErrAbortHandler {
				log.ExtractLogger(r.Context()).
					Errorw("recover", "panic", errors.Errorf("%v", rvr)) // log the stacktrace
				webresp.RenderError(w, r, response.NewError(http.StatusInternalServerError, response.MsgSmoke))
			}
		}()

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
