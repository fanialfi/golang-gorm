package middleware

import "net/http"

const (
	USERNAME = "fanialfi"
	PASSWORD = "saichiopy"
)

func MiddlewareBasicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "somethin wrong", http.StatusUnauthorized)
			return
		}

		isValid := (username == USERNAME) && (password == PASSWORD)
		if !isValid {
			http.Error(w, "username or password wrong", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
