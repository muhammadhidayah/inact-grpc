package middleware

import "net/http"

type GoMiddleware struct {
	http.ServeMux
	middleware []func(next http.Handler) http.Handler
}

func (d *GoMiddleware) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	d.middleware = append(d.middleware, next)
}

func (d *GoMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &d.ServeMux

	for _, next := range d.middleware {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}
