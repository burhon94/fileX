package router

import (
	"github.com/burhon94/fileX/cmd/fileX/handlers"
	newRouter "github.com/burhon94/fileX/pkg/router"
	"net/http"
)

func InitRoute(addr string) {
	router := newRouter.NewRouter()

	router.GET("/ping", handlers.PingHandler)

	server := http.Server{Addr: addr, Handler: router.GetServeHTTP()}

	panic(server.ListenAndServe())
}
