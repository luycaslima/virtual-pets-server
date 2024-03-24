package configs

import "net/http"

func CacheControlWrapper(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=3600")
		next.ServeHTTP(w, r)
	}
}
