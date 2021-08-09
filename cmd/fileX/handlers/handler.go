package handlers

import (
	"github.com/burhon94/fileX/internals/health"
	"github.com/burhon94/fileX/internals/structs"
	"github.com/burhon94/fileX/pkg/reply"
	"net/http"
)

//PingHandler
// @Tags Ping Service
// @Summary check service status
// @ID ping
// @Param key header string false "key Request"
// @Produce  json
// @Success 200 {object} structs.Response "OK"
// @Failure 400 {object} structs.Response "Неверные данные"
// @Failure 500 {object} structs.Response "Внутренняя ошибка сервера"
// @Router /ping [get]
func PingHandler(w http.ResponseWriter, r *http.Request) {
	var resp structs.Response

	key := r.Header.Get("key")

	defer reply.Json(w, &resp)

	resp = health.Pong(key)
}
