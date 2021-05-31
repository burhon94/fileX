package main

import (
	"github.com/burhon94/fileX/cmd/fileX/router"
	"net"
)

func main() {

	var (
		host = "0.0.0.0"
		port = "9999"
	)
	addr := net.JoinHostPort(host, port)

	router.InitRoute(addr)
}
