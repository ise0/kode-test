package middlewares

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/ise0/kode-test/src/shared/authjwt"
	"github.com/ise0/kode-test/src/shared/logger"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, _ := strings.CutPrefix(r.Header.Get("Authorization"), "Bearer ")
		userId, err := authjwt.ParseAuthJwt(jwt)
		if err != nil {
			http.Error(w, "something went wrong", 500)
			return
		}
		ctx := context.WithValue(r.Context(), "userId", userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Logger(h http.Handler) http.Handler {
	loggingFn := func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		ww := middleware.NewWrapResponseWriter(rw, req.ProtoMajor)

		h.ServeHTTP(ww, req)

		duration := time.Since(start)
		logger.L.Infow("request completed",
			"method", req.Method,
			"uri", req.RequestURI,
			"status", ww.Status(),
			"duration", duration,
			"size", ww.BytesWritten(),
		)
	}
	return http.HandlerFunc(loggingFn)
}
