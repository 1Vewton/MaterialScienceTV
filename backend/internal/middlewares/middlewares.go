package middlewares

import (
	"context"
	"net/http"

	"github.com/1Vewton/MaterialScienceTV/backend/internal/ctxkey"
)

// SetCookieMiddleware allows graphqls to set cookie
func SetCookieMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				ctx := context.WithValue(
					r.Context(),
					ctxkey.ResponseWriterKey,
					w,
				)
				next.ServeHTTP(
					w,
					r.WithContext(ctx),
				)
			},
		)
	}
}
