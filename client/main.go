package main

import (
	"net/http"

	"github.com/muhammadhidayah/inact-grpc/client/delivery/user"
	"github.com/muhammadhidayah/inact-grpc/client/middleware"
)

func main() {
	customMux := new(middleware.GoMiddleware)
	mwFunc := new(middleware.MiddlewareFunc)

	customMux.RegisterMiddleware(mwFunc.CORS)
	customMux.RegisterMiddleware(middleware.MiddlewareJWTAuth)

	user.NewUserHandler(customMux)

	server := new(http.Server)
	server.Addr = ":9001"
	server.Handler = customMux
	server.ListenAndServe()
}
