package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler  {
	return http.HandlerFunc(
		func (write http.ResponseWriter, request *http.Request) {
			write.Header().Set("Content-type", "application/json")
			next.ServeHTTP(write, request)
		},
	)
}