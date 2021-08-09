package router

import (
	"github.com/burhon94/fileX/cmd/fileX/handlers"
	_ "github.com/burhon94/fileX/docs"
	newRouter "github.com/burhon94/fileX/pkg/router"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

//InitRoute
// @title file-X Project API
// @version 0.0.2
// @host localhost:9999
// @BasePath /
func InitRoute(addr string) {
	router := newRouter.NewRouter()

	router.GET("/ping", handlers.PingHandler)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler).Methods("GET")

	server := http.Server{Addr: addr, Handler: router.GetServeHTTP()}

	panic(server.ListenAndServe())
}
