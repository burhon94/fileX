package health

import (
	"github.com/burhon94/fileX/internals/structs"
	"github.com/burhon94/fileX/internals/structs/responses"
	"strings"
)

func Pong(key string) (resp structs.Response) {

	if strings.TrimSpace(key) == "" {
		resp = responses.BadRequest
		return
	}

	resp = responses.Success
	resp.Payload = "your key: " + key

	return
}
