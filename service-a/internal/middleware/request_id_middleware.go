package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type CtxKey int

const ridKey CtxKey = CtxKey(0)

func RequestId(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		rid := r.Header.Get("X-Request-ID")
		if rid == "" {
			rid = uuid.New().String()
		}
		ctx := context.WithValue(r.Context(), ridKey, rid)
		rw.Header().Add("X-Request-ID", rid)

		next.ServeHTTP(rw, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}
