package middleware

import (
	"context"
	"net/http"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/ctxkey"
)

// SetCookieMiddleware allows graphqls to set cookie
func SetCookieMiddleware(
	handler http.Handler,
) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(
				r.Context(),
				ctxkey.ResponseWriterKey,
				w,
			)
			handler.ServeHTTP(
				w,
				r.WithContext(ctx),
			)
		},
	)
}
