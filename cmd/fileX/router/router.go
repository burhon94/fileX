package router

import (
	"github.com/burhon94/fileX/cmd/fileX/handlers"
	"net/http"
)

func InitRoute(addr string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/ping", handlers.PingHandler)

	panic(http.ListenAndServe(addr, mux))
}
