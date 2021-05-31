package handlers

import (
	"github.com/burhon94/fileX/internals/health"
	"github.com/burhon94/fileX/internals/structs"
	"github.com/burhon94/fileX/reply"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	var resp structs.Response

	key := r.Header.Get("key")

	defer reply.Json(w, resp)

	resp = health.Pong(key)
}
