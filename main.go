package requestcontext

import (
	"context"
	"net/http"
)

type requestID struct{}

var requestKey = &requestID{}

func Middleware(handler http.Handler, key string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(key)
		if id != "" {
			r = r.WithContext(context.WithValue(r.Context(), requestKey, id))
		}
		handler.ServeHTTP(w, r)
	})
}

func IDFromContext(ctx context.Context) string {
	v := ctx.Value(requestKey)
	if v == nil {
		return ""
	}
	return v.(string)
}
