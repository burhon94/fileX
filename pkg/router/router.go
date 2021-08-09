package http

import (
	"github.com/burhon94/fileX/internals/structs"
	"github.com/burhon94/fileX/internals/structs/responses"
	"github.com/burhon94/fileX/pkg/reply"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() Router {
	return &HttpRouter{
		router: mux.NewRouter(),
	}
}

type Mw func(http.HandlerFunc) http.HandlerFunc

type Router interface {
	GET(pattern string, handler http.HandlerFunc, mw ...Mw)
	POST(pattern string, handler http.HandlerFunc, mw ...Mw)
	PUT(pattern string, handler http.HandlerFunc, mw ...Mw)
	DELETE(pattern string, handler http.HandlerFunc, mw ...Mw)

	PathPrefix(pattern string) *mux.Route
	GetServeHTTP() http.HandlerFunc
}

type HttpRouter struct {
	router *mux.Router
}

func panicCatcher(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var response structs.Response

		defer func() {
			if err := recover(); err != nil {
				response = responses.InternalErr
				response.Message = err.(string)
				reply.Json(w, &response)
				return
			}
		}()

		h(w, r)
	}
}

func (r *HttpRouter) ChainsMiddlewares(method string, handler http.HandlerFunc, mw ...Mw) http.HandlerFunc {

	//middleware chains
	mws := make([]Mw, 0)
	mws = append(mws, panicCatcher)
	mws = append(mws, mw...)

	for _, httpMW := range mws {
		handler = httpMW(handler)
	}
	return handler
}

func (r *HttpRouter) Handle(method string, pattern string, handler http.HandlerFunc, mw ...Mw) {
	handler = r.ChainsMiddlewares(method, handler, mw...)
	r.router.HandleFunc(pattern, handler).Methods(method)
}

func (r *HttpRouter) GET(pattern string, handler http.HandlerFunc, mw ...Mw) {
	r.Handle("GET", pattern, handler, mw...)
}

func (r *HttpRouter) POST(pattern string, handler http.HandlerFunc, mw ...Mw) {
	r.Handle("POST", pattern, handler, mw...)
}

func (r *HttpRouter) PUT(pattern string, handler http.HandlerFunc, mw ...Mw) {
	r.Handle("PUT", pattern, handler, mw...)
}

func (r *HttpRouter) DELETE(pattern string, handler http.HandlerFunc, mw ...Mw) {
	r.Handle("DELETE", pattern, handler, mw...)
}

func (r *HttpRouter) PathPrefix(pattern string) *mux.Route {
	return r.router.NewRoute().PathPrefix(pattern)
}

func (r *HttpRouter) GetServeHTTP() http.HandlerFunc {
	return r.router.ServeHTTP
}
