package main

import (
	"net"
	"net/http"
)

func main() {

	var (
		host = "0.0.0.0"
		port = "9999"
	)
	addr := net.JoinHostPort(host, port)

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", PingHandler)

	panic(http.ListenAndServe(addr, mux))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	w.WriteHeader(200)
}
