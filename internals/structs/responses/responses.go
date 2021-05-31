package responses

import "github.com/burhon94/fileX/internals/structs"

func newResponse(code int, message string) structs.Response {
	return structs.Response{Code: code, Message: message}
}

const (
	OkCode  = 200
	BadCode = 400
)

var (
	Success    = newResponse(OkCode, "Success")
	BadRequest = newResponse(BadCode, "BadRequest")
)
