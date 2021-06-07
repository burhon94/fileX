package reply

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data interface{}) {
	reply, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(reply)

}
